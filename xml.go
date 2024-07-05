package main

import "encoding/xml"

// Structs to define the XML structure
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Language    string `xml:"language"`
	WXRVersion  string `xml:"wp:wxr_version"`
	BaseSiteURL string `xml:"wp:base_site_url"`
	BaseBlogURL string `xml:"wp:base_blog_url"`
	Generator   string `xml:"generator"`
	Author      Author `xml:"wp:author"`
	Items       []Item `xml:"item"`
}

type Author struct {
	AuthorID          string `xml:"wp:author_id"`
	AuthorLogin       string `xml:"wp:author_login"`
	AuthorEmail       string `xml:"wp:author_email"`
	AuthorDisplayName string `xml:"wp:author_display_name"`
	AuthorFirstName   string `xml:"wp:author_first_name"`
	AuthorLastName    string `xml:"wp:author_last_name"`
}

type Item struct {
	Title           string `xml:"title"`
	Link            string `xml:"link"`
	PubDate         string `xml:"pubDate"`
	Creator         string `xml:"dc:creator"`
	GUID            string `xml:"guid"`
	Description     string `xml:"description"`
	Content         string `xml:"content:encoded"`
	Excerpt         string `xml:"excerpt:encoded"`
	PostID          string `xml:"wp:post_id"`
	PostDate        string `xml:"wp:post_date"`
	PostDateGMT     string `xml:"wp:post_date_gmt"`
	PostModified    string `xml:"wp:post_modified"`
	PostModifiedGMT string `xml:"wp:post_modified_gmt"`
	CommentStatus   string `xml:"wp:comment_status"`
	PingStatus      string `xml:"wp:ping_status"`
	PostName        string `xml:"wp:post_name"`
	Status          string `xml:"wp:status"`
	PostParent      string `xml:"wp:post_parent"`
	MenuOrder       string `xml:"wp:menu_order"`
	PostType        string `xml:"wp:post_type"`
	PostPassword    string `xml:"wp:post_password"`
	IsSticky        string `xml:"wp:is_sticky"`
}
