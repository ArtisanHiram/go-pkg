package __obf_5441fcd9a319cf59

import (
	"fmt"
	"io"
)

// urlFormat := fmt.Sprintf("http://localhost/segment/c7ghedehscu246733tl0/%d/%d\n", Resolution, SegmentIndex)
func (__obf_bf443cef12bfef60 *HlsServer) GetPlaylist(__obf_769a794ca28221f1 string, __obf_94dd6b9a9338ee98 float32, __obf_9e22338370faf798 io.Writer) {
	fmt.Fprint(__obf_9e22338370faf798, "#EXTM3U\n")
	fmt.Fprint(__obf_9e22338370faf798, "#EXT-X-VERSION:3\n")
	fmt.Fprint(__obf_9e22338370faf798, "#EXT-X-MEDIA-SEQUENCE:0\n")
	fmt.Fprint(__obf_9e22338370faf798, "#EXT-X-ALLOW-CACHE:YES\n")
	fmt.Fprintf(__obf_9e22338370faf798, "#EXT-X-TARGETDURATION:%.f\n", __obf_bf443cef12bfef60.Option.SegmentLen)
	fmt.Fprint(__obf_9e22338370faf798, "#EXT-X-PLAYLIST-TYPE:VOD\n")
	__obf_5c639997e0c33563 := 0

	for __obf_94dd6b9a9338ee98 > 0 {
		if __obf_94dd6b9a9338ee98 > __obf_bf443cef12bfef60.Option.SegmentLen {
			fmt.Fprintf(__obf_9e22338370faf798, "#EXTINF: %.3f,\n", __obf_bf443cef12bfef60.Option.SegmentLen)
		} else {
			fmt.Fprintf(__obf_9e22338370faf798, "#EXTINF: %.3f,\n", __obf_94dd6b9a9338ee98)
		}
		fmt.Fprintf(__obf_9e22338370faf798, __obf_769a794ca28221f1, __obf_5c639997e0c33563)
		__obf_5c639997e0c33563++
		__obf_94dd6b9a9338ee98 = __obf_94dd6b9a9338ee98 - __obf_bf443cef12bfef60.Option.SegmentLen
	}
	fmt.Fprint(__obf_9e22338370faf798, "#EXT-X-ENDLIST\n")
}
