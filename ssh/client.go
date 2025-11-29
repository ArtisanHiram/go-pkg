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
func NewSSHClient(__obf_82c9fb3a5fe80324, __obf_f721d8a7327061d3 string, __obf_b7498e79180a6248 *ssh.ClientConfig) (__obf_830b4d8a76aa0c7e *SSHClient) {
	// c.Client, err = Dial("tcp", addr, port)
	return &SSHClient{
		Addr:   net.JoinHostPort(__obf_82c9fb3a5fe80324, __obf_f721d8a7327061d3),
		Config: __obf_b7498e79180a6248,
	}
}

func (__obf_830b4d8a76aa0c7e *SSHClient) init() error {
	if __obf_830b4d8a76aa0c7e.Client == nil {
		var __obf_6b7ad169ad557bb8 error
		__obf_830b4d8a76aa0c7e.
			Client, __obf_6b7ad169ad557bb8 = ssh.Dial("tcp", __obf_830b4d8a76aa0c7e.Addr, __obf_830b4d8a76aa0c7e.Config)
		if __obf_6b7ad169ad557bb8 != nil {
			return __obf_6b7ad169ad557bb8
		}
	}
	return nil
}

// Run starts a new SSH session and runs the cmd, it returns CombinedOutput and err if some.
func (__obf_830b4d8a76aa0c7e SSHClient) Run(__obf_90a8cd7abc28791b string) ([]byte, error) {

	var __obf_6b7ad169ad557bb8 = __obf_830b4d8a76aa0c7e.init()
	if __obf_6b7ad169ad557bb8 != nil {
		return nil, __obf_6b7ad169ad557bb8
	}

	var __obf_418d26baabe0e8ca *ssh.Session
	if __obf_418d26baabe0e8ca, __obf_6b7ad169ad557bb8 = __obf_830b4d8a76aa0c7e.NewSession(); __obf_6b7ad169ad557bb8 != nil {
		return nil, __obf_6b7ad169ad557bb8
	}

	defer __obf_418d26baabe0e8ca.Close()

	return __obf_418d26baabe0e8ca.CombinedOutput(__obf_90a8cd7abc28791b)
}

// NewSftp returns new sftp client and error if some.
func (__obf_830b4d8a76aa0c7e SSHClient) NewSftp(__obf_4de90efe0c010113 ...sftp.ClientOption) (*sftp.Client, error) {
	if __obf_6b7ad169ad557bb8 := __obf_830b4d8a76aa0c7e.init(); __obf_6b7ad169ad557bb8 != nil {
		return nil, __obf_6b7ad169ad557bb8
	}
	return sftp.NewClient(__obf_830b4d8a76aa0c7e.Client, __obf_4de90efe0c010113...)
}

// uploadFile a local file to remote server!
func (__obf_830b4d8a76aa0c7e SSHClient) __obf_12f27a7d701c16fe(__obf_169ed16405c19a4b *sftp.Client, __obf_2f1fbf084548a529 string, __obf_2d3d268cf72d1378 string) (__obf_6b7ad169ad557bb8 error) {

	var __obf_a899886b40471dd6 *os.File
	__obf_a899886b40471dd6, __obf_6b7ad169ad557bb8 = os.Open(__obf_2f1fbf084548a529)
	if __obf_6b7ad169ad557bb8 != nil {
		return
	}
	defer __obf_a899886b40471dd6.Close()

	var __obf_b94a00558fe734ed *sftp.File
	__obf_5cf7867ef681864b := path.Join(__obf_2d3d268cf72d1378, filepath.Base(__obf_2f1fbf084548a529))
	__obf_b94a00558fe734ed, __obf_6b7ad169ad557bb8 = __obf_169ed16405c19a4b.Create(__obf_5cf7867ef681864b)
	if __obf_6b7ad169ad557bb8 != nil {
		__obf_6b7ad169ad557bb8 = fmt.Errorf("[Create: %s] %w", __obf_5cf7867ef681864b, __obf_6b7ad169ad557bb8)
		return
	}
	defer __obf_b94a00558fe734ed.Close()

	// size := len(s)
	// info, err := local.Stat()
	// if err != nil {
	// 	return err
	// }

	var __obf_0db395662a4e8f4e int64
	__obf_0db395662a4e8f4e, __obf_6b7ad169ad557bb8 = io.Copy(__obf_b94a00558fe734ed, io.TeeReader(__obf_a899886b40471dd6, &WriteCounter{}))
	fmt.Fprintf(os.Stdout, "Transferred %.2f kb\n", float64(__obf_0db395662a4e8f4e/1024))
	return
}

// Upload a local file or directory to remote server!
func (__obf_830b4d8a76aa0c7e SSHClient) Upload(__obf_2f1fbf084548a529 string, __obf_2d3d268cf72d1378 string) (__obf_6b7ad169ad557bb8 error) {
	__obf_14fb9973675e56ce,

		//获取路径的属性
		__obf_6b7ad169ad557bb8 := os.Stat(__obf_2f1fbf084548a529)
	if __obf_6b7ad169ad557bb8 != nil {
		return
	}
	__obf_25b61651c212a81e, __obf_6b7ad169ad557bb8 := __obf_830b4d8a76aa0c7e.NewSftp()
	if __obf_6b7ad169ad557bb8 != nil {
		return
	}
	defer __obf_25b61651c212a81e.Close()

	// 判断是否是文件夹
	if __obf_14fb9973675e56ce.IsDir() {
		var __obf_09fb191cde6f2c6a []os.DirEntry
		__obf_09fb191cde6f2c6a, __obf_6b7ad169ad557bb8 = os.ReadDir(__obf_2f1fbf084548a529)
		if __obf_6b7ad169ad557bb8 != nil {
			return
		}
		__obf_2d3d268cf72d1378 = // 先创建最外层文件夹
			path.Join(__obf_2d3d268cf72d1378, __obf_14fb9973675e56ce.Name())
		__obf_6b7ad169ad557bb8 = __obf_25b61651c212a81e.Mkdir(__obf_2d3d268cf72d1378)
		if __obf_6b7ad169ad557bb8 != nil {
			return fmt.Errorf("[Mkdir] %s failed: %w", __obf_2d3d268cf72d1378, __obf_6b7ad169ad557bb8)
		}
		// 遍历文件夹内容
		for _, __obf_f4ebdeec3b395b0f := range __obf_09fb191cde6f2c6a {
			__obf_6b7ad169ad557bb8 = // 判断是否是文件,是文件直接上传.是文件夹,先远程创建文件夹,再递归复制内部文件
				__obf_830b4d8a76aa0c7e.Upload(path.Join(__obf_2f1fbf084548a529, __obf_f4ebdeec3b395b0f.Name()), __obf_2d3d268cf72d1378)
			if __obf_6b7ad169ad557bb8 != nil {
				return fmt.Errorf("[Upload to: %s] %w", __obf_2d3d268cf72d1378, __obf_6b7ad169ad557bb8)
			}
		}
	} else {
		__obf_6b7ad169ad557bb8 = __obf_830b4d8a76aa0c7e.__obf_12f27a7d701c16fe(__obf_25b61651c212a81e, __obf_2f1fbf084548a529, __obf_2d3d268cf72d1378)
	}
	return
}

// Download file from remote server!
func (__obf_830b4d8a76aa0c7e SSHClient) Download(__obf_5cf7867ef681864b string, __obf_2f1fbf084548a529 string) error {
	var __obf_7685556c51556aed = path.Join(__obf_2f1fbf084548a529, filepath.Base(__obf_5cf7867ef681864b))
	_, __obf_6b7ad169ad557bb8 := os.Stat(__obf_7685556c51556aed)
	if __obf_6b7ad169ad557bb8 == nil {
		return errors.New("file exists")
	}
	if os.IsNotExist(__obf_6b7ad169ad557bb8) {
		__obf_a899886b40471dd6, __obf_6b7ad169ad557bb8 := os.Create(__obf_7685556c51556aed)
		if __obf_6b7ad169ad557bb8 != nil {
			return fmt.Errorf("[Create: %s] %w", __obf_7685556c51556aed, __obf_6b7ad169ad557bb8)
		}
		defer __obf_a899886b40471dd6.Close()
		__obf_25b61651c212a81e, __obf_6b7ad169ad557bb8 := __obf_830b4d8a76aa0c7e.NewSftp()
		if __obf_6b7ad169ad557bb8 != nil {

			return __obf_6b7ad169ad557bb8
		}
		defer __obf_25b61651c212a81e.Close()
		__obf_b94a00558fe734ed, __obf_6b7ad169ad557bb8 := __obf_25b61651c212a81e.Open(__obf_5cf7867ef681864b)
		if __obf_6b7ad169ad557bb8 != nil {
			return __obf_6b7ad169ad557bb8
		}
		defer __obf_b94a00558fe734ed.Close()

		if _, __obf_6b7ad169ad557bb8 = io.Copy(__obf_a899886b40471dd6, __obf_b94a00558fe734ed); __obf_6b7ad169ad557bb8 != nil {
			return __obf_6b7ad169ad557bb8
		}

		return __obf_a899886b40471dd6.Sync()
	}
	return __obf_6b7ad169ad557bb8
}
