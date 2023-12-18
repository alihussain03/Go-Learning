package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

func main() {
	fmt.Println("Pointer Receiver")
	fileName := "file.txt"

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
	fmt.Println()
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
