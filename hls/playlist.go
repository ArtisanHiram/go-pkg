package __obf_90dd9b56c0f1bd65

import (
	"fmt"
	"io"
)

// urlFormat := fmt.Sprintf("http://localhost/segment/c7ghedehscu246733tl0/%d/%d\n", Resolution, SegmentIndex)
func (__obf_073dee7a39e9dbe6 *HlsServer) GetPlaylist(__obf_9306c338b6a6c013 string, __obf_269e922f6c5f0020 float32, __obf_596cef7dccce398b io.Writer) {
	fmt.Fprint(__obf_596cef7dccce398b, "#EXTM3U\n")
	fmt.Fprint(__obf_596cef7dccce398b, "#EXT-X-VERSION:3\n")
	fmt.Fprint(__obf_596cef7dccce398b, "#EXT-X-MEDIA-SEQUENCE:0\n")
	fmt.Fprint(__obf_596cef7dccce398b, "#EXT-X-ALLOW-CACHE:YES\n")
	fmt.Fprintf(__obf_596cef7dccce398b, "#EXT-X-TARGETDURATION:%.f\n", __obf_073dee7a39e9dbe6.Option.SegmentLen)
	fmt.Fprint(__obf_596cef7dccce398b, "#EXT-X-PLAYLIST-TYPE:VOD\n")
	__obf_c5623bd28785fd95 := 0

	for __obf_269e922f6c5f0020 > 0 {
		if __obf_269e922f6c5f0020 > __obf_073dee7a39e9dbe6.Option.SegmentLen {
			fmt.Fprintf(__obf_596cef7dccce398b, "#EXTINF: %.3f,\n", __obf_073dee7a39e9dbe6.Option.SegmentLen)
		} else {
			fmt.Fprintf(__obf_596cef7dccce398b, "#EXTINF: %.3f,\n", __obf_269e922f6c5f0020)
		}
		fmt.Fprintf(__obf_596cef7dccce398b, __obf_9306c338b6a6c013, __obf_c5623bd28785fd95)
		__obf_c5623bd28785fd95++
		__obf_269e922f6c5f0020 = __obf_269e922f6c5f0020 - __obf_073dee7a39e9dbe6.Option.SegmentLen
	}
	fmt.Fprint(__obf_596cef7dccce398b, "#EXT-X-ENDLIST\n")
}
