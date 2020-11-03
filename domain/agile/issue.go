package agile

type Issue struct {
	size       int
	labels     []string
	doneSprint string
}

func NewIssue(size int, labels []string, doneSprint string) Issue {
	return Issue{
		size:       size,
		labels:     labels,
		doneSprint: doneSprint,
	}
}

func (i Issue) Size() int {
	return i.size
}

func (i Issue) HasDone() bool {
	return i.doneSprint != ""
}

func (i Issue) DoneSprint() string {
	return i.doneSprint
}

func (i Issue) HasCommittedBy(team string) bool {
	for _, l := range i.labels {
		if l == team {
			return true
		}
	}
	return false
}