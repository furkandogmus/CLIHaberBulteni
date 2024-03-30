package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
)
const version = "1.0.0"

func main() {

	showVersion := flag.Bool("v", false, "Versiyon bilgisini göster!")
	showHelp := flag.Bool("h", false, "Yardım bilgisini göster!")
	flag.Parse()

	if *showVersion {
		fmt.Printf("Haberler Versiyon %s\n", version)
		return
	}
	if *showHelp {
		fmt.Println("CLI Haber Bulteni")
		fmt.Println("Linkler ve haber başlıkları gzt.com 'dan alınmıştır.")
		fmt.Println()
		fmt.Println("Seçenekler:")
		fmt.Println("-h  yardım bilgisini verir.")
		fmt.Println("-v  versiyon bilgisini verir.")
		return
	}

	categories := []string{"POLITIKA", "DUNYA", "EKONOMI", "BILIM", "GUNCEL", "AKTUEL-KULTUR", "SAGLIK"}
	fmt.Printf("%sKategoriler:\n", Cyan)
	for i, cat := range categories {
		fmt.Printf("%s%d. %s%s\n", Yellow, i+1, cat, Reset)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%sKategori numarası girin: \n", Cyan)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	categoryNum, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("Sadece 1-7 arasında bir sayı girmelisiniz.")
	}

	if categoryNum < 1 || categoryNum > len(categories) {
		fmt.Println("Geçersiz kategori numarası.")
		return
	}

	category := categories[categoryNum-1]

	url := fmt.Sprintf("https://www.gzt.com/%s", strings.ToLower(category))

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".feed-card-content.news-card-content").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		text = strings.Split(text, "…..devamı")[0] + ": "
		fmt.Printf("%s%s", Red, text)
		firstLink := s.Find("a").First()
		href, exists := firstLink.Attr("href")
		if exists {
			fmt.Printf("%shttps://www.gzt.com%s%s\n", Gray, href, Reset)
		}
	})
}
