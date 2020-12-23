# goutils
collect golang utils


# useage
```golang
package main

import (
  "github.com/iepngs/goutils/timetool"
  "log"
  "time"
)

func main(){
  log.Println(timetool.DateFormart(time.Now(), "YYYY-MM-DD HH:mm:ss"))
}
```
