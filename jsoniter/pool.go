package __obf_030d39f7a8de96c6

import (
	"io"
)

// IteratorPool a thread safe pool of iterators with same configuration
type IteratorPool interface {
	BorrowIterator(__obf_c28f0e30eceb6d4a []byte) *Iterator
	ReturnIterator(__obf_4ab56a99965952e7 *Iterator)
}

// StreamPool a thread safe pool of streams with same configuration
type StreamPool interface {
	BorrowStream(__obf_ae48508e054620bd io.Writer) *Stream
	ReturnStream(__obf_8a2c51fe22775655 *Stream)
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) BorrowStream(__obf_ae48508e054620bd io.Writer) *Stream {
	__obf_8a2c51fe22775655 := __obf_679611dc56ff465b.__obf_e195df3f11517e00.Get().(*Stream)
	__obf_8a2c51fe22775655.
		Reset(__obf_ae48508e054620bd)
	return __obf_8a2c51fe22775655
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) ReturnStream(__obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.__obf_238de3ec07c7e9da = nil
	__obf_8a2c51fe22775655.
		Error = nil
	__obf_8a2c51fe22775655.
		Attachment = nil
	__obf_679611dc56ff465b.__obf_e195df3f11517e00.
		Put(__obf_8a2c51fe22775655)
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) BorrowIterator(__obf_c28f0e30eceb6d4a []byte) *Iterator {
	__obf_4ab56a99965952e7 := __obf_679611dc56ff465b.__obf_4522de1f25b34c0b.Get().(*Iterator)
	__obf_4ab56a99965952e7.
		ResetBytes(__obf_c28f0e30eceb6d4a)
	return __obf_4ab56a99965952e7
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) ReturnIterator(__obf_4ab56a99965952e7 *Iterator) {
	__obf_4ab56a99965952e7.
		Error = nil
	__obf_4ab56a99965952e7.
		Attachment = nil
	__obf_679611dc56ff465b.__obf_4522de1f25b34c0b.
		Put(__obf_4ab56a99965952e7)
}
