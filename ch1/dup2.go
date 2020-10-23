// 輸出檔案中出現超過一次的行
// 讀取stdin或一系列檔案

package main

import(
  "bufio"
  "fmt"
  "os"
)

func main(){
  counts := make( map[string]int )
  files := os.Args[1:]    // 由執行go run時代入參數(檔案名稱)
  if len(files) == 0 {
    countLines(os.Stdin, counts)
  } else {
      for _, arg := range files {
        f, err := os.Open(arg)
        if err != nil {
            fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
            continue
        }
        countLines(f, counts)
        f.Close()
      }
    }

  for line, n := range counts{
    if n > 1 {
      fmt.Printf("%d\t%s\n", n, line)
    }
  }

}

func countLines( f *os.File, counts map[string]int) {
  input := bufio.NewScanner(f)
  for input.Scan() {
    counts[input.Text()]++
  }
}
