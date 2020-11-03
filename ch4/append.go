//實作slice的append如何運作

package main

func appendInt(x []int, y int,) []int {
  var z []int
  zlen := len(x)+1
  if zlen <= cap(x) {
    //有空間供成長 擴張slice
    z = x[:zlen]
  } else {
    // 空間不夠 分配新陣列
    //倍增 平攤線性複雜度
    zcap := zlen
    if zcap < 2*len(x) {
      zcap = 2*len(x)
    }
    z = make([]int, zlen, zcap)
    copy(z, x)
  }
  z[len(x)] = y //新增y
  return z
}
