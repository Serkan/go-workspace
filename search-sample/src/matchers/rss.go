package matchers

import (
	"net/http"
	"search"
	"log"
	"encoding/xml"
	"regexp"
)

type rssMatcher struct{}

func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

type (
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string `xml:"pubDate"`
		Title       string `xml:"title"`
		Description string  `xml:"description"`
		Link        string `xml:"link"`
		GUID        string `xml:"guid"`
		GeoRssPoint string `xml:"georss:point"`
	}

	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string `xml:"url"`
		Title   string `xml:"title"`
		Link    string `xml:"link"`
	}

	// in the rss document.
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string  `xml:"title"`
		Description    string `xml:"description"`
		Link           string `xml:"link"`
		PubDate        string `xml:"pubDate"`
		LastBuildDate  string `xml:"lastBuildDate"`
		TTL            string `xml:"ttl"`
		Language       string `xml:"language"`
		ManagingEditor string `xml:"managingEditor"`
		WebMaster      string `xml:"webMaster"`
		Image          image `xml:"image"`
		Item           []item `xml:"item"`
	}

	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel `xml:"channel"`
	}

)

func (m rssMatcher) Search(feed *search.Feed, term string) ([]*search.Result, error) {
	res, err := http.Get(feed.URI)
	if err != nil {
		log.Fatalf(err)
	}
	defer res.Close()

	if res.Status != 200 {
		log.Fatal("RSS matcher can not reach the site " + feed.URI)
		return nil, nil
	}
	document := rssDocument{}
	xml.NewDecoder(res).Decode(&document)

	results := []*search.Result{}
	for _, item := range document.Channel.Item {
		matched, err := regexp.Match(term, item.Description)
		if err != nil {
			return nil, nil
		}
		if matched {
			results = append(results, search.Result{
				Field : "Description",
				Content: item.Description,
			})
		}
	}
	return results
}
