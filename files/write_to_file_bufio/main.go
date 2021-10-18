package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"my_file.txt",
		os.O_WRONLY|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	bufferedWritter := bufio.NewWriter(file)
	bs := []byte{97, 98, 99}
	bytesWritten, err := bufferedWritter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bytes written to buffer (not file): %d\n", bytesWritten)

	bytesAvailable := bufferedWritter.Available()
	log.Printf("Bytes available in buffer: %d\n", bytesAvailable)

	bytesWritten, err = bufferedWritter.WriteString("\nJust a random string")
	if err != nil {
		log.Fatal(err)
	}

	unflushedBufferSize := bufferedWritter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	bufferedWritter.Flush()

	// bufferedWritter.Reset(bufferedWritter)
}
