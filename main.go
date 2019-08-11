// Package main provides
//
// File:  mian.go
// Author: ymiyamoto
//
// Created on Fri Aug  9 10:31:43 2019
//
package main

import (
	"fmt"
	"time"

	"gocv.io/x/gocv"
)

func isOpen(webcam *gocv.VideoCapture) bool {
	img := gocv.NewMat()
	defer img.Close()

	return webcam.Read(&img)
}

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Printf("can't open webcam.%+v", err)
	}
	defer webcam.Close()

	if !isOpen(webcam) {
		fmt.Printf("Cannot read device %v\n", 0)
	}

	videoName := fmt.Sprintf("video-%s.avi", time.Now().Format(time.RFC3339))
	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		fmt.Printf("cannot read device %v\n", 0)
	}

	writer, err := gocv.VideoWriterFile(videoName, "MJPG", 25, img.Cols(), img.Rows(), true)
	if err != nil {
		fmt.Printf("error opening video writer device: %v.%+v\n", videoName, err)
		return
	}
	defer writer.Close()

	for i := 0; i < 100; i++ {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", videoName)
			return
		}
		if img.Empty() {
			continue
		}

		writer.Write(img)
	}
}
