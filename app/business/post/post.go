package post

import (
	// "errors"
	"fmt"
	"github.com/xjchan/titans/DB/mongoDB"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	database  = mongoDB.GetConfig().Database
	coloction = "post"
)

//Post the struct of post
type Post struct {
	ID   int64  `bson:"ID"`  //unique id of a post
	UID  int64  `bson:"UID"` //the uid of the post owner
	Text string `bson:"Text"`
}

//Reply the Reply of a post
type Reply struct {
	ID     int64
	PostID int64
	FromID int64
	AtID   int64
	Text   string
}

func (p *Post) Get(m bson.M) error {
	session, err := mongoDB.Connnet()
	if err != nil {
		return err
	}

	defer mongoDB.Close(session)

	c := session.DB(database).C(coloction)

	err = c.Find(m).One(p)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p)

	return nil
}

func (p *Post) Add() error {
	session, err := mongoDB.Connnet()
	if err != nil {
		return err
	}
	defer mongoDB.Close(session)

	c := session.DB(database).C(coloction)

	err = c.Insert(p)
	return err
}

func (p *Post) Update(selector bson.M) error {
	session, err := mongoDB.Connnet()
	if err != nil {
		return err
	}
	defer mongoDB.Close(session)

	c := session.DB(database).C(coloction)

	err = c.Update(selector, p)
	return err
}

func (p *Post) Delete(selector bson.M) error {
	session, err := mongoDB.Connnet()
	if err != nil {
		return err
	}
	defer mongoDB.Close(session)

	c := session.DB(database).C(coloction)

	err = c.Remove(selector)

	return err
}
