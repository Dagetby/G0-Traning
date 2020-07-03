package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

var store BookmarkStore
var id string

func init() {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"127.0.0.1:27017"},
		Timeout:  60 * time.Second,
		Username: "root",
		Password: "example",
	})
	if err != nil {
		log.Fatalf("[MongoDb Session]: %s\n", err)
	}
	collection := session.DB("testdb").C("testcollection")

	store = BookmarkStore{
		C: collection,
	}
}

func createUpdate() {
	bookmark := Bookmark{
		Name:        "mgo",
		Description: "Go driver for MongoDb",
		Location:    "https://girhub.com/go-mgo/mgo",
		Priority:    2,
		CreatedOn:   time.Now(),
		Tags:        []string{"go", "nosql", "mongodb"},
	}
	if err := store.Create(&bookmark); err != nil {
		log.Fatalf("[Create]: %s\n", err)
	}
	id = bookmark.ID.Hex()
	fmt.Printf("New bookmark has been inserted with ID: %s\n", id)
	bookmark.Priority = 1
	if err := store.Update(bookmark); err != nil {
		log.Fatalf("[Update]: %s\n", err)
	}
	fmt.Println("The value after update:")
	getByID(id)

	bookmark = Bookmark{
		Name:        "gorethink",
		Description: "Go driver for RethikDB",
		Location:    "https://github.com/dancannon/gorethink",
		Priority:    3,
		CreatedOn:   time.Now(),
		Tags:        []string{"go", "nosql", "rethinkdb"},
	}

	if err := store.Create(&bookmark); err != nil {
		log.Fatalf("[Create]: %s\n", err)
	}

	id = bookmark.ID.Hex()
	fmt.Printf("New bookmark  has been inserted with ID: %s\n", id)
}

func getByID(id string) {
	bookmark, err := store.GetById(id)
	if err != nil {
		log.Fatalf("[GetByID]: %s\n", err)
	}
	fmt.Printf("Name:%s, Description:%s, Priority:%d\n", bookmark.Name,
		bookmark.Description, bookmark.Priority)
}

func getAll() {
	layout := "2006-01-02 15:04:05"
	boomarks := store.GetAll()
	fmt.Println("Read all documents")
	for _, v := range boomarks {
		fmt.Printf("Name:%s, Description:%s, Priority:%d, CreatedOn:%s\n", v.Name,
			v.Description, v.Priority, v.CreatedOn.Format(layout))
	}
}

func getByTags() {
	layout := "2006-01-02 15:04:05"
	fmt.Println("Query with Tags - 'go, nosql'")
	bookmarks := store.GetByTag([]string{"go", "nosql"})
	for _, v := range bookmarks {
		fmt.Printf("Name:%s, Description:%s, Priority:%d, CreatedOn:%s\n", v.Name,
			v.Description, v.Priority, v.CreatedOn.Format(layout))
	}
}

func delete() {
	if err := store.Delete(id); err != nil {
		log.Fatalf("[Delete]: %s\n", err)
	}
	bookmarks := store.GetAll()
	fmt.Printf("Number of documentds in the collection afeter delete: %d\n",
		len(bookmarks))
}

func main() {
	createUpdate()
	getAll()
	getByTags()
	delete()
}
