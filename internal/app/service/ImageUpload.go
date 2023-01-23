package service

import (
	logger "CRUD/internal/app/logs"
	"io"
	"mime/multipart"
	"os"
)

func UploadImageFromForm(file multipart.File, handler *multipart.FileHeader) string {

	storageDist := "internal/app/storage/" + handler.Filename
	defer file.Close()
	f, err := os.OpenFile(storageDist, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logger.Error.Println("Error in parsing form")
		panic(err)
	}
	defer f.Close()
	io.Copy(f, file)

	return storageDist
}