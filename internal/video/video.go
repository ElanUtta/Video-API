package video

import (
	"fmt"

	"gocv.io/x/gocv"
)

func GetVideoStream(rtspURL string) {

	can, err := gocv.VideoCaptureFile(rtspURL)
	if err != nil {
		fmt.Printf("Error opening video capture: %v\n", err)
		return
	}
	defer can.Close()

	// Create a window to display the video feed
	window := gocv.NewWindow("RTSP Video Feed")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	fmt.Printf("Press 'ESC' to exit...\n")

	// Read frames from the RTSP stream and display them in the window
	for {
		if ok := can.Read(&img); !ok {
			fmt.Printf("Cannot read frame from video capture\n")
			break
		}

		if img.Empty() {
			continue
		}

		window.IMShow(img)
		key := window.WaitKey(1)

		// Check for the 'ESC' key to exit the loop
		if key == 27 {
			break
		}
	}

}
