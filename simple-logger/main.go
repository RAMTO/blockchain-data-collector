package simplelogger

import "fmt"

func Log(label string, value interface{}) {
	fmt.Println("")
	fmt.Println("==== Logger ====")
	fmt.Println(label+"-> ", value)
	fmt.Println("================")
	fmt.Println("")
}

func LogValue(label string, value interface{}) {
	fmt.Println("")
	fmt.Println("==== Logger ====")
	fmt.Println(label + ":")
	fmt.Printf("%+v\n", value)
	fmt.Println("================")
	fmt.Println("")
}
