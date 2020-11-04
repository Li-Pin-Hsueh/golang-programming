// 透過encoding/json套件將資料結構轉換為JSON格式
package main

import(
  "fmt"
  "encoding/json"
  "log"
)

type Movie struct {
  Title     string
  Year      int   `json:"released"`           //宣告後面的字串實字是欄位標籤
  Color     bool  `json:"color,omitempty"`
  Actors    []string
}

var movies = []Movie{
  {Title: "Casablanca", Year: 1942, Color: false,
    Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
  {Title: "Cool Hand Luke", Year: 1967, Color: true,
    Actors: []string{"Paul Newman"}},
  // ...
}

func main() {

  data, err := json.Marshal(movies)
  // 另一種可以轉換為可閱讀JSON的function
  // 第二和第三個參數是前綴與縮排字串
  // data, err := json.MarshalIndent(movies, "", "  ")
  
  if err != nil {
    log.Fatalf("JSON marshaling failed: %s", err)
  }
  fmt.Printf("%s\n", data)

}
