package parser

import (
	"Go-Reptile/vxiangqin/engine"
	"regexp"
	"strconv"
)

const cityListRe = `<a href='(http://[a-z]+.vxiangqin.com)'[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 1 //只查一个城市
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		// 遍历男女混合页面
		for i := 1; i <= 25; i++ {
			userPageUrl := string(m[1]) + `/p/user.php?p=` + strconv.Itoa(i)
			result.Requests = append(result.Requests, engine.Request{
				Url:        userPageUrl,
				ParserFunc: ParseCity,
			})
		}
		// 遍历男性页面
		for i := 1; i <= 25; i++ {
			userPageUrl := string(m[1]) + `/p/user.php?form_mate_sex=1&t=1&p=` + strconv.Itoa(i)
			result.Requests = append(result.Requests, engine.Request{
				Url:        userPageUrl,
				ParserFunc: ParseCity,
			})
		}
		// 遍历女性页面
		for i := 1; i <= 25; i++ {
			userPageUrl := string(m[1]) + `/p/user.php?form_mate_sex=2&t=1&p=` + strconv.Itoa(i)
			result.Requests = append(result.Requests, engine.Request{
				Url:        userPageUrl,
				ParserFunc: ParseCity,
			})
		}
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
