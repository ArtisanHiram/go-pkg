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
	__obf_6c577d710224ac3e, __obf_77d535ea9cfc3da2 := os.UserHomeDir()
	if __obf_77d535ea9cfc3da2 != nil {
		return "", __obf_77d535ea9cfc3da2
	}

	return fmt.Sprintf("%s/.ssh/known_hosts", __obf_6c577d710224ac3e), __obf_77d535ea9cfc3da2
}

// DefaultKnownHosts returns host key callback from default known hosts path, and error if any.
func DefaultKnownHosts() (ssh.HostKeyCallback, error) {
	__obf_44f8544fdc1ad7c2, __obf_77d535ea9cfc3da2 := DefaultKnownHostsPath()
	if __obf_77d535ea9cfc3da2 != nil {
		return nil, __obf_77d535ea9cfc3da2
	}

	return knownhosts.New(__obf_44f8544fdc1ad7c2)
}

// CheckKnownHost checks is host in known hosts file.
// it returns is the host found in known_hosts file and error, if the host found in
// known_hosts file and error not nil that means public key mismatch, maybe MAN IN THE MIDDLE ATTACK! you should not handshake.
func CheckKnownHost(__obf_f212d97459fd403e string, __obf_028d1134ee14a61e net.Addr, __obf_082db7478b31219c ssh.PublicKey, __obf_6d9dcf149e6088c5 string) (__obf_e9e8258c33836d87 bool, __obf_77d535ea9cfc3da2 error) {

	var __obf_216a702b9d17d1a3 *knownhosts.KeyError

	// Fallback to default known_hosts file
	if __obf_6d9dcf149e6088c5 == "" {
		__obf_44f8544fdc1ad7c2, __obf_77d535ea9cfc3da2 := DefaultKnownHostsPath()
		if __obf_77d535ea9cfc3da2 != nil {
			return false, __obf_77d535ea9cfc3da2
		}

		__obf_6d9dcf149e6088c5 = __obf_44f8544fdc1ad7c2
	}

	// Get host key callback
	__obf_17f97d86b99e63d7, __obf_77d535ea9cfc3da2 := knownhosts.New(__obf_6d9dcf149e6088c5)

	if __obf_77d535ea9cfc3da2 != nil {
		return false, __obf_77d535ea9cfc3da2
	}

	// check if host already exists.
	__obf_77d535ea9cfc3da2 = __obf_17f97d86b99e63d7(__obf_f212d97459fd403e, __obf_028d1134ee14a61e, __obf_082db7478b31219c)

	// Known host already exists.
	if __obf_77d535ea9cfc3da2 == nil {
		return true, nil
	}

	// Make sure that the error returned from the callback is host not in file error.
	// If keyErr.Want is greater than 0 length, that means host is in file with different key.
	if errors.As(__obf_77d535ea9cfc3da2, &__obf_216a702b9d17d1a3) && len(__obf_216a702b9d17d1a3.Want) > 0 {
		return true, __obf_216a702b9d17d1a3
	}

	// Key is not trusted because it is not in the file.
	return false, nil
}

// AddKnownHost add a a host to known hosts file.
func AddKnownHost(__obf_f212d97459fd403e string, __obf_028d1134ee14a61e net.Addr, __obf_082db7478b31219c ssh.PublicKey, __obf_6d9dcf149e6088c5 string) (__obf_77d535ea9cfc3da2 error) {

	// Fallback to default known_hosts file
	if __obf_6d9dcf149e6088c5 == "" {
		__obf_44f8544fdc1ad7c2, __obf_77d535ea9cfc3da2 := DefaultKnownHostsPath()
		if __obf_77d535ea9cfc3da2 != nil {
			return __obf_77d535ea9cfc3da2
		}

		__obf_6d9dcf149e6088c5 = __obf_44f8544fdc1ad7c2
	}

	__obf_23e7bbb689522332, __obf_77d535ea9cfc3da2 := os.OpenFile(__obf_6d9dcf149e6088c5, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if __obf_77d535ea9cfc3da2 != nil {
		return __obf_77d535ea9cfc3da2
	}

	defer __obf_23e7bbb689522332.Close()

	__obf_a503dd376051ec43 := knownhosts.Normalize(__obf_028d1134ee14a61e.String())
	__obf_5622d04b15ceffde := knownhosts.Normalize(__obf_f212d97459fd403e)
	__obf_b0e6648016c1ea4f := []string{__obf_a503dd376051ec43}

	if __obf_5622d04b15ceffde != __obf_a503dd376051ec43 {
		__obf_b0e6648016c1ea4f = append(__obf_b0e6648016c1ea4f, __obf_5622d04b15ceffde)
	}

	_, __obf_77d535ea9cfc3da2 = __obf_23e7bbb689522332.WriteString(knownhosts.Line(__obf_b0e6648016c1ea4f, __obf_082db7478b31219c) + "\n")

	return __obf_77d535ea9cfc3da2
}

// PrivateKey Loads a private and public key from "path" and returns a SSH ClientConfig to authenticate with the server
func PrivateKey(__obf_e340a497b13e63e3 string, __obf_44f8544fdc1ad7c2 string, __obf_5a010411bd86de3d int) (*ssh.ClientConfig, error) {
	__obf_9631b177d158b066, __obf_77d535ea9cfc3da2 := os.ReadFile(__obf_44f8544fdc1ad7c2)

	if __obf_77d535ea9cfc3da2 != nil {
		return nil, __obf_77d535ea9cfc3da2
	}

	__obf_431e5436f2b16cac, __obf_77d535ea9cfc3da2 := ssh.ParsePrivateKey(__obf_9631b177d158b066)

	if __obf_77d535ea9cfc3da2 != nil {
		return nil, __obf_77d535ea9cfc3da2
	}

	__obf_17f97d86b99e63d7, __obf_77d535ea9cfc3da2 := DefaultKnownHosts()
	if __obf_77d535ea9cfc3da2 != nil {
		return nil, __obf_77d535ea9cfc3da2
	}

	return &ssh.ClientConfig{
		User:    __obf_e340a497b13e63e3,
		Timeout: time.Duration(__obf_5a010411bd86de3d) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_431e5436f2b16cac),
		},
		HostKeyCallback: __obf_17f97d86b99e63d7,
	}, nil
}

// Creates the configuration for a client that authenticates with a password protected private key
func PrivateKeyWithPassphrase(__obf_e340a497b13e63e3 string, __obf_5a010411bd86de3d int, __obf_2db4c5edf2659ef4 []byte, __obf_44f8544fdc1ad7c2 string) (*ssh.ClientConfig, error) {
	__obf_9631b177d158b066, __obf_77d535ea9cfc3da2 := os.ReadFile(__obf_44f8544fdc1ad7c2)

	if __obf_77d535ea9cfc3da2 != nil {
		return nil, __obf_77d535ea9cfc3da2
	}
	__obf_431e5436f2b16cac, __obf_77d535ea9cfc3da2 := ssh.ParsePrivateKeyWithPassphrase(__obf_9631b177d158b066, __obf_2db4c5edf2659ef4)

	if __obf_77d535ea9cfc3da2 != nil {
		return nil, __obf_77d535ea9cfc3da2
	}

	__obf_17f97d86b99e63d7, __obf_77d535ea9cfc3da2 := DefaultKnownHosts()
	if __obf_77d535ea9cfc3da2 != nil {
		return nil, __obf_77d535ea9cfc3da2
	}

	return &ssh.ClientConfig{
		User:    __obf_e340a497b13e63e3,
		Timeout: time.Duration(__obf_5a010411bd86de3d) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_431e5436f2b16cac),
		},
		HostKeyCallback: __obf_17f97d86b99e63d7,
	}, nil
}

// Creates a configuration for a client that fetches public-private key from the SSH agent for authentication
func SSHAgent(__obf_e340a497b13e63e3 string, __obf_5a010411bd86de3d int, __obf_076bc4cc228010e4 ssh.HostKeyCallback) (*ssh.ClientConfig, error) {
	__obf_d41f34950de20fce := os.Getenv("SSH_AUTH_SOCK")
	__obf_b25dce559f6fd610, __obf_77d535ea9cfc3da2 := net.Dial("unix", __obf_d41f34950de20fce)
	if __obf_77d535ea9cfc3da2 != nil {
		return &ssh.ClientConfig{
			Timeout: time.Duration(__obf_5a010411bd86de3d) * time.Second,
		}, __obf_77d535ea9cfc3da2
	}

	return &ssh.ClientConfig{
		User:    __obf_e340a497b13e63e3,
		Timeout: time.Duration(__obf_5a010411bd86de3d) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agent.NewClient(__obf_b25dce559f6fd610).Signers),
		},
		HostKeyCallback: __obf_076bc4cc228010e4,
	}, nil
}

// Creates a configuration for a client that authenticates using username and password
func PasswordKey(__obf_e340a497b13e63e3 string, __obf_2429c62cdaeccfe5 string, __obf_5a010411bd86de3d int) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User:    __obf_e340a497b13e63e3,
		Timeout: time.Duration(__obf_5a010411bd86de3d) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.Password(__obf_2429c62cdaeccfe5),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

type WriteCounter struct {
	Total int64 // Total # of bytes written
}

func (__obf_099cdbac2cca91cd *WriteCounter) Write(__obf_e334805d30ee892b []byte) (int, error) {
	__obf_276bf15d6f932826 := len(__obf_e334805d30ee892b)
	__obf_099cdbac2cca91cd.Total += int64(__obf_276bf15d6f932826)
	fmt.Fprintf(os.Stdout, "%.2f kb transferred\n", float64(__obf_099cdbac2cca91cd.Total/1024))
	return __obf_276bf15d6f932826, nil
}
