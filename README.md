# 🤖 Gap - WIP

Gap is a modular REST API with which you can interact to view, add and remove content.
It's a project for practicing my Go skills so there's nothing really fancy, just me learning...

## 📑 Features

You can launch it from the command line and add some arguments.
```bash
$ gap -h
Usage: gap [-p PORT | -h]
$ gap -p 8080
Server is running at 'http://localhost:8080'
```

## 🔨 Building

#### 🐁 Go
```bash
$ go build .
```

#### ❄️ Nix
```bash
$ nix build .#
```
