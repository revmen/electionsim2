package config

import (
	"fmt"
	"math"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Contests   []ContestGen   `yaml:"contests"`
	Candidates []CandidateGen `yaml:"candidates"`
	Voters     []VoterGen     `yaml:"voters"`
	Plurality  bool           `yaml:"plurality"` // Evaluate using plurality voting
	Approval   bool           `yaml:"approval"`  // Evaluate using approval voting
	IRV        bool           `yaml:"irv"`       // Evaluate using instant-runoff voting
}

type ContestGen struct {
	Count        int     `yaml:"count"`        // Number of iterations to generate in this block of contests
	MinStrategic float64 `yaml:"minStrategic"` // Lowest fraction of strategic voters, 0 to 1
	MaxStrategic float64 `yaml:"maxStrategic"` // Highest fraction of strategic voters, 0 to 1
}

type CandidateGen struct {
	MinDirection float64 `yaml:"minDirection"` // Minimum angle from the origin, 0 to 2π
	MaxDirection float64 `yaml:"maxDirection"` // Maximum angle from the origin, 0 to 2π
	MinDistance  float64 `yaml:"minDistance"`  // Minimum distance from the origin, 0 to 1
	MaxDistance  float64 `yaml:"maxDistance"`  // Maximum distance from the origin, 0 to 1
	MinNotoriety float64 `yaml:"minNotoriety"` // How well known the candidate is, 0 to 1
	MaxNotoriety float64 `yaml:"maxNotoriety"` // How well known the candidate is, 0 to 1
}

type VoterGen struct {
	Count          int     `yaml:"count"`          // Number of voters to generate in this block
	MinDirection   float64 `yaml:"minDirection"`   // Minimum angle from the origin, 0 to 2π
	MaxDirection   float64 `yaml:"maxDirection"`   // Maximum angle from the origin, 0 to 2π
	MinDistance    float64 `yaml:"minDistance"`    // Minimum distance from the origin, 0 to 1
	MaxDistance    float64 `yaml:"maxDistance"`    // Maximum distance from the origin, 0 to 1
	MinInformation float64 `yaml:"minInformation"` // Minimum information level, 0 to 1
	MaxInformation float64 `yaml:"maxInformation"` // Maximum information level, 0 to 1
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %v", err)
	}

	// Convert angles from degrees to radians
	for i := range config.Candidates {
		config.Candidates[i].MinDirection = config.Candidates[i].MinDirection * math.Pi / 180.0
		config.Candidates[i].MaxDirection = config.Candidates[i].MaxDirection * math.Pi / 180.0
	}

	for i := range config.Voters {
		config.Voters[i].MinDirection = config.Voters[i].MinDirection * math.Pi / 180.0
		config.Voters[i].MaxDirection = config.Voters[i].MaxDirection * math.Pi / 180.0
	}

	// Create Config struct from config
	result := &Config{
		Contests:   config.Contests,
		Candidates: config.Candidates,
		Voters:     config.Voters,
	}

	return result, nil
}
