package main

import (
    "fmt"
    "os"
    "strings"

    "github.com/ahmdrz/goinsta/v2"
)

func main() {
    // Read the username and password from environment variables
    username := os.Getenv("INSTAGRAM_USERNAME")
    password := os.Getenv("INSTAGRAM_PASSWORD")

    // Create a new Instagram instance
    insta := goinsta.New(username, password)

    // Log in to Instagram
    if err := insta.Login(); err != nil {
        fmt.Println("Error logging in:", err)
        return
    }

    // Read the post URL from the command line argument
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <post_url>")
        return
    }
    postURL := os.Args[1]

    // Extract the post ID from the URL
    postID := strings.TrimPrefix(postURL, "https://www.instagram.com/p/")
    postID = strings.TrimSuffix(postID, "/")

    // Report the post
    post, err := insta.ReportPost(postID, "spam")
    if err != nil {
        fmt.Println("Error reporting post:", err)
        return
    }

    fmt.Println("Post reported successfully.")
}
