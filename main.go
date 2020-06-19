package main

import (
	"github.com/dmskdlghs213/go_authAPI/app/presentation"
)

func main() {

	e := presentation.Router()
	e.Logger.Fatal(e.Start(":8081")) // listen and serve on :8080

}
