package __obf_5028a0a829ddcdab

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

func (__obf_883d14d3bb9e62d2 *HlsServer) GetSegment(__obf_8ddbcc776441ebf5 string, __obf_553b07955299cf31, __obf_6f8081e033cbdc4a int, __obf_f0018e81daaf1fa9 io.Writer) error {
	__obf_386c4f77dd2b9648 := float32(__obf_553b07955299cf31) * __obf_883d14d3bb9e62d2.Option.SegmentLen
	// see http://superuser.com/questions/908280/what-is-the-correct-way-to-fix-keyframes-in-ffmpeg-for-dash
	__obf_76b5a3882b8f075c := []string{
		// Prevent encoding to run longer than 30 seonds
		"-timelimit", "45",

		// TODO: Some stuff to investigate
		// "-probesize", "524288",
		// "-fpsprobesize", "10",
		// "-analyzeduration", "2147483647",
		// "-hwaccel:0", "vda",

		// The start time
		// important: needs to be before -i to do input seeking
		"-ss", fmt.Sprintf("%.2f", __obf_386c4f77dd2b9648),

		// The source file
		"-i", __obf_8ddbcc776441ebf5,

		// Put all streams to output
		// "-map", "0",

		// The duration
		"-t", fmt.Sprintf("%.2f", __obf_883d14d3bb9e62d2.Option.SegmentLen),

		// TODO: Find out what it does
		//"-strict", "-2",

		// Synchronize audio
		"-async", "1",

		// 720p
		"-vf", fmt.Sprintf("scale=-2:%d", __obf_6f8081e033cbdc4a),

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

		"-force_key_frames", fmt.Sprintf("expr:gte(t,n_forced*%.2f)", __obf_883d14d3bb9e62d2.Option.SegmentLen),

		//"-force_key_frames", "00:00:00.00",
		//"-x264opts", "keyint=25:min-keyint=25:scenecut=-1",

		//"-f", "mpegts",

		"-f", "ssegment",
		"-segment_time", fmt.Sprintf("%.2f", __obf_883d14d3bb9e62d2.Option.SegmentLen),
		"-initial_offset", fmt.Sprintf("%.2f", __obf_386c4f77dd2b9648),

		"pipe:out%03d.ts",
	}

	return __obf_883d14d3bb9e62d2.__obf_afca1d674bf7806a.Serve(FFMpegPath, __obf_76b5a3882b8f075c, __obf_f0018e81daaf1fa9)
}
