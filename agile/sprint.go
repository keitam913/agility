package agile

import "math"

type Sprint struct {
	issues []Issue
	done   bool
}

func NewSprint() Sprint {
	return Sprint{
		issues: []Issue{},
		done:   false,
	}
}

func (s Sprint) AddIssue(issue Issue) {
	s.issues = append(s.issues, issue)
}

func (s Sprint) SetDone(done bool) {
	s.done = done
}

func (s Sprint) AllCommitment() int {
	c := 0
	for _, i := range s.issues {
		c += i.Size()
	}
	return c
}

func (s Sprint) AllDone() int {
	c := 0
	for _, i := range s.issues {
		if i.HasDone() {
			c += i.Size()
		}
	}
	return c
}

func (s Sprint) AllVelocity(lastSprints []Sprint) int {
	if !s.done {
		return -1
	}
	sum := s.AllDone()
	for i := len(lastSprints); 0 < i && len(lastSprints)-2 < i; i-- {
		sum += lastSprints[i].AllDone()
	}
	return int(math.Round(float64(sum) / 3))
}

func (s Sprint) Commitment(team string) int {
	c := 0
	for _, i := range s.issues {
		if i.HasCommittedBy(team) {
			c += i.Size()
		}
	}
	return c
}

func (s Sprint) Done(team string) int {
	c := 0
	for _, i := range s.issues {
		if i.HasCommittedBy(team) && i.HasDone() {
			c += i.Size()
		}
	}
	return c
}

func (s Sprint) Velocity(team string, lastSprints []Sprint) int {
	if !s.done {
		return -1
	}
	sum := s.Done(team)
	for i := len(lastSprints); 0 < i && len(lastSprints)-2 < i; i-- {
		sum += lastSprints[i-1].Done(team)
	}
	return int(math.Round(float64(sum) / 3))
}