package main

import (
    "fmt"
    "github.com/shezadkhan137/go-qrcode/qrcode"
)

func main() {
    results, err := qrcode.GetDataFromPNG("/Go/src/github.com/GoGrid/IMG_3068.png")
    if err != nil {
        panic(err)
    }

    for _, result := range results{
        fmt.Println(result)
        //fmt.Printf("Symbol Type: %s, Data %s", result.SymbolType, result.Data )
    }

}
