package main

import (
	"log"

	"github.com/appleboy/pyroscope/pkg/agent/profiler"
)

func work(n int) {
	// revive:disable:empty-block this is fine because this is a example app, not real production code
	for i := 0; i < n; i++ {

	}
	// revive:enable:empty-block
}

func fastFunction() {
	work(2000)
}

func slowFunction() {
	work(8000)
}

func main() {
	profiler.Start(profiler.Config{
		ApplicationName: "simple.golang.app",
		ServerAddress:   "http://pyroscope:4040", // this will run inside docker-compose, hence `pyroscope` for hostname
	})

	log.Println("test")
	for {
		fastFunction()
		slowFunction()
	}
}
