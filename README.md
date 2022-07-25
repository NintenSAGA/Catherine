![image-20220725132035374](assets/title.png)

# Catherine 👠

![Go](https://img.shields.io/badge/-Golang-087CFA?style=flat&logo=go&logoColor=white)
![Platform](https://img.shields.io/badge/Platform-macOS_|_Linux-white)

A simple TUI block pushing game.

## 1. Introduction 🤔

This game was first written in C in my freshman year (2020).

Recently, in order to backup my pre-github projects, I dug it out again. But besides puting the original one in my repo, I also decided to rewrite it in Go to get more familiar with the language.

### Why Catherine? 

Well, obviously it’s named after the game _Catherine_ (Atlus, 2011), and even the name of the hero and the background story are same as its. 

This is mainly because _Catherine_ is the first game I can call to mind when thinking about block pushing game (though I wrote the game in plain TUI 2D).

## 2. Build & Run 🛠

The build scripts of both versions are written in Makefile.

### [Root](.)

| Target   | Description        |
| -------- | ------------------ |
| `run_go` | Build and run the Go version. |
| `run_c`  | Build and run the C version.  |

### [C Version](./c_version/) / [Go Version](./go_version/)

| Target            | Description    |
| ----------------- | -------------- |
| `(default) / all` | Build only. |
| `run`             | Build and run the game.   |

## 3. Demo 🎥

https://user-images.githubusercontent.com/72867349/180706054-056b555d-18f2-461a-8fa6-7ded9fef1b40.mov
