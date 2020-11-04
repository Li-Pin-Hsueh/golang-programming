// 以struct實現樹狀結構

package main

import(
  "fmt"
)

type tree struct {
  value       int
  left, right *tree   //必須使用指標型別 (C base思維)
}

//原址排序
func Sort(values []int) {
  var root *tree
  for _, v := range values{
    root = add(root, v)
  }
  appendValues(values[:0], root)    // value[:0] 實際上就是整個slice
}

//依序插入元素t
//回傳結果slice
func appendValues(values []int, t *tree) []int {
  if t != nil {
    values = appendValues(values, t.left)   //先取最左(最小)
    values = append(values, t.value)        //append到slice
    values = appendValues(values, t.right)
  }
  return values
}

func add(t *tree, value int) *tree {
  if t == nil {
    // 等同回傳&tree{value: value}
    // 遞迴會一直return回tree的根結點t
    t = new(tree)
    t.value = value
    return t
  }

  if value < t.value {
    t.left = add(t.left, value)
  } else {
    t.right = add(t.right, value)
  }
  return t
}

func main() {
  arr := [5]int{2,1,5,4,3}
  slice := arr[:]
  fmt.Println("Before sorted: ", slice)
  Sort(slice)
  fmt.Println("After sorted: ", slice)

}
