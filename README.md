# titleshot

This is program that takes a screenshot of a window given the title.

```go
mike@bullseye ~/projects/titleshot⭐  ./titleshot -l
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
xgbutil/main.go at master · BurntSushi/xgbutil · GitHub - Mozilla Firefox
titleshot
mike@bullseye ~/projects/titleshot⭐  ./titleshot -title  'fatih/vim-go: Go development plugin for Vim - Chromium'
mike@bullseye ~/projects/titleshot⭐  ls
README.md  screenshot.png  titleshot  titleshot.go
mike@bullseye ~/projects/titleshot⭐  ls *.png
screenshot.png
```

## TODO

* Allow user to pass a name/path to save the image
* Is any of this possible on MacOS?
* Is any of this possible on Windows?
* Is any of this possible under Wayland?
