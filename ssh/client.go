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
func NewSSHClient(__obf_369951218214c1ea, __obf_b6906968613850c9 string, __obf_55b88d738cf620b1 *ssh.ClientConfig) (__obf_1512f9d2953c5e07 *SSHClient) {
	// c.Client, err = Dial("tcp", addr, port)
	return &SSHClient{
		Addr:   net.JoinHostPort(__obf_369951218214c1ea, __obf_b6906968613850c9),
		Config: __obf_55b88d738cf620b1,
	}
}

func (__obf_1512f9d2953c5e07 *SSHClient) init() error {
	if __obf_1512f9d2953c5e07.Client == nil {
		var __obf_649f0be20cc0940a error
		__obf_1512f9d2953c5e07.
			Client, __obf_649f0be20cc0940a = ssh.Dial("tcp", __obf_1512f9d2953c5e07.Addr, __obf_1512f9d2953c5e07.Config)
		if __obf_649f0be20cc0940a != nil {
			return __obf_649f0be20cc0940a
		}
	}
	return nil
}

// Run starts a new SSH session and runs the cmd, it returns CombinedOutput and err if some.
func (__obf_1512f9d2953c5e07 SSHClient) Run(__obf_810974c702390dfa string) ([]byte, error) {

	var __obf_649f0be20cc0940a = __obf_1512f9d2953c5e07.init()
	if __obf_649f0be20cc0940a != nil {
		return nil, __obf_649f0be20cc0940a
	}

	var __obf_4ddc34b770e5920c *ssh.Session
	if __obf_4ddc34b770e5920c, __obf_649f0be20cc0940a = __obf_1512f9d2953c5e07.NewSession(); __obf_649f0be20cc0940a != nil {
		return nil, __obf_649f0be20cc0940a
	}

	defer __obf_4ddc34b770e5920c.Close()

	return __obf_4ddc34b770e5920c.CombinedOutput(__obf_810974c702390dfa)
}

// NewSftp returns new sftp client and error if some.
func (__obf_1512f9d2953c5e07 SSHClient) NewSftp(__obf_22973e86d5aca2de ...sftp.ClientOption) (*sftp.Client, error) {
	if __obf_649f0be20cc0940a := __obf_1512f9d2953c5e07.init(); __obf_649f0be20cc0940a != nil {
		return nil, __obf_649f0be20cc0940a
	}
	return sftp.NewClient(__obf_1512f9d2953c5e07.Client, __obf_22973e86d5aca2de...)
}

// uploadFile a local file to remote server!
func (__obf_1512f9d2953c5e07 SSHClient) __obf_54503739cac070cd(__obf_d1d6d070488647d0 *sftp.Client, __obf_42348af07e5e7f7b string, __obf_c6767ec1ae1733c4 string) (__obf_649f0be20cc0940a error) {

	var __obf_103bc072fcf6c676 *os.File
	__obf_103bc072fcf6c676, __obf_649f0be20cc0940a = os.Open(__obf_42348af07e5e7f7b)
	if __obf_649f0be20cc0940a != nil {
		return
	}
	defer __obf_103bc072fcf6c676.Close()

	var __obf_a6114148307b6fb2 *sftp.File
	__obf_eb21b7bb9b2c4504 := path.Join(__obf_c6767ec1ae1733c4, filepath.Base(__obf_42348af07e5e7f7b))
	__obf_a6114148307b6fb2, __obf_649f0be20cc0940a = __obf_d1d6d070488647d0.Create(__obf_eb21b7bb9b2c4504)
	if __obf_649f0be20cc0940a != nil {
		__obf_649f0be20cc0940a = fmt.Errorf("[Create: %s] %w", __obf_eb21b7bb9b2c4504, __obf_649f0be20cc0940a)
		return
	}
	defer __obf_a6114148307b6fb2.Close()

	// size := len(s)
	// info, err := local.Stat()
	// if err != nil {
	// 	return err
	// }

	var __obf_7804310ba06b6026 int64
	__obf_7804310ba06b6026, __obf_649f0be20cc0940a = io.Copy(__obf_a6114148307b6fb2, io.TeeReader(__obf_103bc072fcf6c676, &WriteCounter{}))
	fmt.Fprintf(os.Stdout, "Transferred %.2f kb\n", float64(__obf_7804310ba06b6026/1024))
	return
}

// Upload a local file or directory to remote server!
func (__obf_1512f9d2953c5e07 SSHClient) Upload(__obf_42348af07e5e7f7b string, __obf_c6767ec1ae1733c4 string) (__obf_649f0be20cc0940a error) {
	__obf_da2340aa1a0ef5c0,

		//获取路径的属性
		__obf_649f0be20cc0940a := os.Stat(__obf_42348af07e5e7f7b)
	if __obf_649f0be20cc0940a != nil {
		return
	}
	__obf_8238dc36bf37c8e4, __obf_649f0be20cc0940a := __obf_1512f9d2953c5e07.NewSftp()
	if __obf_649f0be20cc0940a != nil {
		return
	}
	defer __obf_8238dc36bf37c8e4.Close()

	// 判断是否是文件夹
	if __obf_da2340aa1a0ef5c0.IsDir() {
		var __obf_44e2ca82b183f43a []os.DirEntry
		__obf_44e2ca82b183f43a, __obf_649f0be20cc0940a = os.ReadDir(__obf_42348af07e5e7f7b)
		if __obf_649f0be20cc0940a != nil {
			return
		}
		__obf_c6767ec1ae1733c4 = // 先创建最外层文件夹
			path.Join(__obf_c6767ec1ae1733c4, __obf_da2340aa1a0ef5c0.Name())
		__obf_649f0be20cc0940a = __obf_8238dc36bf37c8e4.Mkdir(__obf_c6767ec1ae1733c4)
		if __obf_649f0be20cc0940a != nil {
			return fmt.Errorf("[Mkdir] %s failed: %w", __obf_c6767ec1ae1733c4, __obf_649f0be20cc0940a)
		}
		// 遍历文件夹内容
		for _, __obf_ef9fc665dd86e5e5 := range __obf_44e2ca82b183f43a {
			__obf_649f0be20cc0940a = // 判断是否是文件,是文件直接上传.是文件夹,先远程创建文件夹,再递归复制内部文件
				__obf_1512f9d2953c5e07.Upload(path.Join(__obf_42348af07e5e7f7b, __obf_ef9fc665dd86e5e5.Name()), __obf_c6767ec1ae1733c4)
			if __obf_649f0be20cc0940a != nil {
				return fmt.Errorf("[Upload to: %s] %w", __obf_c6767ec1ae1733c4, __obf_649f0be20cc0940a)
			}
		}
	} else {
		__obf_649f0be20cc0940a = __obf_1512f9d2953c5e07.__obf_54503739cac070cd(__obf_8238dc36bf37c8e4, __obf_42348af07e5e7f7b, __obf_c6767ec1ae1733c4)
	}
	return
}

// Download file from remote server!
func (__obf_1512f9d2953c5e07 SSHClient) Download(__obf_eb21b7bb9b2c4504 string, __obf_42348af07e5e7f7b string) error {
	var __obf_18d861720376d6cb = path.Join(__obf_42348af07e5e7f7b, filepath.Base(__obf_eb21b7bb9b2c4504))
	_, __obf_649f0be20cc0940a := os.Stat(__obf_18d861720376d6cb)
	if __obf_649f0be20cc0940a == nil {
		return errors.New("file exists")
	}
	if os.IsNotExist(__obf_649f0be20cc0940a) {
		__obf_103bc072fcf6c676, __obf_649f0be20cc0940a := os.Create(__obf_18d861720376d6cb)
		if __obf_649f0be20cc0940a != nil {
			return fmt.Errorf("[Create: %s] %w", __obf_18d861720376d6cb, __obf_649f0be20cc0940a)
		}
		defer __obf_103bc072fcf6c676.Close()
		__obf_8238dc36bf37c8e4, __obf_649f0be20cc0940a := __obf_1512f9d2953c5e07.NewSftp()
		if __obf_649f0be20cc0940a != nil {

			return __obf_649f0be20cc0940a
		}
		defer __obf_8238dc36bf37c8e4.Close()
		__obf_a6114148307b6fb2, __obf_649f0be20cc0940a := __obf_8238dc36bf37c8e4.Open(__obf_eb21b7bb9b2c4504)
		if __obf_649f0be20cc0940a != nil {
			return __obf_649f0be20cc0940a
		}
		defer __obf_a6114148307b6fb2.Close()

		if _, __obf_649f0be20cc0940a = io.Copy(__obf_103bc072fcf6c676, __obf_a6114148307b6fb2); __obf_649f0be20cc0940a != nil {
			return __obf_649f0be20cc0940a
		}

		return __obf_103bc072fcf6c676.Sync()
	}
	return __obf_649f0be20cc0940a
}
