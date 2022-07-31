package invoice

import "image/color"

type Config struct {
	EventName      string
	BackgroundPath string
	TextProperties []Properties
}

type Properties struct {
	FontPath  string
	FontSize  float64
	PositionX float64
	PositionY float64
	Color     color.RGBA
}

type InvoiceDetails struct {
	Date            string
	TeamName        string
	LeaderName      string
	MemberOne       string
	MemberTwo       string
	CompetitionName string
	Price           string
}
