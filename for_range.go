package main

import "fmt"

func main() {
    //
    list := [...]rune{'a', 'b', 'c'}
    for i, v := range list {
        fmt.Println(i, v)
    }

    for i := range list {
        fmt.Println(i, list[i])
    }

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
