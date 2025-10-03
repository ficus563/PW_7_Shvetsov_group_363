package main

import (
	"fmt"
	"time"
)

type handler func(data interface{})

type EventBus struct {
	subs map[string][]handler
}

func (eb *EventBus) on(event string, h handler) {
	eb.subs[event] = append(eb.subs[event], h)
}

func (eb *EventBus) send(event string, data interface{}) {
	for _, h := range eb.subs[event] {
		go h(data)
	}
}

func main() {
	bus := EventBus{
		subs: make(map[string][]handler),
	}

	h1 := func(data interface{}) {
		fmt.Println("Событие:", data)
	}

	bus.on("msg", h1)
	bus.send("msg", "тест")

	time.Sleep(10 * time.Millisecond)
}
