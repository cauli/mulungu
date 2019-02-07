package storage

import (
	"fmt"
	"time"

	postgres "github.com/cauli/mulungu/storage/driver"
)

func Save(resource, id string, value interface{}) error {

	err := postgres.DB.Save(resource, id, value.(string))
	if err != nil {
		return err
	}

	key := generateKey(resource, id)
	GetGlobalCache().Mcache.Set(key, value, time.Hour)

	return nil

}

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
