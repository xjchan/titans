package app

import (
	"github.com/go-macaron/oauth2"
	"github.com/go-macaron/session"
	"github.com/xjchan/titans/app/routers"
	goauth2 "golang.org/x/oauth2"
	"gopkg.in/macaron.v1"
	"os"
)

// Init 初始化服务器
func Init() {
	// fmt.Println("dir:" + Dir())

	m := macaron.Classic()
	m.Use(session.Sessioner())

	m.Use(oauth2.Google(
		&goauth2.Config{
			ClientID:     "client_id",
			ClientSecret: "client_secret",
			Scopes:       []string{"https://www.googleapis.com/auth/drive"},
			RedirectURL:  "redirect_url",
		},
	))

	m.Use(func() string {
		return "test_midware"
	})

	// m.MapTo("inject_test", 1)
	routers.IinitRouters(m)
	m.Run()
}

//Dir to get app dir
func Dir() string {
	dir, _ := os.Getwd()
	return dir
}
