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
	__obf_40e3a2f9e1b7ab67, __obf_6b7ad169ad557bb8 := os.UserHomeDir()
	if __obf_6b7ad169ad557bb8 != nil {
		return "", __obf_6b7ad169ad557bb8
	}

	return fmt.Sprintf("%s/.ssh/known_hosts", __obf_40e3a2f9e1b7ab67), __obf_6b7ad169ad557bb8
}

// DefaultKnownHosts returns host key callback from default known hosts path, and error if some.
func DefaultKnownHosts() (ssh.HostKeyCallback, error) {
	__obf_4da76925c58c56c3, __obf_6b7ad169ad557bb8 := DefaultKnownHostsPath()
	if __obf_6b7ad169ad557bb8 != nil {
		return nil, __obf_6b7ad169ad557bb8
	}

	return knownhosts.New(__obf_4da76925c58c56c3)
}

// CheckKnownHost checks is host in known hosts file.
// it returns is the host found in known_hosts file and error, if the host found in
// known_hosts file and error not nil that means public key mismatch, maybe MAN IN THE MIDDLE ATTACK! you should not handshake.
func CheckKnownHost(__obf_466f45cac1c44ab0 string, __obf_b94a00558fe734ed net.Addr, __obf_2a72b1f52a769dce ssh.PublicKey, __obf_aff9e2e428787858 string) (__obf_c543ebdeec5136c9 bool, __obf_6b7ad169ad557bb8 error) {

	var __obf_a57c71906604472b *knownhosts.KeyError

	// Fallback to default known_hosts file
	if __obf_aff9e2e428787858 == "" {
		__obf_4da76925c58c56c3, __obf_6b7ad169ad557bb8 := DefaultKnownHostsPath()
		if __obf_6b7ad169ad557bb8 != nil {
			return false, __obf_6b7ad169ad557bb8
		}
		__obf_aff9e2e428787858 = __obf_4da76925c58c56c3
	}
	__obf_c93e4171338de24f,

		// Get host key callback
		__obf_6b7ad169ad557bb8 := knownhosts.New(__obf_aff9e2e428787858)

	if __obf_6b7ad169ad557bb8 != nil {
		return false, __obf_6b7ad169ad557bb8
	}
	__obf_6b7ad169ad557bb8 = // check if host already exists.
		__obf_c93e4171338de24f(__obf_466f45cac1c44ab0, __obf_b94a00558fe734ed,

			// Known host already exists.
			__obf_2a72b1f52a769dce)

	if __obf_6b7ad169ad557bb8 == nil {
		return true, nil
	}

	// Make sure that the error returned from the callback is host not in file error.
	// If keyErr.Want is greater than 0 length, that means host is in file with different key.
	if errors.As(__obf_6b7ad169ad557bb8, &__obf_a57c71906604472b) && len(__obf_a57c71906604472b.Want) > 0 {
		return true, __obf_a57c71906604472b
	}

	// Key is not trusted because it is not in the file.
	return false, nil
}

// AddKnownHost add a a host to known hosts file.
func AddKnownHost(__obf_466f45cac1c44ab0 string, __obf_b94a00558fe734ed net.Addr, __obf_2a72b1f52a769dce ssh.PublicKey, __obf_aff9e2e428787858 string) (__obf_6b7ad169ad557bb8 error) {

	// Fallback to default known_hosts file
	if __obf_aff9e2e428787858 == "" {
		__obf_4da76925c58c56c3, __obf_6b7ad169ad557bb8 := DefaultKnownHostsPath()
		if __obf_6b7ad169ad557bb8 != nil {
			return __obf_6b7ad169ad557bb8
		}
		__obf_aff9e2e428787858 = __obf_4da76925c58c56c3
	}
	__obf_a0a1e10b9a0250a0, __obf_6b7ad169ad557bb8 := os.OpenFile(__obf_aff9e2e428787858, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if __obf_6b7ad169ad557bb8 != nil {
		return __obf_6b7ad169ad557bb8
	}

	defer __obf_a0a1e10b9a0250a0.Close()
	__obf_f90a3eef049983e9 := knownhosts.Normalize(__obf_b94a00558fe734ed.String())
	__obf_7b906e1a4b5ab614 := knownhosts.Normalize(__obf_466f45cac1c44ab0)
	__obf_5c6b0977930dac60 := []string{__obf_f90a3eef049983e9}

	if __obf_7b906e1a4b5ab614 != __obf_f90a3eef049983e9 {
		__obf_5c6b0977930dac60 = append(__obf_5c6b0977930dac60, __obf_7b906e1a4b5ab614)
	}

	_, __obf_6b7ad169ad557bb8 = __obf_a0a1e10b9a0250a0.WriteString(knownhosts.Line(__obf_5c6b0977930dac60, __obf_2a72b1f52a769dce) + "\n")

	return __obf_6b7ad169ad557bb8
}

// PrivateKey Loads a private and public key from "path" and returns a SSH ClientConfig to authenticate with the server
func PrivateKey(__obf_dc48463f77ccf261 string, __obf_4da76925c58c56c3 string, __obf_ca0b5b69d2f4a82a int) (*ssh.ClientConfig, error) {
	__obf_dcfe1a6106432ea8, __obf_6b7ad169ad557bb8 := os.ReadFile(__obf_4da76925c58c56c3)

	if __obf_6b7ad169ad557bb8 != nil {
		return nil, __obf_6b7ad169ad557bb8
	}
	__obf_27a7650912a63f68, __obf_6b7ad169ad557bb8 := ssh.ParsePrivateKey(__obf_dcfe1a6106432ea8)

	if __obf_6b7ad169ad557bb8 != nil {
		return nil, __obf_6b7ad169ad557bb8
	}
	__obf_c93e4171338de24f, __obf_6b7ad169ad557bb8 := DefaultKnownHosts()
	if __obf_6b7ad169ad557bb8 != nil {
		return nil, __obf_6b7ad169ad557bb8
	}

	return &ssh.ClientConfig{
		User:    __obf_dc48463f77ccf261,
		Timeout: time.Duration(__obf_ca0b5b69d2f4a82a) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_27a7650912a63f68),
		},
		HostKeyCallback: __obf_c93e4171338de24f,
	}, nil
}

// Creates the configuration for a client that authenticates with a password protected private key
func PrivateKeyWithPassphrase(__obf_dc48463f77ccf261 string, __obf_ca0b5b69d2f4a82a int, __obf_a95facae45895f41 []byte, __obf_4da76925c58c56c3 string) (*ssh.ClientConfig, error) {
	__obf_dcfe1a6106432ea8, __obf_6b7ad169ad557bb8 := os.ReadFile(__obf_4da76925c58c56c3)

	if __obf_6b7ad169ad557bb8 != nil {
		return nil, __obf_6b7ad169ad557bb8
	}
	__obf_27a7650912a63f68, __obf_6b7ad169ad557bb8 := ssh.ParsePrivateKeyWithPassphrase(__obf_dcfe1a6106432ea8, __obf_a95facae45895f41)

	if __obf_6b7ad169ad557bb8 != nil {
		return nil, __obf_6b7ad169ad557bb8
	}
	__obf_c93e4171338de24f, __obf_6b7ad169ad557bb8 := DefaultKnownHosts()
	if __obf_6b7ad169ad557bb8 != nil {
		return nil, __obf_6b7ad169ad557bb8
	}

	return &ssh.ClientConfig{
		User:    __obf_dc48463f77ccf261,
		Timeout: time.Duration(__obf_ca0b5b69d2f4a82a) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_27a7650912a63f68),
		},
		HostKeyCallback: __obf_c93e4171338de24f,
	}, nil
}

// Creates a configuration for a client that fetches public-private key from the SSH agent for authentication
func SSHAgent(__obf_dc48463f77ccf261 string, __obf_ca0b5b69d2f4a82a int, __obf_622ea150945c4238 ssh.HostKeyCallback) (*ssh.ClientConfig, error) {
	__obf_3bf6db4ca9423625 := os.Getenv("SSH_AUTH_SOCK")
	__obf_f01e4e4b7dc7e812, __obf_6b7ad169ad557bb8 := net.Dial("unix", __obf_3bf6db4ca9423625)
	if __obf_6b7ad169ad557bb8 != nil {
		return &ssh.ClientConfig{
			Timeout: time.Duration(__obf_ca0b5b69d2f4a82a) * time.Second,
		}, __obf_6b7ad169ad557bb8
	}

	return &ssh.ClientConfig{
		User:    __obf_dc48463f77ccf261,
		Timeout: time.Duration(__obf_ca0b5b69d2f4a82a) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agent.NewClient(__obf_f01e4e4b7dc7e812).Signers),
		},
		HostKeyCallback: __obf_622ea150945c4238,
	}, nil
}

// Creates a configuration for a client that authenticates using username and password
func PasswordKey(__obf_dc48463f77ccf261 string, __obf_1dfd3b16ffa6e253 string, __obf_ca0b5b69d2f4a82a int) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User:    __obf_dc48463f77ccf261,
		Timeout: time.Duration(__obf_ca0b5b69d2f4a82a) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.Password(__obf_1dfd3b16ffa6e253),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

type WriteCounter struct {
	Total int64 // Total # of bytes written
}

func (__obf_920e8b86db62c495 *WriteCounter) Write(__obf_7bb0ae9a5229a5f8 []byte) (int, error) {
	__obf_3f743908cd2c99a1 := len(__obf_7bb0ae9a5229a5f8)
	__obf_920e8b86db62c495.
		Total += int64(__obf_3f743908cd2c99a1)
	fmt.Fprintf(os.Stdout, "%.2f kb transferred\n", float64(__obf_920e8b86db62c495.Total/1024))
	return __obf_3f743908cd2c99a1, nil
}
