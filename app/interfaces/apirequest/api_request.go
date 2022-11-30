package apirequest

import (
	"encoding/json"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type ApiRequestRepository struct {
	ApiRequestHandler
}

/*
middleClassCode
smallClassCode
detailedClassCode
checkinDate
checkoutDate
adultNum
maxCharge
minCharge

*/
func (request *ApiRequestRepository) FindRoom() domain.VacantHotels {

	url := "https://app.rakuten.co.jp/services/api/Travel/VacantHotelSearch/20170426?applicationId=1019851050958250331&format=json&largeClassCode=japan&middleClassCode=akita&smallClassCode=tazawa&checkinDate=2022-12-04&checkoutDate=2022-12-05&adultNum=1"

	vacantHotel := domain.VacantHotels{}
	body, err := request.Get(url)
	if err != nil {

	}
	json.Unmarshal(body, &vacantHotel)

	return vacantHotel
}
