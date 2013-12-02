/*
github.com/twitter/snowflake in golang
*/

package main

import (
	"github.com/simonz05/snowflake"
	"fmt"
)

func main() {
	v, err := snowflake.Default()
	if err != nil {
		panic(err)
	}
	//v := snowflake.NewSnowflake(100)
	for i := 0; i < 10; i++ {
		id, err := v.Next()
		if err != nil {
			fmt.Printf("err: %v", err)
		} else {
			fmt.Println(id)
		}
	}
}
