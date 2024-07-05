package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFile := "/path/to/your/file.csv"
	xmlFile := "/path/to/save/wordpress_posts.xml"

	// Read the CSV file
	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	// Create the RSS feed
	rss := RSS{
		Version: "2.0",
		Channel: Channel{
			Title:       "Your Site Title",
			Link:        "https://www.yoursite.com",
			Description: "Your Site Description",
			PubDate:     time.Now().Format(time.RFC1123Z),
			Language:    "en-US",
			WXRVersion:  "1.2",
			BaseSiteURL: "https://www.yoursite.com",
			BaseBlogURL: "https://www.yoursite.com",
			Generator:   "https://wordpress.org/?v=6.5.4",
			Author: Author{
				AuthorID:          "1",
				AuthorLogin:       "author_login",
				AuthorEmail:       "author@example.com",
				AuthorDisplayName: "Author Name",
				AuthorFirstName:   "Author",
				AuthorLastName:    "Name",
			},
		},
	}

	// Add posts to the RSS feed
	for i, record := range records[1:] { // Skip header row
		title := record[0]
		content := record[1]

		item := Item{
			Title:           title,
			Link:            fmt.Sprintf("https://www.yoursite.com/%s/", strings.Replace(strings.ToLower(title), " ", "-", -1)),
			PubDate:         time.Now().Format(time.RFC1123Z),
			Creator:         "author_login",
			GUID:            fmt.Sprintf("https://www.yoursite.com/?post_type=post&#038;p=%d", i),
			Description:     "",
			Content:         fmt.Sprintf("<![CDATA[<p>%s</p>]]>", content),
			Excerpt:         "",
			PostID:          fmt.Sprintf("%d", i),
			PostDate:        time.Now().Format("2006-01-02 15:04:05"),
			PostDateGMT:     time.Now().UTC().Format("2006-01-02 15:04:05"),
			PostModified:    time.Now().Format("2006-01-02 15:04:05"),
			PostModifiedGMT: time.Now().UTC().Format("2006-01-02 15:04:05"),
			CommentStatus:   "closed",
			PingStatus:      "closed",
			PostName:        strings.Replace(strings.ToLower(title), " ", "-", -1),
			Status:          "publish",
			PostParent:      "0",
			MenuOrder:       "0",
			PostType:        "post",
			PostPassword:    "",
			IsSticky:        "0",
		}

		rss.Channel.Items = append(rss.Channel.Items, item)
	}

	// Create the XML file
	output, err := xml.MarshalIndent(rss, "", "    ")
	if err != nil {
		fmt.Println("Error generating XML:", err)
		return
	}

	// Write to file
	err = os.WriteFile(xmlFile, output, 0644)
	if err != nil {
		fmt.Println("Error writing XML file:", err)
		return
	}

	fmt.Println("XML file created successfully.")
}
