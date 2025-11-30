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
func NewSSHClient(__obf_d116fb544d775b2d, __obf_2f91711c8ec963e3 string, __obf_8141ac645a681b8e *ssh.ClientConfig) (__obf_67195a8151e0c64a *SSHClient) {
	// c.Client, err = Dial("tcp", addr, port)
	return &SSHClient{
		Addr:   net.JoinHostPort(__obf_d116fb544d775b2d, __obf_2f91711c8ec963e3),
		Config: __obf_8141ac645a681b8e,
	}
}

func (__obf_67195a8151e0c64a *SSHClient) init() error {
	if __obf_67195a8151e0c64a.Client == nil {
		var __obf_7f6f4ab0f2c744c3 error
		__obf_67195a8151e0c64a.
			Client, __obf_7f6f4ab0f2c744c3 = ssh.Dial("tcp", __obf_67195a8151e0c64a.Addr, __obf_67195a8151e0c64a.Config)
		if __obf_7f6f4ab0f2c744c3 != nil {
			return __obf_7f6f4ab0f2c744c3
		}
	}
	return nil
}

// Run starts a new SSH session and runs the cmd, it returns CombinedOutput and err if some.
func (__obf_67195a8151e0c64a SSHClient) Run(__obf_540515e864f82561 string) ([]byte, error) {

	var __obf_7f6f4ab0f2c744c3 = __obf_67195a8151e0c64a.init()
	if __obf_7f6f4ab0f2c744c3 != nil {
		return nil, __obf_7f6f4ab0f2c744c3
	}

	var __obf_82c02e30a8448e4a *ssh.Session
	if __obf_82c02e30a8448e4a, __obf_7f6f4ab0f2c744c3 = __obf_67195a8151e0c64a.NewSession(); __obf_7f6f4ab0f2c744c3 != nil {
		return nil, __obf_7f6f4ab0f2c744c3
	}

	defer __obf_82c02e30a8448e4a.Close()

	return __obf_82c02e30a8448e4a.CombinedOutput(__obf_540515e864f82561)
}

// NewSftp returns new sftp client and error if some.
func (__obf_67195a8151e0c64a SSHClient) NewSftp(__obf_079235b9fc5a5a64 ...sftp.ClientOption) (*sftp.Client, error) {
	if __obf_7f6f4ab0f2c744c3 := __obf_67195a8151e0c64a.init(); __obf_7f6f4ab0f2c744c3 != nil {
		return nil, __obf_7f6f4ab0f2c744c3
	}
	return sftp.NewClient(__obf_67195a8151e0c64a.Client, __obf_079235b9fc5a5a64...)
}

// uploadFile a local file to remote server!
func (__obf_67195a8151e0c64a SSHClient) __obf_527bd2433cb23cc2(__obf_dc430c471dd12613 *sftp.Client, __obf_321a70496679c815 string, __obf_941abe6dce7e870c string) (__obf_7f6f4ab0f2c744c3 error) {

	var __obf_e86848930264c52f *os.File
	__obf_e86848930264c52f, __obf_7f6f4ab0f2c744c3 = os.Open(__obf_321a70496679c815)
	if __obf_7f6f4ab0f2c744c3 != nil {
		return
	}
	defer __obf_e86848930264c52f.Close()

	var __obf_83d70bf93c31aff3 *sftp.File
	__obf_7897a1786c6083a1 := path.Join(__obf_941abe6dce7e870c, filepath.Base(__obf_321a70496679c815))
	__obf_83d70bf93c31aff3, __obf_7f6f4ab0f2c744c3 = __obf_dc430c471dd12613.Create(__obf_7897a1786c6083a1)
	if __obf_7f6f4ab0f2c744c3 != nil {
		__obf_7f6f4ab0f2c744c3 = fmt.Errorf("[Create: %s] %w", __obf_7897a1786c6083a1, __obf_7f6f4ab0f2c744c3)
		return
	}
	defer __obf_83d70bf93c31aff3.Close()

	// size := len(s)
	// info, err := local.Stat()
	// if err != nil {
	// 	return err
	// }

	var __obf_72ce6ca7a2237474 int64
	__obf_72ce6ca7a2237474, __obf_7f6f4ab0f2c744c3 = io.Copy(__obf_83d70bf93c31aff3, io.TeeReader(__obf_e86848930264c52f, &WriteCounter{}))
	fmt.Fprintf(os.Stdout, "Transferred %.2f kb\n", float64(__obf_72ce6ca7a2237474/1024))
	return
}

// Upload a local file or directory to remote server!
func (__obf_67195a8151e0c64a SSHClient) Upload(__obf_321a70496679c815 string, __obf_941abe6dce7e870c string) (__obf_7f6f4ab0f2c744c3 error) {
	__obf_6d5c1f9b948f912e,

		//获取路径的属性
		__obf_7f6f4ab0f2c744c3 := os.Stat(__obf_321a70496679c815)
	if __obf_7f6f4ab0f2c744c3 != nil {
		return
	}
	__obf_c44777f527dbbd3d, __obf_7f6f4ab0f2c744c3 := __obf_67195a8151e0c64a.NewSftp()
	if __obf_7f6f4ab0f2c744c3 != nil {
		return
	}
	defer __obf_c44777f527dbbd3d.Close()

	// 判断是否是文件夹
	if __obf_6d5c1f9b948f912e.IsDir() {
		var __obf_5ccbaee70972cde4 []os.DirEntry
		__obf_5ccbaee70972cde4, __obf_7f6f4ab0f2c744c3 = os.ReadDir(__obf_321a70496679c815)
		if __obf_7f6f4ab0f2c744c3 != nil {
			return
		}
		__obf_941abe6dce7e870c = // 先创建最外层文件夹
			path.Join(__obf_941abe6dce7e870c, __obf_6d5c1f9b948f912e.Name())
		__obf_7f6f4ab0f2c744c3 = __obf_c44777f527dbbd3d.Mkdir(__obf_941abe6dce7e870c)
		if __obf_7f6f4ab0f2c744c3 != nil {
			return fmt.Errorf("[Mkdir] %s failed: %w", __obf_941abe6dce7e870c, __obf_7f6f4ab0f2c744c3)
		}
		// 遍历文件夹内容
		for _, __obf_ac067049358b447d := range __obf_5ccbaee70972cde4 {
			__obf_7f6f4ab0f2c744c3 = // 判断是否是文件,是文件直接上传.是文件夹,先远程创建文件夹,再递归复制内部文件
				__obf_67195a8151e0c64a.Upload(path.Join(__obf_321a70496679c815, __obf_ac067049358b447d.Name()), __obf_941abe6dce7e870c)
			if __obf_7f6f4ab0f2c744c3 != nil {
				return fmt.Errorf("[Upload to: %s] %w", __obf_941abe6dce7e870c, __obf_7f6f4ab0f2c744c3)
			}
		}
	} else {
		__obf_7f6f4ab0f2c744c3 = __obf_67195a8151e0c64a.__obf_527bd2433cb23cc2(__obf_c44777f527dbbd3d, __obf_321a70496679c815, __obf_941abe6dce7e870c)
	}
	return
}

// Download file from remote server!
func (__obf_67195a8151e0c64a SSHClient) Download(__obf_7897a1786c6083a1 string, __obf_321a70496679c815 string) error {
	var __obf_e76b5842365b8831 = path.Join(__obf_321a70496679c815, filepath.Base(__obf_7897a1786c6083a1))
	_, __obf_7f6f4ab0f2c744c3 := os.Stat(__obf_e76b5842365b8831)
	if __obf_7f6f4ab0f2c744c3 == nil {
		return errors.New("file exists")
	}
	if os.IsNotExist(__obf_7f6f4ab0f2c744c3) {
		__obf_e86848930264c52f, __obf_7f6f4ab0f2c744c3 := os.Create(__obf_e76b5842365b8831)
		if __obf_7f6f4ab0f2c744c3 != nil {
			return fmt.Errorf("[Create: %s] %w", __obf_e76b5842365b8831, __obf_7f6f4ab0f2c744c3)
		}
		defer __obf_e86848930264c52f.Close()
		__obf_c44777f527dbbd3d, __obf_7f6f4ab0f2c744c3 := __obf_67195a8151e0c64a.NewSftp()
		if __obf_7f6f4ab0f2c744c3 != nil {

			return __obf_7f6f4ab0f2c744c3
		}
		defer __obf_c44777f527dbbd3d.Close()
		__obf_83d70bf93c31aff3, __obf_7f6f4ab0f2c744c3 := __obf_c44777f527dbbd3d.Open(__obf_7897a1786c6083a1)
		if __obf_7f6f4ab0f2c744c3 != nil {
			return __obf_7f6f4ab0f2c744c3
		}
		defer __obf_83d70bf93c31aff3.Close()

		if _, __obf_7f6f4ab0f2c744c3 = io.Copy(__obf_e86848930264c52f, __obf_83d70bf93c31aff3); __obf_7f6f4ab0f2c744c3 != nil {
			return __obf_7f6f4ab0f2c744c3
		}

		return __obf_e86848930264c52f.Sync()
	}
	return __obf_7f6f4ab0f2c744c3
}
