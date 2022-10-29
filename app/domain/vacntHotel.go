package domain

import "time"

//domain 層
//Entity の定義を行う
type VacantHotels struct {
	PagingInfo PagingInfo  `json:"pagingInfo"`
	Hotel      interface{} `json:"hotels"`
}

type PagingInfo struct {
	RecordCount int `json:"recordCount"`
	PageCount   int `json:"pageCount"`
	Page        int `json:"page"`
	First       int `json:"first"`
	Last        int `json:"last"`
}

type HotelBasicInfo struct {
	HotelNo            uint16  `json:"hotelNo"`
	HotelName          string  `json:"hotelInformationUrl"`
	PlanListUrl        string  `json:"planListUrl"`
	DpPlanListUrl      string  `json:"dpPlanListUrl"`
	ReviewUrl          string  `json:"reviewUrl"`
	HotelKanaName      string  `json:"hotelKanaName"`
	HotelSpecial       string  `json:"hotelSpecial"`
	HotelMinCharge     string  `json:"hotelMinCharge"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	PostalCode         string  `json:"postalCode"`
	Address1           string  `json:"address1"`
	Address2           string  `json:"address2"`
	TelephoneNo        string  `json:"telephoneNo"`
	FaxNo              string  `json:"faxNo"`
	Access             string  `json:"access"`
	ParkingInformation string  `json:"parkingInformation"`
	NearestStation     string  `json:"nearestStation"`
	HotelImageUrl      string  `json:"hotelImageUrl"`
	HotelThumbnailUrl  string  `json:"hotelThumbnailUrl"`
	RoomImageUrl       string  `json:"roomImageUrl"`
	RoomThumbnailUrl   string  `json:"roomThumbnailUrl"`
	HotelMapImageUrl   string  `json:"hotelMapImageUrl"`
	ReviewCount        uint16  `json:"reviewCount"`
	ReviewAverage      float64 `json:"reviewAverage"`
	UserReview         string  `json:"userReview"`
}

type RoomBasicInfo struct {
	RoomClass           string `json:"roomClass"`
	RoomName            string `json:"roomName"`
	PlanId              string `json:"planId"`
	PlanName            string `json:"planName"`
	PointRate           string `json:"pointRate"`
	WithDinnerFlag      int    `json:"withDinnerFlag"`
	DinnerSelectFlag    int    `json:"dinnerSelectFlag"`
	WithBreakfastFlag   int    `json:"withBreakfastFlag"`
	BreakfastSelectFlag int    `json:"breakfastSelectFlag"`
	Payment             string `json:"payment"`
	ReserveUrl          string `json:"reserveUrl"`
	SalesformFlag       int    `json:"salesformFlag"`
}

type DailyCharge struct {
	StayDate      time.Time `json:"stayDate"` //べつの方のほうがいいかも？
	RakutenCharge uint16    `json:"rakutenCharge"`
	Total         uint16    `json:"total"`
	ChargeFlag    int       `json:"chargeFlag"`
}
