package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

var (
	bufferedData []string
	bufferMutex  sync.Mutex
)

const (
	maxBufferSize   = 0.5 * 1024 * 16
	defaultFileName = "received-jsonl"
	specPath        = "/blackhole/"
)

func main() {
	// http.HandleFunc("/blackhole/", handleRequest)
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		w.Write([]byte("Hello, GET request!"))

	case http.MethodPost:

		fmt.Println("Post Request Received!")
		bufferWriterOOD(w, r)

	case http.MethodPut:

		w.Write([]byte("Hello, PUT request!"))

	case http.MethodDelete:

		w.Write([]byte("Hello, DELETE request!"))

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func bufferWriterOOD(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	fileName, err := getFileName(r.URL.Path)
	if err != nil {
		http.Error(w, "Bad Path!", http.StatusPartialContent)
		return
	}

	bufferMutex.Lock()
	bufferedData = append(bufferedData, string(body))
	bufferSize := calculateBufferSize()
	bufferMutex.Unlock()
	if bufferSize >= maxBufferSize {
		go writeToFile(fileName)
	}
	w.Write([]byte("JSON received and buffered successfully!"))
}

func getFileName(path string) (string, error) {
	// Get the filename for the bufferedData, return error if anything is incorrect
	if len(path) < len(specPath) {
		return defaultFileName, nil
	} else if path[0:len(specPath)] != specPath {
		return defaultFileName, nil
	} else if path[0:len(specPath)] == specPath && isProperReqPath(path[len(specPath):]) {
		return path[len(specPath):], nil
	} else {
		return "", errors.New("Broken Path!")
	}
}

func isProperReqPath(str string) bool {
	if str == "" {
		return false
	}
	regex := regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]*$")
	return regex.MatchString(str)
}

func calculateBufferSize() int {
	size := 0
	for _, data := range bufferedData {
		size += len(data)
	}
	return size
}

func writeToFile(fileName string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %v", err)
	}

	filePath := filepath.Join(currentDir, "data", fileName)

	existingData, err := ioutil.ReadFile(filePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to read data from file %s: %v", filePath, err)
	}

	for _, item := range bufferedData {
		existingData = append(existingData, []byte(item+"\n")...)
	}

	err = ioutil.WriteFile(filePath, existingData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write data to file %s: %v", filePath, err)
	}

	log.Printf("Data written to file: %s", filePath)

	bufferedData = nil

	return nil
}
