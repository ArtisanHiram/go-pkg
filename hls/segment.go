package __obf_f60326fd90eb13d9

import (
	"fmt"
	"io"
)

// var encodeWorker = NewWorker(
// 	2, filepath.Join(".hls", "segments"),
// 	func(i any, w io.Writer) error {
// 		return nil
// 	},
// )

func (__obf_dd04e55ec17d0ebf *HlsServer) GetSegment(__obf_2c7a65b2dac56e8a string, __obf_5326656dca618466, __obf_7c59036812aec7d9 int, __obf_811b158c28965ee0 io.Writer) error {
	__obf_de70286adac56f34 := float32(__obf_5326656dca618466) * __obf_dd04e55ec17d0ebf.Option.SegmentLen
	__obf_4790050703cb3251 := // see http://superuser.com/questions/908280/what-is-the-correct-way-to-fix-keyframes-in-ffmpeg-for-dash
		[]string{
			// Prevent encoding to run longer than 30 seonds
			"-timelimit", "45",

			// TODO: Some stuff to investigate
			// "-probesize", "524288",
			// "-fpsprobesize", "10",
			// "-analyzeduration", "2147483647",
			// "-hwaccel:0", "vda",

			// The start time
			// important: needs to be before -i to do input seeking
			"-ss", fmt.Sprintf("%.2f", __obf_de70286adac56f34),

			// The source file
			"-i", __obf_2c7a65b2dac56e8a, // Put all streams to output
			// "-map", "0",

			// The duration
			"-t", fmt.Sprintf("%.2f", __obf_dd04e55ec17d0ebf.Option.SegmentLen),

			// TODO: Find out what it does
			//"-strict", "-2",

			// Synchronize audio
			"-async", "1",

			// 720p
			"-vf", fmt.Sprintf("scale=-2:%d", __obf_7c59036812aec7d9),

			// x264 video codec
			"-vcodec", "libx264",

			// x264 preset
			"-preset", "veryfast",

			// aac audio codec
			"-c:a", "aac",
			"-b:a", "128k",
			"-ac", "2",

			// TODO
			"-pix_fmt", "yuv420p",

			//"-r", "25", // fixed framerate

			"-force_key_frames", fmt.Sprintf("expr:gte(t,n_forced*%.2f)", __obf_dd04e55ec17d0ebf.Option.SegmentLen),

			//"-force_key_frames", "00:00:00.00",
			//"-x264opts", "keyint=25:min-keyint=25:scenecut=-1",

			//"-f", "mpegts",

			"-f", "ssegment",
			"-segment_time", fmt.Sprintf("%.2f", __obf_dd04e55ec17d0ebf.Option.SegmentLen),
			"-initial_offset", fmt.Sprintf("%.2f", __obf_de70286adac56f34),

			"pipe:out%03d.ts",
		}

	return __obf_dd04e55ec17d0ebf.__obf_69915dd1bdb5a610.Serve(FFMpegPath, __obf_4790050703cb3251, __obf_811b158c28965ee0)
}
