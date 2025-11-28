package __obf_acea4ab24a824c18

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

func (__obf_c971f1d170edee24 *HlsServer) GetSegment(__obf_944727d4e31f21db string, __obf_36a58ee8d6cc5b0b, __obf_7922e1e47dcc332e int, __obf_7bfc44ad58c48031 io.Writer) error {
	__obf_10074c180b084d32 := float32(__obf_36a58ee8d6cc5b0b) * __obf_c971f1d170edee24.Option.SegmentLen
	// see http://superuser.com/questions/908280/what-is-the-correct-way-to-fix-keyframes-in-ffmpeg-for-dash
	__obf_4e7062996d563de2 := []string{
		// Prevent encoding to run longer than 30 seonds
		"-timelimit", "45",

		// TODO: Some stuff to investigate
		// "-probesize", "524288",
		// "-fpsprobesize", "10",
		// "-analyzeduration", "2147483647",
		// "-hwaccel:0", "vda",

		// The start time
		// important: needs to be before -i to do input seeking
		"-ss", fmt.Sprintf("%.2f", __obf_10074c180b084d32),

		// The source file
		"-i", __obf_944727d4e31f21db,

		// Put all streams to output
		// "-map", "0",

		// The duration
		"-t", fmt.Sprintf("%.2f", __obf_c971f1d170edee24.Option.SegmentLen),

		// TODO: Find out what it does
		//"-strict", "-2",

		// Synchronize audio
		"-async", "1",

		// 720p
		"-vf", fmt.Sprintf("scale=-2:%d", __obf_7922e1e47dcc332e),

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

		"-force_key_frames", fmt.Sprintf("expr:gte(t,n_forced*%.2f)", __obf_c971f1d170edee24.Option.SegmentLen),

		//"-force_key_frames", "00:00:00.00",
		//"-x264opts", "keyint=25:min-keyint=25:scenecut=-1",

		//"-f", "mpegts",

		"-f", "ssegment",
		"-segment_time", fmt.Sprintf("%.2f", __obf_c971f1d170edee24.Option.SegmentLen),
		"-initial_offset", fmt.Sprintf("%.2f", __obf_10074c180b084d32),

		"pipe:out%03d.ts",
	}

	return __obf_c971f1d170edee24.__obf_4abf18d2926be19a.Serve(FFMpegPath, __obf_4e7062996d563de2, __obf_7bfc44ad58c48031)
}
