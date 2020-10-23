// 能夠處理接受到之標頭與表單資料的迷你"echo"伺服器

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

  //將回應串流交給代表瀏覽器的 w(http.ResponseWriter)
  fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
  for k, v := range r.Header {
    fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
  }
  fmt.Fprintf(w, "Host = %q\n", r.Host)
  fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

  if err := r.ParseForm(); err != nil {
    log.Print(err)
  }
  for k, v := range r.Form {
    fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
  }
}

//計算回應數量
func counter(w http.ResponseWriter, r *http.Request) {
  mu.Lock()
  fmt.Fprintf(w, "Count %d\n", count)
  mu.Unlock()
}
