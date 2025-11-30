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
	__obf_69430ba9c26d8252, __obf_7f6f4ab0f2c744c3 := os.UserHomeDir()
	if __obf_7f6f4ab0f2c744c3 != nil {
		return "", __obf_7f6f4ab0f2c744c3
	}

	return fmt.Sprintf("%s/.ssh/known_hosts", __obf_69430ba9c26d8252), __obf_7f6f4ab0f2c744c3
}

// DefaultKnownHosts returns host key callback from default known hosts path, and error if some.
func DefaultKnownHosts() (ssh.HostKeyCallback, error) {
	__obf_00b878bbf8da5796, __obf_7f6f4ab0f2c744c3 := DefaultKnownHostsPath()
	if __obf_7f6f4ab0f2c744c3 != nil {
		return nil, __obf_7f6f4ab0f2c744c3
	}

	return knownhosts.New(__obf_00b878bbf8da5796)
}

// CheckKnownHost checks is host in known hosts file.
// it returns is the host found in known_hosts file and error, if the host found in
// known_hosts file and error not nil that means public key mismatch, maybe MAN IN THE MIDDLE ATTACK! you should not handshake.
func CheckKnownHost(__obf_d55af7c5d3d84f3d string, __obf_83d70bf93c31aff3 net.Addr, __obf_b104c952b502a6cb ssh.PublicKey, __obf_e63c83e71371d8d1 string) (__obf_b4a7526b6775833f bool, __obf_7f6f4ab0f2c744c3 error) {

	var __obf_938663f99bc6650c *knownhosts.KeyError

	// Fallback to default known_hosts file
	if __obf_e63c83e71371d8d1 == "" {
		__obf_00b878bbf8da5796, __obf_7f6f4ab0f2c744c3 := DefaultKnownHostsPath()
		if __obf_7f6f4ab0f2c744c3 != nil {
			return false, __obf_7f6f4ab0f2c744c3
		}
		__obf_e63c83e71371d8d1 = __obf_00b878bbf8da5796
	}
	__obf_ba6bfbe62b222d79,

		// Get host key callback
		__obf_7f6f4ab0f2c744c3 := knownhosts.New(__obf_e63c83e71371d8d1)

	if __obf_7f6f4ab0f2c744c3 != nil {
		return false, __obf_7f6f4ab0f2c744c3
	}
	__obf_7f6f4ab0f2c744c3 = // check if host already exists.
		__obf_ba6bfbe62b222d79(__obf_d55af7c5d3d84f3d, __obf_83d70bf93c31aff3,

			// Known host already exists.
			__obf_b104c952b502a6cb)

	if __obf_7f6f4ab0f2c744c3 == nil {
		return true, nil
	}

	// Make sure that the error returned from the callback is host not in file error.
	// If keyErr.Want is greater than 0 length, that means host is in file with different key.
	if errors.As(__obf_7f6f4ab0f2c744c3, &__obf_938663f99bc6650c) && len(__obf_938663f99bc6650c.Want) > 0 {
		return true, __obf_938663f99bc6650c
	}

	// Key is not trusted because it is not in the file.
	return false, nil
}

// AddKnownHost add a a host to known hosts file.
func AddKnownHost(__obf_d55af7c5d3d84f3d string, __obf_83d70bf93c31aff3 net.Addr, __obf_b104c952b502a6cb ssh.PublicKey, __obf_e63c83e71371d8d1 string) (__obf_7f6f4ab0f2c744c3 error) {

	// Fallback to default known_hosts file
	if __obf_e63c83e71371d8d1 == "" {
		__obf_00b878bbf8da5796, __obf_7f6f4ab0f2c744c3 := DefaultKnownHostsPath()
		if __obf_7f6f4ab0f2c744c3 != nil {
			return __obf_7f6f4ab0f2c744c3
		}
		__obf_e63c83e71371d8d1 = __obf_00b878bbf8da5796
	}
	__obf_31fdd89e34f6d6ed, __obf_7f6f4ab0f2c744c3 := os.OpenFile(__obf_e63c83e71371d8d1, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if __obf_7f6f4ab0f2c744c3 != nil {
		return __obf_7f6f4ab0f2c744c3
	}

	defer __obf_31fdd89e34f6d6ed.Close()
	__obf_ca5ef72eec45bf9e := knownhosts.Normalize(__obf_83d70bf93c31aff3.String())
	__obf_ba969b69fcb8f496 := knownhosts.Normalize(__obf_d55af7c5d3d84f3d)
	__obf_ff11d3777bbc76a4 := []string{__obf_ca5ef72eec45bf9e}

	if __obf_ba969b69fcb8f496 != __obf_ca5ef72eec45bf9e {
		__obf_ff11d3777bbc76a4 = append(__obf_ff11d3777bbc76a4, __obf_ba969b69fcb8f496)
	}

	_, __obf_7f6f4ab0f2c744c3 = __obf_31fdd89e34f6d6ed.WriteString(knownhosts.Line(__obf_ff11d3777bbc76a4, __obf_b104c952b502a6cb) + "\n")

	return __obf_7f6f4ab0f2c744c3
}

// PrivateKey Loads a private and public key from "path" and returns a SSH ClientConfig to authenticate with the server
func PrivateKey(__obf_8396f3f955a12116 string, __obf_00b878bbf8da5796 string, __obf_4cc150c81db7ed0b int) (*ssh.ClientConfig, error) {
	__obf_d12f070c694d97b8, __obf_7f6f4ab0f2c744c3 := os.ReadFile(__obf_00b878bbf8da5796)

	if __obf_7f6f4ab0f2c744c3 != nil {
		return nil, __obf_7f6f4ab0f2c744c3
	}
	__obf_30bdb2325ce01286, __obf_7f6f4ab0f2c744c3 := ssh.ParsePrivateKey(__obf_d12f070c694d97b8)

	if __obf_7f6f4ab0f2c744c3 != nil {
		return nil, __obf_7f6f4ab0f2c744c3
	}
	__obf_ba6bfbe62b222d79, __obf_7f6f4ab0f2c744c3 := DefaultKnownHosts()
	if __obf_7f6f4ab0f2c744c3 != nil {
		return nil, __obf_7f6f4ab0f2c744c3
	}

	return &ssh.ClientConfig{
		User:    __obf_8396f3f955a12116,
		Timeout: time.Duration(__obf_4cc150c81db7ed0b) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_30bdb2325ce01286),
		},
		HostKeyCallback: __obf_ba6bfbe62b222d79,
	}, nil
}

// Creates the configuration for a client that authenticates with a password protected private key
func PrivateKeyWithPassphrase(__obf_8396f3f955a12116 string, __obf_4cc150c81db7ed0b int, __obf_27eb4f81fccc9708 []byte, __obf_00b878bbf8da5796 string) (*ssh.ClientConfig, error) {
	__obf_d12f070c694d97b8, __obf_7f6f4ab0f2c744c3 := os.ReadFile(__obf_00b878bbf8da5796)

	if __obf_7f6f4ab0f2c744c3 != nil {
		return nil, __obf_7f6f4ab0f2c744c3
	}
	__obf_30bdb2325ce01286, __obf_7f6f4ab0f2c744c3 := ssh.ParsePrivateKeyWithPassphrase(__obf_d12f070c694d97b8, __obf_27eb4f81fccc9708)

	if __obf_7f6f4ab0f2c744c3 != nil {
		return nil, __obf_7f6f4ab0f2c744c3
	}
	__obf_ba6bfbe62b222d79, __obf_7f6f4ab0f2c744c3 := DefaultKnownHosts()
	if __obf_7f6f4ab0f2c744c3 != nil {
		return nil, __obf_7f6f4ab0f2c744c3
	}

	return &ssh.ClientConfig{
		User:    __obf_8396f3f955a12116,
		Timeout: time.Duration(__obf_4cc150c81db7ed0b) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_30bdb2325ce01286),
		},
		HostKeyCallback: __obf_ba6bfbe62b222d79,
	}, nil
}

// Creates a configuration for a client that fetches public-private key from the SSH agent for authentication
func SSHAgent(__obf_8396f3f955a12116 string, __obf_4cc150c81db7ed0b int, __obf_f1033d588653ae36 ssh.HostKeyCallback) (*ssh.ClientConfig, error) {
	__obf_daa7c51f05118c67 := os.Getenv("SSH_AUTH_SOCK")
	__obf_2dffdfb1e34c3517, __obf_7f6f4ab0f2c744c3 := net.Dial("unix", __obf_daa7c51f05118c67)
	if __obf_7f6f4ab0f2c744c3 != nil {
		return &ssh.ClientConfig{
			Timeout: time.Duration(__obf_4cc150c81db7ed0b) * time.Second,
		}, __obf_7f6f4ab0f2c744c3
	}

	return &ssh.ClientConfig{
		User:    __obf_8396f3f955a12116,
		Timeout: time.Duration(__obf_4cc150c81db7ed0b) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agent.NewClient(__obf_2dffdfb1e34c3517).Signers),
		},
		HostKeyCallback: __obf_f1033d588653ae36,
	}, nil
}

// Creates a configuration for a client that authenticates using username and password
func PasswordKey(__obf_8396f3f955a12116 string, __obf_348a357e35399f7b string, __obf_4cc150c81db7ed0b int) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User:    __obf_8396f3f955a12116,
		Timeout: time.Duration(__obf_4cc150c81db7ed0b) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.Password(__obf_348a357e35399f7b),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

type WriteCounter struct {
	Total int64 // Total # of bytes written
}

func (__obf_8303b448bb45dd23 *WriteCounter) Write(__obf_cad45b7713c2b6dd []byte) (int, error) {
	__obf_77c3cd0ff127b595 := len(__obf_cad45b7713c2b6dd)
	__obf_8303b448bb45dd23.
		Total += int64(__obf_77c3cd0ff127b595)
	fmt.Fprintf(os.Stdout, "%.2f kb transferred\n", float64(__obf_8303b448bb45dd23.Total/1024))
	return __obf_77c3cd0ff127b595, nil
}
