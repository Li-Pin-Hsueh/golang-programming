# 第四章節 組合型別
GO語言中共有四種組合型別:
- Array
- Struct
- Slice
- Map

其中Array與Struct為靜態集合型別，在記憶體中的大小固定。反之Slice與Map為動態資料結構，較具彈性。

## 4.1 Array
Array的初始化：
```
var q [3]int{1, 2, 3}

var p [...]int{1, 2, 3, 4} //陣列長度由初始化數量決定
```

**陣列的大小也是其型別的一部分**，因此 [3]int 與 [4]int是不同的型別，因此以下程式將會出現錯誤：
```
q := [3]int{1, 2, 3}
q = [4]int {1, 2, 3, 4} //編譯錯誤
```

若以陣列作為參數，函式收到的是複製品而不是來源，以這種方式傳遞陣列是非常沒有效率且無法更改來源的，因此在Go語言中，通常會以陣列指標作為參數，但這種行為與**以參考**(call by reference)傳遞陣列的語言有些許不同。

## 4.2 Slice
Slice的型別寫作[]T，其中T為元素型別。Array和Slice緊密相關，slice是輕量化結構，可存取**底層陣列**的部分或全部元素，主要有三個元件：指標、長度、容量。
- 指標：指向slice第一個元素，但不一定是底層陣列的第一個元素。
- 長度：可用len()來取得。
- 容量：可用cap()來取得。
**切割超過cap將造成panic，但切割超過len則會擴張slice**

由於slice帶有陣列元素的指標，傳遞slice給函式可以修改底層陣列元素。但要注意slice不能比較，只能夠自行撰寫。原因之一是slice的元素是間接的，隨著底層陣列被修改，slice的值在不同時間有可能會有差異。

## 4.3 Map
雜湊是最靈活的資料結構，在Go語言中，map是雜湊表的參考，map的型別寫做map[K]V，其中K與V為鍵與值的型別，所有的鍵為相同型別，所有值亦然，**鍵值K型別必須能夠用==來執行比較**。

內建的make函數可用來建構map：
```
ages := make(map[string]int) //字串為鍵值
```
帶有初始值的map：
```
ages := map[string]int {
  "alice"   : 34,
  "charlie" : 31,
}
```
上述等同於：
```
ages := make[map[string]int]
ages["alice"]   = 34
ages["charlie"] = 31
```

**map元素不是變數**，不能取得它的位址：
```
_ = &ages["alice"] //編譯錯誤
```
原因是一個成長中的map可能導致現有元素重新計算雜湊而放在新的位置，導致位址無效。

map型別的零值為nil

## 4.4 Struct
見範例程式 treesort.go

## 4.5 JSON
JSON是JavaScript值，字串、數字、布林、陣列，與物件的**Unicode文字編碼**。

JSON陣列是有序的一系列值，寫成以方括號包圍，逗號分隔的清單；JSON陣列用於Go的array與slice的編碼。

JSON物件是字串與值的對應，寫成一系列的name:value對，由逗號分隔，以括號包圍；JSON物件用於Go的map與struct的編碼。

將Go程式之資料結構轉換為JSON，見範例movie.go
