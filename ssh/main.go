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
	__obf_94648c337f52b972, __obf_dbce9f361de2d092 := os.UserHomeDir()
	if __obf_dbce9f361de2d092 != nil {
		return "", __obf_dbce9f361de2d092
	}

	return fmt.Sprintf("%s/.ssh/known_hosts", __obf_94648c337f52b972), __obf_dbce9f361de2d092
}

// DefaultKnownHosts returns host key callback from default known hosts path, and error if any.
func DefaultKnownHosts() (ssh.HostKeyCallback, error) {
	__obf_1f6d09fe2d1b2ea9, __obf_dbce9f361de2d092 := DefaultKnownHostsPath()
	if __obf_dbce9f361de2d092 != nil {
		return nil, __obf_dbce9f361de2d092
	}

	return knownhosts.New(__obf_1f6d09fe2d1b2ea9)
}

// CheckKnownHost checks is host in known hosts file.
// it returns is the host found in known_hosts file and error, if the host found in
// known_hosts file and error not nil that means public key mismatch, maybe MAN IN THE MIDDLE ATTACK! you should not handshake.
func CheckKnownHost(__obf_e89052812679b819 string, __obf_22f1a0bdee77e6e9 net.Addr, __obf_89c1ca092bd43ab1 ssh.PublicKey, __obf_cba95ac1131a7410 string) (__obf_73da32db24a2798a bool, __obf_dbce9f361de2d092 error) {

	var __obf_c3783feb927ca606 *knownhosts.KeyError

	// Fallback to default known_hosts file
	if __obf_cba95ac1131a7410 == "" {
		__obf_1f6d09fe2d1b2ea9, __obf_dbce9f361de2d092 := DefaultKnownHostsPath()
		if __obf_dbce9f361de2d092 != nil {
			return false, __obf_dbce9f361de2d092
		}

		__obf_cba95ac1131a7410 = __obf_1f6d09fe2d1b2ea9
	}

	// Get host key callback
	__obf_3132c55bc037d689, __obf_dbce9f361de2d092 := knownhosts.New(__obf_cba95ac1131a7410)

	if __obf_dbce9f361de2d092 != nil {
		return false, __obf_dbce9f361de2d092
	}

	// check if host already exists.
	__obf_dbce9f361de2d092 = __obf_3132c55bc037d689(__obf_e89052812679b819, __obf_22f1a0bdee77e6e9, __obf_89c1ca092bd43ab1)

	// Known host already exists.
	if __obf_dbce9f361de2d092 == nil {
		return true, nil
	}

	// Make sure that the error returned from the callback is host not in file error.
	// If keyErr.Want is greater than 0 length, that means host is in file with different key.
	if errors.As(__obf_dbce9f361de2d092, &__obf_c3783feb927ca606) && len(__obf_c3783feb927ca606.Want) > 0 {
		return true, __obf_c3783feb927ca606
	}

	// Key is not trusted because it is not in the file.
	return false, nil
}

// AddKnownHost add a a host to known hosts file.
func AddKnownHost(__obf_e89052812679b819 string, __obf_22f1a0bdee77e6e9 net.Addr, __obf_89c1ca092bd43ab1 ssh.PublicKey, __obf_cba95ac1131a7410 string) (__obf_dbce9f361de2d092 error) {

	// Fallback to default known_hosts file
	if __obf_cba95ac1131a7410 == "" {
		__obf_1f6d09fe2d1b2ea9, __obf_dbce9f361de2d092 := DefaultKnownHostsPath()
		if __obf_dbce9f361de2d092 != nil {
			return __obf_dbce9f361de2d092
		}

		__obf_cba95ac1131a7410 = __obf_1f6d09fe2d1b2ea9
	}

	__obf_7dedd21070ae08f6, __obf_dbce9f361de2d092 := os.OpenFile(__obf_cba95ac1131a7410, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if __obf_dbce9f361de2d092 != nil {
		return __obf_dbce9f361de2d092
	}

	defer __obf_7dedd21070ae08f6.Close()

	__obf_a6f6590910887c40 := knownhosts.Normalize(__obf_22f1a0bdee77e6e9.String())
	__obf_eb2527af98e416c8 := knownhosts.Normalize(__obf_e89052812679b819)
	__obf_04133251a3e2cec3 := []string{__obf_a6f6590910887c40}

	if __obf_eb2527af98e416c8 != __obf_a6f6590910887c40 {
		__obf_04133251a3e2cec3 = append(__obf_04133251a3e2cec3, __obf_eb2527af98e416c8)
	}

	_, __obf_dbce9f361de2d092 = __obf_7dedd21070ae08f6.WriteString(knownhosts.Line(__obf_04133251a3e2cec3, __obf_89c1ca092bd43ab1) + "\n")

	return __obf_dbce9f361de2d092
}

// PrivateKey Loads a private and public key from "path" and returns a SSH ClientConfig to authenticate with the server
func PrivateKey(__obf_8fc7e79d550ae422 string, __obf_1f6d09fe2d1b2ea9 string, __obf_efad4d41979f60e8 int) (*ssh.ClientConfig, error) {
	__obf_0f6933b72d3e89b5, __obf_dbce9f361de2d092 := os.ReadFile(__obf_1f6d09fe2d1b2ea9)

	if __obf_dbce9f361de2d092 != nil {
		return nil, __obf_dbce9f361de2d092
	}

	__obf_aef07006393d1a6c, __obf_dbce9f361de2d092 := ssh.ParsePrivateKey(__obf_0f6933b72d3e89b5)

	if __obf_dbce9f361de2d092 != nil {
		return nil, __obf_dbce9f361de2d092
	}

	__obf_3132c55bc037d689, __obf_dbce9f361de2d092 := DefaultKnownHosts()
	if __obf_dbce9f361de2d092 != nil {
		return nil, __obf_dbce9f361de2d092
	}

	return &ssh.ClientConfig{
		User:    __obf_8fc7e79d550ae422,
		Timeout: time.Duration(__obf_efad4d41979f60e8) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_aef07006393d1a6c),
		},
		HostKeyCallback: __obf_3132c55bc037d689,
	}, nil
}

// Creates the configuration for a client that authenticates with a password protected private key
func PrivateKeyWithPassphrase(__obf_8fc7e79d550ae422 string, __obf_efad4d41979f60e8 int, __obf_aff1eab6970b23ff []byte, __obf_1f6d09fe2d1b2ea9 string) (*ssh.ClientConfig, error) {
	__obf_0f6933b72d3e89b5, __obf_dbce9f361de2d092 := os.ReadFile(__obf_1f6d09fe2d1b2ea9)

	if __obf_dbce9f361de2d092 != nil {
		return nil, __obf_dbce9f361de2d092
	}
	__obf_aef07006393d1a6c, __obf_dbce9f361de2d092 := ssh.ParsePrivateKeyWithPassphrase(__obf_0f6933b72d3e89b5, __obf_aff1eab6970b23ff)

	if __obf_dbce9f361de2d092 != nil {
		return nil, __obf_dbce9f361de2d092
	}

	__obf_3132c55bc037d689, __obf_dbce9f361de2d092 := DefaultKnownHosts()
	if __obf_dbce9f361de2d092 != nil {
		return nil, __obf_dbce9f361de2d092
	}

	return &ssh.ClientConfig{
		User:    __obf_8fc7e79d550ae422,
		Timeout: time.Duration(__obf_efad4d41979f60e8) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_aef07006393d1a6c),
		},
		HostKeyCallback: __obf_3132c55bc037d689,
	}, nil
}

// Creates a configuration for a client that fetches public-private key from the SSH agent for authentication
func SSHAgent(__obf_8fc7e79d550ae422 string, __obf_efad4d41979f60e8 int, __obf_0a978b4b52f689fb ssh.HostKeyCallback) (*ssh.ClientConfig, error) {
	__obf_ccf7af57338a64d0 := os.Getenv("SSH_AUTH_SOCK")
	__obf_c5e4060e1bb02ae1, __obf_dbce9f361de2d092 := net.Dial("unix", __obf_ccf7af57338a64d0)
	if __obf_dbce9f361de2d092 != nil {
		return &ssh.ClientConfig{
			Timeout: time.Duration(__obf_efad4d41979f60e8) * time.Second,
		}, __obf_dbce9f361de2d092
	}

	return &ssh.ClientConfig{
		User:    __obf_8fc7e79d550ae422,
		Timeout: time.Duration(__obf_efad4d41979f60e8) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agent.NewClient(__obf_c5e4060e1bb02ae1).Signers),
		},
		HostKeyCallback: __obf_0a978b4b52f689fb,
	}, nil
}

// Creates a configuration for a client that authenticates using username and password
func PasswordKey(__obf_8fc7e79d550ae422 string, __obf_75f0ccd92e720efc string, __obf_efad4d41979f60e8 int) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User:    __obf_8fc7e79d550ae422,
		Timeout: time.Duration(__obf_efad4d41979f60e8) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.Password(__obf_75f0ccd92e720efc),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

type WriteCounter struct {
	Total int64 // Total # of bytes written
}

func (__obf_a1a5472c3b1b2ce9 *WriteCounter) Write(__obf_1fe752a7fd0e9bc7 []byte) (int, error) {
	__obf_e0d9627cbe800863 := len(__obf_1fe752a7fd0e9bc7)
	__obf_a1a5472c3b1b2ce9.Total += int64(__obf_e0d9627cbe800863)
	fmt.Fprintf(os.Stdout, "%.2f kb transferred\n", float64(__obf_a1a5472c3b1b2ce9.Total/1024))
	return __obf_e0d9627cbe800863, nil
}
