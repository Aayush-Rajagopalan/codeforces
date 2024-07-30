// this is the code behind the api server located at latte.cf.aayus.me
// this server is used to scrape the codeforces website and store the problems in a redis database

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"

	"github.com/PuerkitoBio/goquery"
	"github.com/redis/go-redis/v9"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type problem struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Url     string `json:"url"`
}

var ctx = context.Background()

func addProblem(ID string, Title string, Content string, url string) {
	opt, _ := redis.ParseURL("REDIS URL")
	client := redis.NewClient(opt)
	problem := problem{ID: ID, Title: Title, Content: Content, Url: url}
	problemJSON, _ := json.Marshal(problem)
	client.HSet(ctx, "problems", ID, problemJSON)
}

func getProblemFromDb(ID string) (*problem, error) {
	opt, _ := redis.ParseURL("REDIS URL")
	client := redis.NewClient(opt)
	val, err := client.HGet(ctx, "problems", ID).Result()

	// If there's an error or the problem is not found, return nil and the error
	if err != nil {
		return nil, err
	}

	var problem problem
	if err := json.Unmarshal([]byte(val), &problem); err != nil {
		return nil, err
	}

	return &problem, nil
}

func getProblems(ctx *gin.Context) {
	id := ctx.Param("id")
	num := ctx.Param("num")
	url := fmt.Sprintf("https://codeforces.com/problemset/problem/%s/%s", id, num)

	// Check if the problem is already in the database

	val, err := getProblemFromDb(fmt.Sprintf("%s/%s", id, num))

	if err == nil && val != nil {
		ctx.IndentedJSON(http.StatusOK, val)
		return
	}
	// If the problem is not in the database, scrape the website
	// and add it to the database

	c := colly.NewCollector(
		colly.AllowedDomains("codeforces.com"),
	)

	c.OnHTML("div.problem-statement", func(e *colly.HTMLElement) {
		title := e.ChildText("div.title")
		title = strings.ReplaceAll(title, "Input", "")
		title = strings.ReplaceAll(title, "Output", "")
		content, _ := e.DOM.Html()
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse the HTML content"})
			return
		}
		doc.Find("div.time-limit, div.header, div.memory-limit, div.input-file, div.output-file").Remove()
		var cleanedContent strings.Builder
		doc.Find("div").Each(func(i int, s *goquery.Selection) {
			html, err := s.Html()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse the HTML content"})
				return
			}
			cleanedContent.WriteString(html)
		})

		println("Content found:", cleanedContent.String())
		addProblem(fmt.Sprintf("%s/%s", id, num), title, cleanedContent.String(), url)
	})

	c.OnRequest(func(r *colly.Request) {
		println("Visiting", r.URL.String())
	})

	c.OnScraped(func(r *colly.Response) {
		val, err := getProblemFromDb(fmt.Sprintf("%s/%s", id, num))
		if err == nil && val != nil {
			ctx.IndentedJSON(http.StatusOK, val)

		}
	})

	err = c.Visit(url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to visit the site"})
		return
	}
}

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://cf.aayus.me"}
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	router.GET("/:id/:num", getProblems)

	router.Run(":8080")
}
