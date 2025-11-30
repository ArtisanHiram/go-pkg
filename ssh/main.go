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
	__obf_9efc7b77e0fb7039, __obf_fe854faece812b24 := os.UserHomeDir()
	if __obf_fe854faece812b24 != nil {
		return "", __obf_fe854faece812b24
	}

	return fmt.Sprintf("%s/.ssh/known_hosts", __obf_9efc7b77e0fb7039), __obf_fe854faece812b24
}

// DefaultKnownHosts returns host key callback from default known hosts path, and error if some.
func DefaultKnownHosts() (ssh.HostKeyCallback, error) {
	__obf_00d6aa5c6f1a3f83, __obf_fe854faece812b24 := DefaultKnownHostsPath()
	if __obf_fe854faece812b24 != nil {
		return nil, __obf_fe854faece812b24
	}

	return knownhosts.New(__obf_00d6aa5c6f1a3f83)
}

// CheckKnownHost checks is host in known hosts file.
// it returns is the host found in known_hosts file and error, if the host found in
// known_hosts file and error not nil that means public key mismatch, maybe MAN IN THE MIDDLE ATTACK! you should not handshake.
func CheckKnownHost(__obf_b32f9ae2b92fa9c4 string, __obf_4dea84270d4357ce net.Addr, __obf_34a6760e4823bcc4 ssh.PublicKey, __obf_7d93f0370fdc9c0d string) (__obf_5b02867880f68f70 bool, __obf_fe854faece812b24 error) {

	var __obf_937dc6d0e2235461 *knownhosts.KeyError

	// Fallback to default known_hosts file
	if __obf_7d93f0370fdc9c0d == "" {
		__obf_00d6aa5c6f1a3f83, __obf_fe854faece812b24 := DefaultKnownHostsPath()
		if __obf_fe854faece812b24 != nil {
			return false, __obf_fe854faece812b24
		}
		__obf_7d93f0370fdc9c0d = __obf_00d6aa5c6f1a3f83
	}
	__obf_8a489fd67f4fd8b7,

		// Get host key callback
		__obf_fe854faece812b24 := knownhosts.New(__obf_7d93f0370fdc9c0d)

	if __obf_fe854faece812b24 != nil {
		return false, __obf_fe854faece812b24
	}
	__obf_fe854faece812b24 = // check if host already exists.
		__obf_8a489fd67f4fd8b7(__obf_b32f9ae2b92fa9c4, __obf_4dea84270d4357ce,

			// Known host already exists.
			__obf_34a6760e4823bcc4)

	if __obf_fe854faece812b24 == nil {
		return true, nil
	}

	// Make sure that the error returned from the callback is host not in file error.
	// If keyErr.Want is greater than 0 length, that means host is in file with different key.
	if errors.As(__obf_fe854faece812b24, &__obf_937dc6d0e2235461) && len(__obf_937dc6d0e2235461.Want) > 0 {
		return true, __obf_937dc6d0e2235461
	}

	// Key is not trusted because it is not in the file.
	return false, nil
}

// AddKnownHost add a a host to known hosts file.
func AddKnownHost(__obf_b32f9ae2b92fa9c4 string, __obf_4dea84270d4357ce net.Addr, __obf_34a6760e4823bcc4 ssh.PublicKey, __obf_7d93f0370fdc9c0d string) (__obf_fe854faece812b24 error) {

	// Fallback to default known_hosts file
	if __obf_7d93f0370fdc9c0d == "" {
		__obf_00d6aa5c6f1a3f83, __obf_fe854faece812b24 := DefaultKnownHostsPath()
		if __obf_fe854faece812b24 != nil {
			return __obf_fe854faece812b24
		}
		__obf_7d93f0370fdc9c0d = __obf_00d6aa5c6f1a3f83
	}
	__obf_160f003b48872623, __obf_fe854faece812b24 := os.OpenFile(__obf_7d93f0370fdc9c0d, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if __obf_fe854faece812b24 != nil {
		return __obf_fe854faece812b24
	}

	defer __obf_160f003b48872623.Close()
	__obf_0218e39a5da60977 := knownhosts.Normalize(__obf_4dea84270d4357ce.String())
	__obf_2e4397fafe88a57e := knownhosts.Normalize(__obf_b32f9ae2b92fa9c4)
	__obf_2661b8d352a1b3d4 := []string{__obf_0218e39a5da60977}

	if __obf_2e4397fafe88a57e != __obf_0218e39a5da60977 {
		__obf_2661b8d352a1b3d4 = append(__obf_2661b8d352a1b3d4, __obf_2e4397fafe88a57e)
	}

	_, __obf_fe854faece812b24 = __obf_160f003b48872623.WriteString(knownhosts.Line(__obf_2661b8d352a1b3d4, __obf_34a6760e4823bcc4) + "\n")

	return __obf_fe854faece812b24
}

// PrivateKey Loads a private and public key from "path" and returns a SSH ClientConfig to authenticate with the server
func PrivateKey(__obf_0a38f9669a1e0be7 string, __obf_00d6aa5c6f1a3f83 string, __obf_9f302ccafea5dbad int) (*ssh.ClientConfig, error) {
	__obf_9fbc51d32a918c61, __obf_fe854faece812b24 := os.ReadFile(__obf_00d6aa5c6f1a3f83)

	if __obf_fe854faece812b24 != nil {
		return nil, __obf_fe854faece812b24
	}
	__obf_b2c2b1cf02bb1176, __obf_fe854faece812b24 := ssh.ParsePrivateKey(__obf_9fbc51d32a918c61)

	if __obf_fe854faece812b24 != nil {
		return nil, __obf_fe854faece812b24
	}
	__obf_8a489fd67f4fd8b7, __obf_fe854faece812b24 := DefaultKnownHosts()
	if __obf_fe854faece812b24 != nil {
		return nil, __obf_fe854faece812b24
	}

	return &ssh.ClientConfig{
		User:    __obf_0a38f9669a1e0be7,
		Timeout: time.Duration(__obf_9f302ccafea5dbad) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_b2c2b1cf02bb1176),
		},
		HostKeyCallback: __obf_8a489fd67f4fd8b7,
	}, nil
}

// Creates the configuration for a client that authenticates with a password protected private key
func PrivateKeyWithPassphrase(__obf_0a38f9669a1e0be7 string, __obf_9f302ccafea5dbad int, __obf_3de1c879794c2fe1 []byte, __obf_00d6aa5c6f1a3f83 string) (*ssh.ClientConfig, error) {
	__obf_9fbc51d32a918c61, __obf_fe854faece812b24 := os.ReadFile(__obf_00d6aa5c6f1a3f83)

	if __obf_fe854faece812b24 != nil {
		return nil, __obf_fe854faece812b24
	}
	__obf_b2c2b1cf02bb1176, __obf_fe854faece812b24 := ssh.ParsePrivateKeyWithPassphrase(__obf_9fbc51d32a918c61, __obf_3de1c879794c2fe1)

	if __obf_fe854faece812b24 != nil {
		return nil, __obf_fe854faece812b24
	}
	__obf_8a489fd67f4fd8b7, __obf_fe854faece812b24 := DefaultKnownHosts()
	if __obf_fe854faece812b24 != nil {
		return nil, __obf_fe854faece812b24
	}

	return &ssh.ClientConfig{
		User:    __obf_0a38f9669a1e0be7,
		Timeout: time.Duration(__obf_9f302ccafea5dbad) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_b2c2b1cf02bb1176),
		},
		HostKeyCallback: __obf_8a489fd67f4fd8b7,
	}, nil
}

// Creates a configuration for a client that fetches public-private key from the SSH agent for authentication
func SSHAgent(__obf_0a38f9669a1e0be7 string, __obf_9f302ccafea5dbad int, __obf_51ff25ff488838c3 ssh.HostKeyCallback) (*ssh.ClientConfig, error) {
	__obf_d9d768c7f56be1ba := os.Getenv("SSH_AUTH_SOCK")
	__obf_bb04ef02d54f927c, __obf_fe854faece812b24 := net.Dial("unix", __obf_d9d768c7f56be1ba)
	if __obf_fe854faece812b24 != nil {
		return &ssh.ClientConfig{
			Timeout: time.Duration(__obf_9f302ccafea5dbad) * time.Second,
		}, __obf_fe854faece812b24
	}

	return &ssh.ClientConfig{
		User:    __obf_0a38f9669a1e0be7,
		Timeout: time.Duration(__obf_9f302ccafea5dbad) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agent.NewClient(__obf_bb04ef02d54f927c).Signers),
		},
		HostKeyCallback: __obf_51ff25ff488838c3,
	}, nil
}

// Creates a configuration for a client that authenticates using username and password
func PasswordKey(__obf_0a38f9669a1e0be7 string, __obf_17f15b8197ad1ece string, __obf_9f302ccafea5dbad int) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User:    __obf_0a38f9669a1e0be7,
		Timeout: time.Duration(__obf_9f302ccafea5dbad) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.Password(__obf_17f15b8197ad1ece),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

type WriteCounter struct {
	Total int64 // Total # of bytes written
}

func (__obf_4e563c90eae509e7 *WriteCounter) Write(__obf_e7bff70265c42906 []byte) (int, error) {
	__obf_4ffd377394debb36 := len(__obf_e7bff70265c42906)
	__obf_4e563c90eae509e7.
		Total += int64(__obf_4ffd377394debb36)
	fmt.Fprintf(os.Stdout, "%.2f kb transferred\n", float64(__obf_4e563c90eae509e7.Total/1024))
	return __obf_4ffd377394debb36, nil
}
