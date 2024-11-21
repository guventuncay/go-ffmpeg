package main

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {
	//ffmpeg -i sample_data/input.mp4 -i sample_data/music.mp3 -c:v copy -map 0:v -map 1:a -shortest sample_data/output_with_music.mp4 -y
	err := ffmpeg.
		Input("sample_data/music.mp3", ffmpeg.KwArgs{
			"i": "sample_data/input.mp4"}).
		Output("sample_data/output_with_music.mp4", ffmpeg.KwArgs{
			"c:v":      "copy",
			"map":      []string{"0:v", "1:a"},
			"shortest": ""}).
		OverWriteOutput().ErrorToStdOut().Run()

	if err != nil {
		panic(err)
	}
}

//ffmpeg -i sample_data/music.mp3 -i sample_data/input.mp4 -c:v copy -map 0:v -map 1:a -shortest sample_data/output_with_music.mp4 -y
//ffmpeg -i sample_data/input.mp4 -i sample_data/music.mp3 -c:v copy -map 0:v -map 1:a -shortest -y sample_data/output_with_music.mp4
