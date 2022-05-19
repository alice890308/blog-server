package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("status: ok"))
}

func (s *Service) Upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Println("upload file error")

		return
	}
	fmt.Println(header.Size)
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}

	var tempFile *os.File
	resp := make(map[string]string)

	contentType := http.DetectContentType(buffer)
	fmt.Println(contentType)
	switch contentType {
	case "image/jpg":
		tempFile, err = ioutil.TempFile("static", "*.jpg")
		if err != nil {
			fmt.Println(err)
		}
	case "image/png":
		tempFile, err = ioutil.TempFile("static", "*.png")
		if err != nil {
			fmt.Println(err)
		}
	case "image/jpeg":
		tempFile, err = ioutil.TempFile("static", "*.jpeg")
		if err != nil {
			fmt.Println(err)
		}
	default:
		resp["message"] = "wrong image type"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}

	fileName := tempFile.Name()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	resp["message"] = "success"
	resp["file path"] = fileName
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(jsonResp)
}
