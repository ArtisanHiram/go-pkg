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
func NewSSHClient(__obf_06c42060b2c0c374, __obf_abf4e9548e3a9b2c string, __obf_712de178ead485e9 *ssh.ClientConfig) (__obf_948f47c64b517453 *SSHClient) {
	// c.Client, err = Dial("tcp", addr, port)
	return &SSHClient{
		Addr:   net.JoinHostPort(__obf_06c42060b2c0c374, __obf_abf4e9548e3a9b2c),
		Config: __obf_712de178ead485e9,
	}
}

func (__obf_948f47c64b517453 *SSHClient) __obf_8f2f239fa0f5f4a1() error {
	if __obf_948f47c64b517453.Client == nil {
		var __obf_6eb8d6029fd3be00 error
		__obf_948f47c64b517453.Client, __obf_6eb8d6029fd3be00 = ssh.Dial("tcp", __obf_948f47c64b517453.Addr, __obf_948f47c64b517453.Config)
		if __obf_6eb8d6029fd3be00 != nil {
			return __obf_6eb8d6029fd3be00
		}
	}
	return nil
}

// Run starts a new SSH session and runs the cmd, it returns CombinedOutput and err if any.
func (__obf_948f47c64b517453 SSHClient) Run(__obf_b71fdb506a8909e5 string) ([]byte, error) {

	var __obf_6eb8d6029fd3be00 = __obf_948f47c64b517453.__obf_8f2f239fa0f5f4a1()
	if __obf_6eb8d6029fd3be00 != nil {
		return nil, __obf_6eb8d6029fd3be00
	}

	var __obf_a7c84478eb9c2947 *ssh.Session
	if __obf_a7c84478eb9c2947, __obf_6eb8d6029fd3be00 = __obf_948f47c64b517453.NewSession(); __obf_6eb8d6029fd3be00 != nil {
		return nil, __obf_6eb8d6029fd3be00
	}

	defer __obf_a7c84478eb9c2947.Close()

	return __obf_a7c84478eb9c2947.CombinedOutput(__obf_b71fdb506a8909e5)
}

// NewSftp returns new sftp client and error if any.
func (__obf_948f47c64b517453 SSHClient) NewSftp(__obf_de552880423f4c8b ...sftp.ClientOption) (*sftp.Client, error) {
	if __obf_6eb8d6029fd3be00 := __obf_948f47c64b517453.__obf_8f2f239fa0f5f4a1(); __obf_6eb8d6029fd3be00 != nil {
		return nil, __obf_6eb8d6029fd3be00
	}
	return sftp.NewClient(__obf_948f47c64b517453.Client, __obf_de552880423f4c8b...)
}

// uploadFile a local file to remote server!
func (__obf_948f47c64b517453 SSHClient) __obf_7572c18460bdf666(__obf_2b9ee3e623c3237f *sftp.Client, __obf_a4235e886f7e90d9 string, __obf_5b82ce4e631943d9 string) (__obf_6eb8d6029fd3be00 error) {

	var __obf_838beb0c1955599c *os.File
	__obf_838beb0c1955599c, __obf_6eb8d6029fd3be00 = os.Open(__obf_a4235e886f7e90d9)
	if __obf_6eb8d6029fd3be00 != nil {
		return
	}
	defer __obf_838beb0c1955599c.Close()

	var __obf_b9d3a70bdbc1c449 *sftp.File
	__obf_f3a735d573dd9e2b := path.Join(__obf_5b82ce4e631943d9, filepath.Base(__obf_a4235e886f7e90d9))
	__obf_b9d3a70bdbc1c449, __obf_6eb8d6029fd3be00 = __obf_2b9ee3e623c3237f.Create(__obf_f3a735d573dd9e2b)
	if __obf_6eb8d6029fd3be00 != nil {
		__obf_6eb8d6029fd3be00 = fmt.Errorf("[Create: %s] %w", __obf_f3a735d573dd9e2b, __obf_6eb8d6029fd3be00)
		return
	}
	defer __obf_b9d3a70bdbc1c449.Close()

	// size := len(s)
	// info, err := local.Stat()
	// if err != nil {
	// 	return err
	// }

	var __obf_4a2184a1397948a5 int64
	__obf_4a2184a1397948a5, __obf_6eb8d6029fd3be00 = io.Copy(__obf_b9d3a70bdbc1c449, io.TeeReader(__obf_838beb0c1955599c, &WriteCounter{}))
	fmt.Fprintf(os.Stdout, "Transferred %.2f kb\n", float64(__obf_4a2184a1397948a5/1024))
	return
}

// Upload a local file or directory to remote server!
func (__obf_948f47c64b517453 SSHClient) Upload(__obf_a4235e886f7e90d9 string, __obf_5b82ce4e631943d9 string) (__obf_6eb8d6029fd3be00 error) {

	//获取路径的属性
	__obf_a3ba964874e0249b, __obf_6eb8d6029fd3be00 := os.Stat(__obf_a4235e886f7e90d9)
	if __obf_6eb8d6029fd3be00 != nil {
		return
	}

	__obf_25c45af65c4abecf, __obf_6eb8d6029fd3be00 := __obf_948f47c64b517453.NewSftp()
	if __obf_6eb8d6029fd3be00 != nil {
		return
	}
	defer __obf_25c45af65c4abecf.Close()

	// 判断是否是文件夹
	if __obf_a3ba964874e0249b.IsDir() {
		var __obf_3b8bf59ce1f4e62b []os.DirEntry
		__obf_3b8bf59ce1f4e62b, __obf_6eb8d6029fd3be00 = os.ReadDir(__obf_a4235e886f7e90d9)
		if __obf_6eb8d6029fd3be00 != nil {
			return
		}
		// 先创建最外层文件夹
		__obf_5b82ce4e631943d9 = path.Join(__obf_5b82ce4e631943d9, __obf_a3ba964874e0249b.Name())
		__obf_6eb8d6029fd3be00 = __obf_25c45af65c4abecf.Mkdir(__obf_5b82ce4e631943d9)
		if __obf_6eb8d6029fd3be00 != nil {
			return fmt.Errorf("[Mkdir] %s failed: %w", __obf_5b82ce4e631943d9, __obf_6eb8d6029fd3be00)
		}
		// 遍历文件夹内容
		for _, __obf_11bb2e873fd37028 := range __obf_3b8bf59ce1f4e62b {
			// 判断是否是文件,是文件直接上传.是文件夹,先远程创建文件夹,再递归复制内部文件
			__obf_6eb8d6029fd3be00 = __obf_948f47c64b517453.Upload(path.Join(__obf_a4235e886f7e90d9, __obf_11bb2e873fd37028.Name()), __obf_5b82ce4e631943d9)
			if __obf_6eb8d6029fd3be00 != nil {
				return fmt.Errorf("[Upload to: %s] %w", __obf_5b82ce4e631943d9, __obf_6eb8d6029fd3be00)
			}
		}
	} else {
		__obf_6eb8d6029fd3be00 = __obf_948f47c64b517453.__obf_7572c18460bdf666(__obf_25c45af65c4abecf, __obf_a4235e886f7e90d9, __obf_5b82ce4e631943d9)
	}
	return
}

// Download file from remote server!
func (__obf_948f47c64b517453 SSHClient) Download(__obf_f3a735d573dd9e2b string, __obf_a4235e886f7e90d9 string) error {
	var __obf_8af93cc10ba02fe0 = path.Join(__obf_a4235e886f7e90d9, filepath.Base(__obf_f3a735d573dd9e2b))
	_, __obf_6eb8d6029fd3be00 := os.Stat(__obf_8af93cc10ba02fe0)
	if __obf_6eb8d6029fd3be00 == nil {
		return errors.New("file exists")
	}
	if os.IsNotExist(__obf_6eb8d6029fd3be00) {
		__obf_838beb0c1955599c, __obf_6eb8d6029fd3be00 := os.Create(__obf_8af93cc10ba02fe0)
		if __obf_6eb8d6029fd3be00 != nil {
			return fmt.Errorf("[Create: %s] %w", __obf_8af93cc10ba02fe0, __obf_6eb8d6029fd3be00)
		}
		defer __obf_838beb0c1955599c.Close()

		__obf_25c45af65c4abecf, __obf_6eb8d6029fd3be00 := __obf_948f47c64b517453.NewSftp()
		if __obf_6eb8d6029fd3be00 != nil {

			return __obf_6eb8d6029fd3be00
		}
		defer __obf_25c45af65c4abecf.Close()

		__obf_b9d3a70bdbc1c449, __obf_6eb8d6029fd3be00 := __obf_25c45af65c4abecf.Open(__obf_f3a735d573dd9e2b)
		if __obf_6eb8d6029fd3be00 != nil {
			return __obf_6eb8d6029fd3be00
		}
		defer __obf_b9d3a70bdbc1c449.Close()

		if _, __obf_6eb8d6029fd3be00 = io.Copy(__obf_838beb0c1955599c, __obf_b9d3a70bdbc1c449); __obf_6eb8d6029fd3be00 != nil {
			return __obf_6eb8d6029fd3be00
		}

		return __obf_838beb0c1955599c.Sync()
	}
	return __obf_6eb8d6029fd3be00
}
