/*
练习 2.2： 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的话则
是从标准输入读取参数，然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对应英尺和
米，重量单位可以对应磅和公斤等。
 */
package main

import (
	"ch2/exercise/lenconv"
	"fmt"
)

func main() {
	fmt.Println(lenconv.MtoDM(lenconv.OneM))
	fmt.Println(lenconv.DMtoCM(lenconv.MtoDM(lenconv.TenM)))
}

