package app

import (
	"fmt"
	"github.com/xjchan/titans/app/handlers"
	"gopkg.in/macaron.v1"
	"os"
	// "path/filepath"
)

// Init 初始化服务器
func Init() {
	fmt.Println("dir:" + Dir())

	m := macaron.Classic()

	m.Get("/", func() string {
		return "Hello World！"
	})

	m.Get("/post", handlers.GetPost)
	m.Run()
}

//Dir to get app dir
func Dir() string {
	dir, _ := os.Getwd()
	return dir
}
