package parser

import (
	"Go-Reptile/vxiangqin/engine"
	"Go-Reptile/vxiangqin/model"
	"log"
	"regexp"
)

const baseInfoRe = `<div class="udata">[\s]*<li>([^<]+)</li>[\s]*<li>([^<]+)</li>[\s]*<li>([^<]+)</li>[\s]*<li>([^<]+)</li>[\s]*<li>([^<]+)</li>[\s]*<li>([^<]+)</li>[\s]*<li>([^<]+)</li>[\s]*<li>([^<]+)</li>[\s]*</div>`

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
