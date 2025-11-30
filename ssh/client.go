// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package ssh

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"path/filepath"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// SSHClient represents Goph client.
type SSHClient struct {
	Addr   string
	Config *ssh.ClientConfig
	*ssh.Client
}

// NewConn returns new client and error if some.
func NewSSHClient(__obf_3442d542a5a038c5, __obf_f64d00c862cf3392 string, __obf_3e68bbead4ecc947 *ssh.ClientConfig) (__obf_efb6a09747bd1703 *SSHClient) {
	// c.Client, err = Dial("tcp", addr, port)
	return &SSHClient{
		Addr:   net.JoinHostPort(__obf_3442d542a5a038c5, __obf_f64d00c862cf3392),
		Config: __obf_3e68bbead4ecc947,
	}
}

func (__obf_efb6a09747bd1703 *SSHClient) init() error {
	if __obf_efb6a09747bd1703.Client == nil {
		var __obf_977f2a9e62be94d2 error
		__obf_efb6a09747bd1703.
			Client, __obf_977f2a9e62be94d2 = ssh.Dial("tcp", __obf_efb6a09747bd1703.Addr, __obf_efb6a09747bd1703.Config)
		if __obf_977f2a9e62be94d2 != nil {
			return __obf_977f2a9e62be94d2
		}
	}
	return nil
}

// Run starts a new SSH session and runs the cmd, it returns CombinedOutput and err if some.
func (__obf_efb6a09747bd1703 SSHClient) Run(__obf_5d03f47a0a0a6768 string) ([]byte, error) {

	var __obf_977f2a9e62be94d2 = __obf_efb6a09747bd1703.init()
	if __obf_977f2a9e62be94d2 != nil {
		return nil, __obf_977f2a9e62be94d2
	}

	var __obf_a0da51cb0d75b7e7 *ssh.Session
	if __obf_a0da51cb0d75b7e7, __obf_977f2a9e62be94d2 = __obf_efb6a09747bd1703.NewSession(); __obf_977f2a9e62be94d2 != nil {
		return nil, __obf_977f2a9e62be94d2
	}

	defer __obf_a0da51cb0d75b7e7.Close()

	return __obf_a0da51cb0d75b7e7.CombinedOutput(__obf_5d03f47a0a0a6768)
}

// NewSftp returns new sftp client and error if some.
func (__obf_efb6a09747bd1703 SSHClient) NewSftp(__obf_eaa3d955db051788 ...sftp.ClientOption) (*sftp.Client, error) {
	if __obf_977f2a9e62be94d2 := __obf_efb6a09747bd1703.init(); __obf_977f2a9e62be94d2 != nil {
		return nil, __obf_977f2a9e62be94d2
	}
	return sftp.NewClient(__obf_efb6a09747bd1703.Client, __obf_eaa3d955db051788...)
}

// uploadFile a local file to remote server!
func (__obf_efb6a09747bd1703 SSHClient) __obf_f6e73a479738f418(__obf_db10cd1b56bfe7c9 *sftp.Client, __obf_57c08a37393f9fbc string, __obf_a71b4b218489f05c string) (__obf_977f2a9e62be94d2 error) {

	var __obf_019f3e36e7030bc1 *os.File
	__obf_019f3e36e7030bc1, __obf_977f2a9e62be94d2 = os.Open(__obf_57c08a37393f9fbc)
	if __obf_977f2a9e62be94d2 != nil {
		return
	}
	defer __obf_019f3e36e7030bc1.Close()

	var __obf_37760f54a739cb2c *sftp.File
	__obf_2e3f259d3d12f102 := path.Join(__obf_a71b4b218489f05c, filepath.Base(__obf_57c08a37393f9fbc))
	__obf_37760f54a739cb2c, __obf_977f2a9e62be94d2 = __obf_db10cd1b56bfe7c9.Create(__obf_2e3f259d3d12f102)
	if __obf_977f2a9e62be94d2 != nil {
		__obf_977f2a9e62be94d2 = fmt.Errorf("[Create: %s] %w", __obf_2e3f259d3d12f102, __obf_977f2a9e62be94d2)
		return
	}
	defer __obf_37760f54a739cb2c.Close()

	// size := len(s)
	// info, err := local.Stat()
	// if err != nil {
	// 	return err
	// }

	var __obf_6f0b078822cbc6e7 int64
	__obf_6f0b078822cbc6e7, __obf_977f2a9e62be94d2 = io.Copy(__obf_37760f54a739cb2c, io.TeeReader(__obf_019f3e36e7030bc1, &WriteCounter{}))
	fmt.Fprintf(os.Stdout, "Transferred %.2f kb\n", float64(__obf_6f0b078822cbc6e7/1024))
	return
}

// Upload a local file or directory to remote server!
func (__obf_efb6a09747bd1703 SSHClient) Upload(__obf_57c08a37393f9fbc string, __obf_a71b4b218489f05c string) (__obf_977f2a9e62be94d2 error) {
	__obf_7095fe63080466c1,

		//获取路径的属性
		__obf_977f2a9e62be94d2 := os.Stat(__obf_57c08a37393f9fbc)
	if __obf_977f2a9e62be94d2 != nil {
		return
	}
	__obf_4dc256ce1a8111ed, __obf_977f2a9e62be94d2 := __obf_efb6a09747bd1703.NewSftp()
	if __obf_977f2a9e62be94d2 != nil {
		return
	}
	defer __obf_4dc256ce1a8111ed.Close()

	// 判断是否是文件夹
	if __obf_7095fe63080466c1.IsDir() {
		var __obf_c69d17637d4c24ab []os.DirEntry
		__obf_c69d17637d4c24ab, __obf_977f2a9e62be94d2 = os.ReadDir(__obf_57c08a37393f9fbc)
		if __obf_977f2a9e62be94d2 != nil {
			return
		}
		__obf_a71b4b218489f05c = // 先创建最外层文件夹
			path.Join(__obf_a71b4b218489f05c, __obf_7095fe63080466c1.Name())
		__obf_977f2a9e62be94d2 = __obf_4dc256ce1a8111ed.Mkdir(__obf_a71b4b218489f05c)
		if __obf_977f2a9e62be94d2 != nil {
			return fmt.Errorf("[Mkdir] %s failed: %w", __obf_a71b4b218489f05c, __obf_977f2a9e62be94d2)
		}
		// 遍历文件夹内容
		for _, __obf_0ee3628cf1bb91b6 := range __obf_c69d17637d4c24ab {
			__obf_977f2a9e62be94d2 = // 判断是否是文件,是文件直接上传.是文件夹,先远程创建文件夹,再递归复制内部文件
				__obf_efb6a09747bd1703.Upload(path.Join(__obf_57c08a37393f9fbc, __obf_0ee3628cf1bb91b6.Name()), __obf_a71b4b218489f05c)
			if __obf_977f2a9e62be94d2 != nil {
				return fmt.Errorf("[Upload to: %s] %w", __obf_a71b4b218489f05c, __obf_977f2a9e62be94d2)
			}
		}
	} else {
		__obf_977f2a9e62be94d2 = __obf_efb6a09747bd1703.__obf_f6e73a479738f418(__obf_4dc256ce1a8111ed, __obf_57c08a37393f9fbc, __obf_a71b4b218489f05c)
	}
	return
}

// Download file from remote server!
func (__obf_efb6a09747bd1703 SSHClient) Download(__obf_2e3f259d3d12f102 string, __obf_57c08a37393f9fbc string) error {
	var __obf_2ec5a8c6a5ad7f2c = path.Join(__obf_57c08a37393f9fbc, filepath.Base(__obf_2e3f259d3d12f102))
	_, __obf_977f2a9e62be94d2 := os.Stat(__obf_2ec5a8c6a5ad7f2c)
	if __obf_977f2a9e62be94d2 == nil {
		return errors.New("file exists")
	}
	if os.IsNotExist(__obf_977f2a9e62be94d2) {
		__obf_019f3e36e7030bc1, __obf_977f2a9e62be94d2 := os.Create(__obf_2ec5a8c6a5ad7f2c)
		if __obf_977f2a9e62be94d2 != nil {
			return fmt.Errorf("[Create: %s] %w", __obf_2ec5a8c6a5ad7f2c, __obf_977f2a9e62be94d2)
		}
		defer __obf_019f3e36e7030bc1.Close()
		__obf_4dc256ce1a8111ed, __obf_977f2a9e62be94d2 := __obf_efb6a09747bd1703.NewSftp()
		if __obf_977f2a9e62be94d2 != nil {

			return __obf_977f2a9e62be94d2
		}
		defer __obf_4dc256ce1a8111ed.Close()
		__obf_37760f54a739cb2c, __obf_977f2a9e62be94d2 := __obf_4dc256ce1a8111ed.Open(__obf_2e3f259d3d12f102)
		if __obf_977f2a9e62be94d2 != nil {
			return __obf_977f2a9e62be94d2
		}
		defer __obf_37760f54a739cb2c.Close()

		if _, __obf_977f2a9e62be94d2 = io.Copy(__obf_019f3e36e7030bc1, __obf_37760f54a739cb2c); __obf_977f2a9e62be94d2 != nil {
			return __obf_977f2a9e62be94d2
		}

		return __obf_019f3e36e7030bc1.Sync()
	}
	return __obf_977f2a9e62be94d2
}
