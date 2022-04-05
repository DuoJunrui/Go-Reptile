package parser

import (
	"Go-Reptile/crawier/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "1141781127", "百里千寻")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
		Data: model.Data{
			Id:         1141781127,
			Name:       "百里千寻",
			Gender:     "男士",
			Age:        40,
			Height:     "168cm",
			Weight:     "65kg",
			Income:     "5001-8000元",
			Marriage:   "离异",
			Education:  "大学本科",
			Occupation: "生产/制造",
			Jiguan:     "安徽淮南",
			WorkCity:   "淮南凤台",
			Xinzuo:     "魔羯座(12.22-01.19)",
			House:      "已购房",
			Car:        "已买车",
		},
	}
	if profile != expected {
		t.Errorf("excepted %v: but was %v", expected, profile)
	}
}
