package utils

import "strings"

type ContestParser struct {
	Platform string
	Contest  string
}

// https://atcoder.jp/contests/abc377
func NewContestParser(link string) *ContestParser {
	splitLink := strings.Split(link, "/")
	return &ContestParser{
		Platform: splitLink[2],
		Contest:  splitLink[4],
	}
}
