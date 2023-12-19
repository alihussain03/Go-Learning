package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	ID          string      `json:"id"`
	Items       []Item      `json:"items"`
	DateOrdered RFC822ZTime `json:"date_ordered"`
	CustomerID  string      `json:"customer_id"`
}

type RFC822ZTime struct {
	time.Time
}

func main() {

	//IO and friends

	fmt.Println("Pointer Receiver")
	fileName := "file.txt"

	fmt.Println(fileName)

	//create empty file
	createEmptyFile(fileName)

	//write into empty file
	writeToFile(fileName)

	// read from file
	readerFromFile(fileName)

	//example from book
	err := simpleCountLetters()
	if err != nil {
		slog.Error("error with simpleCountLetters", "msg", err)
	}

	err = gzipCountLetters()
	if err != nil {
		slog.Error("error with gzipCountLetters", "msg", err)
	}

	fmt.Println()
	//Seek interface
	readerFromFileUsingSeek(fileName)
	//End IO and friends

	//time
	timefunc()

	//json encoding
	jsonConvert()
}

func jsonConvert() {
	fmt.Println("Json convert Start \n\n")
	data := `
	{
		"id": "12345",
		"items": [
			{
				"id": "xyz123",
				"name": "Thing 1"
			},
			{
				"id": "abc789",
				"name": "Thing 2"
			}
		],
		"date_ordered": "01 May 20 13:01 +0000",
		"customer_id": "3"
	}`

	var o Order
	err := json.Unmarshal([]byte(data), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
	fmt.Println(o.DateOrdered.Month())
}

func timefunc() {
	p := fmt.Println
	p("\n\nStart Date function")
	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	p("End Date function\n\n")

}

func gzipCountLetters() error {
	r, closer, err := buildGZipReader("my_data.txt.gz")
	if err != nil {
		return err
	}
	defer closer()
	counts, err := countLetters(r)
	if err != nil {
		return err
	}
	fmt.Println(counts)
	return nil
}

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func simpleCountLetters() error {
	s := "The quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		return err
	}
	fmt.Println(counts)
	return nil
}

func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func createEmptyFile(fileName string) {
	file, _ := os.Create(fileName)
	file.Close()
}

func writeToFile(fileName string) {
	file, _ := os.Create("file.txt")
	writer := io.Writer(file)
	n, err := writer.Write([]byte("Hello ali"))

	n, err = io.WriteString(writer, "!")
	fmt.Println(n, err)
	file.Close()
}

func readerFromFile(fileName string) {

	file, _ := os.Open(fileName)
	reader := io.Reader(file)
	buffer := make([]byte, 1)
	// buffer := make([]byte, 1000) for larger file, instead of for loop
	for {
		// for read first method
		n, err := reader.Read(buffer)
		fmt.Printf("Read n={%v}, err={%v}, buffer={%v}\n", n, err, string(buffer))

		if err != nil {
			break
		}
	}
	fmt.Println()
	// for read second method
	buffer, err := io.ReadAll(file)
	fmt.Printf("Read n={%v}, err={%v}", string(buffer), err)
	file.Close()

}

func readerFromFileUsingSeek(fileName string) {

	file, _ := os.Open(fileName)
	reader := io.Reader(file)
	seeker := reader.(io.Seeker)
	seeker.Seek(0, io.SeekStart)

	// for read second method
	buffer, err := io.ReadAll(file)
	fmt.Printf("Read n={%v}, err={%v}", string(buffer), err)
	file.Close()

}
