package ssh

import (
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"golang.org/x/crypto/ssh/knownhosts"
)

// DefaultKnownHostsPath returns default user knows hosts file.
func DefaultKnownHostsPath() (string, error) {
	__obf_f689976da6a746cc, __obf_977f2a9e62be94d2 := os.UserHomeDir()
	if __obf_977f2a9e62be94d2 != nil {
		return "", __obf_977f2a9e62be94d2
	}

	return fmt.Sprintf("%s/.ssh/known_hosts", __obf_f689976da6a746cc), __obf_977f2a9e62be94d2
}

// DefaultKnownHosts returns host key callback from default known hosts path, and error if some.
func DefaultKnownHosts() (ssh.HostKeyCallback, error) {
	__obf_487fba6b4a924660, __obf_977f2a9e62be94d2 := DefaultKnownHostsPath()
	if __obf_977f2a9e62be94d2 != nil {
		return nil, __obf_977f2a9e62be94d2
	}

	return knownhosts.New(__obf_487fba6b4a924660)
}

// CheckKnownHost checks is host in known hosts file.
// it returns is the host found in known_hosts file and error, if the host found in
// known_hosts file and error not nil that means public key mismatch, maybe MAN IN THE MIDDLE ATTACK! you should not handshake.
func CheckKnownHost(__obf_ec889ac3933152c2 string, __obf_37760f54a739cb2c net.Addr, __obf_477ac5dcccce5559 ssh.PublicKey, __obf_32027a9e22484f58 string) (__obf_1eac8943340ec9ec bool, __obf_977f2a9e62be94d2 error) {

	var __obf_09c8b7b63f9ffe17 *knownhosts.KeyError

	// Fallback to default known_hosts file
	if __obf_32027a9e22484f58 == "" {
		__obf_487fba6b4a924660, __obf_977f2a9e62be94d2 := DefaultKnownHostsPath()
		if __obf_977f2a9e62be94d2 != nil {
			return false, __obf_977f2a9e62be94d2
		}
		__obf_32027a9e22484f58 = __obf_487fba6b4a924660
	}
	__obf_86e2394e990481dd,

		// Get host key callback
		__obf_977f2a9e62be94d2 := knownhosts.New(__obf_32027a9e22484f58)

	if __obf_977f2a9e62be94d2 != nil {
		return false, __obf_977f2a9e62be94d2
	}
	__obf_977f2a9e62be94d2 = // check if host already exists.
		__obf_86e2394e990481dd(__obf_ec889ac3933152c2, __obf_37760f54a739cb2c,

			// Known host already exists.
			__obf_477ac5dcccce5559)

	if __obf_977f2a9e62be94d2 == nil {
		return true, nil
	}

	// Make sure that the error returned from the callback is host not in file error.
	// If keyErr.Want is greater than 0 length, that means host is in file with different key.
	if errors.As(__obf_977f2a9e62be94d2, &__obf_09c8b7b63f9ffe17) && len(__obf_09c8b7b63f9ffe17.Want) > 0 {
		return true, __obf_09c8b7b63f9ffe17
	}

	// Key is not trusted because it is not in the file.
	return false, nil
}

// AddKnownHost add a a host to known hosts file.
func AddKnownHost(__obf_ec889ac3933152c2 string, __obf_37760f54a739cb2c net.Addr, __obf_477ac5dcccce5559 ssh.PublicKey, __obf_32027a9e22484f58 string) (__obf_977f2a9e62be94d2 error) {

	// Fallback to default known_hosts file
	if __obf_32027a9e22484f58 == "" {
		__obf_487fba6b4a924660, __obf_977f2a9e62be94d2 := DefaultKnownHostsPath()
		if __obf_977f2a9e62be94d2 != nil {
			return __obf_977f2a9e62be94d2
		}
		__obf_32027a9e22484f58 = __obf_487fba6b4a924660
	}
	__obf_509a944eb949a8ff, __obf_977f2a9e62be94d2 := os.OpenFile(__obf_32027a9e22484f58, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if __obf_977f2a9e62be94d2 != nil {
		return __obf_977f2a9e62be94d2
	}

	defer __obf_509a944eb949a8ff.Close()
	__obf_d0a9204669c39fd0 := knownhosts.Normalize(__obf_37760f54a739cb2c.String())
	__obf_de63d1d294a6aca3 := knownhosts.Normalize(__obf_ec889ac3933152c2)
	__obf_278b06b0dea33d85 := []string{__obf_d0a9204669c39fd0}

	if __obf_de63d1d294a6aca3 != __obf_d0a9204669c39fd0 {
		__obf_278b06b0dea33d85 = append(__obf_278b06b0dea33d85, __obf_de63d1d294a6aca3)
	}

	_, __obf_977f2a9e62be94d2 = __obf_509a944eb949a8ff.WriteString(knownhosts.Line(__obf_278b06b0dea33d85, __obf_477ac5dcccce5559) + "\n")

	return __obf_977f2a9e62be94d2
}

// PrivateKey Loads a private and public key from "path" and returns a SSH ClientConfig to authenticate with the server
func PrivateKey(__obf_a20d55265c935340 string, __obf_487fba6b4a924660 string, __obf_98642ab04448ec31 int) (*ssh.ClientConfig, error) {
	__obf_8260d857dc5431ea, __obf_977f2a9e62be94d2 := os.ReadFile(__obf_487fba6b4a924660)

	if __obf_977f2a9e62be94d2 != nil {
		return nil, __obf_977f2a9e62be94d2
	}
	__obf_c78f0fc94a9c1265, __obf_977f2a9e62be94d2 := ssh.ParsePrivateKey(__obf_8260d857dc5431ea)

	if __obf_977f2a9e62be94d2 != nil {
		return nil, __obf_977f2a9e62be94d2
	}
	__obf_86e2394e990481dd, __obf_977f2a9e62be94d2 := DefaultKnownHosts()
	if __obf_977f2a9e62be94d2 != nil {
		return nil, __obf_977f2a9e62be94d2
	}

	return &ssh.ClientConfig{
		User:    __obf_a20d55265c935340,
		Timeout: time.Duration(__obf_98642ab04448ec31) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_c78f0fc94a9c1265),
		},
		HostKeyCallback: __obf_86e2394e990481dd,
	}, nil
}

// Creates the configuration for a client that authenticates with a password protected private key
func PrivateKeyWithPassphrase(__obf_a20d55265c935340 string, __obf_98642ab04448ec31 int, __obf_793f32cc3c055fb4 []byte, __obf_487fba6b4a924660 string) (*ssh.ClientConfig, error) {
	__obf_8260d857dc5431ea, __obf_977f2a9e62be94d2 := os.ReadFile(__obf_487fba6b4a924660)

	if __obf_977f2a9e62be94d2 != nil {
		return nil, __obf_977f2a9e62be94d2
	}
	__obf_c78f0fc94a9c1265, __obf_977f2a9e62be94d2 := ssh.ParsePrivateKeyWithPassphrase(__obf_8260d857dc5431ea, __obf_793f32cc3c055fb4)

	if __obf_977f2a9e62be94d2 != nil {
		return nil, __obf_977f2a9e62be94d2
	}
	__obf_86e2394e990481dd, __obf_977f2a9e62be94d2 := DefaultKnownHosts()
	if __obf_977f2a9e62be94d2 != nil {
		return nil, __obf_977f2a9e62be94d2
	}

	return &ssh.ClientConfig{
		User:    __obf_a20d55265c935340,
		Timeout: time.Duration(__obf_98642ab04448ec31) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_c78f0fc94a9c1265),
		},
		HostKeyCallback: __obf_86e2394e990481dd,
	}, nil
}

// Creates a configuration for a client that fetches public-private key from the SSH agent for authentication
func SSHAgent(__obf_a20d55265c935340 string, __obf_98642ab04448ec31 int, __obf_29dce3a005fec24f ssh.HostKeyCallback) (*ssh.ClientConfig, error) {
	__obf_308b5d7aeb6d10f5 := os.Getenv("SSH_AUTH_SOCK")
	__obf_0b123847a66ece7a, __obf_977f2a9e62be94d2 := net.Dial("unix", __obf_308b5d7aeb6d10f5)
	if __obf_977f2a9e62be94d2 != nil {
		return &ssh.ClientConfig{
			Timeout: time.Duration(__obf_98642ab04448ec31) * time.Second,
		}, __obf_977f2a9e62be94d2
	}

	return &ssh.ClientConfig{
		User:    __obf_a20d55265c935340,
		Timeout: time.Duration(__obf_98642ab04448ec31) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agent.NewClient(__obf_0b123847a66ece7a).Signers),
		},
		HostKeyCallback: __obf_29dce3a005fec24f,
	}, nil
}

// Creates a configuration for a client that authenticates using username and password
func PasswordKey(__obf_a20d55265c935340 string, __obf_b85999ea9ceec719 string, __obf_98642ab04448ec31 int) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User:    __obf_a20d55265c935340,
		Timeout: time.Duration(__obf_98642ab04448ec31) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.Password(__obf_b85999ea9ceec719),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

type WriteCounter struct {
	Total int64 // Total # of bytes written
}

func (__obf_8701428cc870ec5e *WriteCounter) Write(__obf_8834a1506f637ab0 []byte) (int, error) {
	__obf_c6c390b4558e8c1e := len(__obf_8834a1506f637ab0)
	__obf_8701428cc870ec5e.
		Total += int64(__obf_c6c390b4558e8c1e)
	fmt.Fprintf(os.Stdout, "%.2f kb transferred\n", float64(__obf_8701428cc870ec5e.Total/1024))
	return __obf_c6c390b4558e8c1e, nil
}
