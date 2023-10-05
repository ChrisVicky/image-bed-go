# Github Image Host

## Feature and Functionality
* Input: Image Path (Stored Locally)
* Output: Image URL (Stored in Github Repo)

## Dependency
* Github Api: github.com/google/go-github/v55/github
* Toml: github.com/BurntSushi/toml

## Usage

### 1. Github Setup

1. Github Repo (Public) Setup
2. Token Generation
* https://github.com/settings/tokens
* ![image_2023-09-25-20-11-38](https://raw.githubusercontent.com/ChrisVicky/image-bed/main/2023-09/image_2023-09-25-20-11-38.png)

### 2. Program Configuration
1. Copy Configurations
```shell
cp config.toml.example config.toml
```
2. Configurations
```toml
# Owner of the Repo
owner       = "xxxxx"

# Repo for storing
repo        = "github-image-bed"

# Token generated from Github
token       = "xxxxxxxxxxxxxxxx"

# Baseurl for fetching image
baseURL     = "https://raw.githubusercontent.com"
```

* Configuration Default Location: `~/.config/upload-img-github/config.toml`

### 3. Run
```shell
go run imagebed.go ./img/blured.png
```

## Vim-Integration
![vim-image-bed-integration](./img/vim-image-bed-integration.gif)

