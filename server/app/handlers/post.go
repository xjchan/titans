package handlers

import (
	"fmt"
	"github.com/xjchan/titans/server/app/business/post"
	"gopkg.in/macaron.v1"
	// "net/http"
)

//GetPost get post
func GetPost(ctx *macaron.Context) string {
	post, _ := post.GetPost(1)
	return fmt.Sprintln(post) + ctx.Req.RequestURI
}

func Test(p macaron.ReturnHandler) macaron.ReturnHandler {
	return p
}
