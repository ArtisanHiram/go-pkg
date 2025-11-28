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
	__obf_272399742fd75d7d, __obf_6eb8d6029fd3be00 := os.UserHomeDir()
	if __obf_6eb8d6029fd3be00 != nil {
		return "", __obf_6eb8d6029fd3be00
	}

	return fmt.Sprintf("%s/.ssh/known_hosts", __obf_272399742fd75d7d), __obf_6eb8d6029fd3be00
}

// DefaultKnownHosts returns host key callback from default known hosts path, and error if any.
func DefaultKnownHosts() (ssh.HostKeyCallback, error) {
	__obf_31c40f41410e07e1, __obf_6eb8d6029fd3be00 := DefaultKnownHostsPath()
	if __obf_6eb8d6029fd3be00 != nil {
		return nil, __obf_6eb8d6029fd3be00
	}

	return knownhosts.New(__obf_31c40f41410e07e1)
}

// CheckKnownHost checks is host in known hosts file.
// it returns is the host found in known_hosts file and error, if the host found in
// known_hosts file and error not nil that means public key mismatch, maybe MAN IN THE MIDDLE ATTACK! you should not handshake.
func CheckKnownHost(__obf_601a0a443b0788cd string, __obf_b9d3a70bdbc1c449 net.Addr, __obf_149eb49612a947d1 ssh.PublicKey, __obf_854af891088d842c string) (__obf_47647716efce58fb bool, __obf_6eb8d6029fd3be00 error) {

	var __obf_eed7ce38672bf610 *knownhosts.KeyError

	// Fallback to default known_hosts file
	if __obf_854af891088d842c == "" {
		__obf_31c40f41410e07e1, __obf_6eb8d6029fd3be00 := DefaultKnownHostsPath()
		if __obf_6eb8d6029fd3be00 != nil {
			return false, __obf_6eb8d6029fd3be00
		}

		__obf_854af891088d842c = __obf_31c40f41410e07e1
	}

	// Get host key callback
	__obf_82ef3aa63811c2bc, __obf_6eb8d6029fd3be00 := knownhosts.New(__obf_854af891088d842c)

	if __obf_6eb8d6029fd3be00 != nil {
		return false, __obf_6eb8d6029fd3be00
	}

	// check if host already exists.
	__obf_6eb8d6029fd3be00 = __obf_82ef3aa63811c2bc(__obf_601a0a443b0788cd, __obf_b9d3a70bdbc1c449, __obf_149eb49612a947d1)

	// Known host already exists.
	if __obf_6eb8d6029fd3be00 == nil {
		return true, nil
	}

	// Make sure that the error returned from the callback is host not in file error.
	// If keyErr.Want is greater than 0 length, that means host is in file with different key.
	if errors.As(__obf_6eb8d6029fd3be00, &__obf_eed7ce38672bf610) && len(__obf_eed7ce38672bf610.Want) > 0 {
		return true, __obf_eed7ce38672bf610
	}

	// Key is not trusted because it is not in the file.
	return false, nil
}

// AddKnownHost add a a host to known hosts file.
func AddKnownHost(__obf_601a0a443b0788cd string, __obf_b9d3a70bdbc1c449 net.Addr, __obf_149eb49612a947d1 ssh.PublicKey, __obf_854af891088d842c string) (__obf_6eb8d6029fd3be00 error) {

	// Fallback to default known_hosts file
	if __obf_854af891088d842c == "" {
		__obf_31c40f41410e07e1, __obf_6eb8d6029fd3be00 := DefaultKnownHostsPath()
		if __obf_6eb8d6029fd3be00 != nil {
			return __obf_6eb8d6029fd3be00
		}

		__obf_854af891088d842c = __obf_31c40f41410e07e1
	}

	__obf_407d8c375b09b323, __obf_6eb8d6029fd3be00 := os.OpenFile(__obf_854af891088d842c, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if __obf_6eb8d6029fd3be00 != nil {
		return __obf_6eb8d6029fd3be00
	}

	defer __obf_407d8c375b09b323.Close()

	__obf_193c7c86d4e3461b := knownhosts.Normalize(__obf_b9d3a70bdbc1c449.String())
	__obf_dba12a730951a15f := knownhosts.Normalize(__obf_601a0a443b0788cd)
	__obf_8f5dabe66ccfa58e := []string{__obf_193c7c86d4e3461b}

	if __obf_dba12a730951a15f != __obf_193c7c86d4e3461b {
		__obf_8f5dabe66ccfa58e = append(__obf_8f5dabe66ccfa58e, __obf_dba12a730951a15f)
	}

	_, __obf_6eb8d6029fd3be00 = __obf_407d8c375b09b323.WriteString(knownhosts.Line(__obf_8f5dabe66ccfa58e, __obf_149eb49612a947d1) + "\n")

	return __obf_6eb8d6029fd3be00
}

// PrivateKey Loads a private and public key from "path" and returns a SSH ClientConfig to authenticate with the server
func PrivateKey(__obf_cf27e6b5e0104ecc string, __obf_31c40f41410e07e1 string, __obf_1eea5d91d967dee8 int) (*ssh.ClientConfig, error) {
	__obf_66eeaeb0d48c981f, __obf_6eb8d6029fd3be00 := os.ReadFile(__obf_31c40f41410e07e1)

	if __obf_6eb8d6029fd3be00 != nil {
		return nil, __obf_6eb8d6029fd3be00
	}

	__obf_8089f89a8cf428de, __obf_6eb8d6029fd3be00 := ssh.ParsePrivateKey(__obf_66eeaeb0d48c981f)

	if __obf_6eb8d6029fd3be00 != nil {
		return nil, __obf_6eb8d6029fd3be00
	}

	__obf_82ef3aa63811c2bc, __obf_6eb8d6029fd3be00 := DefaultKnownHosts()
	if __obf_6eb8d6029fd3be00 != nil {
		return nil, __obf_6eb8d6029fd3be00
	}

	return &ssh.ClientConfig{
		User:    __obf_cf27e6b5e0104ecc,
		Timeout: time.Duration(__obf_1eea5d91d967dee8) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_8089f89a8cf428de),
		},
		HostKeyCallback: __obf_82ef3aa63811c2bc,
	}, nil
}

// Creates the configuration for a client that authenticates with a password protected private key
func PrivateKeyWithPassphrase(__obf_cf27e6b5e0104ecc string, __obf_1eea5d91d967dee8 int, __obf_28dd7412e5846285 []byte, __obf_31c40f41410e07e1 string) (*ssh.ClientConfig, error) {
	__obf_66eeaeb0d48c981f, __obf_6eb8d6029fd3be00 := os.ReadFile(__obf_31c40f41410e07e1)

	if __obf_6eb8d6029fd3be00 != nil {
		return nil, __obf_6eb8d6029fd3be00
	}
	__obf_8089f89a8cf428de, __obf_6eb8d6029fd3be00 := ssh.ParsePrivateKeyWithPassphrase(__obf_66eeaeb0d48c981f, __obf_28dd7412e5846285)

	if __obf_6eb8d6029fd3be00 != nil {
		return nil, __obf_6eb8d6029fd3be00
	}

	__obf_82ef3aa63811c2bc, __obf_6eb8d6029fd3be00 := DefaultKnownHosts()
	if __obf_6eb8d6029fd3be00 != nil {
		return nil, __obf_6eb8d6029fd3be00
	}

	return &ssh.ClientConfig{
		User:    __obf_cf27e6b5e0104ecc,
		Timeout: time.Duration(__obf_1eea5d91d967dee8) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_8089f89a8cf428de),
		},
		HostKeyCallback: __obf_82ef3aa63811c2bc,
	}, nil
}

// Creates a configuration for a client that fetches public-private key from the SSH agent for authentication
func SSHAgent(__obf_cf27e6b5e0104ecc string, __obf_1eea5d91d967dee8 int, __obf_8f855facca4364f5 ssh.HostKeyCallback) (*ssh.ClientConfig, error) {
	__obf_f5fd1bf3112a6ca1 := os.Getenv("SSH_AUTH_SOCK")
	__obf_86c40e1a6abe2d34, __obf_6eb8d6029fd3be00 := net.Dial("unix", __obf_f5fd1bf3112a6ca1)
	if __obf_6eb8d6029fd3be00 != nil {
		return &ssh.ClientConfig{
			Timeout: time.Duration(__obf_1eea5d91d967dee8) * time.Second,
		}, __obf_6eb8d6029fd3be00
	}

	return &ssh.ClientConfig{
		User:    __obf_cf27e6b5e0104ecc,
		Timeout: time.Duration(__obf_1eea5d91d967dee8) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agent.NewClient(__obf_86c40e1a6abe2d34).Signers),
		},
		HostKeyCallback: __obf_8f855facca4364f5,
	}, nil
}

// Creates a configuration for a client that authenticates using username and password
func PasswordKey(__obf_cf27e6b5e0104ecc string, __obf_64d311e22f2e8698 string, __obf_1eea5d91d967dee8 int) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User:    __obf_cf27e6b5e0104ecc,
		Timeout: time.Duration(__obf_1eea5d91d967dee8) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.Password(__obf_64d311e22f2e8698),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

type WriteCounter struct {
	Total int64 // Total # of bytes written
}

func (__obf_853d24f93a091fa8 *WriteCounter) Write(__obf_1b391f7f704d806d []byte) (int, error) {
	__obf_ef825296c0896fe1 := len(__obf_1b391f7f704d806d)
	__obf_853d24f93a091fa8.Total += int64(__obf_ef825296c0896fe1)
	fmt.Fprintf(os.Stdout, "%.2f kb transferred\n", float64(__obf_853d24f93a091fa8.Total/1024))
	return __obf_ef825296c0896fe1, nil
}
