package main

import "fmt"

func main() {
    //list slice map 均可使用for range
    list := [...]rune{'a', 'b', 'c'}
    //i:v 可以获取序号和值
    for i, v := range list {
        fmt.Println(i, v)
    }
    //可以只获取序号,然后在循环内获取值
    for i := range list {
        fmt.Println(i, list[i])
    }
    //也可以不获取序号和值,仅仅循环len()次
    for range list {
        fmt.Println("*")
    }

    slice := []byte("string")
    for i, v := range slice {
        fmt.Println(i, slice[i], v)
    }

    //map的key:alue是无序的
    map_ := map[string]string{"monday":"1", "sundy":"0"}
    for k, v := range map_ {
        fmt.Println(k, v)
    }

    for k := range map_ {
        fmt.Println(k, map_[k])
    }

    for range map_ {
        fmt.Println("*")
    }

}
