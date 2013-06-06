package main

import (
	// "bytes"
	"fmt"
	"image"
	"image/png"
	// "io"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func cutterHandler(res http.ResponseWriter, req *http.Request) {
	println(1)
	defer req.Body.Close()

	reqImg, err := http.Get("http://fluidapp.com/images/dock_large.png")
	if err != nil {
		fmt.Fprintf(res, "Error %d", err)
		return
	}
	println(2)
	// raw := make([]byte, reqImg.ContentLength)
	// io.ReadFull(reqImg.Body, raw)

	// buffer := bytes.NewBuffer(raw)
	img, _, err := image.Decode(reqImg.Body)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	println(3)

	res.Header().Set("Content-Type", reqImg.Header.Get("Content-Type"))

	png.Encode(res, img)

}

func cutterHandler2(res http.ResponseWriter, req *http.Request) {
	println(1)
	reqImg, err := http.Get("http://fluidapp.com/images/dock_large.png")
	if err != nil {
		fmt.Fprintf(res, "Error %d", err)
		return
	}
	println(2)
	raw := make([]byte, reqImg.ContentLength)
	io.ReadFull(reqImg.Body, raw)

	buffer := bytes.NewBuffer(raw)
	println(3)
	res.Header().Set("Content-Type", reqImg.Header.Get("Content-Type"))

	io.Copy(res, buffer)
}

func cutterHandler3(res http.ResponseWriter, req *http.Request) {
	println(1)
	reqImg, err := http.Get("http://qortex.net/theplant.jp/file/509b74873c58161d1d00a022/Screen%20Shot%202012-11-08%20at%2017.55.47.PNG")
	if err != nil {
		// give default image
		return
	}

	println(2)
	content, err := ioutil.ReadAll(reqImg.Body)
	if err != nil {
		// give default image
		return
	}

	println(3)
	// reqImg.Body.Close()

	res.Header().Set("Status", "302")
	io.Copy(res, bytes.NewBuffer(content))
}

func cutterHandler4(w http.ResponseWriter, req *http.Request) {
	reqImg, err := http.Get("http://files.gamebanana.com/bitpit/vayne.jpg")
	if err != nil {
		// give default image
		return
	}
	defer reqImg.Body.Close()

	raw := make([]byte, reqImg.ContentLength)
	io.ReadFull(reqImg.Body, raw)

	w.Header().Set("Content-Length", fmt.Sprintf("%d", reqImg.ContentLength))
	w.Header().Set("Expires", FormatDays(30))
	w.Header().Set("Cache-Control", "max-age="+FormatDayToSec(30))

	w.Write(raw)
}

// the time format used for HTTP headers
const HTTP_TIME_FORMAT = "Mon, 02 Jan 2006 15:04:05 GMT"

func FormatHour(hours string) string {
	d, _ := time.ParseDuration(hours + "h")
	return time.Now().Add(d).Format(HTTP_TIME_FORMAT)
}

func FormatDays(day int) string {
	return FormatHour(fmt.Sprintf("%d", day*24))
}

func FormatDayToSec(day int) string {
	return fmt.Sprintf("%d", day*60*60*24)
}

func main() {
	http.HandleFunc("/cut", cutterHandler)
	http.HandleFunc("/cut2", cutterHandler2)
	http.HandleFunc("/cut3", cutterHandler3)
	http.HandleFunc("/cut4", cutterHandler4)
	http.ListenAndServe(":8080", nil) /* TODO Configurable */
}
