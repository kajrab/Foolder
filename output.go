package main

import "github.com/fatih/color"

func printBanner() {
	color.Cyan(`
  ______           _     _
 |  ____|         | |   | |
 | |__ ___   ___  | | __| | ___ _ __
 |  __/ _ \ / _ \ | |/ _  |/ _ \ '__|
 | | | (_) | (_) || | (_| |  __/ |
 |_|  \___/ \___/ |_|\__,_|\___|_|
`)
	color.White(" \n Fooling around in folders you shouldn't be in.\n")
	color.HiBlack(" ----------------------------------------------\n\n")
}
