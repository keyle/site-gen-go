package main

import (
	"log"
	"sort"
	"strings"
	"time"
)

func ParseSelector(selector, haystack string) string {
	openingTag := "<" + selector + ">"
	closingTag := "</" + selector + ">"

	// Splitting by the opening tag and getting the last part
	parts := strings.Split(haystack, openingTag)
	if len(parts) < 2 {
		return "" // Tag not found
	}

	// Splitting the result by the closing tag and getting the first part
	contentParts := strings.Split(parts[1], closingTag)
	if len(contentParts) == 0 {
		return "" // Malformed HTML
	}

	return contentParts[0]
}

func PubDateFromNow() string {
	return time.Now().Format("2006-01-02 15:04")
}

func PubdateForBlogIndex(date string) string {
	dt, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatal("Could not format date for Blog Index: ", date)
	}

	return dt.Format("Jan 02, 2006")
}

func PubDateForSitemap(date string) string {
	parts := strings.Split(date, " ")
	if len(parts) > 0 {
		return parts[0]
	}
	log.Fatal("Could not extract date from pubDate for Sitemap: ", date)
	return ""
}

func GetBlogsSorted(articles []Article) []Article {
	var blogs []Article

	for _, article := range articles {
		if article.IsBlog {
			blogs = append(blogs, article)
		}
	}
	sort.Slice(blogs, func(i, j int) bool {
		timeI, errI := time.Parse("2006-01-02", blogs[i].PubDate)
		timeJ, errJ := time.Parse("2006-01-02", blogs[j].PubDate)
		if errI != nil || errJ != nil {
			log.Fatal("Could not sort blogs by pubDate desc")
		}
		return timeI.After(timeJ)
	})

	return blogs
}
