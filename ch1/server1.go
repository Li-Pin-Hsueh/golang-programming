// 迷你"echo"伺服器

package main

import (
  "fmt"
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/", handler)  //每個請求呼叫處理程序 (HandleFunc register the handler function for the given  pattern)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//處理程序回應所請求url的path元件
func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w,"URL.Path = %q\n", r.URL.Path)
}
