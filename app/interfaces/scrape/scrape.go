package scrape

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/Kantaro0829/clean-architecture-in-go/domain"
	"github.com/PuerkitoBio/goquery"
)

type ScrapeRepository struct {
	ScrapeHandler
}

func (scrape *ScrapeRepository) FindMeaning(word string) string {

	meaning := scrape.Find(word)
	fmt.Println(meaning)
	return meaning
}

func (scrape *ScrapeRepository) FindTransports() string {
	text := scrape.FindTransportation()
	return text
}

func (scrape *ScrapeRepository) FindOneTransport(
	orvStation string,
	dnvStation string,
	year string,
	month string,
	day string,
	hour string,
	minute string,
) domain.TransportsJson {
	url := scrape.CreateUrl(orvStation, dnvStation, year, month, day, hour, minute)
	doc := scrape.GetDoc(url)

	json := domain.TransportsJson{}
	charge, leftToArrive := getTimeAndCharge(doc)
	json.MaxCharge, json.LeftToArrive = charge, leftToArrive

	routeDetails := doc.Find("li.route_detail")
	first := routeDetails.First()

	detailFrame := first.Find("div.section_detail_frame")
	routeInfo := detailFrame.Children()

	station := false
	train := false
	stationMarche := false
	walk := false
	aroundInfo := false
	headerLink := false

	routeInfo.Next().Each(func(i int, s *goquery.Selection) {
		timeInfo := ""
		nameInfo := ""
		urlInfo := ""
		plathomeInfo := ""

		//どのクラスに属しているのかを判定
		aroundInfo = s.HasClass("around_info")
		headerLink = s.HasClass("section-header-links")
		station = s.HasClass("section_station_frame")
		train = s.HasClass("detail_frame train")
		walk = s.HasClass("detail_frame walk")
		stationMarche = s.HasClass("section_station_frame marche")
		if aroundInfo {
			fmt.Println("ルート検索終了")
			aroundInfo = false
		}

		if headerLink {
			fmt.Println("最初のヘッダー")
			headerLink = false
		}

		if station {
			station = false
			if stationMarche {
				fmt.Println("stationMarche---------------------------------------------------------------------------------------------------")
				stationMarche = false
				timeInfo, nameInfo = getStationInfo(i, s)
			} else {
				fmt.Println("station---------------------------------------------------------------------------------------------------")
				timeInfo, nameInfo = getStationInfo(i, s)
			}
			transport := CreateTransports("station", nameInfo, timeInfo, "", "")
			json.Transports = append(json.Transports, transport)
		}

		if train {
			fmt.Println("train---------------------------------------------------------------------------------------------------")
			train = false
			timeInfo, plathomeInfo, nameInfo, urlInfo = getTrainInfo(s)
			transport := CreateTransports("train", nameInfo, timeInfo, plathomeInfo, urlInfo)
			json.Transports = append(json.Transports, transport)
		}

		if walk {
			fmt.Println("walk---------------------------------------------------------------------------------------------------")
			walk = false
			fmt.Println(getWalkInfo(s))
			timeInfo = getWalkInfo(s)
			transport := CreateTransports("walk", nameInfo, timeInfo, plathomeInfo, urlInfo)
			json.Transports = append(json.Transports, transport)
		}
	})

	fmt.Println(json)
	return json

}

func (scrape *ScrapeRepository) ByteStringForURlParam(str string, param string) string {
	b := []byte(str)
	encodedStr := hex.EncodeToString(b)
	msg := encodedStr
	splitlen := 2
	runes := []rune(msg)

	url_slice := []string{param}

	for i := 0; i < len(runes); i += splitlen {
		if i+splitlen < len(runes) {
			url_slice = append(url_slice, string(runes[i:(i+splitlen)]))
		} else {
			fmt.Println(string(runes[i:]))
			url_slice = append(url_slice, string(runes[i:]))
		}
	}
	fmt.Println(url_slice)
	finalUrl := strings.Join(url_slice, "%")

	return finalUrl
}

func (scrape *ScrapeRepository) CreateUrl(
	orvStation string,
	dnvStation string,
	year string,
	month string,
	day string,
	hour string,
	minute string,
) string {

	monthParam := strings.Join([]string{"month=", year, "%2f", month}, "")
	dayParam := strings.Join([]string{"day=", day}, "")
	hourParam := strings.Join([]string{"hour=", hour}, "")
	minuteParam := strings.Join([]string{"minute=", minute}, "")

	//URLを作成する（各ParamをJoinでつなげる）

	orvParam := "https://www.navitime.co.jp/transfer/searchlist?orvStationName="
	dnvParam := "dnvStationName="
	orvStationName := scrape.ByteStringForURlParam(orvStation, orvParam)
	dnvStationName := scrape.ByteStringForURlParam(dnvStation, dnvParam)
	fmt.Println(dnvStationName)

	urlSlice := []string{
		orvStationName,
		dnvStationName,
		monthParam,
		dayParam,
		hourParam,
		minuteParam,
		"basis=0&from=view.transfer.searchlist&freePass=0&sort=0&wspeed=100&airplane=1&sprexprs=1&utrexprs=1&othexprs=1&mtrplbus=1&intercitybus=1&ferry=1&accidentRailCode=&accidentRailName=&isrec=",
	}

	url := strings.Join(urlSlice, "&")
	return url
}

func getWalkInfo(s *goquery.Selection) string {
	timeInfo := s.Find("ul.time_info")
	walkTimeText := timeInfo.Find("li").Text()
	walktime := DeleteRedundantSpaceFromString(walkTimeText)
	return walktime
}

func getTimeAndCharge(d *goquery.Document) (string, string) {
	summaryList := d.Find("ol.summary_list")
	firstList := summaryList.Find("li.list_border").First()
	cash := firstList.Find("div.cash").Text()
	time := firstList.Find("dt.left").Text()
	return cash, time
}

func getTrainInfo(s *goquery.Selection) (string, string, string, string) {
	timeInfo := s.Find("ul.time_info")
	railRoadInfo := s.Find("div.railroad_info")
	railRoad := s.Find("dl.railroad")
	aShinkansenLink := railRoad.Find("a.shinkansen_link")
	attr, exist := aShinkansenLink.Attr("href")

	homeText := railRoadInfo.Find("li.left").Text()
	timeOnBoad := timeInfo.Find("li").First().Text()
	distanceOnBoad := timeInfo.Find("li.distance").Text()
	railRoadAreaText := railRoad.Find("div.railroad-area").Text()
	url := ""
	if exist {
		url = AddHttpsString(attr)

	}

	time := DeleteRedundantSpaceFromString(timeOnBoad)
	distance := DeleteRedundantSpaceFromString(distanceOnBoad)
	plathome := DeleteRedundantSpaceFromString(homeText)
	trainName := DeleteRedundantSpaceFromString(railRoadAreaText)
	fmt.Println(time)
	fmt.Println(distance)
	fmt.Println(plathome)
	fmt.Println(trainName)
	fmt.Println(url)
	return time, plathome, trainName, url
}

func getStationInfo(i int, s *goquery.Selection) (string, string) {
	time := ""
	if i == 1 {
		fmt.Println(i)
		//出発時の駅の時
		//timeText := s.Find("dd.left").Text()
		timeText := s.Find("dd.sgk_time").Text()
		time = DeleteRedundantSpaceFromString(timeText)

		fmt.Println(time)
	}
	if s.Find("dd.left").Text() == "" && s.Find("dd.sgk_time").Text() == "" {
		//乗り換え時の駅
		time = s.Find("div.sgk_time").Text()
		fmt.Println(s.Find("div.sgk_time").Text())
	}

	if s.Find("dd.sgk_time").Text() != "" && i != 1 {
		//最後の駅
		time = s.Find("dd.sgk_time").Text()
		fmt.Println(s.Find("dd.sgk_time").Text())
	}

	stationText := s.Find("dl.station_frame").Find("dt.left").Text()
	station := DeleteRedundantSpaceFromString(stationText)
	fmt.Println(station)
	return time, station
}

func CreateTransports(
	ty string,
	name string,
	time string,
	plathome string,
	url string,
) domain.Transport {
	transport := domain.Transport{}
	transport.Name = name
	transport.Type = ty
	transport.Time = time
	transport.Plathome = plathome
	transport.Url = url
	return transport
}

func DeleteRedundantSpaceFromString(s string) string {
	deletedEmpty := strings.Replace(s, " ", "", -1)
	deleteNewLine := strings.Replace(deletedEmpty, "\n", "", -1)
	str := strings.TrimSpace(deleteNewLine)

	return str
}

func AddHttpsString(s string) string {
	urlSlice := []string{
		"https:",
		s,
	}
	url := strings.Join(urlSlice, "")
	return url
}
