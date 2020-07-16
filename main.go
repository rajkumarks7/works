package main

import (
	"fmt"
	"os"
	"time"

	"github.com/equinox-io/equinox"
	hook "github.com/robotn/gohook"
)

const (
	appID   = "app_bCi34JJ6Anp"
	version = "1.0.0"
)

var publicKey = []byte(`
-----BEGIN ECDSA PUBLIC KEY-----
MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAE7FVdU31H9TJbVmHBmUObVCwq97J7iD0z
mSPvNyhtf44XElnQqgUNUy9enCI449LlMipGBen+WgltYVKgqSlRhWEhhfXcK7tY
c6S/HtXArrEAOYLLeFe9V5545vbhuQvf
-----END ECDSA PUBLIC KEY-----
`)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "update" {
		equinoxUpdate()
	}
	mousecount()
}

func mousecount() {
	for {
		EvChan := hook.Start()
		defer hook.End()
		var c int
		b := 0
		j := 0
		var h uint16
		start := time.Now()

		starttime := fmt.Sprint(start.Format("15:04:05"))
		fmt.Println(starttime)
		var end string
		for ev := range EvChan {

			fmt.Sprint("hook: ", ev)
			fmt.Println(ev.When)

			c = c + int(ev.Direction)/3
			if ev.Clicks == 1 {
				j++
				fmt.Println(j)
			} else if ev.Rawcode >= 1 {
				fmt.Println(ev.Rawcode)
				if ev.Rawcode != h {
					h = ev.Rawcode
					b = b + 1
				} else {
					fmt.Println("it is pressed already")
				}
				fmt.Println(b)
			}
			if time.Since(start) >= time.Second*60 {
				endtime := time.Now()
				end = fmt.Sprint(endtime.Format("15:04:05"))
				fmt.Println(end)
				break
			}
		}
		Event := fmt.Sprint(start.Format("02-01-2006"))
		fmt.Println(Event)
		Eventtime := fmt.Sprint(start.Format("15:04:05"))
		fmt.Println(Eventtime)
		events := 0
		if j == 0 {
			events = 0
		} else {
			events = (j - c) / 3
		}
		fmt.Println(events)
		fmt.Println(b)
		fmt.Println(b / 3)
		time.Sleep(250 * time.Millisecond)
	}
}

// this function will make selfupdate of program
func equinoxUpdate() error {
	var opts equinox.Options
	if err := opts.SetPublicKeyPEM(publicKey); err != nil {
		return err
	}

	// check for the update
	resp, err := equinox.Check(appID, opts)
	switch {
	case err == equinox.NotAvailableErr:
		fmt.Println("No update available, already at the latest version!")
		return nil
	case err != nil:
		fmt.Println("Update failed:", err)
		return err
	}

	// fetch the update and apply it
	err = resp.Apply()
	if err != nil {
		return err
	}

	fmt.Printf("Updated to new version: %s!\n", resp.ReleaseVersion)
	return nil
}
