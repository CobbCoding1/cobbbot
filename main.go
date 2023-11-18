package main

import (
    "fmt"
    "github.com/CobbCoding1/goirc"
    "os"
    "log"
    "strings"
    "time"
)

const (
    domain = "irc.twitch.tv"
    port = "6667"
    server = "cobbcoding"
    name = "cobbbot"
    todayfilename = "/.config/cobbbot/todaydata"
    discord = "https://discord.gg/3SkpwrRxpA"
)

func handle_error(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func read_string_from_file(filename string)(string) {
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


func main() {
    home_dir, home_err := os.UserHomeDir()
    handle_error(home_err)

    commands := map[string]string{
        "help": "List all possible commands",
        "ping": "PONG",
        "today": "List the current project for today",
        "socials": "List all my social media accounts",
        "yt": "Get URL for my YouTube channel",
        "surprise": "Run for a fun surpise",
        "specs": "List PC specs and OS information",
        "say": "Make the bot say something",
        "time": "List the current time in the streamers' area",
        "sub": "to subscribe",
        "69": "nice",
        "nice": "69",
        "discord": "discord",
    }

    var client goirc.Client
    token := read_string_from_file(".client")
    client.Init(domain, port, token, "", name) 
    client.Connect()
    defer client.Disconnect()

    client.Join(server)
    for {
        msg := client.GetData()
        if msg.IsPrivateMessage() {
            if(msg.Message[0] == '!') {
                command := msg.Message[1:]
                stripped_command := strings.Fields(command)
                fmt.Println(stripped_command)
                switch(stripped_command[0]){
                case "ping":
                    client.Say("Pong!")
                case "settoday":
                    if(msg.Username != "cobbcoding") {
                        client.Say("You're not allowed to do that")
                    } else {
                        client.Say("Set the topic of today's stream successfully")
                        write_data_to_file(home_dir + todayfilename, strings.Join(stripped_command[1:], " "))
                    }
                case "shoutout":
                    if(msg.Username != "cobbcoding") {
                        client.Say("You're not allowed to do that")
                    } else {
                        streamer := strings.Join(stripped_command[1:], " ")
                        client.Say(fmt.Sprintf("Go follow %s on Twitch! twitch.tv/%s", streamer, streamer))
                    }
                case "today":
                    client.Say(read_string_from_file(home_dir + todayfilename))
                case "socials":
                    client.Say("X (Twitter): https://x.com/cobbcoding, YouTube: https://youtube.com/@cobbcoding, GitHub: https://github.com/CobbCoding1")
                case "yt":
                    client.Say("https://youtube.com/@cobbcoding")
                case "specs":
                    client.Say("specs: i5 2400, 12GB RAM, Radeon R7 260X. Arch Linux BTW")
                case "surprise":
                    client.Say("Enjoy your surpise: https://www.youtube.com/watch?v=xvFZjo5PgG0")
                case "say":
                    client.Say(strings.Join(stripped_command[1:], " "))
                case "time":
                    client.Say(time.Now().Format(time.RFC850))
                case "sub":
                    client.Say("https://www.twitch.tv/subs/cobbcoding")
                case "69":
                    client.Say("nice")
                case "nice":
                    client.Say("69")
                case "discord":
                    client.Say("The Discord is: " + discord)
                case "help":
                    output := ""
                    for key, data := range commands {
                        output += key + ": "
                        output += data + " --- "
                    }
                    client.Say(output)
                default:
                    client.Say("unknown command: !" + command)
                }
            }
        }

        client.HandlePong()
    }
}
