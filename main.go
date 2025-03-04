package main

import (
	"fmt"
	"givemesomething/mcok/cmd/mcok"
	"givemesomething/mcok/config"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}

	mcok.MockCommand.Execute()
}
