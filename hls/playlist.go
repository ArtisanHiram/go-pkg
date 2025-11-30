package __obf_08ee9322ff6adb83

import (
	"fmt"
	"io"
)

// urlFormat := fmt.Sprintf("http://localhost/segment/c7ghedehscu246733tl0/%d/%d\n", Resolution, SegmentIndex)
func (__obf_504413fedee71524 *HlsServer) GetPlaylist(__obf_7af77bd167c09009 string, __obf_f2c85f285cd03b85 float32, __obf_63e18a6991b3d10e io.Writer) {
	fmt.Fprint(__obf_63e18a6991b3d10e, "#EXTM3U\n")
	fmt.Fprint(__obf_63e18a6991b3d10e, "#EXT-X-VERSION:3\n")
	fmt.Fprint(__obf_63e18a6991b3d10e, "#EXT-X-MEDIA-SEQUENCE:0\n")
	fmt.Fprint(__obf_63e18a6991b3d10e, "#EXT-X-ALLOW-CACHE:YES\n")
	fmt.Fprintf(__obf_63e18a6991b3d10e, "#EXT-X-TARGETDURATION:%.f\n", __obf_504413fedee71524.Option.SegmentLen)
	fmt.Fprint(__obf_63e18a6991b3d10e, "#EXT-X-PLAYLIST-TYPE:VOD\n")
	__obf_3f66d1e8f6ee15d8 := 0

	for __obf_f2c85f285cd03b85 > 0 {
		if __obf_f2c85f285cd03b85 > __obf_504413fedee71524.Option.SegmentLen {
			fmt.Fprintf(__obf_63e18a6991b3d10e, "#EXTINF: %.3f,\n", __obf_504413fedee71524.Option.SegmentLen)
		} else {
			fmt.Fprintf(__obf_63e18a6991b3d10e, "#EXTINF: %.3f,\n", __obf_f2c85f285cd03b85)
		}
		fmt.Fprintf(__obf_63e18a6991b3d10e, __obf_7af77bd167c09009, __obf_3f66d1e8f6ee15d8)
		__obf_3f66d1e8f6ee15d8++
		__obf_f2c85f285cd03b85 = __obf_f2c85f285cd03b85 - __obf_504413fedee71524.Option.SegmentLen
	}
	fmt.Fprint(__obf_63e18a6991b3d10e, "#EXT-X-ENDLIST\n")
}
