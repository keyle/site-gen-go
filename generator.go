package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func generateBlogIndex(articles []Article, settings *Settings) {
	var builder strings.Builder

	blogs := GetBlogsSorted(articles)

	builder.WriteString("<table>")

	for _, blog := range blogs {
		builder.WriteString(fmt.Sprintf("<tr><td>%s</td><td><a href='%s'>%s</a></td><td>&nbsp;</td></tr>\n", PubdateForBlogIndex(blog.PubDate), blog.Url, blog.Title))
	}

	builder.WriteString("</table>")

	table := builder.String()

	indexLocation := settings.WorkDir + "/index.html"
	indexb, err := os.ReadFile(indexLocation)
	if err != nil {
		log.Fatal("Could not load index.html")
	}

	index := string(indexb)
	index = strings.Replace(index, "<x-blog-index/>", table, 1)
	err = os.WriteFile(indexLocation, []byte(index), 0644)
	if err != nil {
		log.Fatal("Could not save Blog Index")
	}
}

func generateSitemap(articles []Article, settings *Settings) {
	var builder strings.Builder

	blogs := GetBlogsSorted(articles)

	builder.WriteString(`<?xml version="1.0" encoding="utf-8" standalone="yes"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
  xmlns:xhtml="http://www.w3.org/1999/xhtml">`)

	for _, blog := range blogs {
		builder.WriteString(fmt.Sprintf("\n<url><loc>%s</loc><lastmod>%s</lastmod></url>", blog.Url, PubDateForSitemap(blog.PubDate)))
	}
	builder.WriteString("\n</urlset>")

	sitemapLocation := settings.WorkDir + "/sitemap.xml"

	err := os.WriteFile(sitemapLocation, []byte(string(builder.String())), 0644)
	if err != nil {
		log.Fatal("Could not save sitemap.xml")
	}
}

func generateRSSFeed(articles []Article, settings *Settings) {
	var builder strings.Builder

	blogs := GetBlogsSorted(articles)

	builder.WriteString(`<?xml version="1.0" encoding="utf-8" standalone="yes"?>
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
  <channel>
    <title>NobenLog</title>
    <link>https://noben.org/blog/</link>
    <description>Recent content on NobenLog</description>
    <generator>site-gen-go -- https://github.com/keyle/site-gen-go</generator>
    <language>en-us</language>`)

	for _, blog := range blogs {
		builder.WriteString(fmt.Sprintf("<item><title>%s</title><link>%s</link><pubDate>%s</pubDate><guid>%s</guid><description><![CDATA[ %s ]]></description></item>\n", blog.Title, blog.Url, blog.PubDate, blog.Url, blog.Description))
	}
	builder.WriteString("</channel></rss>\n")

	FeedLocation := settings.WorkDir + "/index.xml"

	err := os.WriteFile(FeedLocation, []byte(string(builder.String())), 0644)
	if err != nil {
		log.Fatal("Could not save RSS feed")
	}
}
