package main

type Settings struct {
	WorkDir        string `json:"workdir"`
	WebRoot        string `json:"webroot"`
	Template       string `json:"template"`
	TemplateIndex  string `json:"templateindex"`
	ContentTag     string `json:"contenttag"`
	TitleTag       string `json:"titletag"`
	DescriptionTag string `json:"descriptiontag"`
	KeywordsTag    string `json:"keywordstag"`
}

type Article struct {
	Path        string
	File        string
	Markdown    string
	Html        string
	Title       string
	IsBlog      bool
	Url         string
	PubDate     string
	Description string
	Tags        string
}
