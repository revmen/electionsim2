package plurality

import "revmen/electionsim2/contest"

type PluralityContests struct {
	Contests []PluralityContest
}

type PluralityBallot struct {
	CandidateIndex int
}

type PluralityContest struct {
	Candidates        []contest.Candidate
	HonestVoters      []PluralityHonestVoter
	StrategicVoters   []PluralityStrategicVoter
	StrategicFraction float64
	WinnerIndex       int
	HonestUtility     float64
	StrategicUtility  float64
}

func GeneratePluralityContest(c *contest.Contest) PluralityContest {
	candidates := c.Candidates
	honestVoters := make([]PluralityHonestVoter, 0)
	strategicVoters := make([]PluralityStrategicVoter, 0)

	for _, v := range c.Voters {
		if v.Strategic {
			sv := PluralityStrategicVoter{Voter: v}
			strategicVoters = append(strategicVoters, sv)
		} else {
			hv := PluralityHonestVoter{Voter: v}
			honestVoters = append(honestVoters, hv)
		}
	}

	return PluralityContest{
		Candidates:        candidates,
		HonestVoters:      honestVoters,
		StrategicVoters:   strategicVoters,
		StrategicFraction: c.StrategicFraction,
		WinnerIndex:       -1,
		HonestUtility:     0,
		StrategicUtility:  0,
	}
}

type PluralityHonestVoter struct {
	Voter contest.Voter
}

func (v PluralityHonestVoter) Vote(candidates []contest.Candidate) PluralityBallot {
	var bestCandidateIndex int
	var bestUtility float64
	for i, u := range v.Voter.Utilities {
		if u.Known && u.Utility > bestUtility {
			bestCandidateIndex = i
			bestUtility = u.Utility
		}
	}
	return PluralityBallot{CandidateIndex: bestCandidateIndex}
}

type PluralityStrategicVoter struct {
	Voter contest.Voter
}

type PluralityStrategy struct {
}
