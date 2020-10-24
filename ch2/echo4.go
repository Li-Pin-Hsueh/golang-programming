// 輸出命令列參數

package main

import (
  "flag"
  "fmt"
  "strings"
)

//flag.Bool函式建構新的bool型別flag變數
// n與sep是flag變數的指標，需要間接存取
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "seperator")

func main() {

  flag.Parse()
  fmt.Print(strings.Join(flag.Args(), *sep))
  if !*n {
    fmt.Println()
  }

}
