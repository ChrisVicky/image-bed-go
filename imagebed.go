package main

import (
	"context"
	"fmt"
	"imagebed/config"
	"log"
	"os"
	"path"
	"time"

	"github.com/google/go-github/v55/github"
	"golang.org/x/oauth2"
)

var conf = "~/.config/upload-img-github/config.toml"

func main() {
    c, err := config.ReadConfig(conf)
    if err != nil {
        log.Printf("Cannot find Configuration at %v", conf)
        return
    }
    
    if len(os.Args) < 2 {
        log.Printf("%v <path to image>\n", os.Args[0])
        return
    }

    filepath := os.Args[1]
    filename := path.Base(filepath)
    // filename := time.Now().Format("2006-01-02T15:04:05") + "-" + strings.TrimPrefix(filepath, "/")
    dirname := time.Now().Format("2006-01")

    content, err := os.ReadFile(filepath)
    if err != nil {
        panic(err)
    }

    opts := &github.RepositoryContentFileOptions{
        Message: github.String(fmt.Sprintf("[img] %s", filename)),
        Content: []byte(content),
    }

    ctx := context.Background()
    ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: c.Token})
    client := github.NewClient(oauth2.NewClient(ctx, ts))

    _, _, _, err = client.Repositories.GetContents(ctx, c.Owner, c.Repo, dirname, &github.RepositoryContentGetOptions{})
    if err != nil {
        readmeOpts := &github.RepositoryContentFileOptions{
            Message: github.String(fmt.Sprintf("[dir] %v", dirname)),
            Content: []byte(fmt.Sprintf("# %v", dirname)),
        }
        _, _, err = client.Repositories.CreateFile(ctx, c.Owner, c.Repo, dirname + "/README.md", readmeOpts)
        if err != nil {
            log.Fatalf("%v", err)
            return
        }
    }

    // Upload the image to GitHub
    _, _, err = client.Repositories.CreateFile(ctx, c.Owner, c.Repo, dirname+"/"+filename, opts)
    if err != nil {
        panic(err)
    }

    url := fmt.Sprintf("%s/%s/%s/main/%s/%s", c.BaseURL, c.Owner, c.Repo, dirname, filename)

    fmt.Printf("%v", url)
}

