package main

import "fmt"

func Template(x, y, z string) string{
    return fmt.Sprintf("%s時の%sは%s", x, y, z)
}

func main(){
    fmt.Println(Template("12", "気温", "24"))
}
