package file

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fs := http.FileServer(http.Dir("./static"))
		fs.ServeHTTP(w, r)
	case http.MethodPost:
		if isMultipart(r) {
			handleMultipart(w, r)
		} else {
			handleRaw(w, r)
		}
		_, _ = w.Write([]byte("upload successful\n"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleMultipart(w http.ResponseWriter, r *http.Request) {
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
			dst, err := os.Create(header.Filename)
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
}
func isMultipart(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data")
}

func handleRaw(w http.ResponseWriter, r *http.Request) {
	filename := r.Header.Get("filename")
	if filename == "" {
		http.Error(w, "filename header not found", http.StatusBadRequest)
		return
	}
	// Create a new file on the server to save the uploaded content.
	dst, err := os.Create(filename)
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
}
