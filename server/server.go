package server

import (
  "fmt"
  "net"
)

type Cache struct {
  Stores map[string][]map[string]interface{}
}

// Not sure if a gobal cache var is very good but it works for now
var cache Cache = Cache{ Stores: make(map[string][]map[string]interface{}, 0)}

func ServerRun(port string) {
  fmt.Println("** Welcome to the Johnnny Cache Server **")
  
  serverAddr, err := net.ResolveUDPAddr("udp", port)
  HandleError(err, "SERVER START")

  serverConnection, err := net.ListenUDP("udp", serverAddr)
  HandleError(err, "SERVER LISTEN")

  packet := make([]byte, 1024)

  for {
    n, addr, err := serverConnection.ReadFromUDP(packet)
    HandleError(err, "SERVER READ FROM UDP")

    dstore := FromJSON(packet[0:n])

    if dstore.Method == "SET_NAMESPACE" {
      response := make(chan StatusResponse)
      // @TODO: make into a func
      go func() {
        name := dstore.NameSpace
        cache.Stores[name] = make([]map[string]interface{}, 0)
        response <- StatusResponse{ Status: "DB MADE" }
      }()
      res := <-response
      serverConnection.WriteToUDP(ToJSON(res), addr)
    }

    if dstore.Method == "SET" {
      response := make(chan StatusResponse)
      // @TODO: make into a func
      go func() {
        stores := cache.Stores[dstore.NameSpace]
        cache.Stores[dstore.NameSpace] = append(stores, dstore.Data)
        response <- StatusResponse{ Status: "SAVED" }
      }()
      res := <-response
      serverConnection.WriteToUDP(ToJSON(res), addr)
    }

    if dstore.Method == "GET" {
      response := make(chan []map[string]interface{})
      go func() { response <- HandleGet(dstore) }()
      res := <-response
      serverConnection.WriteToUDP(ToJSON(res), addr)
    }
  }
}

