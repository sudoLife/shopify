package shopify

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"

	//"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DateFormat represents the date format used by shopify
const DateFormat = "January 2, 2006"

// Review represents the structure of a single review on apps.shopify.com; compatible with json and bson
type Review struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username,omitempty" bson:"username,omitempty"`
	Rating    int                `json:"rating,omitempty" bson:"rating,omitempty"`
	Date      int64              `json:"time,omitempty" bson:"time,omitempty"`
	Content   string             `json:"content,omitempty" bson:"content,omitempty"`
	Helpful   int                `json:"helpful,omitempty" bson:"helpful,omitempty"`
	Reply     string             `json:"reply,omitempty" bson:"reply,omitempty"`
	ReplyDate int64              `json:"replydate,omitempty" bson:"replydate,omitempty"`
}

// Parse parses shopify
func Parse(url string) *[]Review {
	reviews := []Review{}

	c := colly.NewCollector(
		colly.Async(true),
	)

	extensions.RandomUserAgent(c)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 4,
		RandomDelay: time.Second,
	})

	c.OnHTML("div.review-listing", func(e *colly.HTMLElement) {
		fmt.Printf("=")
		var err error
		rating, _ := strconv.Atoi(e.ChildAttr("div[data-review-id] div.review-metadata div:nth-child(1) div.review-metadata__item-value div[data-rating]", "data-rating"))
		helpful, _ := strconv.Atoi(e.ChildText("div.review-footer div.review-helpfulness form button span.review-helpfulness__helpful-count"))
		date, err := time.Parse(DateFormat, strings.TrimSpace(e.ChildText("div[data-review-id] div.review-metadata div:nth-child(2) div.review-metadata__item-value")))

		if err != nil {
			log.Fatal(err)
		}

		replyDate, err := time.Parse(DateFormat, strings.TrimSpace(e.ChildText("div.review-reply div.review-reply__header div.review-reply__header-item")))

		if err != nil {
			replyDate, _ = time.Parse(DateFormat, "January 1, 1970")
		}

		//		replyDate := strings.TrimSpace(e.ChildText("div.review-reply div.review-reply__header div.review-reply__header-item"))

		reviews = append(reviews, Review{
			Username:  e.ChildText("div[data-review-id] div.review-listing-header h3"),
			Rating:    rating,
			Date:      date.Unix(),
			Content:   strings.TrimSpace(e.ChildText("div[data-review-id] div.review-content div.truncate-content-copy")),
			Helpful:   helpful,
			Reply:     strings.TrimSpace(e.ChildText("div.review-reply div.review-content div.truncate-content-copy p")),
			ReplyDate: replyDate.Unix(),
		})

	})

	c.OnError(func(resp *colly.Response, err error) {
		fmt.Println(err)
	})

	c.OnHTML("a.search-pagination__next-page-text", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})

	// Basic loading bar to understand that the process is going
	fmt.Printf("[")
	c.Visit(url)
	c.Wait()
	fmt.Printf("]")

	return &reviews
}
