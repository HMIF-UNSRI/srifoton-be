package invoice

import (
	"strconv"
	"time"

	teamDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"
)

func CreateInvoiceDetails(t teamDomain.Team) InvoiceDetails {
	var price string
	switch string(t.Competition) {
	case "CP":
		price = "Rp100000"
	case "UI/UX":
		price = "Rp80000"
	case "WEB":
		price = "Rp80000"
	case "ESPORT":
		price = "Rp100000"
	}

	invoiceDetails := InvoiceDetails{
		Date:            time.Now().Format("02 January 2006"),
		TeamName:        t.Name,
		LeaderName:      t.Leader.Name,
		MemberOne:       t.Member1.Name,
		MemberTwo:       t.Member2.Name,
		MemberThree:     t.Member3.Name,
		MemberFour:      t.Member4.Name,
		MemberFive:      "Anggota 5   : " + t.Member5.Name,
		CompetitionName: t.GetUCompetitionTypeString(),
		Price:           price,
	}
	return invoiceDetails
}

func (i InvoiceDetails) MapToArray() []string {
	arr := []string{
		i.Date,
		i.TeamName,
		i.LeaderName,
		i.MemberOne,
		i.MemberTwo,
		i.MemberThree,
		i.MemberFour,
		i.MemberFive,
		i.CompetitionName,
		i.Price,
		i.Price,
		i.Price,
	}
	return arr
}

func StringToUint8(numberStr string) uint8 {
	numberInt, err := strconv.Atoi(numberStr)
	if err != nil {
		panic(err)
	}
	return uint8(numberInt)
}
