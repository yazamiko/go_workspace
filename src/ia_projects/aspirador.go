package main

import (
	"fmt"
)

type VacuumCleaner struct {
	initialRect *Rectangle
	actualRect  *Rectangle
}

func (vc *VacuumCleaner) exec() {
	fmt.Println("\nInitializing:\n")
	for {
		fmt.Print("is actual area clear? ")
		fmt.Println(vc.actualRect.isClean)

		if !vc.actualRect.isClean {
			fmt.Println("Cleaning...")
			vc.actualRect.isClean = true
			fmt.Print("Now is clear! ")
		}

		if vc.actualRect.nextRect == vc.initialRect {
			fmt.Println("\n\nAll area is clear, task is done.")
			break
		}

		fmt.Println("Moving...\n")
		vc.actualRect = vc.actualRect.nextRect
	}
}

type Rectangle struct {
	isClean  bool
	nextRect *Rectangle
}

func main() {
	rectA := new(Rectangle)
	rectB := new(Rectangle)
	rectC := new(Rectangle)
	rectD := new(Rectangle)

	rectA.isClean = false
	rectA.nextRect = rectB

	rectB.isClean = true
	rectB.nextRect = rectC

	rectC.isClean = true
	rectC.nextRect = rectD

	rectD.isClean = false
	rectD.nextRect = rectA

	vc := &VacuumCleaner{rectA, rectA}
	vc.exec()
}
