package main

import (
	"fmt"
	"generator/grab"
	"generator/tools"
	"strconv"
)

func processArticle(url string, art chan int64) {
	articles := grab.GetArticles(url)
	for _, t := range articles.Data {
		grab.GenerateStaticArticle(t)
		fmt.Println("ID: " + strconv.FormatInt(t.Id, 10) + " article grab complete")
	}
	art <- 1
}

func main() {
	var config tools.Config = tools.ParseConfig()
	var article_api string
	var category_pi string
	for _, v := range config.Api {
		article_api = v.Article
		category_pi = v.Category
	}

	//抓取分类数据
	cat := make(chan int64)
	category := grab.GetCategorys(category_pi)
	go grab.GenerateStaticCategory(category, cat)

	//抓取文章数据
	var total_page, i int64
	articles := grab.GetArticles(article_api)
	total_page = articles.Total/articles.Per_page + 1
	art := make(chan int64, total_page)
	for i = 1; i <= total_page; i++ {
		go processArticle(article_api+"?page="+strconv.FormatInt(i, 10), art)
	}

	//从channel中获取
	<-cat
	for i = 1; i <= total_page; i++ {
		<-art
	}
	close(cat)
	close(art)
}
