package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	filename := "http://www.xxxxxxx.com/2021.mp4"
	width := 640
	height := 360
	cmd := exec.Command("ffmpeg", "-i", filename, "-vframes", "10", "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	if cmd.Run() != nil {
		panic("could not generate frame")
	}

	ioutil.WriteFile("./output.jpg", []byte(buffer.String()), 0666)
}
