package main

import (
	"fmt"
	"time"
)

type Cache struct {
	data map[string]interface{}
	time map[string]time.Time
}

func (c *Cache) Set(key string, value interface{}, seconds int) {
	if c.data == nil {
		c.data = make(map[string]interface{})
		c.time = make(map[string]time.Time)
	}
	c.data[key] = value
	c.time[key] = time.Now().Add(time.Duration(seconds) * time.Second)
}

func (c *Cache) Get(key string) interface{} {
	if c.data == nil {
		return nil
	}

	expireTime, exists := c.time[key]
	if !exists || time.Now().After(expireTime) {
		delete(c.data, key)
		delete(c.time, key)
		return nil
	}

	return c.data[key]
}

func main() {
	var cache Cache

	cache.Set("data1", "значение1", 2)
	cache.Set("data2", "значение2", 4)

	fmt.Println("Сразу:")
	fmt.Println("data1:", cache.Get("data1"))
	fmt.Println("data2:", cache.Get("data2"))

	time.Sleep(3 * time.Second)

	fmt.Println("После 3 секунд:")
	fmt.Println("data1:", cache.Get("data1"))
	fmt.Println("data2:", cache.Get("data2"))

	time.Sleep(2 * time.Second)

	fmt.Println("После 5 секунд:")
	fmt.Println("data1:", cache.Get("data1"))
	fmt.Println("data2:", cache.Get("data2"))
}
