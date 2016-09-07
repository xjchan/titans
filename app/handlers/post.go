package handlers

import (
	"fmt"
	"github.com/xjchan/titans/app/business/post"
	"github.com/xjchan/titans/tools"
	"gopkg.in/macaron.v1"
	// "net/http"
)

//GetPost get post
func GetPost(ctx *macaron.Context) string {
	post, err := post.GetPost(1)
	if err == nil {
		tools.CheckError(err)
	}
	return fmt.Sprintln(post) + ctx.Req.RequestURI
}

func Test(p macaron.ReturnHandler) macaron.ReturnHandler {
	return p
}
