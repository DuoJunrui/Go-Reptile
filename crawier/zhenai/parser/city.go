package parser

import (
	"Go-Spider/crawier/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		id := string(m[1][26:])
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Request = append(result.Request, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, id, name)
			},
		})
	}

	return result
}
