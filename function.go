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
	//var d struct {
	//	Message string `json:"message"`
	//}
	//
	//if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
	//	switch err {
	//	case io.EOF:
	//		_, _ = fmt.Fprint(w, "Enter a number")
	//		return
	//	default:
	//		log.Printf("json.NewDecoder: %v", err)
	//		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	//		return
	//	}
	//}
	//
	//if d.Message == "" {
	//	_, _ = fmt.Fprint(w, "Error Occurred ")
	//	return
	//}

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
	listofTime := [20]int64{}
	for j := 0; j < 40; j++ {
		start := time.Now().UnixNano()
		// Loop 40 times.
		for i := 0; i <= 40; i++ {
			imageProcessing()
		}
		// End time
		end := time.Now().UnixNano()
		// Results
		if j > 20 {
			difference := end - start
			listofTime[j-20] = difference
		}
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
	throughput := avg / 40

	// Response
	fmt.Fprintf(w, "Time taken by %s function is %v ops/ns \n", funcName, throughput)
}

func imageProcessing() image.Image {
	//file, err := os.Open("image.jpg")
	var file, err = os.OpenFile("image.jpg",os.O_RDWR,0644)
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