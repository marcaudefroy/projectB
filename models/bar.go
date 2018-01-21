package models

import beez "github.com/marcaudefroy/projectC/models"

//Bar bar
type Bar struct {
	B beez.Beez
}

//C c
var C = Bar{
	B: beez.Beez{
		Name: "kk",
	},
}
