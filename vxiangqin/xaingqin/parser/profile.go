package parser

import (
	"Go-Reptile/vxiangqin/engine"
	"Go-Reptile/vxiangqin/model"
	"log"
	"regexp"
)

const baseInfoRe = `<div class="udata">[\s]*<li>([^<]+)</li>[\s]*<li>([^<]+)</li>[\s]*<li>([^<]+)</li>[\s]*<li>([^<]+)</li>[\s]*<li>([^<]+)</li>[\s]*<li>([^<]+)</li>[\s]*<li>([^<]+)</li>[\s]*<li>([^<]+)</li>[\s]*</div>`

var houseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>([^<]+房)</div>`)
var nameRe = regexp.MustCompile(`1616742751"[^>]*>([^<]+)<font class="S14 C999">`)
var sexRe = regexp.MustCompile(`up/p/img/grade([0-9]+)`)
var weightRe = regexp.MustCompile(`<dt>体　　重：</dt><dd>([0-9]+kg)</dd>`)
var carRe = regexp.MustCompile(`<dl><dt>买车情况：</dt><dd>([^>]*[^<]+)</dd></dl>`)

func ParseProfile(contents []byte, id string) engine.ParseResult {
	re := regexp.MustCompile(baseInfoRe)
	matches := re.FindAllSubmatch(contents, -1)

	if 0 == len(matches) {
		log.Printf("profile matches is empty")
		return engine.ParseResult{
			Items: []interface{}{},
		}
	}

	profile := model.Profile{}
	profile.UserId = id
	profile.Name = extractString(contents, nameRe)
	sexResult := extractString(contents, sexRe)
	if "21" == sexResult {
		profile.Sex = "女"
	} else if "11" == sexResult {
		profile.Sex = "男"
	} else {
		profile.Sex = "其他"
	}
	profile.Weight = extractString(contents, weightRe)
	profile.Car = extractString(contents, carRe)

	for i := 0; i < len(matches[0]); i++ {
		switch i {
		case 1:
			profile.Age = string(matches[0][1])
		case 2:
			profile.Marriage = string(matches[0][2])
		case 3:
			profile.Height = string(matches[0][3])
		case 4:
			profile.Income = string(matches[0][4])
		case 5:
			profile.Education = string(matches[0][5])
		case 6:
			profile.Occupation = string(matches[0][6])
		case 7:
			profile.House = string(matches[0][7])
		case 8:
			profile.ExpectedMarriageDate = string(matches[0][8])
		}
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
