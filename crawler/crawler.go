package crawler

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/gocolly/colly"
	"github.com/tuanlh1908developer/webFilmCrawler/todb"
)

type Film struct {
	title       string
	rating      float32
	category    []string
	description string
	director    string
	writers     []string
	stars       []string
}

func Crawl(db *sql.DB) {
	c1 := colly.NewCollector(colly.Async(false))
	
	c2 := colly.NewCollector(colly.Async(false))


	c1.OnHTML("tbody.lister-list tr", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a[href]", "href")
		c2.Visit(e.Request.AbsoluteURL(link))

	})
	
	c2.OnHTML("section.ipc-page-background.ipc-page-background--base.TitlePage__StyledPageBackground-wzlr49-0.dDUGgO", func(e *colly.HTMLElement) {
		title := e.ChildText("h1")
		rating, err := strconv.ParseFloat(e.ChildText("div.Hero__ContentContainer-kvkd64-10.eaUohq span.AggregateRatingButton__RatingScore-sc-1ll29m0-1.iTLWoV"), 32)
		if err != nil {
			rating = 0.0
			log.Printf("Can not get rating of the film %s\n", title)
		}

		category := make([]string, 0)
		stars := make([]string, 0)
		writers := make([]string, 0)
		e.ForEach(".GenresAndPlot__ContentParent-cum89p-8>div>a", func(i int, h *colly.HTMLElement) {
			category = append(category, h.Text)
		})
		description := e.ChildText(".GenresAndPlot__ContentParent-cum89p-8>p")
		director := e.ChildText(".PrincipalCredits__PrincipalCreditsPanelWideScreen-hdn81t-0>ul>li:first-child a")
		e.ForEach(".PrincipalCredits__PrincipalCreditsPanelWideScreen-hdn81t-0>ul>li:nth-child(2) a", func(i int, h *colly.HTMLElement) {
			writers = append(writers, h.Text)
		})
		e.ForEach(".PrincipalCredits__PrincipalCreditsPanelWideScreen-hdn81t-0>ul>li:nth-child(3)>div a ", func(i int, h *colly.HTMLElement) {
			stars = append(stars, h.Text)
		})
		aFilm := Film{
			title:       title,
			rating:      float32(rating),
			category:    category,
			description: description,
			director:    director,
			writers:     writers,
			stars:       stars,
		}
		rows, err := todb.InsertToDB(db, aFilm.title, aFilm.rating, aFilm.category, aFilm.description, aFilm.director, aFilm.writers, aFilm.stars)
		if rows != 0 {
			log.Printf("Insert for film %q suceed, %v affected\n", aFilm.title, rows)
		} else {
			log.Printf("Insert for film %q failed, %v\n", aFilm.title, err)
		}
	})

	c1.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting...", r.URL.String())
		r.Headers.Add("Accept-Language", "en-US")
	})
	c2.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting...", r.URL.String())
		r.Headers.Add("Accept-Language", "en-US")
	})
	c1.OnError(func(r *colly.Response, e error) {
		log.Println(e)
	})
	c1.Visit("https://www.imdb.com/chart/top/?ref_=nv_mv_250")
}
