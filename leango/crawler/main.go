package main

import (
	"leango/crawler/engine"
	"leango/crawler/persist"
	"leango/crawler/scheduler"
	"leango/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{Url: "http://www.zhenai.com/zhenhun", Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList")})
}
