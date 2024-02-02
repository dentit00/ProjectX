package main

import (
	"log"
	"os"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func main() {
	filePath := os.Args                  // define the variable which contains filepath, using command argument
	err := ffmpeg_go.Input(filePath[1]). // specialize the input
						Output("output1.mp3").  // specialize the output file
						OverWriteOutput().Run() // overwrite the output if it already exists
	if err != nil {
		log.Fatalln(err)
	}

}
