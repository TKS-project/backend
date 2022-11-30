package infrastructure

import (
	"io"
	"net/http"

	api "github.com/Kantaro0829/clean-architecture-in-go/interfaces/apirequest"
)

type ApiRequestHandler struct {
	request *http.Response
}

func NewApiRequestHandler() api.ApiRequestHandler {

	apiRequetHandler := new(ApiRequestHandler)
	// rakutenApiHandler.request, _ = goquery.NewDocumentFromReader(res.Body)
	return apiRequetHandler
}

func (handler *ApiRequestHandler) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	handler.request = resp
	defer resp.Body.Close()
	handler.request.Body = resp.Body

	body, err := io.ReadAll(handler.request.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
