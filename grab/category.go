package grab

import (
	"encoding/json"
	"fmt"
	"generator/tools"
	"io"
	"os"
	"strconv"
	"strings"
)

const category_tpl = "./static/tpl/category.tpl"

type Category struct {
	Id           int64
	Title        string
	Articles     []SubArticle
	Sub_category []Category
}

type SubArticle struct {
	Id    int64
	Title string
}

func GetCategorys(api string) (c []Category) {
	result := tools.HttpGet(api)

	err := json.Unmarshal([]byte(result), &c)
	if err != nil {
		fmt.Println("Failed unmarshalling json: %s\n", err)
		return
	}
	return c
}

func GenerateCategoryLevel(c []Category) (string, string) {
	var content string
	var default_page string = ""
	content = "<ul class=\"tree\">"

	for _, v := range c {
		content = content + "<li><span>" + v.Title + "</span>"
		if v.Articles != nil {
			content = content + "<ul>"
			for _, sa := range v.Articles {
				var sid string = strconv.FormatInt(sa.Id, 10)
				content = content + "<li><a title='" + sa.Title + "' href='" + sid + ".html'>" + sa.Title + "</a></li>"
				if default_page == "" {
					default_page = sid
				}
			}
			content = content + "</ul>"
		}
		if v.Sub_category != nil {
			content = content + "<ul>"
			for _, s := range v.Sub_category {
				content = content + "<li><span>" + s.Title + "</span>"
				if s.Articles != nil {
					content = content + "<ul>"
					for _, a := range s.Articles {
						var aid string = strconv.FormatInt(a.Id, 10)
						content = content + "<li><a title='" + a.Title + "' href='" + aid + ".html'>" + a.Title + "</a></li>"
						if default_page == "" {
							default_page = aid
						}
					}
					content = content + "</ul>"
				}
				content = content + "</li>"
			}
			content = content + "</ul>"
		}
		content = content + "</li>"
	}

	content = content + "</ul>"
	return content, default_page
}

func GenerateStaticCategory(c []Category) {
	var static_file string = static_file_path + "index.html"
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

	tpl_content, _ := tools.ReadFile(category_tpl)
	content := string(tpl_content)
	category, default_page := GenerateCategoryLevel(c)
	content = strings.Replace(content, "{{category}}", category, -1)
	content = strings.Replace(content, "{{default_page}}", default_page+".html", -1)
	content = strings.Replace(content, "{{footer}}", "Copyright © 2015 - 2017 markbest.site - 你的指尖有改变世界的力量 - All Rights Reserved.", -1)

	io.WriteString(f, content) //写入文件
}
