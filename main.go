package main

import (
	"log"
)

func main() {
	settings, err := LoadSettings()
	if err != nil {
		log.Fatal("could not load settings.json ", err)
	}
	markdownFiles, err := FindMarkdownFiles(settings.WorkDir)
	if err != nil {
		log.Fatal("could not load markdown files ", err)
	}

	articles := CreateArticles(markdownFiles)
	articles = ProcessArticles(articles, settings)

	generateBlogIndex(articles, settings)
	generateSitemap(articles, settings)
	generateRSSFeed(articles, settings)
}
