package video

import (
	"fmt"
	"log"

	"github.com/thinkski/go-v4l2"
)

func GetVideoStream(rtspURL string) {
	fmt.Println("GetVideoStream")

	device, err := v4l2.Open("/dev/video0")
	if err != nil {
		log.Fatalf("Failed to open video device: %v", err)
	}

	err = device.SetPixelFormat(1280, 720, v4l2.V4L2_PIX_FMT_H264)
	if err != nil {
		log.Fatalf("Failed to set pixel format: %v", err)
	}

	// Remember to close the device when you're done with it
	defer device.Close()
}
