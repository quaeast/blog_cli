package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var blogNameArg = flag.String("n", "New Blog", "添加新博客")

func createNewBlog(blogName *string)  {
	os.Mkdir(*blogName, os.ModePerm)
	os.Create("./" + *blogName + "/main.md")

	f, err := os.OpenFile("main.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("\n### [%s](./%s/main.md)\n", *blogName, *blogName))
}

func main() {
	flag.Parse()
	createNewBlog(blogNameArg)
}
