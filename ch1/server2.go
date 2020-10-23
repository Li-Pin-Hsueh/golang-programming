// 迷你"echo"與計數伺服器

package main

import (
  "fmt"
  "log"
  "net/http"
  "sync"
)

var mu sync.Mutex
var count int

func main() {
  //此伺服器有兩個處理程序
  http.HandleFunc("/", handler)
  http.HandleFunc("/count", counter)

  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//處理程序回應所請求url的pathe元件
func handler(w http.ResponseWriter, r *http.Request) {
  //可能出現race condition (同時想遞增count)
  //必須確保每次最多一個goroutine存取該變數 : 以Lock與Unlock包圍count的存取
  mu.Lock()
  count++
  mu.Unlock()
  fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

//計算回應數量
func counter(w http.ResponseWriter, r *http.Request) {
  mu.Lock()
  fmt.Fprintf(w, "Count %d\n", count)
  mu.Unlock()
}
