package main

import "fmt"

//参考https://github.com/polaris1119/The-Golang-Standard-Library-by-Example/blob/master/chapter01/01.3.md

/*
占位符

普通占位符

占位符                       说明                      举例                                      输出
%v      相应值的默认格式。                             Printf("%v", site)，Printf("%+v", site)    {studygolang}，{Name:studygolang}
        在打印结构体时，“加号”标记（%+v）会添加字段名
%#v     相应值的Go语法表示                          Printf("#v", site)                      main.Website{Name:"studygolang"}
%T      相应值的类型的Go语法表示                     Printf("%T", site)                      main.Website
%%      字面上的百分号，并非值的占位符                   Printf("%%")                            %

布尔占位符

占位符                       说明                      举例                                      输出
%t      单词 true 或 false。                            Printf("%t", true)                      true
整数占位符

占位符                       说明                      举例                                  输出
%b      二进制表示                                 Printf("%b", 5)                     101
%c      相应Unicode码点所表示的字符                   Printf("%c", 0x4E2D)                中
%d      十进制表示                                 Printf("%d", 0x12)                  18
%o      八进制表示                                 Printf("%d", 10)                    12
%q      单引号围绕的字符字面值，由Go语法安全地转义      Printf("%q", 0x4E2D)                '中'
%x      十六进制表示，字母形式为小写 a-f              Printf("%x", 13)                    d
%X      十六进制表示，字母形式为大写 A-F              Printf("%x", 13)                    D
%U      Unicode格式：U+1234，等同于 "U+%04X"         Printf("%U", 0x4E2D)                U+4E2D

浮点数和复数的组成部分（实部和虚部）

占位符                       说明                                              举例                                  输出
%b      无小数部分的，指数为二的幂的科学计数法，与 strconv.FormatFloat
        的 'b' 转换格式一致。例如 -123456p-78
%e      科学计数法，例如 -1234.456e+78                                  Printf("%e", 10.2)                          1.020000e+01
%E      科学计数法，例如 -1234.456E+78                                  Printf("%e", 10.2)                          1.020000E+01
%f      有小数点而无指数，例如 123.456                                   Printf("%f", 10.2)                          10.200000
%g      根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出             Printf("%g", 10.20)                         10.2
%G      根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出             Printf("%G", 10.20+2i)                      (10.2+2i)

字符串与字节切片

占位符                       说明                                              举例                                  输出
%s      输出字符串表示（string类型或[]byte)                          Printf("%s", []byte("Go语言学习园地"))        Go语言学习园地
%q      双引号围绕的字符串，由Go语法安全地转义                            Printf("%q", "Go语言学习园地")                "Go语言学习园地"
%x      十六进制，小写字母，每字节两个字符                             Printf("%x", "golang")                      676f6c616e67
%X      十六进制，大写字母，每字节两个字符                             Printf("%X", "golang")                      676F6C616E67

指针

占位符                       说明                                              举例                                  输出
%p      十六进制表示，前缀 0x                                          Printf("%p", &site)                         0x4f57f0
这里没有 'u' 标记。若整数为无符号类型，他们就会被打印成无符号的。类似地，这里也不需要指定操作数的大小（int8，int64）。

宽度与精度的控制格式以Unicode码点为单位。（这点与C的 printf 不同，它以字节数为单位）二者或其中之一均可用字符 '*' 表示，此时它们的值会从下一个操作数中获取，该操作数的类型必须为 int。

对数值而言，宽度为该数值占用区域的最小宽度；精度为小数点之后的位数。 但对于 %g/%G 而言，精度为所有数字的总数。例如，对于123.45，格式 %6.2f 会打印123.45，而 %.4g 会打印123.5。%e 和 %f 的默认精度为6；但对于 %g 而言，它的默认精度为确定该值所必须的最小位数。

对大多数的值而言，宽度为输出的最小字符数，如果必要的话会为已格式化的形式填充空格。对字符串而言，精度为输出的最大字符数，如果必要的话会直接截断。

其它标记

占位符                       说明                                              举例                                  输出
+       总打印数值的正负号；对于%q（%+q）保证只输出ASCII编码的字符。           Printf("%+q", "中文")                 "\u4e2d\u6587"
-       在右侧而非左侧填充空格（左对齐该区域）
#       备用格式：为八进制添加前导 0（%#o），为十六进制添加前导 0x（%#x）或 Printf("%#U", '中')                        U+4E2D '中'
        0X（%#X），为 %p（%#p）去掉前导 0x；如果可能的话，%q（%#q）会打印原始
        （即反引号围绕的）字符串；如果是可打印字符，%U（%#U）会写出该字符的
        Unicode 编码形式（如字符 x 会被打印成 U+0078 'x'）。
' '     （空格）为数值中省略的正负号留出空白（% d）；
        以十六进制（% x, % X）打印字符串或切片时，在字节之间用空格隔开
0       填充前导的0而非空格；对于数字，这会将填充移到正负号之后

*/
type Website struct {
	Name string
}

func main() {
	P := fmt.Printf
	// 普通占位符
	var site = Website{Name: "studygolang"}
	P("%v\n", site)  //相应值的默认格式
	P("%+v\n", site) //打印结构体时,显示字段名
	P("%#v\n", site) //相应值的go语法表示
	P("%T\n", site)  // 相应值的类型的go语法表示
	P("%%\n")        // 字面的百分号
	// 布尔占位符
	P("%t\n", true) //布尔值true 或 false
	//整数占位符
	P("%b\n", 5)      //二进制表示
	P("%c\n", 0x4E2D) // 相应Unicode码的字符
	P("%d\n", 0x12)   //将数值转换为十进制
	P("%o\n", 9)      //将数值转换为八进制
	P("%q\n", 0x4E2D) //单引号围绕的字符字面值，由Go语法安全地转义
	P("%x\n", 33)     //转换为16进制, 字母小写
	P("%X\n", 33)     //转换为16进制, 字母大写
	P("%U\n", 0x4E2D) //Unicode格式:U+4E2D
	//浮点数及复数
	P("%e\n", 10.2)     //科学计数法,字母小写(精度默认6,末尾0填充)
	P("%E\n", 10.2)     //科学计数法,字母大写(精度默认6,末尾0填充)
	P("%f\n", 10.2)     //带小数点的浮点数,无指数(精度默认6,末尾0填充)
	P("%g\n", 10.20+2i) //根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出
	P("%G\n", 10.20+2i) //根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出
	//字符串与字节切片
	P("%s\n", []byte("中文good"))                                                   //输出字符串表示（支持string类型或[]byte)
	P("%q\n", []byte{0xe4, 0xb8, 0xad, 0xe6, 0x96, 0x87, 0x67, 0x6f, 0x6f, 0x64}) //双引号围绕的字符串，由Go语法安全地转义
	P("%x\n", []byte("中文good"))                                                   //十六进制，小写字母，每字节两个字符
	P("%X\n", []byte("中文good"))                                                   //十六进制，大写字母，每字节两个字符
	//指针
	P("%p\n", &site) // 指针       十六进制表示,前缀0x

}
