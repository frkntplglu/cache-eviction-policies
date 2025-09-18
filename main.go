package main

import (
	"github.com/frkntplglu/cache/pkg/fifo"
	"github.com/frkntplglu/cache/pkg/lfu"
	"github.com/frkntplglu/cache/pkg/lru"
)

type Post struct {
	Id      int
	Title   string
	Content string
}

func main() {
	posts := []Post{
		{Id: 1, Title: "Go Generics", Content: "Go 1.18 ile generics geldi."},
		{Id: 2, Title: "LRU Cache", Content: "Least Recently Used cache algoritması."},
		{Id: 3, Title: "Event-Driven", Content: "Event-driven mimari ile loose coupling."},
		{Id: 4, Title: "Kubernetes", Content: "K8s ile container orkestrasyonu."},
		{Id: 5, Title: "Microservices", Content: "Monolith'ten microservice'e geçiş."},
		{Id: 6, Title: "DDD", Content: "DDD implementasyon."},
	}

	cacheLRU := lru.New[int, Post](4)
	cacheLFU := lfu.New[int, Post](4)
	cacheFIFO := fifo.New[int, Post](4)

	for _, post := range posts {
		cacheLRU.Put(post.Id, post)
		cacheLFU.Put(post.Id, post)
		cacheFIFO.Put(post.Id, post)
	}

	cacheLRU.PrintAll()
	cacheLFU.PrintAll()
	cacheFIFO.PrintAll()

}
