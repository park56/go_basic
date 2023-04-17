// package main

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strings"

// 	"github.com/PuerkitoBio/goquery"
// 	"github.com/djimenez/iconv-go"
// )

// type stockInfo struct {
// 	item_name        string
// 	current_price    string //현제가
// 	then_yesterday   string //전일비
// 	fluctuation_rate string //등략률
// 	trading_volume   string //거래량
// 	market_cap       string //시가총액
// }

// func main() {

// 	fmt.Println("크롤링 시작")

// 	err := scrape()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	defer fmt.Println("크롤링 종료")
// }

// var url string = "https://finance.naver.com/sise/sise_quant.naver"

// func scrape() error {

// 	res, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}

// 	defer res.Body.Close()
// 	if res.StatusCode != 200 {
// 		return errors.New("status code error : " + res.Status)
// 	}

// 	utfBody, err := iconv.NewReader(res.Body, "EUC-KR", "utf-8")
// 	if err != nil {
// 		return err
// 	}

// 	doc, err := goquery.NewDocumentFromReader(utfBody)
// 	if err != nil {
// 		return err
// 	}

// 	html, _ := doc.Find("table.type_2").Find("tbody").Html()

// 	html = strings.ReplaceAll(html, "\t", "")

// 	htmls := strings.Split(html, "\n")

// 	last := ""

// 	for i, v := range htmls {

// 		if strings.Contains(v, `</a>`) {
// 			last += v + "\n"
// 		} else if strings.Contains(v, `<span class="tah p11`) {
// 			v += htmls[i+1]
// 			v += htmls[i+2]
// 			last += v + "\n"
// 		} else if strings.Contains(v, `<td class="number">`) {
// 			last += v + "\n"
// 		}
// 	}

// 	namefile, _ := os.Create("./namefilex.txt")
// 	namefile.WriteString(last)

// 	return nil
// }

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

type stockInfo struct {
	item_name        string
	current_price    string //현제가
	then_yesterday   string //전일비
	fluctuation_rate string //등략률
	trading_volume   string //거래량
	market_cap       string //시가총액
}

func main() {

	fmt.Println("크롤링 시작")

	err := scrape()
	if err != nil {
		log.Println(err)
	}

	defer fmt.Println("크롤링 종료")
}

var url string = "https://finance.naver.com/sise/sise_quant.naver"

func scrape() error {

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return errors.New("status code error : " + res.Status)
	}

	utfBody, err := iconv.NewReader(res.Body, "EUC-KR", "utf-8")
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		return err
	}

	html, _ := doc.Find("table.type_2").Find("tbody").Html()

	html = strings.ReplaceAll(html, "\t", "")

	htmls := strings.Split(html, "\n")

	last := []string{}

	for i, v := range htmls {

		if strings.Contains(v, `</a>`) {
			last = append(last, v)
		} else if strings.Contains(v, `<span class="tah p11`) {
			v += htmls[i+1]
			v += htmls[i+2]
			last = append(last, v)
		} else if strings.Contains(v, `<td class="number">`) {
			last = append(last, v)
		}
	}

	lastString := ""
	lastStringSlice := []string{}
	flag := false

	for _, v := range last {
		temp := strings.Split(v, "")
		for _, v2 := range temp {

			if v2 == ">" {
				flag = true
			}
			if v2 == "<" {
				flag = false
				lastString = strings.TrimSpace(lastString)
				lastStringSlice = append(lastStringSlice, lastString)
				lastString = ""
			}

			if flag {
				if v2 == ">" {
					continue
				} else if v2 == "" {
					continue
				}
				lastString += v2
			}
		}
	}

	for _, v := range lastStringSlice {
		fmt.Print(v)
		time.Sleep(time.Second)
	}

	//namefile, _ := os.Create("./namefilex.txt")
	//namefile.WriteString(last)

	return nil
}
