package upload

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/davidrbourke/ImageUploader/Backend/utils"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)

	fmt.Println("File upload endpoint hit")
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("uploadedFile")
	if err != nil {
		fmt.Println("Error retrieving the file")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded file: %+v\n", handler.Filename)
	fmt.Printf("File size: %+v\n", handler.Size)
	fmt.Printf("MIME header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "Successfully uploaded file\n")
}
