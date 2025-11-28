package __obf_42bbcad92b7de1a8

import (
	"fmt"
	"io"
)

// urlFormat := fmt.Sprintf("http://localhost/segment/c7ghedehscu246733tl0/%d/%d\n", Resolution, SegmentIndex)
func (__obf_c7be2ffa97bb9914 *HlsServer) GetPlaylist(__obf_68f38846823a618c string, __obf_e1336d8cc646c34c float32, __obf_39100c7ef6f93ec8 io.Writer) {
	fmt.Fprint(__obf_39100c7ef6f93ec8, "#EXTM3U\n")
	fmt.Fprint(__obf_39100c7ef6f93ec8, "#EXT-X-VERSION:3\n")
	fmt.Fprint(__obf_39100c7ef6f93ec8, "#EXT-X-MEDIA-SEQUENCE:0\n")
	fmt.Fprint(__obf_39100c7ef6f93ec8, "#EXT-X-ALLOW-CACHE:YES\n")
	fmt.Fprintf(__obf_39100c7ef6f93ec8, "#EXT-X-TARGETDURATION:%.f\n", __obf_c7be2ffa97bb9914.Option.SegmentLen)
	fmt.Fprint(__obf_39100c7ef6f93ec8, "#EXT-X-PLAYLIST-TYPE:VOD\n")

	__obf_68bea0cef28db5d8 := 0

	for __obf_e1336d8cc646c34c > 0 {
		if __obf_e1336d8cc646c34c > __obf_c7be2ffa97bb9914.Option.SegmentLen {
			fmt.Fprintf(__obf_39100c7ef6f93ec8, "#EXTINF: %.3f,\n", __obf_c7be2ffa97bb9914.Option.SegmentLen)
		} else {
			fmt.Fprintf(__obf_39100c7ef6f93ec8, "#EXTINF: %.3f,\n", __obf_e1336d8cc646c34c)
		}
		fmt.Fprintf(__obf_39100c7ef6f93ec8, __obf_68f38846823a618c, __obf_68bea0cef28db5d8)
		__obf_68bea0cef28db5d8++
		__obf_e1336d8cc646c34c = __obf_e1336d8cc646c34c - __obf_c7be2ffa97bb9914.Option.SegmentLen
	}
	fmt.Fprint(__obf_39100c7ef6f93ec8, "#EXT-X-ENDLIST\n")
}
