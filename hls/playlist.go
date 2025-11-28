package __obf_acea4ab24a824c18

import (
	"fmt"
	"io"
)

// urlFormat := fmt.Sprintf("http://localhost/segment/c7ghedehscu246733tl0/%d/%d\n", Resolution, SegmentIndex)
func (__obf_c971f1d170edee24 *HlsServer) GetPlaylist(__obf_cc90f3fb43f00aba string, __obf_259e37abb748ba52 float32, __obf_7bfc44ad58c48031 io.Writer) {
	fmt.Fprint(__obf_7bfc44ad58c48031, "#EXTM3U\n")
	fmt.Fprint(__obf_7bfc44ad58c48031, "#EXT-X-VERSION:3\n")
	fmt.Fprint(__obf_7bfc44ad58c48031, "#EXT-X-MEDIA-SEQUENCE:0\n")
	fmt.Fprint(__obf_7bfc44ad58c48031, "#EXT-X-ALLOW-CACHE:YES\n")
	fmt.Fprintf(__obf_7bfc44ad58c48031, "#EXT-X-TARGETDURATION:%.f\n", __obf_c971f1d170edee24.Option.SegmentLen)
	fmt.Fprint(__obf_7bfc44ad58c48031, "#EXT-X-PLAYLIST-TYPE:VOD\n")

	__obf_0764188b2b362172 := 0

	for __obf_259e37abb748ba52 > 0 {
		if __obf_259e37abb748ba52 > __obf_c971f1d170edee24.Option.SegmentLen {
			fmt.Fprintf(__obf_7bfc44ad58c48031, "#EXTINF: %.3f,\n", __obf_c971f1d170edee24.Option.SegmentLen)
		} else {
			fmt.Fprintf(__obf_7bfc44ad58c48031, "#EXTINF: %.3f,\n", __obf_259e37abb748ba52)
		}
		fmt.Fprintf(__obf_7bfc44ad58c48031, __obf_cc90f3fb43f00aba, __obf_0764188b2b362172)
		__obf_0764188b2b362172++
		__obf_259e37abb748ba52 = __obf_259e37abb748ba52 - __obf_c971f1d170edee24.Option.SegmentLen
	}
	fmt.Fprint(__obf_7bfc44ad58c48031, "#EXT-X-ENDLIST\n")
}
