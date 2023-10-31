package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"

	"maindirectory.com/mysql"
)

type byStocks []*mysql.StockInfo

func (x byStocks) Len() int { return len(x) }
func (x byStocks) Less(i, j int) bool {
	myi, _ := strconv.Atoi(strings.ReplaceAll(x[i].Price, ",", ""))
	myj, _ := strconv.Atoi(strings.ReplaceAll(x[j].Price, ",", ""))
	return myi < myj
}
func (x byStocks) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func main() {
	mysql.Connect()
	stocks := mysql.GetStockInfo()
	sort.Sort(sort.Reverse(byStocks(stocks)))
	printStocks(stocks)
}

func printStocks(stocks []*mysql.StockInfo) {

	// for _, v := range stocks {
	// 	fmt.Printf("%v\t%v\t%v\t%v\t%v\t\n", v.Name, v.Price, v.PDE, v.TV, v.MC)
	// }

	const format = "%s\t\t%s\t\t%s\t\t%s\t\t%s\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 8, 4, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Name", "Price", "PDE", "TV", "MC")
	fmt.Fprintf(tw, format, "------", "------", "------", "------", "------")
	for _, v := range stocks {
		// 데이터를 탭 문자로 구분하여 포맷팅
		fmt.Fprintf(tw, format, v.Name, v.Price, v.PDE, v.TV, v.MC)
	}
	tw.Flush()
}
