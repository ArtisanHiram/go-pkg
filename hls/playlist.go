package __obf_b28324f38df50634

import (
	"fmt"
	"io"
)

// urlFormat := fmt.Sprintf("http://localhost/segment/c7ghedehscu246733tl0/%d/%d\n", Resolution, SegmentIndex)
func (__obf_2dc28af59fafc546 *HlsServer) GetPlaylist(__obf_9932d810e516ab7f string, __obf_446eb74cabf8b530 float32, __obf_eb3c523b7bd5ff61 io.Writer) {
	fmt.Fprint(__obf_eb3c523b7bd5ff61, "#EXTM3U\n")
	fmt.Fprint(__obf_eb3c523b7bd5ff61, "#EXT-X-VERSION:3\n")
	fmt.Fprint(__obf_eb3c523b7bd5ff61, "#EXT-X-MEDIA-SEQUENCE:0\n")
	fmt.Fprint(__obf_eb3c523b7bd5ff61, "#EXT-X-ALLOW-CACHE:YES\n")
	fmt.Fprintf(__obf_eb3c523b7bd5ff61, "#EXT-X-TARGETDURATION:%.f\n", __obf_2dc28af59fafc546.Option.SegmentLen)
	fmt.Fprint(__obf_eb3c523b7bd5ff61, "#EXT-X-PLAYLIST-TYPE:VOD\n")

	__obf_87724c51d1d5de0e := 0

	for __obf_446eb74cabf8b530 > 0 {
		if __obf_446eb74cabf8b530 > __obf_2dc28af59fafc546.Option.SegmentLen {
			fmt.Fprintf(__obf_eb3c523b7bd5ff61, "#EXTINF: %.3f,\n", __obf_2dc28af59fafc546.Option.SegmentLen)
		} else {
			fmt.Fprintf(__obf_eb3c523b7bd5ff61, "#EXTINF: %.3f,\n", __obf_446eb74cabf8b530)
		}
		fmt.Fprintf(__obf_eb3c523b7bd5ff61, __obf_9932d810e516ab7f, __obf_87724c51d1d5de0e)
		__obf_87724c51d1d5de0e++
		__obf_446eb74cabf8b530 = __obf_446eb74cabf8b530 - __obf_2dc28af59fafc546.Option.SegmentLen
	}
	fmt.Fprint(__obf_eb3c523b7bd5ff61, "#EXT-X-ENDLIST\n")
}
