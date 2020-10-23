// 平行抓取多個url並回報時間與大小

package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  "time"
)

func main() {
  start := time.Now()
  ch := make(chan string)
  for _, url := range os.Args[1:] {
    go fetch(url, ch) // 啟動goroutine
  }
  for range os.Args[1:] {
    fmt.Println(<-ch) //從channel ch接收
  }
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch( url string, ch chan<- string) {
  start := time.Now()
  resp, err := http.Get(url)
  if err != nil {
    ch <- fmt.Sprint(err) //發送到channel ch
    return
  }

  nbytes, err := io.Copy(ioutil.Discard, resp.Body) // Copy回應位元組與錯誤
  resp.Body.Close() //預防洩漏資源
  if err != nil {
    ch <- fmt.Sprintf("while reading %s: %v", url, err)
    return
  }
  secs := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
