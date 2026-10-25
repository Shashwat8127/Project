package handlers

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Shashwat8127/Project/file_converter/api"
)

func TestConvertJpegToPng(t *testing.T) {
	var jpegBytes bytes.Buffer
	source := image.NewRGBA(image.Rect(0, 0, 2, 2))
	source.Set(0, 0, color.RGBA{R: 255, A: 255})
	if err := jpeg.Encode(&jpegBytes, source, nil); err != nil {
		t.Fatalf("encode test JPEG: %v", err)
	}

	var requestBody bytes.Buffer
	form := multipart.NewWriter(&requestBody)
	part, err := form.CreateFormFile("file", "input.jpg")
	if err != nil {
		t.Fatalf("create multipart file: %v", err)
	}
	if _, err := part.Write(jpegBytes.Bytes()); err != nil {
		t.Fatalf("write multipart file: %v", err)
	}
	if err := form.Close(); err != nil {
		t.Fatalf("close multipart form: %v", err)
	}

	request := httptest.NewRequest(http.MethodPost, "/api/v1/convert/jpeg-to-png", &requestBody)
	request.Header.Set("Content-Type", form.FormDataContentType())
	response := httptest.NewRecorder()

	api.Handler(Handler{}).ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, response.Code, response.Body.String())
	}
	if got := response.Header().Get("Content-Type"); got != "image/png" {
		t.Fatalf("expected Content-Type image/png, got %q", got)
	}
	if _, err := png.Decode(response.Body); err != nil {
		t.Fatalf("response is not a valid PNG: %v", err)
	}
}
