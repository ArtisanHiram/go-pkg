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
	__obf_4e670c9667d970ea, __obf_2c7c784faaf5a58c := os.UserHomeDir()
	if __obf_2c7c784faaf5a58c != nil {
		return "", __obf_2c7c784faaf5a58c
	}

	return fmt.Sprintf("%s/.ssh/known_hosts", __obf_4e670c9667d970ea), __obf_2c7c784faaf5a58c
}

// DefaultKnownHosts returns host key callback from default known hosts path, and error if any.
func DefaultKnownHosts() (ssh.HostKeyCallback, error) {
	__obf_708ab0926c8f4817, __obf_2c7c784faaf5a58c := DefaultKnownHostsPath()
	if __obf_2c7c784faaf5a58c != nil {
		return nil, __obf_2c7c784faaf5a58c
	}

	return knownhosts.New(__obf_708ab0926c8f4817)
}

// CheckKnownHost checks is host in known hosts file.
// it returns is the host found in known_hosts file and error, if the host found in
// known_hosts file and error not nil that means public key mismatch, maybe MAN IN THE MIDDLE ATTACK! you should not handshake.
func CheckKnownHost(__obf_b354d9520914a042 string, __obf_01edc81f547f2229 net.Addr, __obf_f2c7225dc6291808 ssh.PublicKey, __obf_1ae2c1652a2286bf string) (__obf_bd31ff505933ad47 bool, __obf_2c7c784faaf5a58c error) {

	var __obf_18e496791677a597 *knownhosts.KeyError

	// Fallback to default known_hosts file
	if __obf_1ae2c1652a2286bf == "" {
		__obf_708ab0926c8f4817, __obf_2c7c784faaf5a58c := DefaultKnownHostsPath()
		if __obf_2c7c784faaf5a58c != nil {
			return false, __obf_2c7c784faaf5a58c
		}

		__obf_1ae2c1652a2286bf = __obf_708ab0926c8f4817
	}

	// Get host key callback
	__obf_006ba3c1e4b4ea6c, __obf_2c7c784faaf5a58c := knownhosts.New(__obf_1ae2c1652a2286bf)

	if __obf_2c7c784faaf5a58c != nil {
		return false, __obf_2c7c784faaf5a58c
	}

	// check if host already exists.
	__obf_2c7c784faaf5a58c = __obf_006ba3c1e4b4ea6c(__obf_b354d9520914a042, __obf_01edc81f547f2229, __obf_f2c7225dc6291808)

	// Known host already exists.
	if __obf_2c7c784faaf5a58c == nil {
		return true, nil
	}

	// Make sure that the error returned from the callback is host not in file error.
	// If keyErr.Want is greater than 0 length, that means host is in file with different key.
	if errors.As(__obf_2c7c784faaf5a58c, &__obf_18e496791677a597) && len(__obf_18e496791677a597.Want) > 0 {
		return true, __obf_18e496791677a597
	}

	// Key is not trusted because it is not in the file.
	return false, nil
}

// AddKnownHost add a a host to known hosts file.
func AddKnownHost(__obf_b354d9520914a042 string, __obf_01edc81f547f2229 net.Addr, __obf_f2c7225dc6291808 ssh.PublicKey, __obf_1ae2c1652a2286bf string) (__obf_2c7c784faaf5a58c error) {

	// Fallback to default known_hosts file
	if __obf_1ae2c1652a2286bf == "" {
		__obf_708ab0926c8f4817, __obf_2c7c784faaf5a58c := DefaultKnownHostsPath()
		if __obf_2c7c784faaf5a58c != nil {
			return __obf_2c7c784faaf5a58c
		}

		__obf_1ae2c1652a2286bf = __obf_708ab0926c8f4817
	}

	__obf_497a9ad704dfc3e5, __obf_2c7c784faaf5a58c := os.OpenFile(__obf_1ae2c1652a2286bf, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if __obf_2c7c784faaf5a58c != nil {
		return __obf_2c7c784faaf5a58c
	}

	defer __obf_497a9ad704dfc3e5.Close()

	__obf_0701461635a679ef := knownhosts.Normalize(__obf_01edc81f547f2229.String())
	__obf_a7cc103da5849d2d := knownhosts.Normalize(__obf_b354d9520914a042)
	__obf_f133f5ae749760f3 := []string{__obf_0701461635a679ef}

	if __obf_a7cc103da5849d2d != __obf_0701461635a679ef {
		__obf_f133f5ae749760f3 = append(__obf_f133f5ae749760f3, __obf_a7cc103da5849d2d)
	}

	_, __obf_2c7c784faaf5a58c = __obf_497a9ad704dfc3e5.WriteString(knownhosts.Line(__obf_f133f5ae749760f3, __obf_f2c7225dc6291808) + "\n")

	return __obf_2c7c784faaf5a58c
}

// PrivateKey Loads a private and public key from "path" and returns a SSH ClientConfig to authenticate with the server
func PrivateKey(__obf_e7527eca3b0fb24c string, __obf_708ab0926c8f4817 string, __obf_5c72e19023a8c203 int) (*ssh.ClientConfig, error) {
	__obf_e31de9409acd110d, __obf_2c7c784faaf5a58c := os.ReadFile(__obf_708ab0926c8f4817)

	if __obf_2c7c784faaf5a58c != nil {
		return nil, __obf_2c7c784faaf5a58c
	}

	__obf_c94257b7d7217370, __obf_2c7c784faaf5a58c := ssh.ParsePrivateKey(__obf_e31de9409acd110d)

	if __obf_2c7c784faaf5a58c != nil {
		return nil, __obf_2c7c784faaf5a58c
	}

	__obf_006ba3c1e4b4ea6c, __obf_2c7c784faaf5a58c := DefaultKnownHosts()
	if __obf_2c7c784faaf5a58c != nil {
		return nil, __obf_2c7c784faaf5a58c
	}

	return &ssh.ClientConfig{
		User:    __obf_e7527eca3b0fb24c,
		Timeout: time.Duration(__obf_5c72e19023a8c203) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_c94257b7d7217370),
		},
		HostKeyCallback: __obf_006ba3c1e4b4ea6c,
	}, nil
}

// Creates the configuration for a client that authenticates with a password protected private key
func PrivateKeyWithPassphrase(__obf_e7527eca3b0fb24c string, __obf_5c72e19023a8c203 int, __obf_76f2b6e06a56786b []byte, __obf_708ab0926c8f4817 string) (*ssh.ClientConfig, error) {
	__obf_e31de9409acd110d, __obf_2c7c784faaf5a58c := os.ReadFile(__obf_708ab0926c8f4817)

	if __obf_2c7c784faaf5a58c != nil {
		return nil, __obf_2c7c784faaf5a58c
	}
	__obf_c94257b7d7217370, __obf_2c7c784faaf5a58c := ssh.ParsePrivateKeyWithPassphrase(__obf_e31de9409acd110d, __obf_76f2b6e06a56786b)

	if __obf_2c7c784faaf5a58c != nil {
		return nil, __obf_2c7c784faaf5a58c
	}

	__obf_006ba3c1e4b4ea6c, __obf_2c7c784faaf5a58c := DefaultKnownHosts()
	if __obf_2c7c784faaf5a58c != nil {
		return nil, __obf_2c7c784faaf5a58c
	}

	return &ssh.ClientConfig{
		User:    __obf_e7527eca3b0fb24c,
		Timeout: time.Duration(__obf_5c72e19023a8c203) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_c94257b7d7217370),
		},
		HostKeyCallback: __obf_006ba3c1e4b4ea6c,
	}, nil
}

// Creates a configuration for a client that fetches public-private key from the SSH agent for authentication
func SSHAgent(__obf_e7527eca3b0fb24c string, __obf_5c72e19023a8c203 int, __obf_6afadeca7de5158c ssh.HostKeyCallback) (*ssh.ClientConfig, error) {
	__obf_ca12de958da72ab0 := os.Getenv("SSH_AUTH_SOCK")
	__obf_198c6580163b2289, __obf_2c7c784faaf5a58c := net.Dial("unix", __obf_ca12de958da72ab0)
	if __obf_2c7c784faaf5a58c != nil {
		return &ssh.ClientConfig{
			Timeout: time.Duration(__obf_5c72e19023a8c203) * time.Second,
		}, __obf_2c7c784faaf5a58c
	}

	return &ssh.ClientConfig{
		User:    __obf_e7527eca3b0fb24c,
		Timeout: time.Duration(__obf_5c72e19023a8c203) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agent.NewClient(__obf_198c6580163b2289).Signers),
		},
		HostKeyCallback: __obf_6afadeca7de5158c,
	}, nil
}

// Creates a configuration for a client that authenticates using username and password
func PasswordKey(__obf_e7527eca3b0fb24c string, __obf_7da17bb396d477f6 string, __obf_5c72e19023a8c203 int) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User:    __obf_e7527eca3b0fb24c,
		Timeout: time.Duration(__obf_5c72e19023a8c203) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.Password(__obf_7da17bb396d477f6),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

type WriteCounter struct {
	Total int64 // Total # of bytes written
}

func (__obf_22893f278bf85786 *WriteCounter) Write(__obf_a26e0200a3e6170f []byte) (int, error) {
	__obf_8d5950047e657148 := len(__obf_a26e0200a3e6170f)
	__obf_22893f278bf85786.Total += int64(__obf_8d5950047e657148)
	fmt.Fprintf(os.Stdout, "%.2f kb transferred\n", float64(__obf_22893f278bf85786.Total/1024))
	return __obf_8d5950047e657148, nil
}
