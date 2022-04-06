package parser

import (
	"Go-Reptile/vxiangqin/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(`<a href="(http://[a-z]+.vxiangqin.com/u/[0-9]+)" class="mbox">`)

func ParseCity(contents []byte) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		id := string(m[1][len(m[1])-6:])
		result.Items = append(result.Items, "User Id "+id)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, id)
			},
		})
	}
	return result
}
