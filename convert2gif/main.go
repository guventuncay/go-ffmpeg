package main

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {

	err := ffmpeg.Input("sample_data/input.mp4", ffmpeg.KwArgs{"ss": "1"}).
		Output("sample_data/output.gif", ffmpeg.KwArgs{"s": "320x240", "pix_fmt": "rgb24", "t": "3", "r": "3"}).
		OverWriteOutput().ErrorToStdOut().Run()

	if err != nil {
		panic(err)
	}
}
