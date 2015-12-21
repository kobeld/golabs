package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/copier"

	"github.com/ulule/deepcopier"
	"gopkg.in/mgo.v2/bson"
)

type base struct {
	Id bson.ObjectId
}

func (this *base) IdHex() string {
	return this.Id.Hex()
}

type Club struct {
	base      `,inline`
	Name      string
	CreatedAt time.Time
}

func (this *Club) CreatedAtStr() string {
	return this.CreatedAt.Format("2006-01-02 15:04:02")
}

type ClubResource struct {
	IdHex        string
	Name         string
	CreatedAtStr string
}

func main() {
	club := &Club{
		Name:      "Fun",
		CreatedAt: time.Now(),
	}

	club.Id = bson.NewObjectId()

	resource := &ClubResource{}

	copier.Copy(resource, club)

	deepcopier.Copy(club).To(resource)
	fmt.Printf("%+v\n", club)
	fmt.Printf("%+v\n", resource)

	clubs := []*Club{club}
	resources := []*ClubResource{}
	copier.Copy(&resources, &clubs)
	fmt.Printf("%+v\n", clubs)
	fmt.Printf("%+v\n", resources)

}
