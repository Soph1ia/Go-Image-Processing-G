package p

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"time"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {

	// Call Benchmarking Function
	benchmark("Image Processing", w)

}

/**
Method : Benchmark

This method gets the time taken to execute the factorial 40 times.
In total it loops 80 times.
It takes the last 20 execution times.
Gets the average time
Calculates the throughput as time / 40

Prints out the throughput.

returns: none

*/
func benchmark(funcName string, w http.ResponseWriter) {
	listofTime := [41]int64{}
	for j := 0; j < 40; j++ {
		start := time.Now().UnixNano()
		imageProcessing()
		// End time
		end := time.Now().UnixNano()
		// Results
		difference := end - start
		listofTime[j] = difference
	}
	// Average Time
	sum := int64(0)
	for i := 0; i < len(listofTime); i++ {
		// adding the values of
		// array to the variable sum
		sum += listofTime[i]
	}
	// avg to find the average
	avg := (float64(sum)) / (float64(len(listofTime)))

	// Throughput Rate
	throughput := 40/avg

	// Response
	fmt.Fprintf(w, "Time taken by %s function is %v ops/ns \n", funcName, throughput)
}

func imageProcessing() image.Image {
	file, err := os.Open("./serverless_function_source_code/image.jpg")
	//var file, err = os.OpenFile("image.jpg",os.O_RDWR,0600)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(1024, 1000, img, resize.Lanczos3)

	return m
}
