package parser

import (
	"Go-Spider/vxiangqin/engine"
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
		/*for i := 1; i <= 25; i++ {
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
		}*/ //去掉上面重复请求的数据

		// 遍历年龄范围页面
		//0-20...20-21..21-22...22-23......79-80...80-0
		for i := 19; i <= 80; i++ {
			startAge := 0
			endAge := 0
			if i == 19 {
				startAge = 0
			} else {
				startAge = i
			}

			if i == 80 {
				endAge = 0
			} else {
				endAge = i + 1
			}

			for j := 1; j <= 25; j++ {
				userPageUrl := string(m[1]) + `/p/user.php?form_mate_sex=0&form_mate_age1=` + strconv.Itoa(startAge) + `&form_mate_age2=` + strconv.Itoa(endAge) + `&form_mate_heigh1=0&form_mate_heigh2=0&m1=0&m2=0&m3=0&areatitle=&form_mate_job=0&form_mate_edu=0&form_mate_love=0&form_mate_house=0&form_mate_pay=0&t=1&areakey=&p=` + strconv.Itoa(j)
				result.Requests = append(result.Requests, engine.Request{
					Url:        userPageUrl,
					ParserFunc: ParseCity,
				})
			}
		}
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
