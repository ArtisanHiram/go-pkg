package __obf_dc577bef9ac5c6b8

import (
	"fmt"
	"io"
)

// urlFormat := fmt.Sprintf("http://localhost/segment/c7ghedehscu246733tl0/%d/%d\n", Resolution, SegmentIndex)
func (__obf_fdd7931bd5a61fac *HlsServer) GetPlaylist(__obf_2050259102ce41ec string, __obf_38c3f1e0683291fc float32, __obf_39c0ed11835c4488 io.Writer) {
	fmt.Fprint(__obf_39c0ed11835c4488, "#EXTM3U\n")
	fmt.Fprint(__obf_39c0ed11835c4488, "#EXT-X-VERSION:3\n")
	fmt.Fprint(__obf_39c0ed11835c4488, "#EXT-X-MEDIA-SEQUENCE:0\n")
	fmt.Fprint(__obf_39c0ed11835c4488, "#EXT-X-ALLOW-CACHE:YES\n")
	fmt.Fprintf(__obf_39c0ed11835c4488, "#EXT-X-TARGETDURATION:%.f\n", __obf_fdd7931bd5a61fac.Option.SegmentLen)
	fmt.Fprint(__obf_39c0ed11835c4488, "#EXT-X-PLAYLIST-TYPE:VOD\n")

	__obf_a6dbafe8c2323258 := 0

	for __obf_38c3f1e0683291fc > 0 {
		if __obf_38c3f1e0683291fc > __obf_fdd7931bd5a61fac.Option.SegmentLen {
			fmt.Fprintf(__obf_39c0ed11835c4488, "#EXTINF: %.3f,\n", __obf_fdd7931bd5a61fac.Option.SegmentLen)
		} else {
			fmt.Fprintf(__obf_39c0ed11835c4488, "#EXTINF: %.3f,\n", __obf_38c3f1e0683291fc)
		}
		fmt.Fprintf(__obf_39c0ed11835c4488, __obf_2050259102ce41ec, __obf_a6dbafe8c2323258)
		__obf_a6dbafe8c2323258++
		__obf_38c3f1e0683291fc = __obf_38c3f1e0683291fc - __obf_fdd7931bd5a61fac.Option.SegmentLen
	}
	fmt.Fprint(__obf_39c0ed11835c4488, "#EXT-X-ENDLIST\n")
}
