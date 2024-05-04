package entities

type UrlRecord struct {
	OriginalUrl string `json:"original"`
	Alias       string `json:"alias"`
	Visits      int    `json:"visits"`
}
