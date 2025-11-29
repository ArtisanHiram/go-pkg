package __obf_6ff082bd539c7df0

import (
	"fmt"
	"io"
)

// urlFormat := fmt.Sprintf("http://localhost/segment/c7ghedehscu246733tl0/%d/%d\n", Resolution, SegmentIndex)
func (__obf_dd5f757557da2c8e *HlsServer) GetPlaylist(__obf_da10118bd6f20f01 string, __obf_0b4264c722fa95a5 float32, __obf_42061c65b3e92c28 io.Writer) {
	fmt.Fprint(__obf_42061c65b3e92c28, "#EXTM3U\n")
	fmt.Fprint(__obf_42061c65b3e92c28, "#EXT-X-VERSION:3\n")
	fmt.Fprint(__obf_42061c65b3e92c28, "#EXT-X-MEDIA-SEQUENCE:0\n")
	fmt.Fprint(__obf_42061c65b3e92c28, "#EXT-X-ALLOW-CACHE:YES\n")
	fmt.Fprintf(__obf_42061c65b3e92c28, "#EXT-X-TARGETDURATION:%.f\n", __obf_dd5f757557da2c8e.Option.SegmentLen)
	fmt.Fprint(__obf_42061c65b3e92c28, "#EXT-X-PLAYLIST-TYPE:VOD\n")
	__obf_53fef868f8b12c3c := 0

	for __obf_0b4264c722fa95a5 > 0 {
		if __obf_0b4264c722fa95a5 > __obf_dd5f757557da2c8e.Option.SegmentLen {
			fmt.Fprintf(__obf_42061c65b3e92c28, "#EXTINF: %.3f,\n", __obf_dd5f757557da2c8e.Option.SegmentLen)
		} else {
			fmt.Fprintf(__obf_42061c65b3e92c28, "#EXTINF: %.3f,\n", __obf_0b4264c722fa95a5)
		}
		fmt.Fprintf(__obf_42061c65b3e92c28, __obf_da10118bd6f20f01, __obf_53fef868f8b12c3c)
		__obf_53fef868f8b12c3c++
		__obf_0b4264c722fa95a5 = __obf_0b4264c722fa95a5 - __obf_dd5f757557da2c8e.Option.SegmentLen
	}
	fmt.Fprint(__obf_42061c65b3e92c28, "#EXT-X-ENDLIST\n")
}
