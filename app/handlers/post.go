package handlers

import (
	"fmt"
	"github.com/xjchan/titans/app/business/post"
	"github.com/xjchan/titans/tools"
	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	// "net/http"
)

//GetPost get post
func GetPost(ctx *macaron.Context) string {
	p := post.Post{}
	err := p.Get(bson.M{"ID": 1})
	tools.CheckError(err)
	return fmt.Sprintln(p) + ctx.Req.RequestURI

}

func AddPost(ctx *macaron.Context) string {
	data, _ := strconv.ParseInt(ctx.Req.URL.Query()["ID"][0], 10, 0)

	p := post.Post{ID: data}
	p.Add()
	return "add"
}

func Test(p macaron.ReturnHandler) macaron.ReturnHandler {
	return p
}
