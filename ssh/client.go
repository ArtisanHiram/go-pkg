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
func NewSSHClient(__obf_13d494a5f30a1586, __obf_36a396dedfab28e7 string, __obf_bee0a49334b7deea *ssh.ClientConfig) (__obf_4867138fc4e66c17 *SSHClient) {
	// c.Client, err = Dial("tcp", addr, port)
	return &SSHClient{
		Addr:   net.JoinHostPort(__obf_13d494a5f30a1586, __obf_36a396dedfab28e7),
		Config: __obf_bee0a49334b7deea,
	}
}

func (__obf_4867138fc4e66c17 *SSHClient) init() error {
	if __obf_4867138fc4e66c17.Client == nil {
		var __obf_fe854faece812b24 error
		__obf_4867138fc4e66c17.
			Client, __obf_fe854faece812b24 = ssh.Dial("tcp", __obf_4867138fc4e66c17.Addr, __obf_4867138fc4e66c17.Config)
		if __obf_fe854faece812b24 != nil {
			return __obf_fe854faece812b24
		}
	}
	return nil
}

// Run starts a new SSH session and runs the cmd, it returns CombinedOutput and err if some.
func (__obf_4867138fc4e66c17 SSHClient) Run(__obf_e9d8cd0e7b7a166b string) ([]byte, error) {

	var __obf_fe854faece812b24 = __obf_4867138fc4e66c17.init()
	if __obf_fe854faece812b24 != nil {
		return nil, __obf_fe854faece812b24
	}

	var __obf_da69d9421cf9fbf7 *ssh.Session
	if __obf_da69d9421cf9fbf7, __obf_fe854faece812b24 = __obf_4867138fc4e66c17.NewSession(); __obf_fe854faece812b24 != nil {
		return nil, __obf_fe854faece812b24
	}

	defer __obf_da69d9421cf9fbf7.Close()

	return __obf_da69d9421cf9fbf7.CombinedOutput(__obf_e9d8cd0e7b7a166b)
}

// NewSftp returns new sftp client and error if some.
func (__obf_4867138fc4e66c17 SSHClient) NewSftp(__obf_fdd1e96537e169d8 ...sftp.ClientOption) (*sftp.Client, error) {
	if __obf_fe854faece812b24 := __obf_4867138fc4e66c17.init(); __obf_fe854faece812b24 != nil {
		return nil, __obf_fe854faece812b24
	}
	return sftp.NewClient(__obf_4867138fc4e66c17.Client, __obf_fdd1e96537e169d8...)
}

// uploadFile a local file to remote server!
func (__obf_4867138fc4e66c17 SSHClient) __obf_5cb356ab68ebf381(__obf_7c61ef11d626f697 *sftp.Client, __obf_4b29cd18e7baef72 string, __obf_8a9ec9c8c132ae26 string) (__obf_fe854faece812b24 error) {

	var __obf_0623b6851dfbd00b *os.File
	__obf_0623b6851dfbd00b, __obf_fe854faece812b24 = os.Open(__obf_4b29cd18e7baef72)
	if __obf_fe854faece812b24 != nil {
		return
	}
	defer __obf_0623b6851dfbd00b.Close()

	var __obf_4dea84270d4357ce *sftp.File
	__obf_696d80be43188f97 := path.Join(__obf_8a9ec9c8c132ae26, filepath.Base(__obf_4b29cd18e7baef72))
	__obf_4dea84270d4357ce, __obf_fe854faece812b24 = __obf_7c61ef11d626f697.Create(__obf_696d80be43188f97)
	if __obf_fe854faece812b24 != nil {
		__obf_fe854faece812b24 = fmt.Errorf("[Create: %s] %w", __obf_696d80be43188f97, __obf_fe854faece812b24)
		return
	}
	defer __obf_4dea84270d4357ce.Close()

	// size := len(s)
	// info, err := local.Stat()
	// if err != nil {
	// 	return err
	// }

	var __obf_c55fb29f02d3badd int64
	__obf_c55fb29f02d3badd, __obf_fe854faece812b24 = io.Copy(__obf_4dea84270d4357ce, io.TeeReader(__obf_0623b6851dfbd00b, &WriteCounter{}))
	fmt.Fprintf(os.Stdout, "Transferred %.2f kb\n", float64(__obf_c55fb29f02d3badd/1024))
	return
}

// Upload a local file or directory to remote server!
func (__obf_4867138fc4e66c17 SSHClient) Upload(__obf_4b29cd18e7baef72 string, __obf_8a9ec9c8c132ae26 string) (__obf_fe854faece812b24 error) {
	__obf_677a9cff4da38866,

		//获取路径的属性
		__obf_fe854faece812b24 := os.Stat(__obf_4b29cd18e7baef72)
	if __obf_fe854faece812b24 != nil {
		return
	}
	__obf_fb7340ce395d819e, __obf_fe854faece812b24 := __obf_4867138fc4e66c17.NewSftp()
	if __obf_fe854faece812b24 != nil {
		return
	}
	defer __obf_fb7340ce395d819e.Close()

	// 判断是否是文件夹
	if __obf_677a9cff4da38866.IsDir() {
		var __obf_133c94bd81d9d55c []os.DirEntry
		__obf_133c94bd81d9d55c, __obf_fe854faece812b24 = os.ReadDir(__obf_4b29cd18e7baef72)
		if __obf_fe854faece812b24 != nil {
			return
		}
		__obf_8a9ec9c8c132ae26 = // 先创建最外层文件夹
			path.Join(__obf_8a9ec9c8c132ae26, __obf_677a9cff4da38866.Name())
		__obf_fe854faece812b24 = __obf_fb7340ce395d819e.Mkdir(__obf_8a9ec9c8c132ae26)
		if __obf_fe854faece812b24 != nil {
			return fmt.Errorf("[Mkdir] %s failed: %w", __obf_8a9ec9c8c132ae26, __obf_fe854faece812b24)
		}
		// 遍历文件夹内容
		for _, __obf_e59cb9284512b302 := range __obf_133c94bd81d9d55c {
			__obf_fe854faece812b24 = // 判断是否是文件,是文件直接上传.是文件夹,先远程创建文件夹,再递归复制内部文件
				__obf_4867138fc4e66c17.Upload(path.Join(__obf_4b29cd18e7baef72, __obf_e59cb9284512b302.Name()), __obf_8a9ec9c8c132ae26)
			if __obf_fe854faece812b24 != nil {
				return fmt.Errorf("[Upload to: %s] %w", __obf_8a9ec9c8c132ae26, __obf_fe854faece812b24)
			}
		}
	} else {
		__obf_fe854faece812b24 = __obf_4867138fc4e66c17.__obf_5cb356ab68ebf381(__obf_fb7340ce395d819e, __obf_4b29cd18e7baef72, __obf_8a9ec9c8c132ae26)
	}
	return
}

// Download file from remote server!
func (__obf_4867138fc4e66c17 SSHClient) Download(__obf_696d80be43188f97 string, __obf_4b29cd18e7baef72 string) error {
	var __obf_aaede9339aa251e1 = path.Join(__obf_4b29cd18e7baef72, filepath.Base(__obf_696d80be43188f97))
	_, __obf_fe854faece812b24 := os.Stat(__obf_aaede9339aa251e1)
	if __obf_fe854faece812b24 == nil {
		return errors.New("file exists")
	}
	if os.IsNotExist(__obf_fe854faece812b24) {
		__obf_0623b6851dfbd00b, __obf_fe854faece812b24 := os.Create(__obf_aaede9339aa251e1)
		if __obf_fe854faece812b24 != nil {
			return fmt.Errorf("[Create: %s] %w", __obf_aaede9339aa251e1, __obf_fe854faece812b24)
		}
		defer __obf_0623b6851dfbd00b.Close()
		__obf_fb7340ce395d819e, __obf_fe854faece812b24 := __obf_4867138fc4e66c17.NewSftp()
		if __obf_fe854faece812b24 != nil {

			return __obf_fe854faece812b24
		}
		defer __obf_fb7340ce395d819e.Close()
		__obf_4dea84270d4357ce, __obf_fe854faece812b24 := __obf_fb7340ce395d819e.Open(__obf_696d80be43188f97)
		if __obf_fe854faece812b24 != nil {
			return __obf_fe854faece812b24
		}
		defer __obf_4dea84270d4357ce.Close()

		if _, __obf_fe854faece812b24 = io.Copy(__obf_0623b6851dfbd00b, __obf_4dea84270d4357ce); __obf_fe854faece812b24 != nil {
			return __obf_fe854faece812b24
		}

		return __obf_0623b6851dfbd00b.Sync()
	}
	return __obf_fe854faece812b24
}
