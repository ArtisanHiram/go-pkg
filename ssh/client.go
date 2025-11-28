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
func NewSSHClient(__obf_fc341c76ad6d1fe3, __obf_1732d69a1d131857 string, __obf_73b9014009786972 *ssh.ClientConfig) (__obf_f6f05ebda427eff8 *SSHClient) {
	// c.Client, err = Dial("tcp", addr, port)
	return &SSHClient{
		Addr:   net.JoinHostPort(__obf_fc341c76ad6d1fe3, __obf_1732d69a1d131857),
		Config: __obf_73b9014009786972,
	}
}

func (__obf_f6f05ebda427eff8 *SSHClient) __obf_20833a339d6814f3() error {
	if __obf_f6f05ebda427eff8.Client == nil {
		var __obf_aa4fb85c75454390 error
		__obf_f6f05ebda427eff8.Client, __obf_aa4fb85c75454390 = ssh.Dial("tcp", __obf_f6f05ebda427eff8.Addr, __obf_f6f05ebda427eff8.Config)
		if __obf_aa4fb85c75454390 != nil {
			return __obf_aa4fb85c75454390
		}
	}
	return nil
}

// Run starts a new SSH session and runs the cmd, it returns CombinedOutput and err if any.
func (__obf_f6f05ebda427eff8 SSHClient) Run(__obf_cfebcb1d85ece2a2 string) ([]byte, error) {

	var __obf_aa4fb85c75454390 = __obf_f6f05ebda427eff8.__obf_20833a339d6814f3()
	if __obf_aa4fb85c75454390 != nil {
		return nil, __obf_aa4fb85c75454390
	}

	var __obf_6a0d03e04c483a66 *ssh.Session
	if __obf_6a0d03e04c483a66, __obf_aa4fb85c75454390 = __obf_f6f05ebda427eff8.NewSession(); __obf_aa4fb85c75454390 != nil {
		return nil, __obf_aa4fb85c75454390
	}

	defer __obf_6a0d03e04c483a66.Close()

	return __obf_6a0d03e04c483a66.CombinedOutput(__obf_cfebcb1d85ece2a2)
}

// NewSftp returns new sftp client and error if any.
func (__obf_f6f05ebda427eff8 SSHClient) NewSftp(__obf_cf3710f7d8b8e96c ...sftp.ClientOption) (*sftp.Client, error) {
	if __obf_aa4fb85c75454390 := __obf_f6f05ebda427eff8.__obf_20833a339d6814f3(); __obf_aa4fb85c75454390 != nil {
		return nil, __obf_aa4fb85c75454390
	}
	return sftp.NewClient(__obf_f6f05ebda427eff8.Client, __obf_cf3710f7d8b8e96c...)
}

// uploadFile a local file to remote server!
func (__obf_f6f05ebda427eff8 SSHClient) __obf_8e3d96ea7f118307(__obf_74e0bab4d4e5f8e3 *sftp.Client, __obf_372b48c4b67f060a string, __obf_4dfc89523121de0c string) (__obf_aa4fb85c75454390 error) {

	var __obf_13e932f137283837 *os.File
	__obf_13e932f137283837, __obf_aa4fb85c75454390 = os.Open(__obf_372b48c4b67f060a)
	if __obf_aa4fb85c75454390 != nil {
		return
	}
	defer __obf_13e932f137283837.Close()

	var __obf_19e636e44a7a89d8 *sftp.File
	__obf_e3528e9556027c15 := path.Join(__obf_4dfc89523121de0c, filepath.Base(__obf_372b48c4b67f060a))
	__obf_19e636e44a7a89d8, __obf_aa4fb85c75454390 = __obf_74e0bab4d4e5f8e3.Create(__obf_e3528e9556027c15)
	if __obf_aa4fb85c75454390 != nil {
		__obf_aa4fb85c75454390 = fmt.Errorf("[Create: %s] %w", __obf_e3528e9556027c15, __obf_aa4fb85c75454390)
		return
	}
	defer __obf_19e636e44a7a89d8.Close()

	// size := len(s)
	// info, err := local.Stat()
	// if err != nil {
	// 	return err
	// }

	var __obf_22f5b916eb86c4eb int64
	__obf_22f5b916eb86c4eb, __obf_aa4fb85c75454390 = io.Copy(__obf_19e636e44a7a89d8, io.TeeReader(__obf_13e932f137283837, &WriteCounter{}))
	fmt.Fprintf(os.Stdout, "Transferred %.2f kb\n", float64(__obf_22f5b916eb86c4eb/1024))
	return
}

// Upload a local file or directory to remote server!
func (__obf_f6f05ebda427eff8 SSHClient) Upload(__obf_372b48c4b67f060a string, __obf_4dfc89523121de0c string) (__obf_aa4fb85c75454390 error) {

	//获取路径的属性
	__obf_7eef5e2ba46b5876, __obf_aa4fb85c75454390 := os.Stat(__obf_372b48c4b67f060a)
	if __obf_aa4fb85c75454390 != nil {
		return
	}

	__obf_e4febb5a30770f64, __obf_aa4fb85c75454390 := __obf_f6f05ebda427eff8.NewSftp()
	if __obf_aa4fb85c75454390 != nil {
		return
	}
	defer __obf_e4febb5a30770f64.Close()

	// 判断是否是文件夹
	if __obf_7eef5e2ba46b5876.IsDir() {
		var __obf_c153390023a03ce9 []os.DirEntry
		__obf_c153390023a03ce9, __obf_aa4fb85c75454390 = os.ReadDir(__obf_372b48c4b67f060a)
		if __obf_aa4fb85c75454390 != nil {
			return
		}
		// 先创建最外层文件夹
		__obf_4dfc89523121de0c = path.Join(__obf_4dfc89523121de0c, __obf_7eef5e2ba46b5876.Name())
		__obf_aa4fb85c75454390 = __obf_e4febb5a30770f64.Mkdir(__obf_4dfc89523121de0c)
		if __obf_aa4fb85c75454390 != nil {
			return fmt.Errorf("[Mkdir] %s failed: %w", __obf_4dfc89523121de0c, __obf_aa4fb85c75454390)
		}
		// 遍历文件夹内容
		for _, __obf_9e7564c013f6df57 := range __obf_c153390023a03ce9 {
			// 判断是否是文件,是文件直接上传.是文件夹,先远程创建文件夹,再递归复制内部文件
			__obf_aa4fb85c75454390 = __obf_f6f05ebda427eff8.Upload(path.Join(__obf_372b48c4b67f060a, __obf_9e7564c013f6df57.Name()), __obf_4dfc89523121de0c)
			if __obf_aa4fb85c75454390 != nil {
				return fmt.Errorf("[Upload to: %s] %w", __obf_4dfc89523121de0c, __obf_aa4fb85c75454390)
			}
		}
	} else {
		__obf_aa4fb85c75454390 = __obf_f6f05ebda427eff8.__obf_8e3d96ea7f118307(__obf_e4febb5a30770f64, __obf_372b48c4b67f060a, __obf_4dfc89523121de0c)
	}
	return
}

// Download file from remote server!
func (__obf_f6f05ebda427eff8 SSHClient) Download(__obf_e3528e9556027c15 string, __obf_372b48c4b67f060a string) error {
	var __obf_98d1c0458ca3ea0c = path.Join(__obf_372b48c4b67f060a, filepath.Base(__obf_e3528e9556027c15))
	_, __obf_aa4fb85c75454390 := os.Stat(__obf_98d1c0458ca3ea0c)
	if __obf_aa4fb85c75454390 == nil {
		return errors.New("file exists")
	}
	if os.IsNotExist(__obf_aa4fb85c75454390) {
		__obf_13e932f137283837, __obf_aa4fb85c75454390 := os.Create(__obf_98d1c0458ca3ea0c)
		if __obf_aa4fb85c75454390 != nil {
			return fmt.Errorf("[Create: %s] %w", __obf_98d1c0458ca3ea0c, __obf_aa4fb85c75454390)
		}
		defer __obf_13e932f137283837.Close()

		__obf_e4febb5a30770f64, __obf_aa4fb85c75454390 := __obf_f6f05ebda427eff8.NewSftp()
		if __obf_aa4fb85c75454390 != nil {

			return __obf_aa4fb85c75454390
		}
		defer __obf_e4febb5a30770f64.Close()

		__obf_19e636e44a7a89d8, __obf_aa4fb85c75454390 := __obf_e4febb5a30770f64.Open(__obf_e3528e9556027c15)
		if __obf_aa4fb85c75454390 != nil {
			return __obf_aa4fb85c75454390
		}
		defer __obf_19e636e44a7a89d8.Close()

		if _, __obf_aa4fb85c75454390 = io.Copy(__obf_13e932f137283837, __obf_19e636e44a7a89d8); __obf_aa4fb85c75454390 != nil {
			return __obf_aa4fb85c75454390
		}

		return __obf_13e932f137283837.Sync()
	}
	return __obf_aa4fb85c75454390
}
