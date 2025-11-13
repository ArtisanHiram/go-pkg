package __obf_85f036759f76ec38

// ---------------------------------
// siprng is needed for random functions, keeping it from the original code.
// ---------------------------------
type __obf_a7a546a1a45d28a2 struct {
	__obf_5f57dfa9afa6e3a6, __obf_d607758cdfecc392, __obf_93885b4d59667b69, __obf_285cb1d79879be67 uint64
}

const (
	__obf_ca1e2db74bde116b = 2
	__obf_950d487fe9fc3a70 = 4
)

func (__obf_9f0cd3415eba4a85 *__obf_a7a546a1a45d28a2) Seed(__obf_c0779e1a8dbd2667 [16]byte) {
	__obf_74afddaecb9ed990 := uint64(__obf_c0779e1a8dbd2667[0]) | uint64(__obf_c0779e1a8dbd2667[1])<<8 | uint64(__obf_c0779e1a8dbd2667[2])<<16 |
		uint64(__obf_c0779e1a8dbd2667[3])<<24 | uint64(__obf_c0779e1a8dbd2667[4])<<32 | uint64(__obf_c0779e1a8dbd2667[5])<<40 |
		uint64(__obf_c0779e1a8dbd2667[6])<<48 | uint64(__obf_c0779e1a8dbd2667[7])<<56
	__obf_c5185f2bb73c108d := uint64(__obf_c0779e1a8dbd2667[8]) | uint64(__obf_c0779e1a8dbd2667[9])<<8 | uint64(__obf_c0779e1a8dbd2667[10])<<16 |
		uint64(__obf_c0779e1a8dbd2667[11])<<24 | uint64(__obf_c0779e1a8dbd2667[12])<<32 | uint64(__obf_c0779e1a8dbd2667[13])<<40 |
		uint64(__obf_c0779e1a8dbd2667[14])<<48 | uint64(__obf_c0779e1a8dbd2667[15])<<56
	__obf_9f0cd3415eba4a85.__obf_5f57dfa9afa6e3a6 = __obf_74afddaecb9ed990 ^ 0x736f6d6570736575
	__obf_9f0cd3415eba4a85.__obf_d607758cdfecc392 = __obf_c5185f2bb73c108d ^ 0x646f72616e646f6d
	__obf_9f0cd3415eba4a85.__obf_93885b4d59667b69 = __obf_74afddaecb9ed990 ^ 0x6c7967656e657261
	__obf_9f0cd3415eba4a85.__obf_285cb1d79879be67 = __obf_c5185f2bb73c108d ^ 0x7465646279746573
}

func (__obf_9f0cd3415eba4a85 *__obf_a7a546a1a45d28a2) Uint64() uint64 {
	__obf_9f0cd3415eba4a85.__obf_285cb1d79879be67++
	for range __obf_ca1e2db74bde116b {
		__obf_9f0cd3415eba4a85.__obf_5f57dfa9afa6e3a6 += __obf_9f0cd3415eba4a85.__obf_d607758cdfecc392
		__obf_9f0cd3415eba4a85.__obf_d607758cdfecc392 = __obf_11fd70442e1560c8(__obf_9f0cd3415eba4a85.__obf_d607758cdfecc392, 13)
		__obf_9f0cd3415eba4a85.__obf_d607758cdfecc392 ^= __obf_9f0cd3415eba4a85.__obf_5f57dfa9afa6e3a6
		__obf_9f0cd3415eba4a85.__obf_5f57dfa9afa6e3a6 = __obf_11fd70442e1560c8(__obf_9f0cd3415eba4a85.__obf_5f57dfa9afa6e3a6, 32)
		__obf_9f0cd3415eba4a85.__obf_93885b4d59667b69 += __obf_9f0cd3415eba4a85.__obf_285cb1d79879be67
		__obf_9f0cd3415eba4a85.__obf_285cb1d79879be67 = __obf_11fd70442e1560c8(__obf_9f0cd3415eba4a85.__obf_285cb1d79879be67, 16)
		__obf_9f0cd3415eba4a85.__obf_285cb1d79879be67 ^= __obf_9f0cd3415eba4a85.__obf_93885b4d59667b69
		__obf_9f0cd3415eba4a85.__obf_93885b4d59667b69 += __obf_9f0cd3415eba4a85.__obf_d607758cdfecc392
		__obf_9f0cd3415eba4a85.__obf_d607758cdfecc392 = __obf_11fd70442e1560c8(__obf_9f0cd3415eba4a85.__obf_d607758cdfecc392, 17)
		__obf_9f0cd3415eba4a85.__obf_d607758cdfecc392 ^= __obf_9f0cd3415eba4a85.__obf_93885b4d59667b69
		__obf_9f0cd3415eba4a85.__obf_93885b4d59667b69 = __obf_11fd70442e1560c8(__obf_9f0cd3415eba4a85.__obf_93885b4d59667b69, 32)
		__obf_9f0cd3415eba4a85.__obf_285cb1d79879be67 += __obf_9f0cd3415eba4a85.__obf_5f57dfa9afa6e3a6
		__obf_9f0cd3415eba4a85.__obf_5f57dfa9afa6e3a6 = __obf_11fd70442e1560c8(__obf_9f0cd3415eba4a85.__obf_5f57dfa9afa6e3a6, 21)
		__obf_9f0cd3415eba4a85.__obf_5f57dfa9afa6e3a6 ^= __obf_9f0cd3415eba4a85.__obf_285cb1d79879be67
	}
	__obf_2370489ba64bc277 := __obf_9f0cd3415eba4a85.__obf_5f57dfa9afa6e3a6 ^ __obf_9f0cd3415eba4a85.__obf_d607758cdfecc392 ^ __obf_9f0cd3415eba4a85.__obf_93885b4d59667b69 ^ __obf_9f0cd3415eba4a85.__obf_285cb1d79879be67
	for range __obf_950d487fe9fc3a70 {
		__obf_9f0cd3415eba4a85.__obf_285cb1d79879be67++
	}
	return __obf_2370489ba64bc277
}

func (__obf_9f0cd3415eba4a85 *__obf_a7a546a1a45d28a2) Float64() float64 {
	return float64(__obf_9f0cd3415eba4a85.Uint64()>>11) / (1 << 53)
}

func (__obf_9f0cd3415eba4a85 *__obf_a7a546a1a45d28a2) Intn(__obf_23131025729f7775 int) int {
	return int(__obf_9f0cd3415eba4a85.Uint64() % uint64(__obf_23131025729f7775))
}

func __obf_11fd70442e1560c8(__obf_24ee2b3ca6532f15 uint64, __obf_df540af270232b3d int) uint64 {
	return (__obf_24ee2b3ca6532f15 << __obf_df540af270232b3d) | (__obf_24ee2b3ca6532f15 >> (64 - __obf_df540af270232b3d))
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
