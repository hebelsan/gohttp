// Package file implements utility routines for uploading and downloading files.
package file

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Handler serves files on GET and downloads files on POST requests
func Handler(w http.ResponseWriter, r *http.Request) {
	filesPath, err := getFilesRoot()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		fs := http.FileServer(http.Dir(filesPath))
		fs.ServeHTTP(w, r)
	case http.MethodPost:
		if isMultipart(r) {
			handleMultipart(w, r, filesPath)
		} else {
			handleRaw(w, r, filesPath)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleMultipart(w http.ResponseWriter, r *http.Request, filesPath string) {
	// Parse the form data, which may include uploaded files.
	err := r.ParseMultipartForm(10 << 20) // Set a reasonable memory limit for the form fields
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Iterate through the files in the multipart form.
	for _, headers := range r.MultipartForm.File {
		for _, header := range headers {
			fmt.Println(header.Filename)
			// Open the file from the request.
			file, err := header.Open()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer file.Close()

			// Create a new file on the server to save the uploaded content.
			filePath := filepath.Join(filesPath, header.Filename)
			dst, err := os.Create(filePath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer dst.Close()

			// Copy the file content to the destination file.
			_, err = io.Copy(dst, file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
	_, _ = w.Write([]byte("upload successful\n"))
}

func handleRaw(w http.ResponseWriter, r *http.Request, filesPath string) {
	filename := r.Header.Get("filename")
	if filename == "" {
		http.Error(w, "filename header not found", http.StatusBadRequest)
		return
	}
	// Create a new file on the server to save the uploaded content.
	filePath := filepath.Join(filesPath, filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the file content to the destination file.
	_, err = io.Copy(dst, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write([]byte("upload successful\n"))
}

func isMultipart(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data")
}

func getFilesRoot() (string, error) {
	if envPath := os.Getenv("FILES_ROOT"); envPath != "" {
		return envPath, nil
	}
	return os.Getwd()
}
