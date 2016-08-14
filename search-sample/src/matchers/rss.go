package matchers

import (
	"net/http"
	"search"
	"log"
	"encoding/xml"
	"regexp"
	"errors"
)

type rssMatcher struct{}

func init() {
	// rss matcher must be registered before the "main" event
	log.Println("Rss matcher init called")
	/*
	we declare a rssMatcher (even it does not contain any field this type "decorated" with matcher interface
	in this file, and this is equivalent of "class RssMatcher implements Matcher" in java)
	 */
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
	res, err := http.Get(feed.URI) // make a "GET" HTTP request to given URI
	if err != nil {
		log.Fatal(err.Error())
	}
	// defer calls always called after parent function call
	defer res.Body.Close() // dont forget the call close (something like InputStream.close())

	if res.Status != "200" {
		log.Fatal("RSS matcher can not reach the site " + feed.URI)
		return nil, nil
	}
	document := rssDocument{} // init a empty XML document structure
	xml.NewDecoder(res.Body).Decode(&document) // fill this structure by passing its address

	results := []*search.Result{}
	for _, item := range document.Channel.Item {
		// unused variables in return values MUST be disappeared with "_"
		matched, err := regexp.MatchString(term, item.Description)
		if err != nil {
			return nil, nil
		}
		if matched {
			results = append(results, &search.Result{// object literals are great to use (we wait it in java as well)
				Field : "Description",
				Content: item.Description,
			})
		}
	}
	/*
	There is no exception mechanism, every error-possible call return with an "error" type variable. So
	 */
	err = errors.New("Error occurred while using rss matcher") // create a error structure and return it
	return results, err
}
