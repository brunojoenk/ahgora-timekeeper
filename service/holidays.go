package service

import "time"

const (
	defaultFormat = "2006-01-02"

	//CarnivalOptionalPoint - Ponto facultativo Carnaval
	//CarnivalOptionalPoint = "2020-03-04"

	//Carnival - Carnaval
	Carnival = "2020-02-25"

	//GoodFriday - Sexta-feira Santa
	GoodFriday = "2020-04-10"

	//Tiradentes
	Tiradentes = "2020-04-21"

	//MayDay - Dia do Trabalho
	MayDay = "2020-05-01"

	//CorpusChristi - Corpus Christi
	//CorpusChristi = "2020-06-20"

	//ProclamationOfTheRepublic - Proclamação Da República
	ProclamationOfTheRepublic = "2020-11-15"

	//Festa de encerramento
	//CompanyParty = "2020-12-13"

	//ChristmasEve - Véspera de Natal
	ChristmasEve = "2020-12-24"

	//Christmas - Natal
	Christmas = "2020-12-25"

	//End's year - Fim de ano
	EndsYear = "2020-12-31"

	//New's year - Inicio de ano
	NewsYear = "2021-01-01"
)

func isHoliday(day time.Time) bool {
	switch day.Format(defaultFormat) {
	case Carnival, GoodFriday, Tiradentes, MayDay, ProclamationOfTheRepublic, ChristmasEve, Christmas, EndsYear, NewsYear:
		return true
	default:
		return false
	}
}
