package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func Init(dbPath string) (err error) {
	db, err = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}

	return db.Update(func(tx *bolt.Tx) (err error) {
		_, err = tx.CreateBucketIfNotExists(taskBucket)
		return
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func AllTasks() ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func DeleteTask(k int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(k))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(s []byte) int {
	return int(binary.BigEndian.Uint64(s))
}
