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
	__obf_768db4dbd8722dfc, __obf_649f0be20cc0940a := os.UserHomeDir()
	if __obf_649f0be20cc0940a != nil {
		return "", __obf_649f0be20cc0940a
	}

	return fmt.Sprintf("%s/.ssh/known_hosts", __obf_768db4dbd8722dfc), __obf_649f0be20cc0940a
}

// DefaultKnownHosts returns host key callback from default known hosts path, and error if some.
func DefaultKnownHosts() (ssh.HostKeyCallback, error) {
	__obf_c3aa46a245d158e2, __obf_649f0be20cc0940a := DefaultKnownHostsPath()
	if __obf_649f0be20cc0940a != nil {
		return nil, __obf_649f0be20cc0940a
	}

	return knownhosts.New(__obf_c3aa46a245d158e2)
}

// CheckKnownHost checks is host in known hosts file.
// it returns is the host found in known_hosts file and error, if the host found in
// known_hosts file and error not nil that means public key mismatch, maybe MAN IN THE MIDDLE ATTACK! you should not handshake.
func CheckKnownHost(__obf_b0fc237688628399 string, __obf_a6114148307b6fb2 net.Addr, __obf_6fc9fa2e6d447fb3 ssh.PublicKey, __obf_a8caf1d5839d8ea5 string) (__obf_4ea7bd14445a9fd6 bool, __obf_649f0be20cc0940a error) {

	var __obf_b278c29f82b1c507 *knownhosts.KeyError

	// Fallback to default known_hosts file
	if __obf_a8caf1d5839d8ea5 == "" {
		__obf_c3aa46a245d158e2, __obf_649f0be20cc0940a := DefaultKnownHostsPath()
		if __obf_649f0be20cc0940a != nil {
			return false, __obf_649f0be20cc0940a
		}
		__obf_a8caf1d5839d8ea5 = __obf_c3aa46a245d158e2
	}
	__obf_88833abc8f8ecebb,

		// Get host key callback
		__obf_649f0be20cc0940a := knownhosts.New(__obf_a8caf1d5839d8ea5)

	if __obf_649f0be20cc0940a != nil {
		return false, __obf_649f0be20cc0940a
	}
	__obf_649f0be20cc0940a = // check if host already exists.
		__obf_88833abc8f8ecebb(__obf_b0fc237688628399, __obf_a6114148307b6fb2,

			// Known host already exists.
			__obf_6fc9fa2e6d447fb3)

	if __obf_649f0be20cc0940a == nil {
		return true, nil
	}

	// Make sure that the error returned from the callback is host not in file error.
	// If keyErr.Want is greater than 0 length, that means host is in file with different key.
	if errors.As(__obf_649f0be20cc0940a, &__obf_b278c29f82b1c507) && len(__obf_b278c29f82b1c507.Want) > 0 {
		return true, __obf_b278c29f82b1c507
	}

	// Key is not trusted because it is not in the file.
	return false, nil
}

// AddKnownHost add a a host to known hosts file.
func AddKnownHost(__obf_b0fc237688628399 string, __obf_a6114148307b6fb2 net.Addr, __obf_6fc9fa2e6d447fb3 ssh.PublicKey, __obf_a8caf1d5839d8ea5 string) (__obf_649f0be20cc0940a error) {

	// Fallback to default known_hosts file
	if __obf_a8caf1d5839d8ea5 == "" {
		__obf_c3aa46a245d158e2, __obf_649f0be20cc0940a := DefaultKnownHostsPath()
		if __obf_649f0be20cc0940a != nil {
			return __obf_649f0be20cc0940a
		}
		__obf_a8caf1d5839d8ea5 = __obf_c3aa46a245d158e2
	}
	__obf_c2b198c77d92d796, __obf_649f0be20cc0940a := os.OpenFile(__obf_a8caf1d5839d8ea5, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if __obf_649f0be20cc0940a != nil {
		return __obf_649f0be20cc0940a
	}

	defer __obf_c2b198c77d92d796.Close()
	__obf_4f9a61be17e28363 := knownhosts.Normalize(__obf_a6114148307b6fb2.String())
	__obf_79caf2776f966afd := knownhosts.Normalize(__obf_b0fc237688628399)
	__obf_f0b68d3466acb4de := []string{__obf_4f9a61be17e28363}

	if __obf_79caf2776f966afd != __obf_4f9a61be17e28363 {
		__obf_f0b68d3466acb4de = append(__obf_f0b68d3466acb4de, __obf_79caf2776f966afd)
	}

	_, __obf_649f0be20cc0940a = __obf_c2b198c77d92d796.WriteString(knownhosts.Line(__obf_f0b68d3466acb4de, __obf_6fc9fa2e6d447fb3) + "\n")

	return __obf_649f0be20cc0940a
}

// PrivateKey Loads a private and public key from "path" and returns a SSH ClientConfig to authenticate with the server
func PrivateKey(__obf_85be5df09f3bfe87 string, __obf_c3aa46a245d158e2 string, __obf_ccc767b877765d54 int) (*ssh.ClientConfig, error) {
	__obf_6db75581e4c7be7d, __obf_649f0be20cc0940a := os.ReadFile(__obf_c3aa46a245d158e2)

	if __obf_649f0be20cc0940a != nil {
		return nil, __obf_649f0be20cc0940a
	}
	__obf_abd50081fc2af301, __obf_649f0be20cc0940a := ssh.ParsePrivateKey(__obf_6db75581e4c7be7d)

	if __obf_649f0be20cc0940a != nil {
		return nil, __obf_649f0be20cc0940a
	}
	__obf_88833abc8f8ecebb, __obf_649f0be20cc0940a := DefaultKnownHosts()
	if __obf_649f0be20cc0940a != nil {
		return nil, __obf_649f0be20cc0940a
	}

	return &ssh.ClientConfig{
		User:    __obf_85be5df09f3bfe87,
		Timeout: time.Duration(__obf_ccc767b877765d54) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_abd50081fc2af301),
		},
		HostKeyCallback: __obf_88833abc8f8ecebb,
	}, nil
}

// Creates the configuration for a client that authenticates with a password protected private key
func PrivateKeyWithPassphrase(__obf_85be5df09f3bfe87 string, __obf_ccc767b877765d54 int, __obf_231724c701fb51fe []byte, __obf_c3aa46a245d158e2 string) (*ssh.ClientConfig, error) {
	__obf_6db75581e4c7be7d, __obf_649f0be20cc0940a := os.ReadFile(__obf_c3aa46a245d158e2)

	if __obf_649f0be20cc0940a != nil {
		return nil, __obf_649f0be20cc0940a
	}
	__obf_abd50081fc2af301, __obf_649f0be20cc0940a := ssh.ParsePrivateKeyWithPassphrase(__obf_6db75581e4c7be7d, __obf_231724c701fb51fe)

	if __obf_649f0be20cc0940a != nil {
		return nil, __obf_649f0be20cc0940a
	}
	__obf_88833abc8f8ecebb, __obf_649f0be20cc0940a := DefaultKnownHosts()
	if __obf_649f0be20cc0940a != nil {
		return nil, __obf_649f0be20cc0940a
	}

	return &ssh.ClientConfig{
		User:    __obf_85be5df09f3bfe87,
		Timeout: time.Duration(__obf_ccc767b877765d54) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_abd50081fc2af301),
		},
		HostKeyCallback: __obf_88833abc8f8ecebb,
	}, nil
}

// Creates a configuration for a client that fetches public-private key from the SSH agent for authentication
func SSHAgent(__obf_85be5df09f3bfe87 string, __obf_ccc767b877765d54 int, __obf_536fa713ca86126f ssh.HostKeyCallback) (*ssh.ClientConfig, error) {
	__obf_2efc16620653c75e := os.Getenv("SSH_AUTH_SOCK")
	__obf_32c1d131a468249b, __obf_649f0be20cc0940a := net.Dial("unix", __obf_2efc16620653c75e)
	if __obf_649f0be20cc0940a != nil {
		return &ssh.ClientConfig{
			Timeout: time.Duration(__obf_ccc767b877765d54) * time.Second,
		}, __obf_649f0be20cc0940a
	}

	return &ssh.ClientConfig{
		User:    __obf_85be5df09f3bfe87,
		Timeout: time.Duration(__obf_ccc767b877765d54) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agent.NewClient(__obf_32c1d131a468249b).Signers),
		},
		HostKeyCallback: __obf_536fa713ca86126f,
	}, nil
}

// Creates a configuration for a client that authenticates using username and password
func PasswordKey(__obf_85be5df09f3bfe87 string, __obf_d0719fc09443c108 string, __obf_ccc767b877765d54 int) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User:    __obf_85be5df09f3bfe87,
		Timeout: time.Duration(__obf_ccc767b877765d54) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.Password(__obf_d0719fc09443c108),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

type WriteCounter struct {
	Total int64 // Total # of bytes written
}

func (__obf_58c6ff1752cfd867 *WriteCounter) Write(__obf_6cc5d8a07dedd61b []byte) (int, error) {
	__obf_5f6cd0910620a6eb := len(__obf_6cc5d8a07dedd61b)
	__obf_58c6ff1752cfd867.
		Total += int64(__obf_5f6cd0910620a6eb)
	fmt.Fprintf(os.Stdout, "%.2f kb transferred\n", float64(__obf_58c6ff1752cfd867.Total/1024))
	return __obf_5f6cd0910620a6eb, nil
}
