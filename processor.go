package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday/v2"
)

func FindMarkdownFiles(filePath string) ([]string, error) {
	var files []string

	err := filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".md") {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func CreateArticles(markdownFiles []string) []Article {
	var articles []Article

	for _, file := range markdownFiles {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Println("Error reading file:", file, err)
			continue
		}
		articles = append(articles, Article{
			Path:     filepath.Dir(file),
			File:     filepath.Base(file),
			Markdown: string(content),
		})
	}
	return articles
}

func ProcessArticles(articles []Article, settings *Settings) []Article {

	var processed []Article
	// Create a renderer with specific HTMLFlags
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags: blackfriday.CommonHTMLFlags & ^blackfriday.SmartypantsFractions,
	})

	// Set up the extensions
	extensions := blackfriday.CommonExtensions

	templateContentsIndex, err := os.ReadFile(settings.TemplateIndex)
	if err != nil {
		log.Fatal("Could not load template index")
	}

	templateContents, err := os.ReadFile(settings.Template)
	if err != nil {
		log.Fatal("Could not load template index")
	}

	for _, article := range articles {

		article.IsBlog = strings.Contains(article.Markdown, "<x-blog-title>")

		if strings.Contains(article.Markdown, "<x-index/>") {
			article.Html = string(templateContentsIndex)
		} else {
			article.Html = string(templateContents)
		}

		converted := blackfriday.Run([]byte(article.Markdown), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(extensions))
		article.Html = strings.Replace(article.Html, settings.ContentTag, string(converted), 1)

		var mdTitleTag string
		if article.IsBlog {
			mdTitleTag = "x-blog-title"
			article.Html = strings.Replace(article.Html, "<body>", "<body class='blog'>", 1)
			article.PubDate = ParseSelector("sub", article.Markdown)
		} else {
			mdTitleTag = "x-title"
			article.PubDate = PubDateFromNow()
		}

		article.Title = ParseSelector(mdTitleTag, article.Markdown)
		article.Html = strings.Replace(article.Html, settings.TitleTag, article.Title, 1)

		article.Tags = ParseSelector("x-tags", article.Markdown)
		article.Html = strings.Replace(article.Html, settings.KeywordsTag, article.Tags, 1)

		article.Description = ParseSelector("x-desc", article.Markdown)
		article.Html = strings.Replace(article.Html, settings.DescriptionTag, article.Description, 2)

		article.Url = (settings.WebRoot + article.Path)
		article.Url = strings.Replace(article.Url, settings.WorkDir, "", 1)
		article.Url = strings.Replace(article.Url, article.File, "", 1)

		target := article.Path + "/index.html"

		err := os.WriteFile(target, []byte(article.Html), 0644)
		if err != nil {
			log.Fatal("Could not save article HTML")
		}
		processed = append(processed, article)
	}
	return processed
}
