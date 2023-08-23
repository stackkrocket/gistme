package routes

import (
	"fmt"
	"os"
)

func ServeHTMLPage(path string, responseChan chan<- []byte) {
	htmlBytes, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("Error reading HTML file: ", err)
		responseChan <- []byte("Internal server error")
		return
	}

	responseChan <- htmlBytes
}
