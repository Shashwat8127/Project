package handlers

import (
	"bytes"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"

	"github.com/Shashwat8127/Project/file_converter/api"
)

type Handler struct {
	api.Unimplemented
}

func (Handler) ConvertJpegToPng(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "request must include a multipart file named file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	image, err := jpeg.Decode(file)
	if err != nil {
		http.Error(w, "uploaded file is not a valid JPEG image", http.StatusBadRequest)
		return
	}

	var converted bytes.Buffer
	if err := png.Encode(&converted, image); err != nil {
		log.Printf("encode converted PNG: %v", err)
		http.Error(w, "failed to convert image", http.StatusInternalServerError)
		return
	}

	response := api.ConvertJpegToPng200ImagepngResponse{
		Body:          bytes.NewReader(converted.Bytes()),
		ContentLength: int64(converted.Len()),
	}
	if err := response.VisitConvertJpegToPngResponse(w); err != nil {
		log.Printf("write converted PNG response: %v", err)
	}
}