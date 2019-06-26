# Live Codding : Discovering Go : webview using zserge/webview go-preact-webview
An example of React app working with go like electron js do. 

## Introduction
In order to construct a Desktop APP we have tested the build of an app using Golang , and React with Webview. 
This stack aims to have a powerfull APP with low memory and cpu consumption. 

## Usage
** To refresh assets.go file (when updating react files ) **
```bash
go-bindata -o assets.go js/react/... js/styles.css 
```
** Build **
```bash
# Linux
$ go build -o webview-example && ./webview-example

# MacOS uses app bundles for GUI apps
$ mkdir -p example.app/Contents/MacOS
$ go build -o example.app/Contents/MacOS/example
$ open example.app # Or click on the app in Finder

# Windows requires special linker flags for GUI apps.
# It's also recommended to use TDM-GCC-64 compiler for CGo.
# http://tdm-gcc.tdragon.net/download
$ go build -ldflags="-H windowsgui" -o webview-example.exe
```

## Video Source
https://www.twitch.tv/videos/444165596

## Twitch Channel
https://www.twitch.tv/calyscope

## Links
https://github.com/zserge/webview
https://github.com/zserge/lorca
https://preactjs.com/


Have a nice DAY Best Regards

