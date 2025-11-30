package __obf_f60326fd90eb13d9

import (
	"fmt"
	"io"
)

// urlFormat := fmt.Sprintf("http://localhost/segment/c7ghedehscu246733tl0/%d/%d\n", Resolution, SegmentIndex)
func (__obf_dd04e55ec17d0ebf *HlsServer) GetPlaylist(__obf_8dcda82388970844 string, __obf_9b0f57fb67097ecf float32, __obf_811b158c28965ee0 io.Writer) {
	fmt.Fprint(__obf_811b158c28965ee0, "#EXTM3U\n")
	fmt.Fprint(__obf_811b158c28965ee0, "#EXT-X-VERSION:3\n")
	fmt.Fprint(__obf_811b158c28965ee0, "#EXT-X-MEDIA-SEQUENCE:0\n")
	fmt.Fprint(__obf_811b158c28965ee0, "#EXT-X-ALLOW-CACHE:YES\n")
	fmt.Fprintf(__obf_811b158c28965ee0, "#EXT-X-TARGETDURATION:%.f\n", __obf_dd04e55ec17d0ebf.Option.SegmentLen)
	fmt.Fprint(__obf_811b158c28965ee0, "#EXT-X-PLAYLIST-TYPE:VOD\n")
	__obf_0b9bab109831a42d := 0

	for __obf_9b0f57fb67097ecf > 0 {
		if __obf_9b0f57fb67097ecf > __obf_dd04e55ec17d0ebf.Option.SegmentLen {
			fmt.Fprintf(__obf_811b158c28965ee0, "#EXTINF: %.3f,\n", __obf_dd04e55ec17d0ebf.Option.SegmentLen)
		} else {
			fmt.Fprintf(__obf_811b158c28965ee0, "#EXTINF: %.3f,\n", __obf_9b0f57fb67097ecf)
		}
		fmt.Fprintf(__obf_811b158c28965ee0, __obf_8dcda82388970844, __obf_0b9bab109831a42d)
		__obf_0b9bab109831a42d++
		__obf_9b0f57fb67097ecf = __obf_9b0f57fb67097ecf - __obf_dd04e55ec17d0ebf.Option.SegmentLen
	}
	fmt.Fprint(__obf_811b158c28965ee0, "#EXT-X-ENDLIST\n")
}
