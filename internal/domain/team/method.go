package team

func (t *Team) GetUCompetitionTypeString() string {
	return string(t.Competition)
}

func (t *Team) SetTeamCompetitionString(c string) {
	switch c {
	case string(Cp):
		t.Competition = "Competitive Programming"
	case string(Web):
		t.Competition = "Web Development"
	case string(UiUx):
		t.Competition = "UI/UX Design"
	case string(Esport):
		t.Competition = "E-Sport"
	}
}
