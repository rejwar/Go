// File Name: ChannelVariables.go

package main

func Main() {
    Messages := make(chan string)
    go func() { Messages <- "Hello!" }()
    println(<-Messages) // Output: Hello!
}
