package __obf_4574f91399c0b6ef

// ---------------------------------
// siprng is needed for random functions, keeping it from the original code.
// ---------------------------------
type __obf_3f24088299fc8cc2 struct {
	__obf_ddc0228d360ac3fe, __obf_1d6eb6491d24d298, __obf_83ba1a0484f42b6e, __obf_ab52ff1ee21ec0ba uint64
}

const (
	__obf_bc02efce650a142f = 2
	__obf_d302a2303152a041 = 4
)

func (__obf_5132c4a0b1dc11cf *__obf_3f24088299fc8cc2) Seed(__obf_9c33ed0b58f68eaf [16]byte) {
	__obf_a38d6c2bf49c6bf1 := uint64(__obf_9c33ed0b58f68eaf[0]) | uint64(__obf_9c33ed0b58f68eaf[1])<<8 | uint64(__obf_9c33ed0b58f68eaf[2])<<16 |
		uint64(__obf_9c33ed0b58f68eaf[3])<<24 | uint64(__obf_9c33ed0b58f68eaf[4])<<32 | uint64(__obf_9c33ed0b58f68eaf[5])<<40 |
		uint64(__obf_9c33ed0b58f68eaf[6])<<48 | uint64(__obf_9c33ed0b58f68eaf[7])<<56
	__obf_7509cdc25eb21471 := uint64(__obf_9c33ed0b58f68eaf[8]) | uint64(__obf_9c33ed0b58f68eaf[9])<<8 | uint64(__obf_9c33ed0b58f68eaf[10])<<16 |
		uint64(__obf_9c33ed0b58f68eaf[11])<<24 | uint64(__obf_9c33ed0b58f68eaf[12])<<32 | uint64(__obf_9c33ed0b58f68eaf[13])<<40 |
		uint64(__obf_9c33ed0b58f68eaf[14])<<48 | uint64(__obf_9c33ed0b58f68eaf[15])<<56
	__obf_5132c4a0b1dc11cf.__obf_ddc0228d360ac3fe = __obf_a38d6c2bf49c6bf1 ^ 0x736f6d6570736575
	__obf_5132c4a0b1dc11cf.__obf_1d6eb6491d24d298 = __obf_7509cdc25eb21471 ^ 0x646f72616e646f6d
	__obf_5132c4a0b1dc11cf.__obf_83ba1a0484f42b6e = __obf_a38d6c2bf49c6bf1 ^ 0x6c7967656e657261
	__obf_5132c4a0b1dc11cf.__obf_ab52ff1ee21ec0ba = __obf_7509cdc25eb21471 ^ 0x7465646279746573
}

func (__obf_5132c4a0b1dc11cf *__obf_3f24088299fc8cc2) Uint64() uint64 {
	__obf_5132c4a0b1dc11cf.__obf_ab52ff1ee21ec0ba++
	for range __obf_bc02efce650a142f {
		__obf_5132c4a0b1dc11cf.__obf_ddc0228d360ac3fe += __obf_5132c4a0b1dc11cf.__obf_1d6eb6491d24d298
		__obf_5132c4a0b1dc11cf.__obf_1d6eb6491d24d298 = __obf_862aa9e6604eb238(__obf_5132c4a0b1dc11cf.__obf_1d6eb6491d24d298, 13)
		__obf_5132c4a0b1dc11cf.__obf_1d6eb6491d24d298 ^= __obf_5132c4a0b1dc11cf.__obf_ddc0228d360ac3fe
		__obf_5132c4a0b1dc11cf.__obf_ddc0228d360ac3fe = __obf_862aa9e6604eb238(__obf_5132c4a0b1dc11cf.__obf_ddc0228d360ac3fe, 32)
		__obf_5132c4a0b1dc11cf.__obf_83ba1a0484f42b6e += __obf_5132c4a0b1dc11cf.__obf_ab52ff1ee21ec0ba
		__obf_5132c4a0b1dc11cf.__obf_ab52ff1ee21ec0ba = __obf_862aa9e6604eb238(__obf_5132c4a0b1dc11cf.__obf_ab52ff1ee21ec0ba, 16)
		__obf_5132c4a0b1dc11cf.__obf_ab52ff1ee21ec0ba ^= __obf_5132c4a0b1dc11cf.__obf_83ba1a0484f42b6e
		__obf_5132c4a0b1dc11cf.__obf_83ba1a0484f42b6e += __obf_5132c4a0b1dc11cf.__obf_1d6eb6491d24d298
		__obf_5132c4a0b1dc11cf.__obf_1d6eb6491d24d298 = __obf_862aa9e6604eb238(__obf_5132c4a0b1dc11cf.__obf_1d6eb6491d24d298, 17)
		__obf_5132c4a0b1dc11cf.__obf_1d6eb6491d24d298 ^= __obf_5132c4a0b1dc11cf.__obf_83ba1a0484f42b6e
		__obf_5132c4a0b1dc11cf.__obf_83ba1a0484f42b6e = __obf_862aa9e6604eb238(__obf_5132c4a0b1dc11cf.__obf_83ba1a0484f42b6e, 32)
		__obf_5132c4a0b1dc11cf.__obf_ab52ff1ee21ec0ba += __obf_5132c4a0b1dc11cf.__obf_ddc0228d360ac3fe
		__obf_5132c4a0b1dc11cf.__obf_ddc0228d360ac3fe = __obf_862aa9e6604eb238(__obf_5132c4a0b1dc11cf.__obf_ddc0228d360ac3fe, 21)
		__obf_5132c4a0b1dc11cf.__obf_ddc0228d360ac3fe ^= __obf_5132c4a0b1dc11cf.__obf_ab52ff1ee21ec0ba
	}
	__obf_a3531c8a8e1f49ff := __obf_5132c4a0b1dc11cf.__obf_ddc0228d360ac3fe ^ __obf_5132c4a0b1dc11cf.__obf_1d6eb6491d24d298 ^ __obf_5132c4a0b1dc11cf.__obf_83ba1a0484f42b6e ^ __obf_5132c4a0b1dc11cf.__obf_ab52ff1ee21ec0ba
	for range __obf_d302a2303152a041 {
		__obf_5132c4a0b1dc11cf.__obf_ab52ff1ee21ec0ba++
	}
	return __obf_a3531c8a8e1f49ff
}

func (__obf_5132c4a0b1dc11cf *__obf_3f24088299fc8cc2) Float64() float64 {
	return float64(__obf_5132c4a0b1dc11cf.Uint64()>>11) / (1 << 53)
}

func (__obf_5132c4a0b1dc11cf *__obf_3f24088299fc8cc2) Intn(__obf_27a156eb5f8cea50 int) int {
	return int(__obf_5132c4a0b1dc11cf.Uint64() % uint64(__obf_27a156eb5f8cea50))
}

func __obf_862aa9e6604eb238(__obf_72e56065c0cb51c1 uint64, __obf_7cdb1db04840a397 int) uint64 {
	return (__obf_72e56065c0cb51c1 << __obf_7cdb1db04840a397) | (__obf_72e56065c0cb51c1 >> (64 - __obf_7cdb1db04840a397))
}

// import "encoding/binary"

// // siprng is PRNG based on SipHash-2-4.
// // (Note: it's not safe to use a single siprng from multiple goroutines.)
// type siprng struct {
// 	k0, k1, ctr uint64
// }

// // siphash implements SipHash-2-4, accepting a uint64 as a message.
// func siphash(k0, k1, m uint64) uint64 {
// 	// Initialization.
// 	v0 := k0 ^ 0x736f6d6570736575
// 	v1 := k1 ^ 0x646f72616e646f6d
// 	v2 := k0 ^ 0x6c7967656e657261
// 	v3 := k1 ^ 0x7465646279746573
// 	t := uint64(8) << 56

// 	// Compression.
// 	v3 ^= m

// 	// Round 1.
// 	v0 += v1
// 	v1 = v1<<13 | v1>>(64-13)
// 	v1 ^= v0
// 	v0 = v0<<32 | v0>>(64-32)

// 	v2 += v3
// 	v3 = v3<<16 | v3>>(64-16)
// 	v3 ^= v2

// 	v0 += v3
// 	v3 = v3<<21 | v3>>(64-21)
// 	v3 ^= v0

// 	v2 += v1
// 	v1 = v1<<17 | v1>>(64-17)
// 	v1 ^= v2
// 	v2 = v2<<32 | v2>>(64-32)

// 	// Round 2.
// 	v0 += v1
// 	v1 = v1<<13 | v1>>(64-13)
// 	v1 ^= v0
// 	v0 = v0<<32 | v0>>(64-32)

// 	v2 += v3
// 	v3 = v3<<16 | v3>>(64-16)
// 	v3 ^= v2

// 	v0 += v3
// 	v3 = v3<<21 | v3>>(64-21)
// 	v3 ^= v0

// 	v2 += v1
// 	v1 = v1<<17 | v1>>(64-17)
// 	v1 ^= v2
// 	v2 = v2<<32 | v2>>(64-32)

// 	v0 ^= m

// 	// Compress last block.
// 	v3 ^= t

// 	// Round 1.
// 	v0 += v1
// 	v1 = v1<<13 | v1>>(64-13)
// 	v1 ^= v0
// 	v0 = v0<<32 | v0>>(64-32)

// 	v2 += v3
// 	v3 = v3<<16 | v3>>(64-16)
// 	v3 ^= v2

// 	v0 += v3
// 	v3 = v3<<21 | v3>>(64-21)
// 	v3 ^= v0

// 	v2 += v1
// 	v1 = v1<<17 | v1>>(64-17)
// 	v1 ^= v2
// 	v2 = v2<<32 | v2>>(64-32)

// 	// Round 2.
// 	v0 += v1
// 	v1 = v1<<13 | v1>>(64-13)
// 	v1 ^= v0
// 	v0 = v0<<32 | v0>>(64-32)

// 	v2 += v3
// 	v3 = v3<<16 | v3>>(64-16)
// 	v3 ^= v2

// 	v0 += v3
// 	v3 = v3<<21 | v3>>(64-21)
// 	v3 ^= v0

// 	v2 += v1
// 	v1 = v1<<17 | v1>>(64-17)
// 	v1 ^= v2
// 	v2 = v2<<32 | v2>>(64-32)

// 	v0 ^= t

// 	// Finalization.
// 	v2 ^= 0xff

// 	// Round 1.
// 	v0 += v1
// 	v1 = v1<<13 | v1>>(64-13)
// 	v1 ^= v0
// 	v0 = v0<<32 | v0>>(64-32)

// 	v2 += v3
// 	v3 = v3<<16 | v3>>(64-16)
// 	v3 ^= v2

// 	v0 += v3
// 	v3 = v3<<21 | v3>>(64-21)
// 	v3 ^= v0

// 	v2 += v1
// 	v1 = v1<<17 | v1>>(64-17)
// 	v1 ^= v2
// 	v2 = v2<<32 | v2>>(64-32)

// 	// Round 2.
// 	v0 += v1
// 	v1 = v1<<13 | v1>>(64-13)
// 	v1 ^= v0
// 	v0 = v0<<32 | v0>>(64-32)

// 	v2 += v3
// 	v3 = v3<<16 | v3>>(64-16)
// 	v3 ^= v2

// 	v0 += v3
// 	v3 = v3<<21 | v3>>(64-21)
// 	v3 ^= v0

// 	v2 += v1
// 	v1 = v1<<17 | v1>>(64-17)
// 	v1 ^= v2
// 	v2 = v2<<32 | v2>>(64-32)

// 	// Round 3.
// 	v0 += v1
// 	v1 = v1<<13 | v1>>(64-13)
// 	v1 ^= v0
// 	v0 = v0<<32 | v0>>(64-32)

// 	v2 += v3
// 	v3 = v3<<16 | v3>>(64-16)
// 	v3 ^= v2

// 	v0 += v3
// 	v3 = v3<<21 | v3>>(64-21)
// 	v3 ^= v0

// 	v2 += v1
// 	v1 = v1<<17 | v1>>(64-17)
// 	v1 ^= v2
// 	v2 = v2<<32 | v2>>(64-32)

// 	// Round 4.
// 	v0 += v1
// 	v1 = v1<<13 | v1>>(64-13)
// 	v1 ^= v0
// 	v0 = v0<<32 | v0>>(64-32)

// 	v2 += v3
// 	v3 = v3<<16 | v3>>(64-16)
// 	v3 ^= v2

// 	v0 += v3
// 	v3 = v3<<21 | v3>>(64-21)
// 	v3 ^= v0

// 	v2 += v1
// 	v1 = v1<<17 | v1>>(64-17)
// 	v1 ^= v2
// 	v2 = v2<<32 | v2>>(64-32)

// 	return v0 ^ v1 ^ v2 ^ v3
// }

// // Seed sets a new secret seed for PRNG.
// func (p *siprng) Seed(k [16]byte) {
// 	p.k0 = binary.LittleEndian.Uint64(k[0:8])
// 	p.k1 = binary.LittleEndian.Uint64(k[8:16])
// 	p.ctr = 1
// }

// // Uint64 returns a new pseudorandom uint64.
// func (p *siprng) Uint64() uint64 {
// 	v := siphash(p.k0, p.k1, p.ctr)
// 	p.ctr++
// 	return v
// }

// func (p *siprng) Bytes(n int) []byte {
// 	// Since we don't have a buffer for generated bytes in siprng state,
// 	// we just generate enough 8-byte blocks and then cut the result to the
// 	// required length. Doing it this way, we lose generated bytes, and we
// 	// don't get the strictly sequential deterministic output from PRNG:
// 	// calling Uint64() and then Bytes(3) produces different output than
// 	// when calling them in the reverse order, but for our applications
// 	// this is OK.
// 	numBlocks := (n + 8 - 1) / 8
// 	b := make([]byte, numBlocks*8)
// 	for i := 0; i < len(b); i += 8 {
// 		binary.LittleEndian.PutUint64(b[i:], p.Uint64())
// 	}
// 	return b[:n]
// }

// func (p *siprng) Int63() int64 {
// 	return int64(p.Uint64() & 0x7fffffffffffffff)
// }

// func (p *siprng) Uint32() uint32 {
// 	return uint32(p.Uint64())
// }

// func (p *siprng) Int31() int32 {
// 	return int32(p.Uint32() & 0x7fffffff)
// }

// func (p *siprng) Intn(n int) int {
// 	if n <= 0 {
// 		panic("invalid argument to Intn")
// 	}
// 	if n <= 1<<31-1 {
// 		return int(p.Int31n(int32(n)))
// 	}
// 	return int(p.Int63n(int64(n)))
// }

// func (p *siprng) Int63n(n int64) int64 {
// 	if n <= 0 {
// 		panic("invalid argument to Int63n")
// 	}
// 	max := int64((1 << 63) - 1 - (1<<63)%uint64(n))
// 	v := p.Int63()
// 	for v > max {
// 		v = p.Int63()
// 	}
// 	return v % n
// }

// func (p *siprng) Int31n(n int32) int32 {
// 	if n <= 0 {
// 		panic("invalid argument to Int31n")
// 	}
// 	max := int32((1 << 31) - 1 - (1<<31)%uint32(n))
// 	v := p.Int31()
// 	for v > max {
// 		v = p.Int31()
// 	}
// 	return v % n
// }

// func (p *siprng) Float64() float64 { return float64(p.Int63()) / (1 << 63) }

// // Int returns a pseudorandom int in range [from, to].
// func (p *siprng) Int(from, to int) int {
// 	return p.Intn(to+1-from) + from
// }

// // Float returns a pseudorandom float64 in range [from, to].
// func (p *siprng) Float(from, to float64) float64 {
// 	return (to-from)*p.Float64() + from
// }
