# goutils
collect golang utils


# useage example
```golang
package main

import (
  "github.com/iepngs/goutils/timetool"
  "log"
  "time"
)

func main(){
  log.Println(timetool.DateFormat(time.Now(), "YYYY-MM-DD HH:mm:ss"))
}
```
