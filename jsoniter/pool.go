package __obf_703d23ba520c3295

import (
	"io"
)

// IteratorPool a thread safe pool of iterators with same configuration
type IteratorPool interface {
	BorrowIterator(__obf_c180417ee204f8c5 []byte) *Iterator
	ReturnIterator(__obf_47edab4c16a2d88d *Iterator)
}

// StreamPool a thread safe pool of streams with same configuration
type StreamPool interface {
	BorrowStream(__obf_8d0bd0f6d68b0c59 io.Writer) *Stream
	ReturnStream(__obf_9a34f62857fb1d1d *Stream)
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) BorrowStream(__obf_8d0bd0f6d68b0c59 io.Writer) *Stream {
	__obf_9a34f62857fb1d1d := __obf_25bd4f754a37b862.__obf_27ad6a094b7bcdbf.Get().(*Stream)
	__obf_9a34f62857fb1d1d.
		Reset(__obf_8d0bd0f6d68b0c59)
	return __obf_9a34f62857fb1d1d
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) ReturnStream(__obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.__obf_1746319fdfc0f229 = nil
	__obf_9a34f62857fb1d1d.
		Error = nil
	__obf_9a34f62857fb1d1d.
		Attachment = nil
	__obf_25bd4f754a37b862.__obf_27ad6a094b7bcdbf.
		Put(__obf_9a34f62857fb1d1d)
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) BorrowIterator(__obf_c180417ee204f8c5 []byte) *Iterator {
	__obf_47edab4c16a2d88d := __obf_25bd4f754a37b862.__obf_582fcdbe5914a904.Get().(*Iterator)
	__obf_47edab4c16a2d88d.
		ResetBytes(__obf_c180417ee204f8c5)
	return __obf_47edab4c16a2d88d
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) ReturnIterator(__obf_47edab4c16a2d88d *Iterator) {
	__obf_47edab4c16a2d88d.
		Error = nil
	__obf_47edab4c16a2d88d.
		Attachment = nil
	__obf_25bd4f754a37b862.__obf_582fcdbe5914a904.
		Put(__obf_47edab4c16a2d88d)
}
