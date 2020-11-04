// 以struct實現樹狀結構

package treesort

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
  appendValues(values[:0], root)
}

//依序插入元素t
//回傳結果slice
func appendValues(values []int, t *tree) []int {
  if t != nil {
    values = appendValues(values, t.left)
    values = append(values, t.value)
    values = appendValues(values, t.right)
  }
  return values
}

func add(t *tree, value int) *tree {
  if t == nil {
    // 等同回傳&tree{value: value}
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
