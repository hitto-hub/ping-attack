package main

import (
    "fmt"
    "os"
    "os/exec"
    "sync"
)

func ping(ip string, wg *sync.WaitGroup, logfile *os.File) {
    defer wg.Done()

    // Pingコマンドを実行
    cmd := exec.Command("ping", ip)

    // 出力をファイルにリダイレクト
    cmd.Stdout = logfile
    cmd.Stderr = logfile

    err := cmd.Run()
    if err == nil {
        fmt.Printf("%s is reachable\n", ip)
    } else {
        fmt.Printf("%s is not reachable\n", ip)
    }
}

func main() {
    var wg sync.WaitGroup

    // ログファイルを作成または開く
    logfile, err := os.OpenFile("ping_logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error opening log file:", err)
        return
    }
    defer logfile.Close()

    ip := "192.168.0.30"
    for i := 0; i < 10000; i++ {
        wg.Add(1)
        go ping(ip, &wg, logfile)
    }
    wg.Wait()
}
