package __obf_6ff082bd539c7df0

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

func (__obf_dd5f757557da2c8e *HlsServer) GetSegment(__obf_f6e6713da7daf44f string, __obf_e9447c7b5ab80dc5, __obf_b55259e8f8171da1 int, __obf_42061c65b3e92c28 io.Writer) error {
	__obf_03158a4f33944ebf := float32(__obf_e9447c7b5ab80dc5) * __obf_dd5f757557da2c8e.Option.SegmentLen
	__obf_d4c7dc8800d11aa1 := // see http://superuser.com/questions/908280/what-is-the-correct-way-to-fix-keyframes-in-ffmpeg-for-dash
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
			"-ss", fmt.Sprintf("%.2f", __obf_03158a4f33944ebf),

			// The source file
			"-i", __obf_f6e6713da7daf44f, // Put all streams to output
			// "-map", "0",

			// The duration
			"-t", fmt.Sprintf("%.2f", __obf_dd5f757557da2c8e.Option.SegmentLen),

			// TODO: Find out what it does
			//"-strict", "-2",

			// Synchronize audio
			"-async", "1",

			// 720p
			"-vf", fmt.Sprintf("scale=-2:%d", __obf_b55259e8f8171da1),

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

			"-force_key_frames", fmt.Sprintf("expr:gte(t,n_forced*%.2f)", __obf_dd5f757557da2c8e.Option.SegmentLen),

			//"-force_key_frames", "00:00:00.00",
			//"-x264opts", "keyint=25:min-keyint=25:scenecut=-1",

			//"-f", "mpegts",

			"-f", "ssegment",
			"-segment_time", fmt.Sprintf("%.2f", __obf_dd5f757557da2c8e.Option.SegmentLen),
			"-initial_offset", fmt.Sprintf("%.2f", __obf_03158a4f33944ebf),

			"pipe:out%03d.ts",
		}

	return __obf_dd5f757557da2c8e.__obf_309e545f7e826626.Serve(FFMpegPath, __obf_d4c7dc8800d11aa1, __obf_42061c65b3e92c28)
}
