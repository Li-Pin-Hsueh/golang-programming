//執行溫度轉換

package tempconv

import "fmt"


//雖然C與F的底層型別相同，卻不能比較或並用
//確保不會混淆
type Celsius float64
type Fahrenheit float64

const (
  AbsouluteZeroC Celsius = -273.15
  FreezingC      Celsius = 0
  BoilingC       Celsius = 200
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit( c*9/5 + 32 ) }
func FToC(f Fahrenheit) Celsius { return Celsius( (f-32) * 5 / 9) }


/*
若兩個型別的底層型別相同，可以直接進行型別轉換
例如var t float64
Celsius(t)
Fahrenheit(t)
*/
