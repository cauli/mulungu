package storage

import (
	"fmt"
	"time"
)

func Save(resourceType string, id string, value interface{}) {
	resourceKey := generateKey(resourceType, id)
	GetGlobalCache().Mcache.Set(resourceKey, value, time.Hour)
}

func GetById(resourceType string, id string) (exists bool, value interface{}) {
	resourceKey := generateKey(resourceType, id)
	item := GetGlobalCache().Mcache.Get(resourceKey)

	if item == nil {
		return false, nil
	}

	return true, (*item).Value()
}

func Delete(resourceType string, id string) bool {
	resourceKey := generateKey(resourceType, id)
	deleted := GetGlobalCache().Mcache.Delete(resourceKey)

	return deleted
}

func generateKey(resourceType string, id string) string {
	return fmt.Sprintf("%s-%s", resourceType, id)
}
