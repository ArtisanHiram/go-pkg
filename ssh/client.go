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

// NewConn returns new client and error if any.
func NewSSHClient(__obf_a3630518afad1edc, __obf_d96c7f14b9ba1795 string, __obf_8b82b34bdcc686e1 *ssh.ClientConfig) (__obf_1571aa912a5dbabc *SSHClient) {
	// c.Client, err = Dial("tcp", addr, port)
	return &SSHClient{
		Addr:   net.JoinHostPort(__obf_a3630518afad1edc, __obf_d96c7f14b9ba1795),
		Config: __obf_8b82b34bdcc686e1,
	}
}

func (__obf_1571aa912a5dbabc *SSHClient) __obf_37245479cc819b77() error {
	if __obf_1571aa912a5dbabc.Client == nil {
		var __obf_0b884480f1bc4644 error
		__obf_1571aa912a5dbabc.Client, __obf_0b884480f1bc4644 = ssh.Dial("tcp", __obf_1571aa912a5dbabc.Addr, __obf_1571aa912a5dbabc.Config)
		if __obf_0b884480f1bc4644 != nil {
			return __obf_0b884480f1bc4644
		}
	}
	return nil
}

// Run starts a new SSH session and runs the cmd, it returns CombinedOutput and err if any.
func (__obf_1571aa912a5dbabc SSHClient) Run(__obf_d4ff50f3c1d374a7 string) ([]byte, error) {

	var __obf_0b884480f1bc4644 = __obf_1571aa912a5dbabc.__obf_37245479cc819b77()
	if __obf_0b884480f1bc4644 != nil {
		return nil, __obf_0b884480f1bc4644
	}

	var __obf_05a53812811a5563 *ssh.Session
	if __obf_05a53812811a5563, __obf_0b884480f1bc4644 = __obf_1571aa912a5dbabc.NewSession(); __obf_0b884480f1bc4644 != nil {
		return nil, __obf_0b884480f1bc4644
	}

	defer __obf_05a53812811a5563.Close()

	return __obf_05a53812811a5563.CombinedOutput(__obf_d4ff50f3c1d374a7)
}

// NewSftp returns new sftp client and error if any.
func (__obf_1571aa912a5dbabc SSHClient) NewSftp(__obf_95d20e97411c3fd8 ...sftp.ClientOption) (*sftp.Client, error) {
	if __obf_0b884480f1bc4644 := __obf_1571aa912a5dbabc.__obf_37245479cc819b77(); __obf_0b884480f1bc4644 != nil {
		return nil, __obf_0b884480f1bc4644
	}
	return sftp.NewClient(__obf_1571aa912a5dbabc.Client, __obf_95d20e97411c3fd8...)
}

// uploadFile a local file to remote server!
func (__obf_1571aa912a5dbabc SSHClient) __obf_28b647e76c41ab0e(__obf_39f074c6ed45e115 *sftp.Client, __obf_fe644580f076be82 string, __obf_09092dcf1a0b6e82 string) (__obf_0b884480f1bc4644 error) {

	var __obf_cc2cdce0448b23fd *os.File
	__obf_cc2cdce0448b23fd, __obf_0b884480f1bc4644 = os.Open(__obf_fe644580f076be82)
	if __obf_0b884480f1bc4644 != nil {
		return
	}
	defer __obf_cc2cdce0448b23fd.Close()

	var __obf_ae71ab6fb9c133f2 *sftp.File
	__obf_419cbbb2d9330362 := path.Join(__obf_09092dcf1a0b6e82, filepath.Base(__obf_fe644580f076be82))
	__obf_ae71ab6fb9c133f2, __obf_0b884480f1bc4644 = __obf_39f074c6ed45e115.Create(__obf_419cbbb2d9330362)
	if __obf_0b884480f1bc4644 != nil {
		__obf_0b884480f1bc4644 = fmt.Errorf("[Create: %s] %w", __obf_419cbbb2d9330362, __obf_0b884480f1bc4644)
		return
	}
	defer __obf_ae71ab6fb9c133f2.Close()

	// size := len(s)
	// info, err := local.Stat()
	// if err != nil {
	// 	return err
	// }

	var __obf_357ed3772c7c376e int64
	__obf_357ed3772c7c376e, __obf_0b884480f1bc4644 = io.Copy(__obf_ae71ab6fb9c133f2, io.TeeReader(__obf_cc2cdce0448b23fd, &WriteCounter{}))
	fmt.Fprintf(os.Stdout, "Transferred %.2f kb\n", float64(__obf_357ed3772c7c376e/1024))
	return
}

// Upload a local file or directory to remote server!
func (__obf_1571aa912a5dbabc SSHClient) Upload(__obf_fe644580f076be82 string, __obf_09092dcf1a0b6e82 string) (__obf_0b884480f1bc4644 error) {

	//获取路径的属性
	__obf_320eed44fc4bb6b4, __obf_0b884480f1bc4644 := os.Stat(__obf_fe644580f076be82)
	if __obf_0b884480f1bc4644 != nil {
		return
	}

	__obf_5d117948005c6217, __obf_0b884480f1bc4644 := __obf_1571aa912a5dbabc.NewSftp()
	if __obf_0b884480f1bc4644 != nil {
		return
	}
	defer __obf_5d117948005c6217.Close()

	// 判断是否是文件夹
	if __obf_320eed44fc4bb6b4.IsDir() {
		var __obf_33299d78b73a21e7 []os.DirEntry
		__obf_33299d78b73a21e7, __obf_0b884480f1bc4644 = os.ReadDir(__obf_fe644580f076be82)
		if __obf_0b884480f1bc4644 != nil {
			return
		}
		// 先创建最外层文件夹
		__obf_09092dcf1a0b6e82 = path.Join(__obf_09092dcf1a0b6e82, __obf_320eed44fc4bb6b4.Name())
		__obf_0b884480f1bc4644 = __obf_5d117948005c6217.Mkdir(__obf_09092dcf1a0b6e82)
		if __obf_0b884480f1bc4644 != nil {
			return fmt.Errorf("[Mkdir] %s failed: %w", __obf_09092dcf1a0b6e82, __obf_0b884480f1bc4644)
		}
		// 遍历文件夹内容
		for _, __obf_ed854092583dbd85 := range __obf_33299d78b73a21e7 {
			// 判断是否是文件,是文件直接上传.是文件夹,先远程创建文件夹,再递归复制内部文件
			__obf_0b884480f1bc4644 = __obf_1571aa912a5dbabc.Upload(path.Join(__obf_fe644580f076be82, __obf_ed854092583dbd85.Name()), __obf_09092dcf1a0b6e82)
			if __obf_0b884480f1bc4644 != nil {
				return fmt.Errorf("[Upload to: %s] %w", __obf_09092dcf1a0b6e82, __obf_0b884480f1bc4644)
			}
		}
	} else {
		__obf_0b884480f1bc4644 = __obf_1571aa912a5dbabc.__obf_28b647e76c41ab0e(__obf_5d117948005c6217, __obf_fe644580f076be82, __obf_09092dcf1a0b6e82)
	}
	return
}

// Download file from remote server!
func (__obf_1571aa912a5dbabc SSHClient) Download(__obf_419cbbb2d9330362 string, __obf_fe644580f076be82 string) error {
	var __obf_d474dede28336e63 = path.Join(__obf_fe644580f076be82, filepath.Base(__obf_419cbbb2d9330362))
	_, __obf_0b884480f1bc4644 := os.Stat(__obf_d474dede28336e63)
	if __obf_0b884480f1bc4644 == nil {
		return errors.New("file exists")
	}
	if os.IsNotExist(__obf_0b884480f1bc4644) {
		__obf_cc2cdce0448b23fd, __obf_0b884480f1bc4644 := os.Create(__obf_d474dede28336e63)
		if __obf_0b884480f1bc4644 != nil {
			return fmt.Errorf("[Create: %s] %w", __obf_d474dede28336e63, __obf_0b884480f1bc4644)
		}
		defer __obf_cc2cdce0448b23fd.Close()

		__obf_5d117948005c6217, __obf_0b884480f1bc4644 := __obf_1571aa912a5dbabc.NewSftp()
		if __obf_0b884480f1bc4644 != nil {

			return __obf_0b884480f1bc4644
		}
		defer __obf_5d117948005c6217.Close()

		__obf_ae71ab6fb9c133f2, __obf_0b884480f1bc4644 := __obf_5d117948005c6217.Open(__obf_419cbbb2d9330362)
		if __obf_0b884480f1bc4644 != nil {
			return __obf_0b884480f1bc4644
		}
		defer __obf_ae71ab6fb9c133f2.Close()

		if _, __obf_0b884480f1bc4644 = io.Copy(__obf_cc2cdce0448b23fd, __obf_ae71ab6fb9c133f2); __obf_0b884480f1bc4644 != nil {
			return __obf_0b884480f1bc4644
		}

		return __obf_cc2cdce0448b23fd.Sync()
	}
	return __obf_0b884480f1bc4644
}
