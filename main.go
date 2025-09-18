package main

import (
	"github.com/frkntplglu/cache/pkg/lru"
)

type Post struct {
	Id      int
	Title   string
	Content string
}

func main() {
	cache := lru.New[int, Post](4)
	posts := []Post{
		{Id: 1, Title: "Go Generics", Content: "Go 1.18 ile generics geldi."},
		{Id: 2, Title: "LRU Cache", Content: "Least Recently Used cache algoritması."},
		{Id: 3, Title: "Event-Driven", Content: "Event-driven mimari ile loose coupling."},
		{Id: 4, Title: "Kubernetes", Content: "K8s ile container orkestrasyonu."},
		{Id: 5, Title: "Microservices", Content: "Monolith'ten microservice'e geçiş."},
		{Id: 6, Title: "DDD", Content: "DDD implementasyon."},
	}

	for _, post := range posts {
		cache.Put(post.Id, post)
	}

	cache.PrintAll()

	cacheFIFO := fifo.New()

	cache.Put(1, "A") // queue: [1]
	cache.Put(2, "B") // queue: [2,1]
	cache.Put(3, "C") // queue: [3,2,1]

	cache.Get(2)      // sıralama değişmez
	cache.Put(4, "D") // 1 silinir, queue: [4,3,2]

	cache.PrintAll()
}
