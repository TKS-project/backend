package infrastructure

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
	api "github.com/Kantaro0829/clean-architecture-in-go/interfaces/rakutenApi"
)

type RakutenApiHandler struct {
	request *http.Response
}

func NewRakutenApiRequestHandler() api.ApiRequestHandler {

	rakutenApiHandler := new(RakutenApiHandler)
	// rakutenApiHandler.request, _ = goquery.NewDocumentFromReader(res.Body)
	return rakutenApiHandler
}

func (handler *RakutenApiHandler) FindAllVacantHotel() domain.VacantHotels {
	resp, err := http.Get("https://app.rakuten.co.jp/services/api/Travel/VacantHotelSearch/20170426?applicationId=1019851050958250331&format=json&largeClassCode=japan&middleClassCode=akita&smallClassCode=tazawa&checkinDate=2022-12-01&checkoutDate=2022-12-02&adultNum=2")
	if err != nil {
		// handle error
	}

	vacantHotels := domain.VacantHotels{}

	handler.request = resp
	defer resp.Body.Close()
	handler.request.Body = resp.Body

	body, err := io.ReadAll(handler.request.Body)
	//body, err := io.ReadAll(resp.Body)
	json.Unmarshal(body, &vacantHotels)
	fmt.Printf(" : \n%v", vacantHotels)
	return vacantHotels
}
