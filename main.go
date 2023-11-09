package main

import (
    "fmt"
    "os"
    "log"
    "strings"
    "time"
    "github.com/gempir/go-twitch-irc"
)

const (
    todayfilename = "todaydata"
)

func handle_error(err error){
    if err != nil {
        log.Fatal(err)
    }
}

func read_token_from_file(filename string)(string) {
    data, err := os.ReadFile(filename)

    handle_error(err)

    str_data := string(data)
    return str_data 
}

func write_data_to_file(filename string, data string) {
    file, err := os.Create(filename)
    
    handle_error(err)

    file.WriteString(data)
}

func main(){
    access_token := read_token_from_file(".client")
    client := twitch.NewClient("cobbcoding", access_token)

    commands := map[string]string{
        "!help": "List all possible commands",
        "!ping": "PONG",
        "!today": "List the current project for today",
        "!socials": "List all my social media accounts",
        "!yt": "Get URL for my YouTube channel",
        "!surprise": "Run for a fun surpise",
        "!specs": "List PC specs and OS information",
        "!say": "Make the bot say something",
        "!time": "List the current time in the streamers' area",
        "!sub": "to subscribe",
    }

    client.OnPrivateMessage(func(message twitch.PrivateMessage) {
        if(message.Message[0] == '!') {
            msg := message.Message[1:]
            stripped_msg := strings.Fields(msg)
            switch(stripped_msg[0]){
                case "ping":
                    client.Say(message.Channel, "Pong!")
                case "settoday":
                    if(message.User.Name != "cobbcoding") {
                        client.Say(message.Channel, "You're not allowed to do that")
                    } else {
                        client.Say(message.Channel, "Set the topic of today's stream successfully")
                        write_data_to_file(todayfilename, strings.Join(stripped_msg[1:], " "))
                    }
                case "today":
                    client.Say(message.Channel, read_token_from_file(todayfilename))
                case "socials":
                    client.Say(message.Channel, "X (Twitter): https://x.com/cobbcoding, YouTube: https://youtube.com/@cobbcoding, GitHub: https://github.com/CobbCoding1")
                case "yt":
                    client.Say(message.Channel, "https://youtube.com/@cobbcoding")
                case "specs":
                    client.Say(message.Channel, "specs: i5 2400, 12GB RAM, Radeon R7 260X. Arch Linux BTW")
                case "surprise":
                    client.Say(message.Channel, "Enjoy your surpise: https://www.youtube.com/watch?v=xvFZjo5PgG0")
                case "say":
                    client.Say(message.Channel, strings.Join(stripped_msg[1:], " "))
                case "time":
                    client.Say(message.Channel, time.Now().Format(time.RFC850))
                case "sub":
                    client.Say(message.Channel, "https://www.twitch.tv/subs/cobbcoding")
                case "help":
                    output := ""
                    for key, data := range commands {
                        output += key + ": "
                        output += data + ",     "
                    }
                    client.Say(message.Channel, output)
                default:
                    client.Say(message.Channel, "unknown command: !" + msg)
            }
        } 
    })

    client.OnUserJoinMessage(func(message twitch.UserJoinMessage) {
        client.Say(message.Channel, fmt.Sprintf("%s: has joined the chat", message.User))
    })
    client.Join("cobbcoding")

    err := client.Connect()

    handle_error(err)
}

