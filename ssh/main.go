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
	__obf_ed10ec24520a2f2d, __obf_3e865493244170fc := os.UserHomeDir()
	if __obf_3e865493244170fc != nil {
		return "", __obf_3e865493244170fc
	}

	return fmt.Sprintf("%s/.ssh/known_hosts", __obf_ed10ec24520a2f2d), __obf_3e865493244170fc
}

// DefaultKnownHosts returns host key callback from default known hosts path, and error if some.
func DefaultKnownHosts() (ssh.HostKeyCallback, error) {
	__obf_264e38acc73f43ba, __obf_3e865493244170fc := DefaultKnownHostsPath()
	if __obf_3e865493244170fc != nil {
		return nil, __obf_3e865493244170fc
	}

	return knownhosts.New(__obf_264e38acc73f43ba)
}

// CheckKnownHost checks is host in known hosts file.
// it returns is the host found in known_hosts file and error, if the host found in
// known_hosts file and error not nil that means public key mismatch, maybe MAN IN THE MIDDLE ATTACK! you should not handshake.
func CheckKnownHost(__obf_2556aa0107f64466 string, __obf_e57dcd828c002c43 net.Addr, __obf_de9b141b8ad27cf5 ssh.PublicKey, __obf_24b7155ec0cc2d6f string) (__obf_bcc189c391308c3f bool, __obf_3e865493244170fc error) {

	var __obf_19e88dd2d6a57d10 *knownhosts.KeyError

	// Fallback to default known_hosts file
	if __obf_24b7155ec0cc2d6f == "" {
		__obf_264e38acc73f43ba, __obf_3e865493244170fc := DefaultKnownHostsPath()
		if __obf_3e865493244170fc != nil {
			return false, __obf_3e865493244170fc
		}
		__obf_24b7155ec0cc2d6f = __obf_264e38acc73f43ba
	}
	__obf_a9052eb0f53f9e9d,

		// Get host key callback
		__obf_3e865493244170fc := knownhosts.New(__obf_24b7155ec0cc2d6f)

	if __obf_3e865493244170fc != nil {
		return false, __obf_3e865493244170fc
	}
	__obf_3e865493244170fc = // check if host already exists.
		__obf_a9052eb0f53f9e9d(__obf_2556aa0107f64466, __obf_e57dcd828c002c43,

			// Known host already exists.
			__obf_de9b141b8ad27cf5)

	if __obf_3e865493244170fc == nil {
		return true, nil
	}

	// Make sure that the error returned from the callback is host not in file error.
	// If keyErr.Want is greater than 0 length, that means host is in file with different key.
	if errors.As(__obf_3e865493244170fc, &__obf_19e88dd2d6a57d10) && len(__obf_19e88dd2d6a57d10.Want) > 0 {
		return true, __obf_19e88dd2d6a57d10
	}

	// Key is not trusted because it is not in the file.
	return false, nil
}

// AddKnownHost add a a host to known hosts file.
func AddKnownHost(__obf_2556aa0107f64466 string, __obf_e57dcd828c002c43 net.Addr, __obf_de9b141b8ad27cf5 ssh.PublicKey, __obf_24b7155ec0cc2d6f string) (__obf_3e865493244170fc error) {

	// Fallback to default known_hosts file
	if __obf_24b7155ec0cc2d6f == "" {
		__obf_264e38acc73f43ba, __obf_3e865493244170fc := DefaultKnownHostsPath()
		if __obf_3e865493244170fc != nil {
			return __obf_3e865493244170fc
		}
		__obf_24b7155ec0cc2d6f = __obf_264e38acc73f43ba
	}
	__obf_2c85d146ccd0f71d, __obf_3e865493244170fc := os.OpenFile(__obf_24b7155ec0cc2d6f, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if __obf_3e865493244170fc != nil {
		return __obf_3e865493244170fc
	}

	defer __obf_2c85d146ccd0f71d.Close()
	__obf_ee90cee872c8fdfc := knownhosts.Normalize(__obf_e57dcd828c002c43.String())
	__obf_91cc51cc75203140 := knownhosts.Normalize(__obf_2556aa0107f64466)
	__obf_9619f54a841785fb := []string{__obf_ee90cee872c8fdfc}

	if __obf_91cc51cc75203140 != __obf_ee90cee872c8fdfc {
		__obf_9619f54a841785fb = append(__obf_9619f54a841785fb, __obf_91cc51cc75203140)
	}

	_, __obf_3e865493244170fc = __obf_2c85d146ccd0f71d.WriteString(knownhosts.Line(__obf_9619f54a841785fb, __obf_de9b141b8ad27cf5) + "\n")

	return __obf_3e865493244170fc
}

// PrivateKey Loads a private and public key from "path" and returns a SSH ClientConfig to authenticate with the server
func PrivateKey(__obf_2165d058b86aaacb string, __obf_264e38acc73f43ba string, __obf_5860b615d409de6d int) (*ssh.ClientConfig, error) {
	__obf_bf09a4d3c36cc1cf, __obf_3e865493244170fc := os.ReadFile(__obf_264e38acc73f43ba)

	if __obf_3e865493244170fc != nil {
		return nil, __obf_3e865493244170fc
	}
	__obf_06f77d7cc665b7e7, __obf_3e865493244170fc := ssh.ParsePrivateKey(__obf_bf09a4d3c36cc1cf)

	if __obf_3e865493244170fc != nil {
		return nil, __obf_3e865493244170fc
	}
	__obf_a9052eb0f53f9e9d, __obf_3e865493244170fc := DefaultKnownHosts()
	if __obf_3e865493244170fc != nil {
		return nil, __obf_3e865493244170fc
	}

	return &ssh.ClientConfig{
		User:    __obf_2165d058b86aaacb,
		Timeout: time.Duration(__obf_5860b615d409de6d) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_06f77d7cc665b7e7),
		},
		HostKeyCallback: __obf_a9052eb0f53f9e9d,
	}, nil
}

// Creates the configuration for a client that authenticates with a password protected private key
func PrivateKeyWithPassphrase(__obf_2165d058b86aaacb string, __obf_5860b615d409de6d int, __obf_8740b45df33f8298 []byte, __obf_264e38acc73f43ba string) (*ssh.ClientConfig, error) {
	__obf_bf09a4d3c36cc1cf, __obf_3e865493244170fc := os.ReadFile(__obf_264e38acc73f43ba)

	if __obf_3e865493244170fc != nil {
		return nil, __obf_3e865493244170fc
	}
	__obf_06f77d7cc665b7e7, __obf_3e865493244170fc := ssh.ParsePrivateKeyWithPassphrase(__obf_bf09a4d3c36cc1cf, __obf_8740b45df33f8298)

	if __obf_3e865493244170fc != nil {
		return nil, __obf_3e865493244170fc
	}
	__obf_a9052eb0f53f9e9d, __obf_3e865493244170fc := DefaultKnownHosts()
	if __obf_3e865493244170fc != nil {
		return nil, __obf_3e865493244170fc
	}

	return &ssh.ClientConfig{
		User:    __obf_2165d058b86aaacb,
		Timeout: time.Duration(__obf_5860b615d409de6d) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(__obf_06f77d7cc665b7e7),
		},
		HostKeyCallback: __obf_a9052eb0f53f9e9d,
	}, nil
}

// Creates a configuration for a client that fetches public-private key from the SSH agent for authentication
func SSHAgent(__obf_2165d058b86aaacb string, __obf_5860b615d409de6d int, __obf_d592d89209c87c9e ssh.HostKeyCallback) (*ssh.ClientConfig, error) {
	__obf_796120a9f8aa93b4 := os.Getenv("SSH_AUTH_SOCK")
	__obf_0e550c109abcd5f0, __obf_3e865493244170fc := net.Dial("unix", __obf_796120a9f8aa93b4)
	if __obf_3e865493244170fc != nil {
		return &ssh.ClientConfig{
			Timeout: time.Duration(__obf_5860b615d409de6d) * time.Second,
		}, __obf_3e865493244170fc
	}

	return &ssh.ClientConfig{
		User:    __obf_2165d058b86aaacb,
		Timeout: time.Duration(__obf_5860b615d409de6d) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeysCallback(agent.NewClient(__obf_0e550c109abcd5f0).Signers),
		},
		HostKeyCallback: __obf_d592d89209c87c9e,
	}, nil
}

// Creates a configuration for a client that authenticates using username and password
func PasswordKey(__obf_2165d058b86aaacb string, __obf_c574efec501a9d88 string, __obf_5860b615d409de6d int) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User:    __obf_2165d058b86aaacb,
		Timeout: time.Duration(__obf_5860b615d409de6d) * time.Second,
		Auth: []ssh.AuthMethod{
			ssh.Password(__obf_c574efec501a9d88),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

type WriteCounter struct {
	Total int64 // Total # of bytes written
}

func (__obf_a12ea930fe78455b *WriteCounter) Write(__obf_b695807ca5d8ea85 []byte) (int, error) {
	__obf_063ac61e0a1a8818 := len(__obf_b695807ca5d8ea85)
	__obf_a12ea930fe78455b.
		Total += int64(__obf_063ac61e0a1a8818)
	fmt.Fprintf(os.Stdout, "%.2f kb transferred\n", float64(__obf_a12ea930fe78455b.Total/1024))
	return __obf_063ac61e0a1a8818, nil
}
