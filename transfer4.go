package main

import (
	"log"
	"os"
	"strconv"
	"time"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func convertion(arg string, index int) {
	num := strconv.Itoa(index)
	num = "output" + num + ".mp3"
	err := ffmpeg_go.Input(arg).
		Output(num).
		OverWriteOutput().Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	filePath := os.Args // define the variable which contains filepath, using command argument
	for i := 1; i < 4; i++ {
		go convertion(filePath[i], i)
	}
	time.Sleep(time.Second * 3)
}

// 	Тут решение через функцию без повторяющихся кусков кода
