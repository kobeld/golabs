package main

import (
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo/bson"
	"log"
)

const (
	ENTRIES = "entries"
)

type Record interface {
	GetEType() (etype string)
}

type base struct {
	Id      bson.ObjectId `bson:"_id"`
	Content string
	Title   string
}

func (this *base) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}

func (this *base) SetTitleAndContent(title, content string) {
	this.Title = title
	this.Content = content
}

type Post struct {
	base       `,inline`
	AuthorName string
}

func (this *Post) GetEType() string {
	return "Post"
}

type Task struct {
	base      `,inline`
	OwnerName string
}

func (this *Task) GetEType() string {
	return "Task"
}

type Entry struct {
	base       `,inline`
	AuthorName string
	OwnerName  string
}

func (this *Entry) GetEType() string {
	return "Entry"
}

func main() {

	db := mgodb.NewDatabase("localhost", "golabs")

	// post := &Post{AuthorName: "Aaron Liang"}
	// post.SetTitleAndContent("This is a post", "This is a post!")
	// err := db.Save(ENTRIES, post)
	// panicErr(err)

	// task := &Task{OwnerName: "Azuma Fan"}
	// task.SetTitleAndContent("This is a task", "This is a task!")
	// err = db.Save(ENTRIES, task)
	// panicErr(err)

	entries := []Record{}

	err := db.FindAll(ENTRIES, bson.M{}, &entries)
	panicErr(err)

	for _, entry := range entries {
		log.Printf("%+v \n", entry)
	}

	log.Println("Done")

}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
