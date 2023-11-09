package main

import (
    "fmt"
    "os"
    "log"
    "github.com/gempir/go-twitch-irc"
)

func read_token_from_file(filename string)(string) {
    data, err := os.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }
    str_data := string(data)
    return str_data 
 }

func main(){
    access_token := read_token_from_file(".client")
    client := twitch.NewClient("cobbcoding", access_token)

    client.OnPrivateMessage(func(message twitch.PrivateMessage) {
        fmt.Println(message.Message)
        if(message.Message == "!ping"){
            client.Say(message.Channel, "Pong!\n")
        }
    })

    client.OnUserJoinMessage(func(message twitch.UserJoinMessage) {
        client.Say(message.Channel, fmt.Sprintf("Welcome: %s", message.User))
    })
    client.Join("cobbcoding")

    err := client.Connect()
    if err != nil {
        panic(err)
    }
}

