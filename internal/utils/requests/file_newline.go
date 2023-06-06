package requests

import (
	"bufio"
	"log"
	"os"
)

func FileContextInit(fn string) (*os.File, *bufio.Writer, error) {
	file, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	datawriter := bufio.NewWriter(file)
	return file, datawriter, err
}

func WriteFileBufferListNewline(datawriter *bufio.Writer, sampledata []string) {
	// file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// if err != nil {
	// 	log.Fatalf("failed creating file: %s", err)
	// }

	// datawriter := bufio.NewWriter(file)

	for _, data := range sampledata {
		_, _ = datawriter.WriteString(data + "\n")
	}

	// datawriter.Flush()
	// file.Close()
}

func FileContextClose(file *os.File, datawriter *bufio.Writer) {
	datawriter.Flush()
	file.Close()
}
