package main

import (
	"github.com/tuanlh1908developer/webFilmCrawler/crawler"
	"github.com/tuanlh1908developer/webFilmCrawler/todb"
)
func main() {
	db := todb.ConnectDB()
	todb.CreateTable(db)
	crawler.Crawl(db)
}
