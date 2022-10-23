package models

import "encoding/xml"

type UrlSet struct {
	XMLName xml.Name `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
	Locations []Url `xml:"loc"`
}
type Url struct {
	U string `xml:"url"`
}