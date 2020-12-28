package control

import "FileCleaner/model"

func Control(control int) {
	if control <= 0 || control > 10 {
		panic("The number of parallel control must between 1 and 10 !")
	}
	for i := 0; i < 10-control; i++ {
		model.ControlCH <- 1
	}
}
