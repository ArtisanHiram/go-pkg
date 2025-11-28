package __obf_5028a0a829ddcdab

import (
	"fmt"
	"io"
)

// urlFormat := fmt.Sprintf("http://localhost/segment/c7ghedehscu246733tl0/%d/%d\n", Resolution, SegmentIndex)
func (__obf_883d14d3bb9e62d2 *HlsServer) GetPlaylist(__obf_a25fb153d77b2e8d string, __obf_428fef9eb22378e9 float32, __obf_f0018e81daaf1fa9 io.Writer) {
	fmt.Fprint(__obf_f0018e81daaf1fa9, "#EXTM3U\n")
	fmt.Fprint(__obf_f0018e81daaf1fa9, "#EXT-X-VERSION:3\n")
	fmt.Fprint(__obf_f0018e81daaf1fa9, "#EXT-X-MEDIA-SEQUENCE:0\n")
	fmt.Fprint(__obf_f0018e81daaf1fa9, "#EXT-X-ALLOW-CACHE:YES\n")
	fmt.Fprintf(__obf_f0018e81daaf1fa9, "#EXT-X-TARGETDURATION:%.f\n", __obf_883d14d3bb9e62d2.Option.SegmentLen)
	fmt.Fprint(__obf_f0018e81daaf1fa9, "#EXT-X-PLAYLIST-TYPE:VOD\n")

	__obf_01d9bf9e519afd77 := 0

	for __obf_428fef9eb22378e9 > 0 {
		if __obf_428fef9eb22378e9 > __obf_883d14d3bb9e62d2.Option.SegmentLen {
			fmt.Fprintf(__obf_f0018e81daaf1fa9, "#EXTINF: %.3f,\n", __obf_883d14d3bb9e62d2.Option.SegmentLen)
		} else {
			fmt.Fprintf(__obf_f0018e81daaf1fa9, "#EXTINF: %.3f,\n", __obf_428fef9eb22378e9)
		}
		fmt.Fprintf(__obf_f0018e81daaf1fa9, __obf_a25fb153d77b2e8d, __obf_01d9bf9e519afd77)
		__obf_01d9bf9e519afd77++
		__obf_428fef9eb22378e9 = __obf_428fef9eb22378e9 - __obf_883d14d3bb9e62d2.Option.SegmentLen
	}
	fmt.Fprint(__obf_f0018e81daaf1fa9, "#EXT-X-ENDLIST\n")
}
