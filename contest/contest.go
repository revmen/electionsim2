package contest

import (
	"revmen/electionsim2/config"
	"revmen/electionsim2/util"
)

type Contest struct {
	Candidates        []Candidate
	Voters            []Voter
	StrategicFraction float64 // Fraction of strategic voters, 0 to 1
}

func GenerateContests(config *config.Config) ([]Contest, error) {
	contests := make([]Contest, 0, len(config.Contests))

	for _, cg := range config.Contests {
		for i := 0; i < cg.Count; i++ {
			contest, err := generateContest(cg, config.Candidates, config.Voters)
			if err != nil {
				return nil, err
			}
			contests = append(contests, contest)
		}
	}

	return contests, nil
}

func generateContest(cg config.ContestGen, candidateGens []config.CandidateGen, voterGens []config.VoterGen) (Contest, error) {

	candidates := make([]Candidate, 0, len(candidateGens))
	for _, cg := range candidateGens {
		candidates = append(candidates, GenerateCandidate(cg))
	}

	strategicFraction := util.RandBetween(cg.MinStrategic, cg.MaxStrategic)
	voters := make([]Voter, 0)

	// Create voters from each voter block
	for _, vg := range voterGens {
		blockVoters, err := GenerateVoters(vg, strategicFraction, candidates)
		if err != nil {
			return Contest{}, err
		}
		voters = append(voters, blockVoters...)
	}

	return Contest{
		Candidates:        candidates,
		Voters:            voters,
		StrategicFraction: strategicFraction,
	}, nil
}
