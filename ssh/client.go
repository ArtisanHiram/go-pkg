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
func NewSSHClient(__obf_b3607560ce9343d4, __obf_a4ecc6ee5badcb4b string, __obf_4f67c5a3c4f89a6d *ssh.ClientConfig) (__obf_4b67250ac9e5dda6 *SSHClient) {
	// c.Client, err = Dial("tcp", addr, port)
	return &SSHClient{
		Addr:   net.JoinHostPort(__obf_b3607560ce9343d4, __obf_a4ecc6ee5badcb4b),
		Config: __obf_4f67c5a3c4f89a6d,
	}
}

func (__obf_4b67250ac9e5dda6 *SSHClient) __obf_d6b07f67f2ecc33a() error {
	if __obf_4b67250ac9e5dda6.Client == nil {
		var __obf_dbce9f361de2d092 error
		__obf_4b67250ac9e5dda6.Client, __obf_dbce9f361de2d092 = ssh.Dial("tcp", __obf_4b67250ac9e5dda6.Addr, __obf_4b67250ac9e5dda6.Config)
		if __obf_dbce9f361de2d092 != nil {
			return __obf_dbce9f361de2d092
		}
	}
	return nil
}

// Run starts a new SSH session and runs the cmd, it returns CombinedOutput and err if any.
func (__obf_4b67250ac9e5dda6 SSHClient) Run(__obf_a308d496cbac4f35 string) ([]byte, error) {

	var __obf_dbce9f361de2d092 = __obf_4b67250ac9e5dda6.__obf_d6b07f67f2ecc33a()
	if __obf_dbce9f361de2d092 != nil {
		return nil, __obf_dbce9f361de2d092
	}

	var __obf_2ed9604fe1c5d441 *ssh.Session
	if __obf_2ed9604fe1c5d441, __obf_dbce9f361de2d092 = __obf_4b67250ac9e5dda6.NewSession(); __obf_dbce9f361de2d092 != nil {
		return nil, __obf_dbce9f361de2d092
	}

	defer __obf_2ed9604fe1c5d441.Close()

	return __obf_2ed9604fe1c5d441.CombinedOutput(__obf_a308d496cbac4f35)
}

// NewSftp returns new sftp client and error if any.
func (__obf_4b67250ac9e5dda6 SSHClient) NewSftp(__obf_193ad0890686bdfe ...sftp.ClientOption) (*sftp.Client, error) {
	if __obf_dbce9f361de2d092 := __obf_4b67250ac9e5dda6.__obf_d6b07f67f2ecc33a(); __obf_dbce9f361de2d092 != nil {
		return nil, __obf_dbce9f361de2d092
	}
	return sftp.NewClient(__obf_4b67250ac9e5dda6.Client, __obf_193ad0890686bdfe...)
}

// uploadFile a local file to remote server!
func (__obf_4b67250ac9e5dda6 SSHClient) __obf_b7c0145aa4f489ea(__obf_98d82f30e785c8bf *sftp.Client, __obf_3b0d2dd9e8dbde33 string, __obf_6a229d6d1d48c255 string) (__obf_dbce9f361de2d092 error) {

	var __obf_b1299453282c6a32 *os.File
	__obf_b1299453282c6a32, __obf_dbce9f361de2d092 = os.Open(__obf_3b0d2dd9e8dbde33)
	if __obf_dbce9f361de2d092 != nil {
		return
	}
	defer __obf_b1299453282c6a32.Close()

	var __obf_22f1a0bdee77e6e9 *sftp.File
	__obf_656dd87ea085b0bc := path.Join(__obf_6a229d6d1d48c255, filepath.Base(__obf_3b0d2dd9e8dbde33))
	__obf_22f1a0bdee77e6e9, __obf_dbce9f361de2d092 = __obf_98d82f30e785c8bf.Create(__obf_656dd87ea085b0bc)
	if __obf_dbce9f361de2d092 != nil {
		__obf_dbce9f361de2d092 = fmt.Errorf("[Create: %s] %w", __obf_656dd87ea085b0bc, __obf_dbce9f361de2d092)
		return
	}
	defer __obf_22f1a0bdee77e6e9.Close()

	// size := len(s)
	// info, err := local.Stat()
	// if err != nil {
	// 	return err
	// }

	var __obf_5cbf9210cbf29393 int64
	__obf_5cbf9210cbf29393, __obf_dbce9f361de2d092 = io.Copy(__obf_22f1a0bdee77e6e9, io.TeeReader(__obf_b1299453282c6a32, &WriteCounter{}))
	fmt.Fprintf(os.Stdout, "Transferred %.2f kb\n", float64(__obf_5cbf9210cbf29393/1024))
	return
}

// Upload a local file or directory to remote server!
func (__obf_4b67250ac9e5dda6 SSHClient) Upload(__obf_3b0d2dd9e8dbde33 string, __obf_6a229d6d1d48c255 string) (__obf_dbce9f361de2d092 error) {

	//获取路径的属性
	__obf_02847525c512d1d8, __obf_dbce9f361de2d092 := os.Stat(__obf_3b0d2dd9e8dbde33)
	if __obf_dbce9f361de2d092 != nil {
		return
	}

	__obf_3edaa0fdcfd8ff4f, __obf_dbce9f361de2d092 := __obf_4b67250ac9e5dda6.NewSftp()
	if __obf_dbce9f361de2d092 != nil {
		return
	}
	defer __obf_3edaa0fdcfd8ff4f.Close()

	// 判断是否是文件夹
	if __obf_02847525c512d1d8.IsDir() {
		var __obf_8973a49b1eeff625 []os.DirEntry
		__obf_8973a49b1eeff625, __obf_dbce9f361de2d092 = os.ReadDir(__obf_3b0d2dd9e8dbde33)
		if __obf_dbce9f361de2d092 != nil {
			return
		}
		// 先创建最外层文件夹
		__obf_6a229d6d1d48c255 = path.Join(__obf_6a229d6d1d48c255, __obf_02847525c512d1d8.Name())
		__obf_dbce9f361de2d092 = __obf_3edaa0fdcfd8ff4f.Mkdir(__obf_6a229d6d1d48c255)
		if __obf_dbce9f361de2d092 != nil {
			return fmt.Errorf("[Mkdir] %s failed: %w", __obf_6a229d6d1d48c255, __obf_dbce9f361de2d092)
		}
		// 遍历文件夹内容
		for _, __obf_386df5a923df47f3 := range __obf_8973a49b1eeff625 {
			// 判断是否是文件,是文件直接上传.是文件夹,先远程创建文件夹,再递归复制内部文件
			__obf_dbce9f361de2d092 = __obf_4b67250ac9e5dda6.Upload(path.Join(__obf_3b0d2dd9e8dbde33, __obf_386df5a923df47f3.Name()), __obf_6a229d6d1d48c255)
			if __obf_dbce9f361de2d092 != nil {
				return fmt.Errorf("[Upload to: %s] %w", __obf_6a229d6d1d48c255, __obf_dbce9f361de2d092)
			}
		}
	} else {
		__obf_dbce9f361de2d092 = __obf_4b67250ac9e5dda6.__obf_b7c0145aa4f489ea(__obf_3edaa0fdcfd8ff4f, __obf_3b0d2dd9e8dbde33, __obf_6a229d6d1d48c255)
	}
	return
}

// Download file from remote server!
func (__obf_4b67250ac9e5dda6 SSHClient) Download(__obf_656dd87ea085b0bc string, __obf_3b0d2dd9e8dbde33 string) error {
	var __obf_54a63f254f6e38c2 = path.Join(__obf_3b0d2dd9e8dbde33, filepath.Base(__obf_656dd87ea085b0bc))
	_, __obf_dbce9f361de2d092 := os.Stat(__obf_54a63f254f6e38c2)
	if __obf_dbce9f361de2d092 == nil {
		return errors.New("file exists")
	}
	if os.IsNotExist(__obf_dbce9f361de2d092) {
		__obf_b1299453282c6a32, __obf_dbce9f361de2d092 := os.Create(__obf_54a63f254f6e38c2)
		if __obf_dbce9f361de2d092 != nil {
			return fmt.Errorf("[Create: %s] %w", __obf_54a63f254f6e38c2, __obf_dbce9f361de2d092)
		}
		defer __obf_b1299453282c6a32.Close()

		__obf_3edaa0fdcfd8ff4f, __obf_dbce9f361de2d092 := __obf_4b67250ac9e5dda6.NewSftp()
		if __obf_dbce9f361de2d092 != nil {

			return __obf_dbce9f361de2d092
		}
		defer __obf_3edaa0fdcfd8ff4f.Close()

		__obf_22f1a0bdee77e6e9, __obf_dbce9f361de2d092 := __obf_3edaa0fdcfd8ff4f.Open(__obf_656dd87ea085b0bc)
		if __obf_dbce9f361de2d092 != nil {
			return __obf_dbce9f361de2d092
		}
		defer __obf_22f1a0bdee77e6e9.Close()

		if _, __obf_dbce9f361de2d092 = io.Copy(__obf_b1299453282c6a32, __obf_22f1a0bdee77e6e9); __obf_dbce9f361de2d092 != nil {
			return __obf_dbce9f361de2d092
		}

		return __obf_b1299453282c6a32.Sync()
	}
	return __obf_dbce9f361de2d092
}
