package utils

import (
	"bufio"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetReader() *bufio.Reader {
	return reader
}
