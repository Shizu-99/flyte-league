package models

import (
	"encoding/xml"
)

type Archers struct {
	XMLName xml.Name `xml:"archers"`
	Archers []Archer `xml:"archer"`
}

type Archer struct {
	XMLName xml.Name `xml:"archer"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	School  string   `xml:"school"`
	Team    string   `xml:"team"`
	Score   Score    `xml:"score"`
}

type Score struct {
	XMLName xml.Name `xml:"score"`
	Num_X   int      `xml:"num_x,attr"`
	Num_10  int      `xml:"num_10,attr"`
	Hits    int      `xml:"hits,attr"`
	Total   int      `xml:"total"`
}
