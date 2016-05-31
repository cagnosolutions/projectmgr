package main

import (
	"fmt"

	"github.com/cagnosolutions/projectmgr"
)

const (
	CODING_PERC        = float64(50)
	BRAINSTORMING_PERC = float64(25)
	TESTING_PERC       = float64(25)
)

func main() {
	p := projectmgr.NewProject()
	p.Name = "New Driver Hire - CNS"
	p.Budget = 7500
	p.Rate = 150
	p.Developers = 1
	p.CodingPercent = CODING_PERC
	p.BrainstormingPercent = BRAINSTORMING_PERC
	p.TestingPercent = TESTING_PERC
	fmt.Println("\n\t+===============================================+")
	fmt.Printf("\t|\tProject Manager Breakdown Utility\t|")
	fmt.Println("\n\t+===============================================+")
	fmt.Printf("\t|       Project Name:  %q\t|\n", p.Name)
	fmt.Printf("\t|     Project Budget:  $%.2f budgeted\t|\n", p.Budget)
	fmt.Printf("\t|     Developer Rate:  $%.2f current rate\t|\n", p.Rate)
	fmt.Printf("\t| Project Developers:  %.3d developer(s)\t\t|\n", int(p.Developers))
	fmt.Println("\t+-----------------------------------------------+")
	fmt.Printf("\t|      Project Hours:  %.3d hrs\t\t\t|\n", int(p.BillableTime()))
	fmt.Printf("\t|       Project Days:  %.3d days\t\t\t|\n", int(p.Days()))
	fmt.Println("\t+-----------------------------------------------+")
	fmt.Printf("\t|       Coding Hours:  %.3d hrs,\t%.3d percent\t|\n", int(p.CodingHrs()), int(p.CodingPercent))
	fmt.Printf("\t|     Planning Hours:  %.3d hrs,\t%.3d percent\t|\n", int(p.BrainstormingHrs()), int(p.BrainstormingPercent))
	fmt.Printf("\t|      Testing Hours:  %.3d hrs,\t%.3d percent\t|\n", int(p.TestingHrs()), int(p.TestingPercent))
	fmt.Println("\t+-----------------------------------------------+\n")
}
