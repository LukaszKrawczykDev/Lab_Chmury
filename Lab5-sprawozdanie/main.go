package main

import (
    "fmt"
    "net"
    "net/http"
    "os"
)

var Version = "undefined"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        hostname, _ := os.Hostname()
        ip := getLocalIP()
        
        fmt.Fprintf(w, "<html><body><h1>Informacje o serwerze</h1>")
        fmt.Fprintf(w, "<p><strong>Server IP:</strong> %s</p>", ip)
        fmt.Fprintf(w, "<p><strong>Hostname:</strong> %s</p>", hostname)
        fmt.Fprintf(w, "<p><strong>Wersja Aplikacji:</strong> %s</p>", Version)
        fmt.Fprintf(w, "</body></html>")
    })

    http.ListenAndServe(":8080", nil)
}

func getLocalIP() string {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return "unknown"
    }
    for _, address := range addrs {
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return "unknown"
}