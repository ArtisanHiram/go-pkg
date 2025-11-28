package __obf_699038f0de0a4a2b

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func RandomStr(__obf_5c981462c1ab6ad0 int) string {
	__obf_36ea44e1be6c9888 := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	__obf_9c17cd3fff53a703 := []byte(__obf_36ea44e1be6c9888)
	__obf_107211937a16fbfb := []byte{}
	__obf_d8fde49203f68fa9 := rand.New(rand.NewSource(time.Now().UnixNano()))
	for range __obf_5c981462c1ab6ad0 {
		__obf_107211937a16fbfb = append(__obf_107211937a16fbfb, __obf_9c17cd3fff53a703[__obf_d8fde49203f68fa9.Intn(len(__obf_9c17cd3fff53a703))])
	}
	return string(__obf_107211937a16fbfb)
}

func StringUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
