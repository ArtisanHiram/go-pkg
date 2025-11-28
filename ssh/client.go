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
func NewSSHClient(__obf_b3031bad8c57f5d5, __obf_21ae3977bac5d238 string, __obf_411b07ac9f15cb2c *ssh.ClientConfig) (__obf_635a2681bc3a8dac *SSHClient) {
	// c.Client, err = Dial("tcp", addr, port)
	return &SSHClient{
		Addr:   net.JoinHostPort(__obf_b3031bad8c57f5d5, __obf_21ae3977bac5d238),
		Config: __obf_411b07ac9f15cb2c,
	}
}

func (__obf_635a2681bc3a8dac *SSHClient) __obf_2eba014ac8293f22() error {
	if __obf_635a2681bc3a8dac.Client == nil {
		var __obf_2c7c784faaf5a58c error
		__obf_635a2681bc3a8dac.Client, __obf_2c7c784faaf5a58c = ssh.Dial("tcp", __obf_635a2681bc3a8dac.Addr, __obf_635a2681bc3a8dac.Config)
		if __obf_2c7c784faaf5a58c != nil {
			return __obf_2c7c784faaf5a58c
		}
	}
	return nil
}

// Run starts a new SSH session and runs the cmd, it returns CombinedOutput and err if any.
func (__obf_635a2681bc3a8dac SSHClient) Run(__obf_9330ab0dcb55965a string) ([]byte, error) {

	var __obf_2c7c784faaf5a58c = __obf_635a2681bc3a8dac.__obf_2eba014ac8293f22()
	if __obf_2c7c784faaf5a58c != nil {
		return nil, __obf_2c7c784faaf5a58c
	}

	var __obf_4e86362180a88afd *ssh.Session
	if __obf_4e86362180a88afd, __obf_2c7c784faaf5a58c = __obf_635a2681bc3a8dac.NewSession(); __obf_2c7c784faaf5a58c != nil {
		return nil, __obf_2c7c784faaf5a58c
	}

	defer __obf_4e86362180a88afd.Close()

	return __obf_4e86362180a88afd.CombinedOutput(__obf_9330ab0dcb55965a)
}

// NewSftp returns new sftp client and error if any.
func (__obf_635a2681bc3a8dac SSHClient) NewSftp(__obf_a779ef4e3d4fb077 ...sftp.ClientOption) (*sftp.Client, error) {
	if __obf_2c7c784faaf5a58c := __obf_635a2681bc3a8dac.__obf_2eba014ac8293f22(); __obf_2c7c784faaf5a58c != nil {
		return nil, __obf_2c7c784faaf5a58c
	}
	return sftp.NewClient(__obf_635a2681bc3a8dac.Client, __obf_a779ef4e3d4fb077...)
}

// uploadFile a local file to remote server!
func (__obf_635a2681bc3a8dac SSHClient) __obf_62eff2eb5f3705d6(__obf_1953d492e34a4d67 *sftp.Client, __obf_212ac50503157386 string, __obf_9be3283a2b485fb9 string) (__obf_2c7c784faaf5a58c error) {

	var __obf_07a074c0d57493fd *os.File
	__obf_07a074c0d57493fd, __obf_2c7c784faaf5a58c = os.Open(__obf_212ac50503157386)
	if __obf_2c7c784faaf5a58c != nil {
		return
	}
	defer __obf_07a074c0d57493fd.Close()

	var __obf_01edc81f547f2229 *sftp.File
	__obf_ef6c4ae5003eeede := path.Join(__obf_9be3283a2b485fb9, filepath.Base(__obf_212ac50503157386))
	__obf_01edc81f547f2229, __obf_2c7c784faaf5a58c = __obf_1953d492e34a4d67.Create(__obf_ef6c4ae5003eeede)
	if __obf_2c7c784faaf5a58c != nil {
		__obf_2c7c784faaf5a58c = fmt.Errorf("[Create: %s] %w", __obf_ef6c4ae5003eeede, __obf_2c7c784faaf5a58c)
		return
	}
	defer __obf_01edc81f547f2229.Close()

	// size := len(s)
	// info, err := local.Stat()
	// if err != nil {
	// 	return err
	// }

	var __obf_271904f038172988 int64
	__obf_271904f038172988, __obf_2c7c784faaf5a58c = io.Copy(__obf_01edc81f547f2229, io.TeeReader(__obf_07a074c0d57493fd, &WriteCounter{}))
	fmt.Fprintf(os.Stdout, "Transferred %.2f kb\n", float64(__obf_271904f038172988/1024))
	return
}

// Upload a local file or directory to remote server!
func (__obf_635a2681bc3a8dac SSHClient) Upload(__obf_212ac50503157386 string, __obf_9be3283a2b485fb9 string) (__obf_2c7c784faaf5a58c error) {

	//获取路径的属性
	__obf_9afb9d412bd83c33, __obf_2c7c784faaf5a58c := os.Stat(__obf_212ac50503157386)
	if __obf_2c7c784faaf5a58c != nil {
		return
	}

	__obf_b1fe6368b8c9b7b8, __obf_2c7c784faaf5a58c := __obf_635a2681bc3a8dac.NewSftp()
	if __obf_2c7c784faaf5a58c != nil {
		return
	}
	defer __obf_b1fe6368b8c9b7b8.Close()

	// 判断是否是文件夹
	if __obf_9afb9d412bd83c33.IsDir() {
		var __obf_277de974815fd550 []os.DirEntry
		__obf_277de974815fd550, __obf_2c7c784faaf5a58c = os.ReadDir(__obf_212ac50503157386)
		if __obf_2c7c784faaf5a58c != nil {
			return
		}
		// 先创建最外层文件夹
		__obf_9be3283a2b485fb9 = path.Join(__obf_9be3283a2b485fb9, __obf_9afb9d412bd83c33.Name())
		__obf_2c7c784faaf5a58c = __obf_b1fe6368b8c9b7b8.Mkdir(__obf_9be3283a2b485fb9)
		if __obf_2c7c784faaf5a58c != nil {
			return fmt.Errorf("[Mkdir] %s failed: %w", __obf_9be3283a2b485fb9, __obf_2c7c784faaf5a58c)
		}
		// 遍历文件夹内容
		for _, __obf_ed13f7b970f49808 := range __obf_277de974815fd550 {
			// 判断是否是文件,是文件直接上传.是文件夹,先远程创建文件夹,再递归复制内部文件
			__obf_2c7c784faaf5a58c = __obf_635a2681bc3a8dac.Upload(path.Join(__obf_212ac50503157386, __obf_ed13f7b970f49808.Name()), __obf_9be3283a2b485fb9)
			if __obf_2c7c784faaf5a58c != nil {
				return fmt.Errorf("[Upload to: %s] %w", __obf_9be3283a2b485fb9, __obf_2c7c784faaf5a58c)
			}
		}
	} else {
		__obf_2c7c784faaf5a58c = __obf_635a2681bc3a8dac.__obf_62eff2eb5f3705d6(__obf_b1fe6368b8c9b7b8, __obf_212ac50503157386, __obf_9be3283a2b485fb9)
	}
	return
}

// Download file from remote server!
func (__obf_635a2681bc3a8dac SSHClient) Download(__obf_ef6c4ae5003eeede string, __obf_212ac50503157386 string) error {
	var __obf_630b9ea7a044617e = path.Join(__obf_212ac50503157386, filepath.Base(__obf_ef6c4ae5003eeede))
	_, __obf_2c7c784faaf5a58c := os.Stat(__obf_630b9ea7a044617e)
	if __obf_2c7c784faaf5a58c == nil {
		return errors.New("file exists")
	}
	if os.IsNotExist(__obf_2c7c784faaf5a58c) {
		__obf_07a074c0d57493fd, __obf_2c7c784faaf5a58c := os.Create(__obf_630b9ea7a044617e)
		if __obf_2c7c784faaf5a58c != nil {
			return fmt.Errorf("[Create: %s] %w", __obf_630b9ea7a044617e, __obf_2c7c784faaf5a58c)
		}
		defer __obf_07a074c0d57493fd.Close()

		__obf_b1fe6368b8c9b7b8, __obf_2c7c784faaf5a58c := __obf_635a2681bc3a8dac.NewSftp()
		if __obf_2c7c784faaf5a58c != nil {

			return __obf_2c7c784faaf5a58c
		}
		defer __obf_b1fe6368b8c9b7b8.Close()

		__obf_01edc81f547f2229, __obf_2c7c784faaf5a58c := __obf_b1fe6368b8c9b7b8.Open(__obf_ef6c4ae5003eeede)
		if __obf_2c7c784faaf5a58c != nil {
			return __obf_2c7c784faaf5a58c
		}
		defer __obf_01edc81f547f2229.Close()

		if _, __obf_2c7c784faaf5a58c = io.Copy(__obf_07a074c0d57493fd, __obf_01edc81f547f2229); __obf_2c7c784faaf5a58c != nil {
			return __obf_2c7c784faaf5a58c
		}

		return __obf_07a074c0d57493fd.Sync()
	}
	return __obf_2c7c784faaf5a58c
}
