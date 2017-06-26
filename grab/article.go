package grab

import (
	"encoding/json"
	"fmt"
	"generator/tools"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const static_file_path = "./static/dist/"
const article_tpl = "./static/tpl/article.tpl"

type Data struct {
	Data     []Article
	Page     int64
	Per_page int64
	Total    int64
}

type Article struct {
	Id         int64
	Title      string
	Slug       string
	Summary    string
	Views      int64
	User       string
	Body       string
	Created_at time.Time
	Updated_at time.Time
}

func GetArticles(api string) (d Data) {
	result := tools.HttpGet(api)

	err := json.Unmarshal([]byte(result), &d)
	if err != nil {
		fmt.Println("Failed unmarshalling json: %s\n", err)
		return
	}
	return d
}

func GenerateStaticArticle(a Article) {
	var id int64 = a.Id
	var static_file string = static_file_path + strconv.FormatInt(id, 10) + ".html"
	var f *os.File
	var err error

	if tools.CheckFileIsExist(static_file) {
		f, err = os.OpenFile(static_file, os.O_RDWR|os.O_TRUNC, 0666) //打开文件
	} else {
		f, err = os.Create(static_file) //创建文件
	}

	defer f.Close()
	if err != nil {
		panic(err)
	}

	tpl_content, _ := tools.ReadFile(article_tpl)
	content := string(tpl_content)
	content = strings.Replace(content, "{{header}}", a.Title+" - markbest.site", -1)
	content = strings.Replace(content, "{{title}}", a.Title, -1)
	content = strings.Replace(content, "{{created_at}}", a.Created_at.Format("01-02"), -1)
	content = strings.Replace(content, "{{author}}", a.User, -1)
	content = strings.Replace(content, "{{views}}", strconv.FormatInt(a.Views, 10), -1)
	content = strings.Replace(content, "{{body}}", a.Body, -1)

	io.WriteString(f, content) //写入文件
}
