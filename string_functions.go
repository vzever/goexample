package main

import s "strings"
import "fmt"

//We alias fmt.Println to a shorter name as we’ll use it a lot below.
var p = fmt.Println
//string中是以byte计算的,如果涉及到unicode的计算,需要转换为rune数组
func main() {
    //Here’s a sample of the functions available in strings. Note that these are all functions from package, not methods on the string object itself. This means that we need pass the string in question as the first argument to the function.
    p("Contains:  ", s.Contains("test", "es"))
    p("ContainsAny:  ", s.ContainsAny("test", "abce"))
    p("ContainsRune:  ", s.ContainsRune("test", 'e'))
    p("ContainsRune:  ", s.ContainsRune("test中文", '中'))
    p("Equalold:  ", s.EqualFold("ABC中文", "abC中文"))
    //以空格split字符串
    p("Fields:    ", s.Fields("  foo bar  baz   "))
    //以函数返回true的rune作为分隔符 分割字符串
    p("FieldsFunc:", s.FieldsFunc("abcdefg", func(r rune) bool {
        if r=='d' {
            return true
        }else {
            return false
        }}))
    p("Count:     ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("中文Index: ", s.Index("中文test", "文"))
    p("Join:      ", s.Join([]string{"ab", "cd"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace:   ", s.Replace("foooo", "o", "0", -1))
    p("Replace:   ", s.Replace("foooo", "o", "0", 2))
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
    p()
    //You can find more functions in the strings package docs.
    //Not part of strings but worth mentioning here are the mechanisms for getting the length of a string and getting a character by index.
    p("Len: ", len("hello"))
    p("Char:", "hello"[1])
}
/*
Contains:   true
Count:      2
HasPrefix:  true
HasSuffix:  true
Index:      1
Join:       ab-cd
Repeat:     aaaaa
Replace:    f00
Replace:    f0o
Split:      [a b c d e]
ToLower:    test
ToUpper:    TEST

Len:  5
Char: 101
*/