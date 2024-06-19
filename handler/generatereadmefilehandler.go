package handler

import (
	"encoding/json"
	"net/http"

	requestDto "github.com/ChrisMinKhant/megoyougo_framework/dto/request"
	responseDto "github.com/ChrisMinKhant/megoyougo_framework/dto/response"
	"github.com/ChrisMinKhant/megoyougo_framework/service/generatereadmefileservice"
)

type generateReadmeFileHandler struct {
	request  *requestDto.GenerateReadmeFileRequest
	response *responseDto.GenereteReadmeFileResponse
}

func NewGenerateReadmeFileHandler() *generateReadmeFileHandler {
	return &generateReadmeFileHandler{
		request: &requestDto.GenerateReadmeFileRequest{},
	}
}

func (generateReadmeFileHandler *generateReadmeFileHandler) Handle(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(generateReadmeFileHandler.request)

	if err != nil {
		panic("Error occurred at reading request body ::: " + err.Error())
	}

	service := generatereadmefileservice.New()

	service.GenerateReadmeFile(generateReadmeFileHandler.request)
}
