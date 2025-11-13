package __obf_ab25ed2437cd567a

import (
	"fmt"
	"io"
)

// urlFormat := fmt.Sprintf("http://localhost/segment/c7ghedehscu246733tl0/%d/%d\n", Resolution, SegmentIndex)
func (__obf_8d3b3856596fc4f4 *HlsServer) GetPlaylist(__obf_01b79f16e0bda2ce string, __obf_bd4a31e6943ce6fe float32, __obf_98c2f72bee615664 io.Writer) {
	fmt.Fprint(__obf_98c2f72bee615664, "#EXTM3U\n")
	fmt.Fprint(__obf_98c2f72bee615664, "#EXT-X-VERSION:3\n")
	fmt.Fprint(__obf_98c2f72bee615664, "#EXT-X-MEDIA-SEQUENCE:0\n")
	fmt.Fprint(__obf_98c2f72bee615664, "#EXT-X-ALLOW-CACHE:YES\n")
	fmt.Fprintf(__obf_98c2f72bee615664, "#EXT-X-TARGETDURATION:%.f\n", __obf_8d3b3856596fc4f4.Option.SegmentLen)
	fmt.Fprint(__obf_98c2f72bee615664, "#EXT-X-PLAYLIST-TYPE:VOD\n")

	__obf_393a61ce09e5d489 := 0

	for __obf_bd4a31e6943ce6fe > 0 {
		if __obf_bd4a31e6943ce6fe > __obf_8d3b3856596fc4f4.Option.SegmentLen {
			fmt.Fprintf(__obf_98c2f72bee615664, "#EXTINF: %.3f,\n", __obf_8d3b3856596fc4f4.Option.SegmentLen)
		} else {
			fmt.Fprintf(__obf_98c2f72bee615664, "#EXTINF: %.3f,\n", __obf_bd4a31e6943ce6fe)
		}
		fmt.Fprintf(__obf_98c2f72bee615664, __obf_01b79f16e0bda2ce, __obf_393a61ce09e5d489)
		__obf_393a61ce09e5d489++
		__obf_bd4a31e6943ce6fe = __obf_bd4a31e6943ce6fe - __obf_8d3b3856596fc4f4.Option.SegmentLen
	}
	fmt.Fprint(__obf_98c2f72bee615664, "#EXT-X-ENDLIST\n")
}
