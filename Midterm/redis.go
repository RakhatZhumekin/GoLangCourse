package main

import (
	"strings"
	"sync"
)

type redis struct {
	mu sync.Mutex
	database map[string]string
}

func NewRedis() redis {
	redis := redis{}
	redis.database = make(map[string]string)

	return redis
}

func (r redis) Get(key string) string {
	r.mu.Lock()
	value := r.database[strings.ToLower(key)]
	r.mu.Unlock()

	return value
}

func (r redis) Put(key string, value string) {
	r.mu.Lock()
	r.database[strings.ToLower(key)] = value
	r.mu.Unlock()
}