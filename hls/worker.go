package __obf_ab25ed2437cd567a

import (
	"bytes"
	"context"
	"crypto/sha1"
	"fmt"
	"io"
	"os/exec"
	"time"
)

// type Handler func(any, io.Writer) error

type Worker struct {
	__obf_a501203efe32a7db Cache
	__obf_f77c32471d383438 int
	__obf_b130575b6ae10f38 chan struct{}
}

func NewWorker(__obf_898122e6cdf30aad, __obf_f77c32471d383438 int, __obf_5671672528628a6a string) *Worker {
	__obf_b130575b6ae10f38 := make(chan struct{}, __obf_898122e6cdf30aad)
	for __obf_be8d4d616fa7a880 := __obf_898122e6cdf30aad; __obf_be8d4d616fa7a880 > 0; __obf_be8d4d616fa7a880-- {
		__obf_b130575b6ae10f38 <- struct{}{}
	}
	return &Worker{NewDirCache(__obf_5671672528628a6a), __obf_f77c32471d383438, __obf_b130575b6ae10f38}
}

// TODO timeout & context
func (__obf_8d3b3856596fc4f4 *Worker) Serve(__obf_7428e3ac679e65fe string, __obf_5b69e94b87fb5441 []string, __obf_98c2f72bee615664 io.Writer) error {

	__obf_9f32223e1acaa774 := sha1.New()
	__obf_9f32223e1acaa774.Write([]byte(__obf_7428e3ac679e65fe))
	for _, __obf_2172876de9966b00 := range __obf_5b69e94b87fb5441 {
		__obf_9f32223e1acaa774.Write([]byte(__obf_2172876de9966b00))
	}
	__obf_f7b675362b70b512 := __obf_9f32223e1acaa774.Sum(nil)
	__obf_10900f54d35c1067 := fmt.Sprintf("%x", __obf_f7b675362b70b512)

	__obf_2225771e74950257, __obf_9d170493b73636ca := __obf_8d3b3856596fc4f4.__obf_a501203efe32a7db.Get(context.Background(), __obf_10900f54d35c1067)
	// If error getting item, return not served with error
	if __obf_9d170493b73636ca != nil {
		return __obf_9d170493b73636ca
	}
	// If no item found, return not served with no error
	if __obf_2225771e74950257 != nil {
		// If copying fails, return served with error
		if _, __obf_9d170493b73636ca = io.Copy(__obf_98c2f72bee615664, bytes.NewReader(__obf_2225771e74950257)); __obf_9d170493b73636ca != nil {
			return __obf_9d170493b73636ca
		}
		// Everything worked, return served with no error
		return nil
	}

	// Wait for token
	__obf_8d5b84246c0c48cb := <-__obf_8d3b3856596fc4f4.__obf_b130575b6ae10f38
	defer func() {
		__obf_8d3b3856596fc4f4.__obf_b130575b6ae10f38 <- __obf_8d5b84246c0c48cb
	}()

	__obf_0e8aa16629ed5bab := new(bytes.Buffer)
	__obf_d71ce65cd2837717 := io.MultiWriter(__obf_0e8aa16629ed5bab, __obf_98c2f72bee615664)

	__obf_378e35ef817a6881, __obf_032b931ab588608d := context.WithTimeout(context.Background(), time.Duration(__obf_8d3b3856596fc4f4.__obf_f77c32471d383438)*time.Second)
	defer __obf_032b931ab588608d()

	__obf_6319c2b6d1f4ed90 := exec.CommandContext(__obf_378e35ef817a6881, __obf_7428e3ac679e65fe, __obf_5b69e94b87fb5441...)

	if __obf_9d170493b73636ca := ExecWriteStdout(__obf_6319c2b6d1f4ed90, __obf_d71ce65cd2837717); __obf_9d170493b73636ca != nil {
		return __obf_9d170493b73636ca
	}

	return __obf_8d3b3856596fc4f4.__obf_a501203efe32a7db.Set(context.Background(), __obf_10900f54d35c1067, __obf_0e8aa16629ed5bab.Bytes())
}
