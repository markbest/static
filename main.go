package main

import (
	"generator/grab"
	"generator/tools"
	"strconv"
)

func processArticle(url string, p int64, d grab.Data) {
	var api_url string = url + "?page=" + strconv.FormatInt(p, 10)
	var per_page int64 = d.Per_page
	var page int64 = d.Page
	var total int64 = d.Total
	var total_page int64 = total/per_page + 1

	if page <= total_page {
		articles := grab.GetArticles(api_url)
		for _, t := range articles.Data {
			grab.GenerateStaticArticle(t)
		}

		//嵌套调用
		processArticle(url, articles.Page+1, articles)
	}
}

func main() {
	config := tools.ParseConfig()
	for _, v := range config.Api {

		//抓取文章数据
		articles := grab.GetArticles(v.Article)
		processArticle(v.Article, articles.Page, articles)

		//抓取分类数据
		category := grab.GetCategorys(v.Category)
		grab.GenerateStaticCategory(category)
	}
}
