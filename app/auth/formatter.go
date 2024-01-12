package auth

import "time"

type LoginFormatter struct {
	Token   string     `json:"token"`
	Expires *time.Time `json:"expires"`
}

func LoginFormat(token string, expires time.Time) LoginFormatter {
	var formatter LoginFormatter

	formatter.Token = token
	formatter.Expires = &expires

	return formatter
}
