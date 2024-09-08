package main

import "sync"

type Cache struct {
	isCourseExists       sync.Map // course_id -> bool
	isRegistrationExists sync.Map // cource_id + user_id -> bool
}

var cache *Cache

func newCache() *Cache {
	return &Cache{}
}
