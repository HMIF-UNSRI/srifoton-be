package team

func (t *Team) GetUCompetitionTypeString() string {
	switch string(t.Competition) {
	case string(Cp):
		return "Competitive Programming"
	case string(Web):
		return "Web Development"
	case string(UiUx):
		return "UI/UX Design"
	case string(Esport):
		return "E-Sport"
	}

	return ""
}

func (t *Team) SetTeamCompetitionString(c string) {
	switch c {
	case string(Cp):
		t.Competition = "CP"
	case string(Web):
		t.Competition = "WEB"
	case string(UiUx):
		t.Competition = "UI/UX"
	case string(Esport):
		t.Competition = "ESPORT"
	}
}
