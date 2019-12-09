package service

import "time"

const (
	defaultFormat = "2006-01-02"

	//CarnivalOptionalPoint - Ponto facultativo Carnaval
	CarnivalOptionalPoint = "2019-03-04"

	//Carnival - Carnaval
	Carnival = "2019-03-05"

	//GoodFriday - Sexta-feira Santa
	GoodFriday = "2019-04-19"

	//MayDay - Dia do Trabalho
	MayDay = "2019-05-01"

	//CorpusChristi - Corpus Christi
	CorpusChristi = "2019-06-20"

	//ProclamationOfTheRepublic - Proclamação Da República
	ProclamationOfTheRepublic = "2019-11-15"

	//Festa de encerramento
	CompanyParty = "2019-12-13"

	//ChristmasEve - Véspera de Natal
	ChristmasEve = "2019-12-24"

	//Christmas - Natal
	Christmas = "2019-12-25"

	//End's year
	EndsYear = "2019-12-31"

	//New's year
	NewsYear = "2020-01-01"
)

func isHoliday(day time.Time) bool {
	switch day.Format(defaultFormat) {
	case CarnivalOptionalPoint, Carnival, GoodFriday, MayDay, CorpusChristi, ProclamationOfTheRepublic, CompanyParty, ChristmasEve, Christmas, EndsYear, NewsYear:
		return true
	default:
		return false
	}
}
