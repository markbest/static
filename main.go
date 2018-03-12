package main

import (
	"fmt"
	"github.com/markbest/static/grab"
	"github.com/markbest/static/tools"
	"runtime"
	"strconv"
	"sync"
)

func processArticle(url string, wgp *sync.WaitGroup) {
	articles := grab.GetArticles(url)
	for _, t := range articles.Data {
		grab.GenerateStaticArticle(t)
		fmt.Println("ID: " + strconv.FormatInt(t.Id, 10) + " article grab complete")
	}
	wgp.Done()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wgp := &sync.WaitGroup{}

	var config = tools.ParseConfig()
	var articleApi, categoryApi string
	for _, v := range config.Api {
		articleApi = v.Article
		categoryApi = v.Category
	}

	//抓取分类数据
	wgp.Add(1)
	category := grab.GetCategorys(categoryApi)
	go grab.GenerateStaticCategory(category, wgp)

	//抓取文章数据
	var totalPage, i int64
	articles := grab.GetArticles(articleApi)
	totalPage = articles.Total/articles.Per_page + 1
	for i = 1; i <= totalPage; i++ {
		wgp.Add(1)
		go processArticle(articleApi+"?page="+strconv.FormatInt(i, 10), wgp)
	}

	//等待抓取完成
	wgp.Wait()
}
