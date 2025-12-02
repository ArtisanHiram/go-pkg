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
func NewSSHClient(__obf_09e44d27170eae26, __obf_a9e00bcaa820b9d1 string, __obf_982850f5d32940f7 *ssh.ClientConfig) (__obf_1ad498ac8051ef26 *SSHClient) {
	// c.Client, err = Dial("tcp", addr, port)
	return &SSHClient{
		Addr:   net.JoinHostPort(__obf_09e44d27170eae26, __obf_a9e00bcaa820b9d1),
		Config: __obf_982850f5d32940f7,
	}
}

func (__obf_1ad498ac8051ef26 *SSHClient) init() error {
	if __obf_1ad498ac8051ef26.Client == nil {
		var __obf_3e865493244170fc error
		__obf_1ad498ac8051ef26.
			Client, __obf_3e865493244170fc = ssh.Dial("tcp", __obf_1ad498ac8051ef26.Addr, __obf_1ad498ac8051ef26.Config)
		if __obf_3e865493244170fc != nil {
			return __obf_3e865493244170fc
		}
	}
	return nil
}

// Run starts a new SSH session and runs the cmd, it returns CombinedOutput and err if some.
func (__obf_1ad498ac8051ef26 SSHClient) Run(__obf_2e76b9ea7b526237 string) ([]byte, error) {

	var __obf_3e865493244170fc = __obf_1ad498ac8051ef26.init()
	if __obf_3e865493244170fc != nil {
		return nil, __obf_3e865493244170fc
	}

	var __obf_1934ca0c6afddf72 *ssh.Session
	if __obf_1934ca0c6afddf72, __obf_3e865493244170fc = __obf_1ad498ac8051ef26.NewSession(); __obf_3e865493244170fc != nil {
		return nil, __obf_3e865493244170fc
	}

	defer __obf_1934ca0c6afddf72.Close()

	return __obf_1934ca0c6afddf72.CombinedOutput(__obf_2e76b9ea7b526237)
}

// NewSftp returns new sftp client and error if some.
func (__obf_1ad498ac8051ef26 SSHClient) NewSftp(__obf_68e0f329675ee6ac ...sftp.ClientOption) (*sftp.Client, error) {
	if __obf_3e865493244170fc := __obf_1ad498ac8051ef26.init(); __obf_3e865493244170fc != nil {
		return nil, __obf_3e865493244170fc
	}
	return sftp.NewClient(__obf_1ad498ac8051ef26.Client, __obf_68e0f329675ee6ac...)
}

// uploadFile a local file to remote server!
func (__obf_1ad498ac8051ef26 SSHClient) __obf_609b36b30d350f3b(__obf_55d2a19184100dc6 *sftp.Client, __obf_d130ec92ef4dfb40 string, __obf_ff4ce6d3cd7446ff string) (__obf_3e865493244170fc error) {

	var __obf_a9f32a6f9726a934 *os.File
	__obf_a9f32a6f9726a934, __obf_3e865493244170fc = os.Open(__obf_d130ec92ef4dfb40)
	if __obf_3e865493244170fc != nil {
		return
	}
	defer __obf_a9f32a6f9726a934.Close()

	var __obf_e57dcd828c002c43 *sftp.File
	__obf_ffa8e066b2a595fb := path.Join(__obf_ff4ce6d3cd7446ff, filepath.Base(__obf_d130ec92ef4dfb40))
	__obf_e57dcd828c002c43, __obf_3e865493244170fc = __obf_55d2a19184100dc6.Create(__obf_ffa8e066b2a595fb)
	if __obf_3e865493244170fc != nil {
		__obf_3e865493244170fc = fmt.Errorf("[Create: %s] %w", __obf_ffa8e066b2a595fb, __obf_3e865493244170fc)
		return
	}
	defer __obf_e57dcd828c002c43.Close()

	// size := len(s)
	// info, err := local.Stat()
	// if err != nil {
	// 	return err
	// }

	var __obf_b8799a85ce8bc42f int64
	__obf_b8799a85ce8bc42f, __obf_3e865493244170fc = io.Copy(__obf_e57dcd828c002c43, io.TeeReader(__obf_a9f32a6f9726a934, &WriteCounter{}))
	fmt.Fprintf(os.Stdout, "Transferred %.2f kb\n", float64(__obf_b8799a85ce8bc42f/1024))
	return
}

// Upload a local file or directory to remote server!
func (__obf_1ad498ac8051ef26 SSHClient) Upload(__obf_d130ec92ef4dfb40 string, __obf_ff4ce6d3cd7446ff string) (__obf_3e865493244170fc error) {
	__obf_7e7beabfb01b5be5,

		//获取路径的属性
		__obf_3e865493244170fc := os.Stat(__obf_d130ec92ef4dfb40)
	if __obf_3e865493244170fc != nil {
		return
	}
	__obf_d4c39326fed9d618, __obf_3e865493244170fc := __obf_1ad498ac8051ef26.NewSftp()
	if __obf_3e865493244170fc != nil {
		return
	}
	defer __obf_d4c39326fed9d618.Close()

	// 判断是否是文件夹
	if __obf_7e7beabfb01b5be5.IsDir() {
		var __obf_a99ffb42361c62ae []os.DirEntry
		__obf_a99ffb42361c62ae, __obf_3e865493244170fc = os.ReadDir(__obf_d130ec92ef4dfb40)
		if __obf_3e865493244170fc != nil {
			return
		}
		__obf_ff4ce6d3cd7446ff = // 先创建最外层文件夹
			path.Join(__obf_ff4ce6d3cd7446ff, __obf_7e7beabfb01b5be5.Name())
		__obf_3e865493244170fc = __obf_d4c39326fed9d618.Mkdir(__obf_ff4ce6d3cd7446ff)
		if __obf_3e865493244170fc != nil {
			return fmt.Errorf("[Mkdir] %s failed: %w", __obf_ff4ce6d3cd7446ff, __obf_3e865493244170fc)
		}
		// 遍历文件夹内容
		for _, __obf_2514831d5cb1179b := range __obf_a99ffb42361c62ae {
			__obf_3e865493244170fc = // 判断是否是文件,是文件直接上传.是文件夹,先远程创建文件夹,再递归复制内部文件
				__obf_1ad498ac8051ef26.Upload(path.Join(__obf_d130ec92ef4dfb40, __obf_2514831d5cb1179b.Name()), __obf_ff4ce6d3cd7446ff)
			if __obf_3e865493244170fc != nil {
				return fmt.Errorf("[Upload to: %s] %w", __obf_ff4ce6d3cd7446ff, __obf_3e865493244170fc)
			}
		}
	} else {
		__obf_3e865493244170fc = __obf_1ad498ac8051ef26.__obf_609b36b30d350f3b(__obf_d4c39326fed9d618, __obf_d130ec92ef4dfb40, __obf_ff4ce6d3cd7446ff)
	}
	return
}

// Download file from remote server!
func (__obf_1ad498ac8051ef26 SSHClient) Download(__obf_ffa8e066b2a595fb string, __obf_d130ec92ef4dfb40 string) error {
	var __obf_b9dafa2cc702aead = path.Join(__obf_d130ec92ef4dfb40, filepath.Base(__obf_ffa8e066b2a595fb))
	_, __obf_3e865493244170fc := os.Stat(__obf_b9dafa2cc702aead)
	if __obf_3e865493244170fc == nil {
		return errors.New("file exists")
	}
	if os.IsNotExist(__obf_3e865493244170fc) {
		__obf_a9f32a6f9726a934, __obf_3e865493244170fc := os.Create(__obf_b9dafa2cc702aead)
		if __obf_3e865493244170fc != nil {
			return fmt.Errorf("[Create: %s] %w", __obf_b9dafa2cc702aead, __obf_3e865493244170fc)
		}
		defer __obf_a9f32a6f9726a934.Close()
		__obf_d4c39326fed9d618, __obf_3e865493244170fc := __obf_1ad498ac8051ef26.NewSftp()
		if __obf_3e865493244170fc != nil {

			return __obf_3e865493244170fc
		}
		defer __obf_d4c39326fed9d618.Close()
		__obf_e57dcd828c002c43, __obf_3e865493244170fc := __obf_d4c39326fed9d618.Open(__obf_ffa8e066b2a595fb)
		if __obf_3e865493244170fc != nil {
			return __obf_3e865493244170fc
		}
		defer __obf_e57dcd828c002c43.Close()

		if _, __obf_3e865493244170fc = io.Copy(__obf_a9f32a6f9726a934, __obf_e57dcd828c002c43); __obf_3e865493244170fc != nil {
			return __obf_3e865493244170fc
		}

		return __obf_a9f32a6f9726a934.Sync()
	}
	return __obf_3e865493244170fc
}
