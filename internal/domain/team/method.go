package team

func (t *Team) GetUCompetitionTypeString() string {
	return string(t.Competition)
}

func (t *Team) SetTeamCompetitionString(c string) {
	switch c {
	case string(Cp):
		t.Competition = Cp
	case string(Web):
		t.Competition = Web
	case string(UiUx):
		t.Competition = UiUx
	case string(Esport):
		t.Competition = Esport
	}
}
