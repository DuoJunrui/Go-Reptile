package persist

import (
	"Go-Reptile/vxiangqin/model"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestSaver(t *testing.T) {
	expected := model.Profile{
		UserId:               "123456",
		Name:                 "小马过河",
		Sex:                  "男",
		Age:                  "36岁",
		Marriage:             "离异",
		Height:               "168cm",
		Weight:               "52kg",
		Income:               "8-12千",
		Education:            "大专",
		Occupation:           "机械工",
		House:                "已购房",
		Car:                  "已买车",
		ExpectedMarriageDate: "期待两年内结婚",
	}

	id, err := save(expected)
	if err != nil {
		panic(err)
	}

	//TODO: Try to start up elastic search
	//here using docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index("user_profile").
		Type("vxiangqin").
		Id(id).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)
	var actual model.Profile
	err = json.Unmarshal(resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("get %v; expected %v", actual, expected)
	}

}
