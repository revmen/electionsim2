package contest

import (
	"revmen/electionsim2/config"
	"revmen/electionsim2/util"
)

type Candidate struct {
	Direction float64 // Angle from the origin, 0 to 2π
	Distance  float64 // Distance from the origin, 0 to 1
	Notoriety float64 // How well known the candidate is, 0 to 1
}

func GenerateCandidate(cg config.CandidateGen) Candidate {
	return Candidate{
		Direction: util.RandBetween(cg.MinDirection, cg.MaxDirection),
		Distance:  util.RandBetween(cg.MinDistance, cg.MaxDistance),
		Notoriety: util.RandBetween(cg.MinNotoriety, cg.MaxNotoriety),
	}
}
