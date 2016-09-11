// TODO: Consider placement of package, encoding? encoding/xml? (take in mind atom or other similar schemes)
// TODO: Consider separate rss1 and rss2?
package rss

import (
	"encoding/xml"
)

// TODO: Add error handling: if not rss? supported rss version? missing mandatory elements? (both encode/decode case)
// TODO: Correct any elements that contains structs instead of pure string (e.g. image)
// TODO: Consider renaming TopElement to FeedElement? Or something else? rss.TopElement right now..
type TopElement struct {
	XMLName xml.Name `xml:"rss"`

	Version string         `xml:"version,attr"`
	Channel ChannelElement `xml:"channel"`
}

type ChannelElement struct {
	XMLName xml.Name `xml:"channel"`

	// Mandatory channel elements
	Title       string `xml:"title""`
	Link        string `xml:"link"`
	Description string `xml:"description"`

	// Optional channel elements TODO: Make sure they are indeed optional for encoding!
	Items          []ItemElement `xml:"item"`
	Language       string        `xml:"language"`
	PubDate        string        `xml:"pubDate"`
	LastBuildDate  string        `xml:"lastBuildDate"`
	Docs           string        `xml:"docs"`
	Generator      string        `xml:"generator"`
	ManagingEditor string        `xml:"managingEditor"`
	WebMaster      string        `xml:"webMaster"`
	Category       string        `xml:"category"`
	Copyright      string        `xml:"copyright"`
	Cloud          string        `xml:"cloud"`
	TimeToLive     string        `xml:"ttl"`
	Image          string        `xml:"image"`
	Rating         string        `xml:"rating"`
	TextInput      string        `xml:"textInput"`
	SkipHours      string        `xml:"skipHours"`
	SkipDays       string        `xml:"skipDays"`
}

type ItemElement struct {
	XMLName xml.Name `xml:"item"`

	// ONE of title and description are mandatory
	Title       string `xml:"title"`
	Description string `xml:"description"`

	// Optional item elements TODO: Make sure they are indeed optional for encoding!
	Link      string `xml:"link"`
	Author    string `xml:"author"`
	Category  string `xml:"category"`
	Comments  string `xml:"comments"`
	Enclosure string `xml:"enclosure"`
	Guid      string `xml:"guid"`
	PubDate   string `xml:"pubDate"`
	Source    string `xml:"source"`
}

// FIXME: how to utilize interface? And should it?
func Encode(t TopElement) ([]byte, error) {
	return xml.MarshalIndent(t, "", "\t")
}

// FIXME: Should decode mimic xml interface insteda?
// And thus removing the need to create TopElement here
func Decode(data []byte) (TopElement, error) {
	var top TopElement
	err := xml.Unmarshal(data, &top)

	return top, err
}
