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
func NewSSHClient(__obf_bbadf7259b199153, __obf_17d4344eaba5a5c8 string, __obf_21f2d0f8584526df *ssh.ClientConfig) (__obf_8757b8a7a1a8a795 *SSHClient) {
	// c.Client, err = Dial("tcp", addr, port)
	return &SSHClient{
		Addr:   net.JoinHostPort(__obf_bbadf7259b199153, __obf_17d4344eaba5a5c8),
		Config: __obf_21f2d0f8584526df,
	}
}

func (__obf_8757b8a7a1a8a795 *SSHClient) __obf_f9fc43c807450183() error {
	if __obf_8757b8a7a1a8a795.Client == nil {
		var __obf_77d535ea9cfc3da2 error
		__obf_8757b8a7a1a8a795.Client, __obf_77d535ea9cfc3da2 = ssh.Dial("tcp", __obf_8757b8a7a1a8a795.Addr, __obf_8757b8a7a1a8a795.Config)
		if __obf_77d535ea9cfc3da2 != nil {
			return __obf_77d535ea9cfc3da2
		}
	}
	return nil
}

// Run starts a new SSH session and runs the cmd, it returns CombinedOutput and err if any.
func (__obf_8757b8a7a1a8a795 SSHClient) Run(__obf_edb840588dce9159 string) ([]byte, error) {

	var __obf_77d535ea9cfc3da2 = __obf_8757b8a7a1a8a795.__obf_f9fc43c807450183()
	if __obf_77d535ea9cfc3da2 != nil {
		return nil, __obf_77d535ea9cfc3da2
	}

	var __obf_00cd66afb8f30b6f *ssh.Session
	if __obf_00cd66afb8f30b6f, __obf_77d535ea9cfc3da2 = __obf_8757b8a7a1a8a795.NewSession(); __obf_77d535ea9cfc3da2 != nil {
		return nil, __obf_77d535ea9cfc3da2
	}

	defer __obf_00cd66afb8f30b6f.Close()

	return __obf_00cd66afb8f30b6f.CombinedOutput(__obf_edb840588dce9159)
}

// NewSftp returns new sftp client and error if any.
func (__obf_8757b8a7a1a8a795 SSHClient) NewSftp(__obf_928c820ec9accd99 ...sftp.ClientOption) (*sftp.Client, error) {
	if __obf_77d535ea9cfc3da2 := __obf_8757b8a7a1a8a795.__obf_f9fc43c807450183(); __obf_77d535ea9cfc3da2 != nil {
		return nil, __obf_77d535ea9cfc3da2
	}
	return sftp.NewClient(__obf_8757b8a7a1a8a795.Client, __obf_928c820ec9accd99...)
}

// uploadFile a local file to remote server!
func (__obf_8757b8a7a1a8a795 SSHClient) __obf_3cd23c34733d6bc2(__obf_af64505ef9be63b8 *sftp.Client, __obf_66ec884a0c111474 string, __obf_ed1acdf308289a5c string) (__obf_77d535ea9cfc3da2 error) {

	var __obf_1acf5666f6a8f5b7 *os.File
	__obf_1acf5666f6a8f5b7, __obf_77d535ea9cfc3da2 = os.Open(__obf_66ec884a0c111474)
	if __obf_77d535ea9cfc3da2 != nil {
		return
	}
	defer __obf_1acf5666f6a8f5b7.Close()

	var __obf_028d1134ee14a61e *sftp.File
	__obf_8f3b14716f4bee54 := path.Join(__obf_ed1acdf308289a5c, filepath.Base(__obf_66ec884a0c111474))
	__obf_028d1134ee14a61e, __obf_77d535ea9cfc3da2 = __obf_af64505ef9be63b8.Create(__obf_8f3b14716f4bee54)
	if __obf_77d535ea9cfc3da2 != nil {
		__obf_77d535ea9cfc3da2 = fmt.Errorf("[Create: %s] %w", __obf_8f3b14716f4bee54, __obf_77d535ea9cfc3da2)
		return
	}
	defer __obf_028d1134ee14a61e.Close()

	// size := len(s)
	// info, err := local.Stat()
	// if err != nil {
	// 	return err
	// }

	var __obf_28d77b743eab9402 int64
	__obf_28d77b743eab9402, __obf_77d535ea9cfc3da2 = io.Copy(__obf_028d1134ee14a61e, io.TeeReader(__obf_1acf5666f6a8f5b7, &WriteCounter{}))
	fmt.Fprintf(os.Stdout, "Transferred %.2f kb\n", float64(__obf_28d77b743eab9402/1024))
	return
}

// Upload a local file or directory to remote server!
func (__obf_8757b8a7a1a8a795 SSHClient) Upload(__obf_66ec884a0c111474 string, __obf_ed1acdf308289a5c string) (__obf_77d535ea9cfc3da2 error) {

	//获取路径的属性
	__obf_e91dd95e583e9f8f, __obf_77d535ea9cfc3da2 := os.Stat(__obf_66ec884a0c111474)
	if __obf_77d535ea9cfc3da2 != nil {
		return
	}

	__obf_9394eaed9fe9230b, __obf_77d535ea9cfc3da2 := __obf_8757b8a7a1a8a795.NewSftp()
	if __obf_77d535ea9cfc3da2 != nil {
		return
	}
	defer __obf_9394eaed9fe9230b.Close()

	// 判断是否是文件夹
	if __obf_e91dd95e583e9f8f.IsDir() {
		var __obf_f22823e7913c6f49 []os.DirEntry
		__obf_f22823e7913c6f49, __obf_77d535ea9cfc3da2 = os.ReadDir(__obf_66ec884a0c111474)
		if __obf_77d535ea9cfc3da2 != nil {
			return
		}
		// 先创建最外层文件夹
		__obf_ed1acdf308289a5c = path.Join(__obf_ed1acdf308289a5c, __obf_e91dd95e583e9f8f.Name())
		__obf_77d535ea9cfc3da2 = __obf_9394eaed9fe9230b.Mkdir(__obf_ed1acdf308289a5c)
		if __obf_77d535ea9cfc3da2 != nil {
			return fmt.Errorf("[Mkdir] %s failed: %w", __obf_ed1acdf308289a5c, __obf_77d535ea9cfc3da2)
		}
		// 遍历文件夹内容
		for _, __obf_22b1c92d24474256 := range __obf_f22823e7913c6f49 {
			// 判断是否是文件,是文件直接上传.是文件夹,先远程创建文件夹,再递归复制内部文件
			__obf_77d535ea9cfc3da2 = __obf_8757b8a7a1a8a795.Upload(path.Join(__obf_66ec884a0c111474, __obf_22b1c92d24474256.Name()), __obf_ed1acdf308289a5c)
			if __obf_77d535ea9cfc3da2 != nil {
				return fmt.Errorf("[Upload to: %s] %w", __obf_ed1acdf308289a5c, __obf_77d535ea9cfc3da2)
			}
		}
	} else {
		__obf_77d535ea9cfc3da2 = __obf_8757b8a7a1a8a795.__obf_3cd23c34733d6bc2(__obf_9394eaed9fe9230b, __obf_66ec884a0c111474, __obf_ed1acdf308289a5c)
	}
	return
}

// Download file from remote server!
func (__obf_8757b8a7a1a8a795 SSHClient) Download(__obf_8f3b14716f4bee54 string, __obf_66ec884a0c111474 string) error {
	var __obf_53185449d5a76db5 = path.Join(__obf_66ec884a0c111474, filepath.Base(__obf_8f3b14716f4bee54))
	_, __obf_77d535ea9cfc3da2 := os.Stat(__obf_53185449d5a76db5)
	if __obf_77d535ea9cfc3da2 == nil {
		return errors.New("file exists")
	}
	if os.IsNotExist(__obf_77d535ea9cfc3da2) {
		__obf_1acf5666f6a8f5b7, __obf_77d535ea9cfc3da2 := os.Create(__obf_53185449d5a76db5)
		if __obf_77d535ea9cfc3da2 != nil {
			return fmt.Errorf("[Create: %s] %w", __obf_53185449d5a76db5, __obf_77d535ea9cfc3da2)
		}
		defer __obf_1acf5666f6a8f5b7.Close()

		__obf_9394eaed9fe9230b, __obf_77d535ea9cfc3da2 := __obf_8757b8a7a1a8a795.NewSftp()
		if __obf_77d535ea9cfc3da2 != nil {

			return __obf_77d535ea9cfc3da2
		}
		defer __obf_9394eaed9fe9230b.Close()

		__obf_028d1134ee14a61e, __obf_77d535ea9cfc3da2 := __obf_9394eaed9fe9230b.Open(__obf_8f3b14716f4bee54)
		if __obf_77d535ea9cfc3da2 != nil {
			return __obf_77d535ea9cfc3da2
		}
		defer __obf_028d1134ee14a61e.Close()

		if _, __obf_77d535ea9cfc3da2 = io.Copy(__obf_1acf5666f6a8f5b7, __obf_028d1134ee14a61e); __obf_77d535ea9cfc3da2 != nil {
			return __obf_77d535ea9cfc3da2
		}

		return __obf_1acf5666f6a8f5b7.Sync()
	}
	return __obf_77d535ea9cfc3da2
}
