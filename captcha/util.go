package __obf_85f036759f76ec38

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func RandomStr(__obf_2fbce240c29d891e int) string {
	__obf_afefc3afd8988867 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	__obf_14aee239033ef89a := []byte(__obf_afefc3afd8988867)
	__obf_c1b0a03f1491a7d4 := []byte{}
	__obf_9f0cd3415eba4a85 := rand.New(rand.NewSource(time.Now().UnixNano()))
	for range __obf_2fbce240c29d891e {
		__obf_c1b0a03f1491a7d4 = append(__obf_c1b0a03f1491a7d4, __obf_14aee239033ef89a[__obf_9f0cd3415eba4a85.Intn(len(__obf_14aee239033ef89a))])
	}
	return string(__obf_c1b0a03f1491a7d4)
}

func StringUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
