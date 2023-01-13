package main

import (
	_ "embed"
	"log"

	bob "github.com/wailsapp/wails/exp/pkg/application"
	"github.com/wailsapp/wails/exp/pkg/options"
)

type GreetService struct {
	SomeVariable int
	lowerCase    string
}

func (*GreetService) Greet(name string) string {
	return "Hello " + name
}

func main() {
	app := bob.New(options.Application{
		Bind: []interface{}{
			&GreetService{},
		},
	})

	app.NewWebviewWindow()

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}

}