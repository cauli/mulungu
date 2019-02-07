package storage

import (
	"fmt"
	"time"

	postgres "github.com/cauli/mulungu/storage/driver"
)

// Save will orchesrtate saving from the persistency,
// inserting to the database first, then creating a cache key
func Save(resource, id string, value interface{}) error {
	err := postgres.DB.Save(resource, id, value.(string))
	if err != nil {
		return err
	}

	key := generateKey(resource, id)
	GetGlobalCache().Mcache.Set(key, value, time.Hour)

	return nil
}

// Load will orchestrate fetching from the persistency
// trying to load from cache first, then from database
func Load(resource string, id string) (exists bool, value interface{}) {
	key := generateKey(resource, id)
	item := GetGlobalCache().Mcache.Get(key)

	if item == nil {
		data, notFound, err := postgres.DB.Load(resource, id)
		if notFound {
			return false, nil
		}

		if err != nil {
			fmt.Println(err)
			return false, err
		}

		return true, data
	}

	return true, (*item).Value()
}

// Delete will orchestrate deletion in the persistency
// going to the database and than invalidating cache
func Delete(resource string, id string) (bool, error) {
	key := generateKey(resource, id)

	err := postgres.DB.Delete(resource, id)
	if err != nil {
		return false, err
	}

	deleted := GetGlobalCache().Mcache.Delete(key)
	return deleted, nil
}

func generateKey(resource string, id string) string {
	return fmt.Sprintf("%s-%s", resource, id)
}
