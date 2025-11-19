package __obf_e49f90b5aaf58063

import (
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func RandomStr(__obf_851d19e463cb709f int) string {
	__obf_dcf3c5ed6748658f := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	__obf_4c0bb1a9ee1c072f := []byte(__obf_dcf3c5ed6748658f)
	__obf_23dc79849b35112f := []byte{}
	__obf_153e9b2c25ef2b10 := rand.New(rand.NewSource(time.Now().UnixNano()))
	for range __obf_851d19e463cb709f {
		__obf_23dc79849b35112f = append(__obf_23dc79849b35112f, __obf_4c0bb1a9ee1c072f[__obf_153e9b2c25ef2b10.Intn(len(__obf_4c0bb1a9ee1c072f))])
	}
	return string(__obf_23dc79849b35112f)
}

func StringUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
