package __obf_eb5e805b9fc230e3

import (
	"fmt"
	"io"
)

// urlFormat := fmt.Sprintf("http://localhost/segment/c7ghedehscu246733tl0/%d/%d\n", Resolution, SegmentIndex)
func (__obf_f3546e71ce2e1d63 *HlsServer) GetPlaylist(__obf_d2dde7e169684a70 string, __obf_6944cbfcac39b37e float32, __obf_1cab0e86de8129e5 io.Writer) {
	fmt.Fprint(__obf_1cab0e86de8129e5, "#EXTM3U\n")
	fmt.Fprint(__obf_1cab0e86de8129e5, "#EXT-X-VERSION:3\n")
	fmt.Fprint(__obf_1cab0e86de8129e5, "#EXT-X-MEDIA-SEQUENCE:0\n")
	fmt.Fprint(__obf_1cab0e86de8129e5, "#EXT-X-ALLOW-CACHE:YES\n")
	fmt.Fprintf(__obf_1cab0e86de8129e5, "#EXT-X-TARGETDURATION:%.f\n", __obf_f3546e71ce2e1d63.Option.SegmentLen)
	fmt.Fprint(__obf_1cab0e86de8129e5, "#EXT-X-PLAYLIST-TYPE:VOD\n")
	__obf_ecf021d13e25bea6 := 0

	for __obf_6944cbfcac39b37e > 0 {
		if __obf_6944cbfcac39b37e > __obf_f3546e71ce2e1d63.Option.SegmentLen {
			fmt.Fprintf(__obf_1cab0e86de8129e5, "#EXTINF: %.3f,\n", __obf_f3546e71ce2e1d63.Option.SegmentLen)
		} else {
			fmt.Fprintf(__obf_1cab0e86de8129e5, "#EXTINF: %.3f,\n", __obf_6944cbfcac39b37e)
		}
		fmt.Fprintf(__obf_1cab0e86de8129e5, __obf_d2dde7e169684a70, __obf_ecf021d13e25bea6)
		__obf_ecf021d13e25bea6++
		__obf_6944cbfcac39b37e = __obf_6944cbfcac39b37e - __obf_f3546e71ce2e1d63.Option.SegmentLen
	}
	fmt.Fprint(__obf_1cab0e86de8129e5, "#EXT-X-ENDLIST\n")
}
