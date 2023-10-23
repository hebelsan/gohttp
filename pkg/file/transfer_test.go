package file

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMultiPart(t *testing.T) {
	fileName := "test_multi.txt"
	fileContent := "test multipart upload"

	// Create a buffer to write the multipart form data
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Add a file to the form
	file, err := writer.CreateFormFile("fileField", fileName)
	if err != nil {
		fmt.Printf("Error creating form file: %v\n", err)
		return
	}

	_, err = io.Copy(file, strings.NewReader(fileContent))
	if err != nil {
		fmt.Printf("Error copying file to form: %v\n", err)
		return
	}

	// Close the multipart writer to finish building the form
	writer.Close()

	req, err := http.NewRequest("POST", "/", &requestBody)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()
	Handler(rr, req)

	expectedStatus := http.StatusOK
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, expectedStatus)
	}

	err = cleanupFile(fileName, fileContent)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestRawUpload(t *testing.T) {
	fileName := "test_raw.txt"
	fileContent := "test raw upload"
	body := strings.NewReader(fileContent)

	req, err := http.NewRequest("POST", "/", body)
	req.Header.Set("filename", fileName)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	Handler(rr, req)

	expectedStatus := http.StatusOK
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, expectedStatus)
	}

	err = cleanupFile(fileName, fileContent)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestRawNoFileNameHeader(t *testing.T) {
	body := strings.NewReader("test upload")
	req, err := http.NewRequest("POST", "/", body)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Handler(rr, req)
	// we expect to fail if filename header is missing
	expectedStatus := http.StatusBadRequest
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, expectedStatus)
	}
}

func TestUnsupportedMethod(t *testing.T) {
	req, err := http.NewRequest("PUT", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	Handler(rr, req)

	expectedStatus := http.StatusMethodNotAllowed
	if rr.Code != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, expectedStatus)
	}
}

func cleanupFile(fileName, fileContent string) error {
	// verify file has been created
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return fmt.Errorf("file '%s' does not exist", fileName)
	} else if err != nil {
		return fmt.Errorf("error checking file: '%v'", err)
	}

	contentBytes, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	content := string(contentBytes)
	if content != fileContent {
		return fmt.Errorf("content: %s != expectedContent %s", content, fileContent)
	}

	// cleanup the file
	err = os.Remove(fileName)
	if err != nil {
		return fmt.Errorf("error deleting the file: %v", err)
	}
	return nil
}
