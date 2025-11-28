package __obf_4574f91399c0b6ef

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func RandomStr(__obf_63d3d9594a8f0770 int) string {
	__obf_d2b48024d1914776 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	__obf_ee698b8f2f341d1b := []byte(__obf_d2b48024d1914776)
	__obf_41485b78b867198b := []byte{}
	__obf_5132c4a0b1dc11cf := rand.New(rand.NewSource(time.Now().UnixNano()))
	for range __obf_63d3d9594a8f0770 {
		__obf_41485b78b867198b = append(__obf_41485b78b867198b, __obf_ee698b8f2f341d1b[__obf_5132c4a0b1dc11cf.Intn(len(__obf_ee698b8f2f341d1b))])
	}
	return string(__obf_41485b78b867198b)
}

func StringUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
