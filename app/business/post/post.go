package post

import (
	// "errors"
	"fmt"
	"github.com/xjchan/titans/app/DB/mongoDB"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

//GetPost get Post data
func GetPost(ID int64) (Post, error) {
	// session, _ := mgo.Dial("127.0.0.1")
	session := mongoDB.Connnet()
	defer session.Close()

	c := session.DB("titans").C("post")

	post := Post{}

	err := c.Find(bson.D{{"ID", ID}}).One(&post)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(post)

	return post, nil
}

//GetPieces get pieces of Reply
func (Reply) GetPieces(ParentID int64, form int, rows int) ([]Reply, error) {
	replies := make([]Reply, 1)
	return replies, nil
}
