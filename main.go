package main

import (
	"fmt"
	"os"
	"revmen/electionsim2/config"
	"revmen/electionsim2/contest"
)

func main() {
	fmt.Println("Rev's Election Sim #2")

	config, err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Print configuration struct
	fmt.Printf("Loaded config: %+v\n",
		config)

	fmt.Println("Generating contests...")
	contests, err := contest.GenerateContests(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating contests: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %d contests\n", len(contests))

	// Print summary of the first contest as an example
	if len(contests) > 0 {
		c := contests[0]
		fmt.Println("\nFirst contest example:")
		fmt.Printf("  Strategic voters: %.2f%%\n", c.StrategicFraction*100)
		fmt.Printf("  Candidates: %d\n", len(c.Candidates))
		for i, c := range c.Candidates {
			fmt.Printf("    Candidate %d: direction=%.2f, distance=%.2f, notoriety=%.2f\n", i, c.Direction, c.Distance, c.Notoriety)
		}
		fmt.Printf("  Voters: %d\n", len(c.Voters))
		v := c.Voters[0]
		fmt.Printf("  First voter example:\n")
		fmt.Printf("    Direction: %.2f\n", v.Direction)
		fmt.Printf("    Distance: %.2f\n", v.Distance)
		fmt.Printf("    Information: %.2f\n", v.Information)
		fmt.Printf("    Strategic: %t\n", v.Strategic)

		fmt.Printf("    Utilities: %d\n", len(v.Utilities))
		for i, u := range v.Utilities {
			fmt.Printf("      Candidate %d: known=%t, utility=%.2f\n", i, u.Known, u.Utility)
		}

	}
}
