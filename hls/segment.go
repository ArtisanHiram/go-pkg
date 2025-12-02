package __obf_eb5e805b9fc230e3

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

func (__obf_f3546e71ce2e1d63 *HlsServer) GetSegment(__obf_22d6c43ac51329e5 string, __obf_7aa7d6d9b855386d, __obf_6952caa782f6503d int, __obf_1cab0e86de8129e5 io.Writer) error {
	__obf_0fa3eddf76bd5045 := float32(__obf_7aa7d6d9b855386d) * __obf_f3546e71ce2e1d63.Option.SegmentLen
	__obf_9ddc2c49f56f7ff3 := // see http://superuser.com/questions/908280/what-is-the-correct-way-to-fix-keyframes-in-ffmpeg-for-dash
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
			"-ss", fmt.Sprintf("%.2f", __obf_0fa3eddf76bd5045),

			// The source file
			"-i", __obf_22d6c43ac51329e5, // Put all streams to output
			// "-map", "0",

			// The duration
			"-t", fmt.Sprintf("%.2f", __obf_f3546e71ce2e1d63.Option.SegmentLen),

			// TODO: Find out what it does
			//"-strict", "-2",

			// Synchronize audio
			"-async", "1",

			// 720p
			"-vf", fmt.Sprintf("scale=-2:%d", __obf_6952caa782f6503d),

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

			"-force_key_frames", fmt.Sprintf("expr:gte(t,n_forced*%.2f)", __obf_f3546e71ce2e1d63.Option.SegmentLen),

			//"-force_key_frames", "00:00:00.00",
			//"-x264opts", "keyint=25:min-keyint=25:scenecut=-1",

			//"-f", "mpegts",

			"-f", "ssegment",
			"-segment_time", fmt.Sprintf("%.2f", __obf_f3546e71ce2e1d63.Option.SegmentLen),
			"-initial_offset", fmt.Sprintf("%.2f", __obf_0fa3eddf76bd5045),

			"pipe:out%03d.ts",
		}

	return __obf_f3546e71ce2e1d63.__obf_1dd73b629d0430c7.Serve(FFMpegPath, __obf_9ddc2c49f56f7ff3, __obf_1cab0e86de8129e5)
}
