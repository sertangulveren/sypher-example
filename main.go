package main

import (
	"fmt"
	"github.com/sertangulveren/sypher"
	_ "github.com/sertangulveren/sypher-example/sypher"
)

func main() {
	sypher.Load(sypher.Config{
		Name: "myconfig",
	})

	awsKey := sypher.Get("AWS_ACCESS_KEY")
	awsSecret := sypher.Get("AWS_ACCESS_SECRET_KEY")
	holyKey := sypher.Get("HOLY_API_KEY")

	// print my environment variables
	fmt.Println(awsKey)
	fmt.Println(awsSecret)
	fmt.Println(holyKey)
}
