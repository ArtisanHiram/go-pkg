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
	__obf_a3505cbc5bd37a40, __obf_aa4fb85c75454390 := os.UserHomeDir()
	if __obf_aa4fb85c75454390 != nil {
		return "", __obf_aa4fb85c75454390
	}

	return fmt.Sprintf("%s/.ssh/known_hosts", __obf_a3505cbc5bd37a40), __obf_aa4fb85c75454390
}

// DefaultKnownHosts returns host key callback from default known hosts path, and error if any.
func DefaultKnownHosts() (ssh.HostKeyCallback, error) {
	__obf_6c280def7a3f3be0, __obf_aa4fb85c75454390 := DefaultKnownHostsPath()
	if __obf_aa4fb85c75454390 != nil {
		return nil, __obf_aa4fb85c75454390
	}

	return knownhosts.New(__obf_6c280def7a3f3be0)
}

// CheckKnownHost checks is host in known hosts file.
// it returns is the host found in known_hosts file and error, if the host found in
// known_hosts file and error not nil that means public key mismatch, maybe MAN IN THE MIDDLE ATTACK! you should not handshake.
func CheckKnownHost(__obf_4f7c64b9a5cdccce string, __obf_19e636e44a7a89d8 net.Addr, __obf_a7a746e8874da03c ssh.PublicKey, __obf_868faa5e2229a1c3 string) (__obf_423a42074cf8439f bool, __obf_aa4fb85c75454390 error) {

	var __obf_847e0998c4452e5f *knownhosts.KeyError

	// Fallback to default known_hosts file
	if __obf_868faa5e2229a1c3 == "" {
		__obf_6c280def7a3f3be0, __obf_aa4fb85c75454390 := DefaultKnownHostsPath()
		if __obf_aa4fb85c75454390 != nil {
			return false, __obf_aa4fb85c75454390
		}

		__obf_868faa5e2229a1c3 = __obf_6c280def7a3f3be0
	}

	// Get host key callback
	__obf_079d3af3eb6056ea, __obf_aa4fb85c75454390 := knownhosts.New(__obf_868faa5e2229a1c3)

	if __obf_aa4fb85c75454390 != nil {
		return false, __obf_aa4fb85c75454390
	}

	// check if host already exists.
	__obf_aa4fb85c75454390 = __obf_079d3af3eb6056ea(__obf_4f7c64b9a5cdccce, __obf_19e636e44a7a89d8, __obf_a7a746e8874da03c)

	// Known host already exists.
	if __obf_aa4fb85c75454390 == nil {
		return true, nil
	}

	// Make sure that the error returned from the callback is host not in file error.
	// If keyErr.Want is greater than 0 length, that means host is in file with different key.
	if errors.As(__obf_aa4fb85c75454390, &__obf_847e0998c4452e5f) && len(__obf_847e0998c4452e5f.Want) > 0 {
		return true, __obf_847e0998c4452e5f
	}

	// Key is not trusted because it is not in the file.
	return false, nil
}

// AddKnownHost add a a host to known hosts file.
func AddKnownHost(__obf_4f7c64b9a5cdccce string, __obf_19e636e44a7a89d8 net.Addr, __obf_a7a746e8874da03c ssh.PublicKey, __obf_868faa5e2229a1c3 string) (__obf_aa4fb85c75454390 error) {

	// Fallback to default known_hosts file
	if __obf_868faa5e2229a1c3 == "" {
		__obf_6c280def7a3f3be0, __obf_aa4fb85c75454390 := DefaultKnownHostsPath()
		if __obf_aa4fb85c75454390 != nil {
			return __obf_aa4fb85c75454390
		}

		__obf_868faa5e2229a1c3 = __obf_6c280def7a3f3be0
	}

	__obf_c86885986c344b16, __obf_aa4fb85c75454390 := os.OpenFile(__obf_868faa5e2229a1c3, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if __obf_aa4fb85c75454390 != nil {
		return __obf_aa4fb85c75454390
	}

	defer __obf_c86885986c344b16.Close()

	__obf_1dc0406cd38cedea := knownhosts.Normalize(__obf_19e636e44a7a89d8.String())
	__obf_b05074fae9c5808d := knownhosts.Normalize(__obf_4f7c64b9a5cdccce)
	__obf_69af561f1dfd88ca := []string{__obf_1dc0406cd38cedea}

	if __obf_b05074fae9c5808d != __obf_1dc0406cd38cedea {
		__obf_69af561f1dfd88ca = append(__obf_69af561f1dfd88ca, __obf_b05074fae9c5808d)
	}

	_, __obf_aa4fb85c75454390 = __obf_c86885986c344b16.WriteString(knownhosts.Line(__obf_69af561f1dfd88ca, __obf_a7a746e8874da03c) + "\n")

	return __obf_aa4fb85c75454390
}

// PrivateKey Loads a private and public key from "path" and returns a SSH ClientConfig to authenticate with the server
func PrivateKey(__obf_fec45acffc50fe68 string, __obf_6c280def7a3f3be0 string, __obf_a8385ab5dd82f50b int) (*ssh.ClientConfig, error) {
	__obf_39c8f25078b1f176, __obf_aa4fb85c75454390 := os.ReadFile(__obf_6c280def7a3f3be0)

	if __obf_aa4fb85c75454390 != nil {
		return nil, __obf_aa4fb85c75454390
	}

	__obf_aa441b330d7f6792, __obf_aa4fb85c75454390 := ssh.ParsePrivateKey(__obf_39c8f25078b1f176)

	if __obf_aa4fb85c75454390 != nil {
		return nil, __obf_aa4fb85c75454390
	}

	__obf_079d3af3eb6056ea, __obf_aa4fb85c75454390 := DefaultKnownHosts()
	if __obf_aa4fb85c75454390 != nil {
		return nil, __obf_aa4fb85c75454390
	}

	return &ssh.ClientConfig{
		User:    __obf_fec45acffc50fe68,
		Timeout: time.Duration(__obf_a8385ab5dd82f50b) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_aa441b330d7f6792),
		},
		HostKeyCallback: __obf_079d3af3eb6056ea,
	}, nil
}

// Creates the configuration for a client that authenticates with a password protected private key
func PrivateKeyWithPassphrase(__obf_fec45acffc50fe68 string, __obf_a8385ab5dd82f50b int, __obf_3a0ae00d29cf2acf []byte, __obf_6c280def7a3f3be0 string) (*ssh.ClientConfig, error) {
	__obf_39c8f25078b1f176, __obf_aa4fb85c75454390 := os.ReadFile(__obf_6c280def7a3f3be0)

	if __obf_aa4fb85c75454390 != nil {
		return nil, __obf_aa4fb85c75454390
	}
	__obf_aa441b330d7f6792, __obf_aa4fb85c75454390 := ssh.ParsePrivateKeyWithPassphrase(__obf_39c8f25078b1f176, __obf_3a0ae00d29cf2acf)

	if __obf_aa4fb85c75454390 != nil {
		return nil, __obf_aa4fb85c75454390
	}

	__obf_079d3af3eb6056ea, __obf_aa4fb85c75454390 := DefaultKnownHosts()
	if __obf_aa4fb85c75454390 != nil {
		return nil, __obf_aa4fb85c75454390
	}

	return &ssh.ClientConfig{
		User:    __obf_fec45acffc50fe68,
		Timeout: time.Duration(__obf_a8385ab5dd82f50b) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_aa441b330d7f6792),
		},
		HostKeyCallback: __obf_079d3af3eb6056ea,
	}, nil
}

// Creates a configuration for a client that fetches public-private key from the SSH agent for authentication
func SSHAgent(__obf_fec45acffc50fe68 string, __obf_a8385ab5dd82f50b int, __obf_4e2351eb779d8199 ssh.HostKeyCallback) (*ssh.ClientConfig, error) {
	__obf_b2b3bc3f2d3b2fc7 := os.Getenv("SSH_AUTH_SOCK")
	__obf_10d2966242657c57, __obf_aa4fb85c75454390 := net.Dial("unix", __obf_b2b3bc3f2d3b2fc7)
	if __obf_aa4fb85c75454390 != nil {
		return &ssh.ClientConfig{
			Timeout: time.Duration(__obf_a8385ab5dd82f50b) * time.Second,
		}, __obf_aa4fb85c75454390
	}

	return &ssh.ClientConfig{
		User:    __obf_fec45acffc50fe68,
		Timeout: time.Duration(__obf_a8385ab5dd82f50b) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agent.NewClient(__obf_10d2966242657c57).Signers),
		},
		HostKeyCallback: __obf_4e2351eb779d8199,
	}, nil
}

// Creates a configuration for a client that authenticates using username and password
func PasswordKey(__obf_fec45acffc50fe68 string, __obf_a5bbbad23f971db0 string, __obf_a8385ab5dd82f50b int) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User:    __obf_fec45acffc50fe68,
		Timeout: time.Duration(__obf_a8385ab5dd82f50b) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.Password(__obf_a5bbbad23f971db0),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

type WriteCounter struct {
	Total int64 // Total # of bytes written
}

func (__obf_42776f9ea28ba1bd *WriteCounter) Write(__obf_34276e4ccbcf277c []byte) (int, error) {
	__obf_a81ecf54ba87e38d := len(__obf_34276e4ccbcf277c)
	__obf_42776f9ea28ba1bd.Total += int64(__obf_a81ecf54ba87e38d)
	fmt.Fprintf(os.Stdout, "%.2f kb transferred\n", float64(__obf_42776f9ea28ba1bd.Total/1024))
	return __obf_a81ecf54ba87e38d, nil
}
