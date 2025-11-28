package __obf_42bbcad92b7de1a8

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
	__obf_cda360a7231b944b Cache
	__obf_3b77a3401f9bd567 int
	__obf_92313ab011144ef7 chan struct{}
}

func NewWorker(__obf_ebdfc724cc2ca964, __obf_3b77a3401f9bd567 int, __obf_9f8948d7d5153274 string) *Worker {
	__obf_92313ab011144ef7 := make(chan struct{}, __obf_ebdfc724cc2ca964)
	for __obf_1da2fb76f7e1e655 := __obf_ebdfc724cc2ca964; __obf_1da2fb76f7e1e655 > 0; __obf_1da2fb76f7e1e655-- {
		__obf_92313ab011144ef7 <- struct{}{}
	}
	return &Worker{NewDirCache(__obf_9f8948d7d5153274), __obf_3b77a3401f9bd567, __obf_92313ab011144ef7}
}

// TODO timeout & context
func (__obf_c7be2ffa97bb9914 *Worker) Serve(__obf_14ffdbc17fdc424c string, __obf_0a4d3e5c45b3d81b []string, __obf_39100c7ef6f93ec8 io.Writer) error {

	__obf_4fec35dab792db7d := sha1.New()
	__obf_4fec35dab792db7d.Write([]byte(__obf_14ffdbc17fdc424c))
	for _, __obf_f20b361e19413a01 := range __obf_0a4d3e5c45b3d81b {
		__obf_4fec35dab792db7d.Write([]byte(__obf_f20b361e19413a01))
	}
	__obf_7779bd288e23b1c5 := __obf_4fec35dab792db7d.Sum(nil)
	__obf_7d80d03709c3cc8f := fmt.Sprintf("%x", __obf_7779bd288e23b1c5)

	__obf_15f3c63e736301fe, __obf_59b2ec060775896c := __obf_c7be2ffa97bb9914.__obf_cda360a7231b944b.Get(context.Background(), __obf_7d80d03709c3cc8f)
	// If error getting item, return not served with error
	if __obf_59b2ec060775896c != nil {
		return __obf_59b2ec060775896c
	}
	// If no item found, return not served with no error
	if __obf_15f3c63e736301fe != nil {
		// If copying fails, return served with error
		if _, __obf_59b2ec060775896c = io.Copy(__obf_39100c7ef6f93ec8, bytes.NewReader(__obf_15f3c63e736301fe)); __obf_59b2ec060775896c != nil {
			return __obf_59b2ec060775896c
		}
		// Everything worked, return served with no error
		return nil
	}

	// Wait for token
	__obf_2fb204dc0d4bd93b := <-__obf_c7be2ffa97bb9914.__obf_92313ab011144ef7
	defer func() {
		__obf_c7be2ffa97bb9914.__obf_92313ab011144ef7 <- __obf_2fb204dc0d4bd93b
	}()

	__obf_e4ba02fa67a5077a := new(bytes.Buffer)
	__obf_47c7cfb528351034 := io.MultiWriter(__obf_e4ba02fa67a5077a, __obf_39100c7ef6f93ec8)

	__obf_a8cca90b5a49e090, __obf_93f551cc9971e65e := context.WithTimeout(context.Background(), time.Duration(__obf_c7be2ffa97bb9914.__obf_3b77a3401f9bd567)*time.Second)
	defer __obf_93f551cc9971e65e()

	__obf_a7aad0ff1142655b := exec.CommandContext(__obf_a8cca90b5a49e090, __obf_14ffdbc17fdc424c, __obf_0a4d3e5c45b3d81b...)

	if __obf_59b2ec060775896c := ExecWriteStdout(__obf_a7aad0ff1142655b, __obf_47c7cfb528351034); __obf_59b2ec060775896c != nil {
		return __obf_59b2ec060775896c
	}

	return __obf_c7be2ffa97bb9914.__obf_cda360a7231b944b.Set(context.Background(), __obf_7d80d03709c3cc8f, __obf_e4ba02fa67a5077a.Bytes())
}
