package routers

import (
	"fmt"
	"github.com/go-macaron/oauth2"
	// "github.com/go-macaron/session"
	"github.com/xjchan/titans/app/handlers"
	// goauth2 "golang.org/x/oauth2"
	"gopkg.in/macaron.v1"
)

func init() {
	fmt.Println("routes")
}

func IinitRouters(m *macaron.Macaron) error {
	m.Get("/", func() string {
		return "Hello World！"
	})

	m.Get("/post", handlers.GetPost)

	m.Get("/add_post", handlers.AddPost)

	m.Get("/test", func(tokens oauth2.Tokens) string {
		return fmt.Sprintln(tokens.Expired())
	})

	return nil
}
