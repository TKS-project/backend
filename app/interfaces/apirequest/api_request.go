package apirequest

import (
	"encoding/json"
	"strings"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
)

type ApiRequestRepository struct {
	ApiRequestHandler
}

/*
middleClassCode
smallClassCode
*detailedClassCode

checkinDate
checkoutDate
adultNum
maxCharge
minCharge

*/
func (request *ApiRequestRepository) FindRoom(
	prefecture string,
	city string,
	checkinDate string,
	checkoutDate string,
	adultNum string,
	maxCharge string,
) domain.VacantHotels {
	baseUrl := "https://app.rakuten.co.jp/services/api/Travel/VacantHotelSearch/20170426?applicationId=1019851050958250331&format=json&largeClassCode=japan"

	url_slice := []string{
		baseUrl,
		"&middleClassCode=", prefecture,
		"&smallClassCode=", city,
		"&checkinDate=", checkinDate,
		"&checkoutDate=", checkoutDate,
		"&adultNum=", adultNum,
		"&maxCharge=", maxCharge,
	}
	//url := "https://app.rakuten.co.jp/services/api/Travel/VacantHotelSearch/20170426?applicationId=1019851050958250331&format=json&largeClassCode=japan&middleClassCode=akita&smallClassCode=tazawa&checkinDate=2022-12-01&checkoutDate=2022-12-02&adultNum=1&maxCharge=10000"

	finalUrl := strings.Join(url_slice, "")

	vacantHotel := domain.VacantHotels{}
	body, err := request.Get(finalUrl) //final Url
	if err != nil {

	}
	json.Unmarshal(body, &vacantHotel)

	return vacantHotel
}
