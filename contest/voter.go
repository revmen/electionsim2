package contest

import (
	"revmen/electionsim2/config"
	"revmen/electionsim2/util"
)

type Voter struct {
	Direction   float64            // Angle from the origin, 0 to 2π
	Distance    float64            // Distance from the origin, 0 to 1
	Information float64            // Information level, 0 to 1
	Strategic   bool               // Whether the voter is strategic
	Utilities   []CandidateUtility // Utility of each candidate for this voter
}

type CandidateUtility struct {
	CandidateIndex int     // Index of the candidate in the contest
	Known          bool    // Whether the voter knows this candidate
	Utility        float64 // Utility of the candidate for this voter, whether or not they know them
}

func GenerateVoters(vg config.VoterGen, strategicFraction float64, candidates []Candidate) ([]Voter, error) {
	voters := make([]Voter, 0, vg.Count)
	for i := 0; i < vg.Count; i++ {
		voter, err := generateVoter(vg, strategicFraction, candidates)
		if err != nil {
			return nil, err
		}
		voters = append(voters, voter)
	}
	return voters, nil
}

func generateVoter(vg config.VoterGen, strategicFraction float64, candidates []Candidate) (Voter, error) {
	voter := Voter{
		Direction:   util.RandBetween(vg.MinDirection, vg.MaxDirection),
		Distance:    util.RandBetween(vg.MinDistance, vg.MaxDistance),
		Information: util.RandBetween(vg.MinInformation, vg.MaxInformation),
		Strategic:   util.RandBool(strategicFraction),
		Utilities:   make([]CandidateUtility, len(candidates)),
	}

	// Calculate utility for each candidate
	// A candidate is considered known if the candidate's notoriety exceeds the voter's information level
	// The maximum possible distance between two points is 2.0, so utility is 2.0 - distance / 2.0

	for i, c := range candidates {
		utility := CandidateUtility{
			CandidateIndex: i,
			Known:          c.Notoriety >= voter.Information,
			Utility:        1.0 - util.Distance(voter.Distance, voter.Direction, c.Distance, c.Direction)/2.0,
		}
		voter.Utilities[i] = utility
	}

	return voter, nil
}
