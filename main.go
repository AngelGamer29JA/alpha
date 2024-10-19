package main

import (
	"fmt"

	"github.com/angelgamer29ja/alpha/alpha"
	"github.com/angelgamer29ja/alpha/alpha/color"
)

func main() {
	fmt.Printf("%sHello %sf\n", &color.Blue, &color.Reset)
	fmt.Println(color.Formatter("&ehello"))

	context := alpha.NewContext(&alpha.Script{Name: "hello"})
	context.RegisterVariable("TEST", "Hello", false)
	fmt.Println(context.GetVariable("TEST").Value)

	context.GetVariable("TEST").SetValue("a")
	fmt.Println(context.GetVariable("TEST").Value)
}
