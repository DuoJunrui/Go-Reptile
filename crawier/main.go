package main

import (
	"Go-Reptile/crawier/engine"
	"Go-Reptile/crawier/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
