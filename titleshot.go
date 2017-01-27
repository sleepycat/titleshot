package main

import (
	"fmt"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/icccm"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/alexflint/go-arg"
	"log"
)

func main() {

	var args struct {
		List  bool   `arg:"help:list window titles"`
		Title string `arg:"help:The title of the window you would like to screenshot."`
	}

	args.List = false
	args.Title = ""
	arg.MustParse(&args)

	// Connect to the X server using the DISPLAY environment variable.
	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	// Get a list of all client ids.
	clientids, err := ewmh.ClientListGet(X)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through each client, find its name and find its size.
	for _, clientid := range clientids {
		name, err := ewmh.WmNameGet(X, clientid)
		//pid, err := ewmh.WmPidGet(X, clientid)

		// If there was a problem getting _NET_WM_NAME or if its empty,
		// try the old-school version.
		if err != nil || len(name) == 0 {
			name, err = icccm.WmNameGet(X, clientid)

			// If we still can't find anything, give up.
			if err != nil || len(name) == 0 {
				name = "N/A"
			}
		}

		if args.List == true {
			fmt.Printf("%s\n", name)
			//fmt.Printf("%s\t %d\n", name, pid)
		}
		if args.Title != "" {
			if args.Title == name {
				// Use the "NewDrawable" constructor to create an xgraphics.Image value
				// from a drawable. (Usually this is done with pixmaps, but drawables
				// can also be windows.)
				ximg, err := xgraphics.NewDrawable(X, xproto.Drawable(clientid))
				if err != nil {
					log.Fatal(err)
				}

				err = ximg.SavePng("screenshot.png")
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
