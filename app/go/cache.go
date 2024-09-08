package main

import "sync"

type Cache struct {
	eCourse       sync.Map // course_id -> bool
	eRegistration sync.Map // cource_id + user_id -> bool
}

var cache *Cache

func newCache() *Cache {
	return &Cache{}
}

func (c *Cache) isRegistrationExists(courseID string, userID string) bool {
	e, ok := c.eRegistration.Load(courseID + userID)
	return ok && e.(bool)
}
