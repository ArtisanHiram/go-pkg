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
	__obf_aa10e35947badbe0, __obf_0b884480f1bc4644 := os.UserHomeDir()
	if __obf_0b884480f1bc4644 != nil {
		return "", __obf_0b884480f1bc4644
	}

	return fmt.Sprintf("%s/.ssh/known_hosts", __obf_aa10e35947badbe0), __obf_0b884480f1bc4644
}

// DefaultKnownHosts returns host key callback from default known hosts path, and error if any.
func DefaultKnownHosts() (ssh.HostKeyCallback, error) {
	__obf_5c9e3b1b46cbe5fe, __obf_0b884480f1bc4644 := DefaultKnownHostsPath()
	if __obf_0b884480f1bc4644 != nil {
		return nil, __obf_0b884480f1bc4644
	}

	return knownhosts.New(__obf_5c9e3b1b46cbe5fe)
}

// CheckKnownHost checks is host in known hosts file.
// it returns is the host found in known_hosts file and error, if the host found in
// known_hosts file and error not nil that means public key mismatch, maybe MAN IN THE MIDDLE ATTACK! you should not handshake.
func CheckKnownHost(__obf_ccab29cbe863a6fc string, __obf_ae71ab6fb9c133f2 net.Addr, __obf_c3d38cfe3a3cb36a ssh.PublicKey, __obf_2ca9366a99d088eb string) (__obf_2589f8a3842bc288 bool, __obf_0b884480f1bc4644 error) {

	var __obf_c8942eed04d20fcc *knownhosts.KeyError

	// Fallback to default known_hosts file
	if __obf_2ca9366a99d088eb == "" {
		__obf_5c9e3b1b46cbe5fe, __obf_0b884480f1bc4644 := DefaultKnownHostsPath()
		if __obf_0b884480f1bc4644 != nil {
			return false, __obf_0b884480f1bc4644
		}

		__obf_2ca9366a99d088eb = __obf_5c9e3b1b46cbe5fe
	}

	// Get host key callback
	__obf_9a512a2d81b6ef5b, __obf_0b884480f1bc4644 := knownhosts.New(__obf_2ca9366a99d088eb)

	if __obf_0b884480f1bc4644 != nil {
		return false, __obf_0b884480f1bc4644
	}

	// check if host already exists.
	__obf_0b884480f1bc4644 = __obf_9a512a2d81b6ef5b(__obf_ccab29cbe863a6fc, __obf_ae71ab6fb9c133f2, __obf_c3d38cfe3a3cb36a)

	// Known host already exists.
	if __obf_0b884480f1bc4644 == nil {
		return true, nil
	}

	// Make sure that the error returned from the callback is host not in file error.
	// If keyErr.Want is greater than 0 length, that means host is in file with different key.
	if errors.As(__obf_0b884480f1bc4644, &__obf_c8942eed04d20fcc) && len(__obf_c8942eed04d20fcc.Want) > 0 {
		return true, __obf_c8942eed04d20fcc
	}

	// Key is not trusted because it is not in the file.
	return false, nil
}

// AddKnownHost add a a host to known hosts file.
func AddKnownHost(__obf_ccab29cbe863a6fc string, __obf_ae71ab6fb9c133f2 net.Addr, __obf_c3d38cfe3a3cb36a ssh.PublicKey, __obf_2ca9366a99d088eb string) (__obf_0b884480f1bc4644 error) {

	// Fallback to default known_hosts file
	if __obf_2ca9366a99d088eb == "" {
		__obf_5c9e3b1b46cbe5fe, __obf_0b884480f1bc4644 := DefaultKnownHostsPath()
		if __obf_0b884480f1bc4644 != nil {
			return __obf_0b884480f1bc4644
		}

		__obf_2ca9366a99d088eb = __obf_5c9e3b1b46cbe5fe
	}

	__obf_40ec5923403ff109, __obf_0b884480f1bc4644 := os.OpenFile(__obf_2ca9366a99d088eb, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if __obf_0b884480f1bc4644 != nil {
		return __obf_0b884480f1bc4644
	}

	defer __obf_40ec5923403ff109.Close()

	__obf_188c3d17909eaa09 := knownhosts.Normalize(__obf_ae71ab6fb9c133f2.String())
	__obf_3b598f29eed28ea6 := knownhosts.Normalize(__obf_ccab29cbe863a6fc)
	__obf_aab92c295f7f3587 := []string{__obf_188c3d17909eaa09}

	if __obf_3b598f29eed28ea6 != __obf_188c3d17909eaa09 {
		__obf_aab92c295f7f3587 = append(__obf_aab92c295f7f3587, __obf_3b598f29eed28ea6)
	}

	_, __obf_0b884480f1bc4644 = __obf_40ec5923403ff109.WriteString(knownhosts.Line(__obf_aab92c295f7f3587, __obf_c3d38cfe3a3cb36a) + "\n")

	return __obf_0b884480f1bc4644
}

// PrivateKey Loads a private and public key from "path" and returns a SSH ClientConfig to authenticate with the server
func PrivateKey(__obf_d0cde6c8bb7462fe string, __obf_5c9e3b1b46cbe5fe string, __obf_9b9d3ee10e2917ce int) (*ssh.ClientConfig, error) {
	__obf_5a083e774b7d5934, __obf_0b884480f1bc4644 := os.ReadFile(__obf_5c9e3b1b46cbe5fe)

	if __obf_0b884480f1bc4644 != nil {
		return nil, __obf_0b884480f1bc4644
	}

	__obf_7f9b180a37e982c2, __obf_0b884480f1bc4644 := ssh.ParsePrivateKey(__obf_5a083e774b7d5934)

	if __obf_0b884480f1bc4644 != nil {
		return nil, __obf_0b884480f1bc4644
	}

	__obf_9a512a2d81b6ef5b, __obf_0b884480f1bc4644 := DefaultKnownHosts()
	if __obf_0b884480f1bc4644 != nil {
		return nil, __obf_0b884480f1bc4644
	}

	return &ssh.ClientConfig{
		User:    __obf_d0cde6c8bb7462fe,
		Timeout: time.Duration(__obf_9b9d3ee10e2917ce) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_7f9b180a37e982c2),
		},
		HostKeyCallback: __obf_9a512a2d81b6ef5b,
	}, nil
}

// Creates the configuration for a client that authenticates with a password protected private key
func PrivateKeyWithPassphrase(__obf_d0cde6c8bb7462fe string, __obf_9b9d3ee10e2917ce int, __obf_22659e5758252feb []byte, __obf_5c9e3b1b46cbe5fe string) (*ssh.ClientConfig, error) {
	__obf_5a083e774b7d5934, __obf_0b884480f1bc4644 := os.ReadFile(__obf_5c9e3b1b46cbe5fe)

	if __obf_0b884480f1bc4644 != nil {
		return nil, __obf_0b884480f1bc4644
	}
	__obf_7f9b180a37e982c2, __obf_0b884480f1bc4644 := ssh.ParsePrivateKeyWithPassphrase(__obf_5a083e774b7d5934, __obf_22659e5758252feb)

	if __obf_0b884480f1bc4644 != nil {
		return nil, __obf_0b884480f1bc4644
	}

	__obf_9a512a2d81b6ef5b, __obf_0b884480f1bc4644 := DefaultKnownHosts()
	if __obf_0b884480f1bc4644 != nil {
		return nil, __obf_0b884480f1bc4644
	}

	return &ssh.ClientConfig{
		User:    __obf_d0cde6c8bb7462fe,
		Timeout: time.Duration(__obf_9b9d3ee10e2917ce) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_7f9b180a37e982c2),
		},
		HostKeyCallback: __obf_9a512a2d81b6ef5b,
	}, nil
}

// Creates a configuration for a client that fetches public-private key from the SSH agent for authentication
func SSHAgent(__obf_d0cde6c8bb7462fe string, __obf_9b9d3ee10e2917ce int, __obf_69b0fe53d22f0149 ssh.HostKeyCallback) (*ssh.ClientConfig, error) {
	__obf_01543001d5f5cbae := os.Getenv("SSH_AUTH_SOCK")
	__obf_2caeb842c964723c, __obf_0b884480f1bc4644 := net.Dial("unix", __obf_01543001d5f5cbae)
	if __obf_0b884480f1bc4644 != nil {
		return &ssh.ClientConfig{
			Timeout: time.Duration(__obf_9b9d3ee10e2917ce) * time.Second,
		}, __obf_0b884480f1bc4644
	}

	return &ssh.ClientConfig{
		User:    __obf_d0cde6c8bb7462fe,
		Timeout: time.Duration(__obf_9b9d3ee10e2917ce) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agent.NewClient(__obf_2caeb842c964723c).Signers),
		},
		HostKeyCallback: __obf_69b0fe53d22f0149,
	}, nil
}

// Creates a configuration for a client that authenticates using username and password
func PasswordKey(__obf_d0cde6c8bb7462fe string, __obf_650d19f13f6ec171 string, __obf_9b9d3ee10e2917ce int) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User:    __obf_d0cde6c8bb7462fe,
		Timeout: time.Duration(__obf_9b9d3ee10e2917ce) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.Password(__obf_650d19f13f6ec171),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

type WriteCounter struct {
	Total int64 // Total # of bytes written
}

func (__obf_00f3379af3fda252 *WriteCounter) Write(__obf_3f43d89298fa9006 []byte) (int, error) {
	__obf_0dbe628930d0c4c6 := len(__obf_3f43d89298fa9006)
	__obf_00f3379af3fda252.Total += int64(__obf_0dbe628930d0c4c6)
	fmt.Fprintf(os.Stdout, "%.2f kb transferred\n", float64(__obf_00f3379af3fda252.Total/1024))
	return __obf_0dbe628930d0c4c6, nil
}
