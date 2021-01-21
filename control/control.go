package control

import (
	"FileCleaner/model"
)

func CtlParallel(control int) {
	if control <= 0 || control > 10 {
		panic("[ERROR]: The number of parallel control must between 1 and 10 !")
	}
	for i := 0; i < 10-control; i++ {
		model.ControlCH <- 1
	}
}
