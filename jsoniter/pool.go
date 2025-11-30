package __obf_c3cd04a15c56f32f

import (
	"io"
)

// IteratorPool a thread safe pool of iterators with same configuration
type IteratorPool interface {
	BorrowIterator(__obf_905295f6bd29aafe []byte) *Iterator
	ReturnIterator(__obf_edd9fe48d39445e4 *Iterator)
}

// StreamPool a thread safe pool of streams with same configuration
type StreamPool interface {
	BorrowStream(__obf_b1699a146b20b28e io.Writer) *Stream
	ReturnStream(__obf_2361f5214e84c60f *Stream)
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) BorrowStream(__obf_b1699a146b20b28e io.Writer) *Stream {
	__obf_2361f5214e84c60f := __obf_0c8065c219ccf0ab.__obf_456a3a3023fd9656.Get().(*Stream)
	__obf_2361f5214e84c60f.
		Reset(__obf_b1699a146b20b28e)
	return __obf_2361f5214e84c60f
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) ReturnStream(__obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.__obf_743aba00c5d1ef26 = nil
	__obf_2361f5214e84c60f.
		Error = nil
	__obf_2361f5214e84c60f.
		Attachment = nil
	__obf_0c8065c219ccf0ab.__obf_456a3a3023fd9656.
		Put(__obf_2361f5214e84c60f)
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) BorrowIterator(__obf_905295f6bd29aafe []byte) *Iterator {
	__obf_edd9fe48d39445e4 := __obf_0c8065c219ccf0ab.__obf_7e0b4a46cc5f35d8.Get().(*Iterator)
	__obf_edd9fe48d39445e4.
		ResetBytes(__obf_905295f6bd29aafe)
	return __obf_edd9fe48d39445e4
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) ReturnIterator(__obf_edd9fe48d39445e4 *Iterator) {
	__obf_edd9fe48d39445e4.
		Error = nil
	__obf_edd9fe48d39445e4.
		Attachment = nil
	__obf_0c8065c219ccf0ab.__obf_7e0b4a46cc5f35d8.
		Put(__obf_edd9fe48d39445e4)
}
