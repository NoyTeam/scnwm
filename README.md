# Simplified Chinese Negative Word Marker

> 完善中，歡迎提交負面關鍵字

簡體中文負面網絡用語標記，支援 GOCC 簡繁轉換和諧音字

## Install

```
go get github.com/NoyTeam/scnwm@v1.0.2
```

## Use

```go
package main

import (
    "fmt"

    "github.com/NoyTeam/scnwm"
)

func main() {
    num, err := scnwm.Test("差不多该写好了吧😅")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(num) // 13，推薦標記大於 7 的句子
}
```