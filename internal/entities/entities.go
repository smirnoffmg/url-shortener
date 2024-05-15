package entities

import "gorm.io/gorm"

type UrlRecord struct {
	gorm.Model
	OriginalUrl string `json:"original"`
	Alias       string `json:"alias" gorm:"uniqueIndex"`
}

type Visit struct {
	gorm.Model
	Alias     string `json:"alias" gorm:"index"`
	IpAddr    string `json:"ip_addr"`
	UserAgent string `json:"user_agent"`
}
