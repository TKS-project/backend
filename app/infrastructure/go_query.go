package infrastructure

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Kantaro0829/clean-architecture-in-go/interfaces/scrape"
	"github.com/PuerkitoBio/goquery"
)

type GoQueryHandler struct {
	query *goquery.Document
}

func NewGoQueryHandler() scrape.ScrapeHandler {

	res, err := http.Get("https://ejje.weblio.jp/content/eliminate")
	if err != nil {
		panic(err.Error)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	goQueryHandler := new(GoQueryHandler)
	goQueryHandler.query, _ = goquery.NewDocumentFromReader(res.Body)
	return goQueryHandler
}

func (handler *GoQueryHandler) Find(word string) string {
	baseUrl := "https://ejje.weblio.jp/content/"
	baseUrl += word

	res, err := http.Get(baseUrl)
	if err != nil {
		panic(err.Error)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	//fmt.Printf("url：%v", handler.query.Url.Path)
	handler.query, _ = goquery.NewDocumentFromReader(res.Body)

	selection := handler.query.Find("div#summary")
	innerSeceltion := selection.Find("p")
	text := innerSeceltion.Text()
	result := strings.Replace(text, " ", "", -1)
	arr := strings.Split(result, "\n")
	var final string
	for i, s := range arr {
		if i == 4 {
			final = s
		}
		fmt.Printf("%s\n", s) // 赤 黄 青
		fmt.Println(i)
	}
	fmt.Println(len(arr)) // 3

	fmt.Printf("編集後のテキスト：%v", result)
	return final
}

// func (handler *GoQueryHandler) FindTransportation() interface{} {
// 	baseUrl := "https://www.navitime.co.jp/transfer/searchlist?orvStationName=中崎町&dnvStationName=新宿&month=2022%2F11&day=03&hour=14&minute=17&basis=1&from=view.transfer.searchlist&freePass=0&sort=0&wspeed=100&airplane=1&sprexprs=1&utrexprs=1&othexprs=1&mtrplbus=1&intercitybus=1&ferry=1&accidentRailCode=&accidentRailName=&isrec="

// 	res, err := http.Get(baseUrl)
// 	if err != nil {
// 		panic(err.Error)
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode != 200 {
// 		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
// 	}
// 	//fmt.Printf("url：%v", handler.query.Url.Path)
// 	handler.query, _ = goquery.NewDocumentFromReader(res.Body)
// 	selection := handler.query.Find("li#route_detail")
// 	fmt.Println(selection)
// }
