# titleshot

This is program that takes a screenshot of a window given the title.

Listing the existing windows can be done with the `--list` option:

```go
mike@bullseye ~/projects/titleshot⭐  ./titleshot --list
XdndCollectionWindowImp
unity-launcher
unity-panel
unity-dash
Hud
Desktop
fatih/vim-go: Go development plugin for Vim - Chromium
Rhythmbox
root@bullseye: ~
Transmission
```
To take a screenshot just pass the title of the window. By default it saves `screenshot.png` in the current directory:

```go
mike@bullseye ~/projects/titleshot⭐  ./titleshot --title  'fatih/vim-go: Go development plugin for Vim - Chromium'
mike@bullseye ~/projects/titleshot⭐  ls
README.md  screenshot.png  titleshot  titleshot.go
```
If you want the screenshot to be saved elsewhere:

```go
./titleshot --title "Terminal" --output ~/terminal.png
```

## TODO

* Get a description string working
* Optionally print MD5 hash of screenshot. Possibly to stdout at exit.
* Is any of this possible on MacOS?
* Is any of this possible on Windows?
* Is any of this possible under Wayland?
