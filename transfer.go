// package main

// import (
// 	"log"
// 	"os"

// 	ffmpeg_go "github.com/u2takey/ffmpeg-go"
// )

//	func main() {
//		pathFile := os.Args
//		err := ffmpeg_go.Input("input.mp4").
//			Output(pathFile[1], ffmpeg_go.KwArgs{"c:a": "libvorbis"}).
//			OverWriteOutput().
//			Run()
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
package main

import (
	"log"
	"os"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func main() {
	filePath := os.Args
	input := ffmpeg_go.Input(filePath[1])
	audioStream := input.Audio()
	err := audioStream.Output("output.mp3").Run()
	if err != nil {
		log.Fatalln(err)
	}
} 
/*package main

import (
  "log"
  "os"

  ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func main() {
  filePath := os.Args
  err := ffmpeg_go.Input(filePath[1]).
    Output("output.mp3", ffmpeg_go.KwArgs{"vn": true, "acodec": "libopus", "ab": "32k", "ar": 24000}).
    OverWriteOutput().
    Run()
  if err != nil {
    log.Fatal(err)
  }
}