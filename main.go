package main

import (
	"github.com/tuanlh1908developer/webFilmCrawler/crawler"
)

//write a csv file, write the header to that file
//define a collector
//define onHTML
//define onError
//define onResponse
//define onRequest
//craw the web

func main() {
	crawler.Crawl()
	// log.Printf("Scraping finished, check file %q for results\n", fName)
	// c1.Wait()
	// c2.Wait()
}
