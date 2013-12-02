snowflake
======

snowflake is a Twitter Snowflake implementation in golang.

~~~golang
package main

import (
    "github.com/simonz05/snowflake"
    "fmt"
)

func main() {
    v, err := snowflake.Default()
    
    // Alternatively you can set the worker id if you are running multiple snowflakes
    // v, err := snowflake.NewSnowflake(100)
    
    for i := 0; i < 10; i++ {
        id, err := v.Next()
        fmt.Println(id)
    }
}
~~~
