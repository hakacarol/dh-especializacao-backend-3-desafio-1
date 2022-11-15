package tickets

import (
	"os"
	"strings"
	"strconv"
)

// estrutura que receberá os dados do csv
type Ticket struct {
	Id string
	Name string 
	Email string
	DestinationCountry string
	BoardingTime string
	Price string
}

// Requisito 01: Uma função que calcule a quantidade de pessoas que viajam para um país determinado
func GetTotalTickets(DestinationCountry string) (int, []Ticket, error) {

	file, err := os.ReadFile("./tickets.csv")
	
	if err != nil {
	  panic("Erro ao ler o arquivo")
	}
	
	tickets := strings.Split(string(file), "\n")

	countTickets := 0
	var ticketsResults []Ticket

	for i := 0; i < len(tickets); i++ {
		registersInfo := strings.Split(tickets[i], ",")

		if registersInfo[3] == DestinationCountry {
			countTickets++

			ticket := Ticket{
				Id: registersInfo[0],
				Name: registersInfo[1],
				Email: registersInfo[2],
				DestinationCountry: registersInfo[3],
				BoardingTime: registersInfo[4],
				Price: registersInfo[5],
			}

			ticketsResults = append(ticketsResults, ticket)
		}
	}

	return countTickets, ticketsResults, nil
}

// Requisito 2: Uma ou várias funções que calculam quantas pessoas viajam de madrugada (0 → 6), manhã (7 → 12), tarde (13 → 19), e noite (20 → 23).
const (
	earlyMorning = "earlyMorning"
	morning = "morning"
	afternoon = "afternoon"
	evening = "evening"
)

func GetPeriods(period string) (int, []Ticket, error) {

	var startTime int
	var endTime int

	const (
		earlyMorning = "earlyMorning"
		morning = "morning"
		afternoon = "afternoon"
		evening = "evening"
	)

	switch period {
	case "earlyMorning":
		startTime = 0
		endTime = 6
	
	case "morning":
		startTime = 7 
		endTime = 12

	case "afternoon":
		startTime = 13
		endTime = 19
	
	case "evening":
		startTime = 20
		endTime = 23
	}

	file, err := os.ReadFile("./tickets.csv")
	
	if err != nil {
	  panic(err)
	}
	
	tickets := strings.Split(string(file), "\n")

	countTickets := 0
	var ticketsResults []Ticket

	for i := 0; i < len(tickets); i++ {

		registersInfo := strings.Split(tickets[i], ",")
		timeFormat := strings.Split(registersInfo[4], ":")
		boardingTime, err := strconv.Atoi(timeFormat[0])
		if err != nil {
			panic(err)
		}

		if boardingTime >= startTime && boardingTime <= endTime{
			countTickets++

			ticket := Ticket{
				Id: registersInfo[0],
				Name: registersInfo[1],
				Email: registersInfo[2],
				DestinationCountry: registersInfo[3],
				BoardingTime: registersInfo[4],
				Price: registersInfo[5],
			}

			ticketsResults = append(ticketsResults, ticket)
		}
	}

	return countTickets, ticketsResults, nil
}



// Requisito 03
func AverageDestination() (int, error) {
	file, err := os.ReadFile("./tickets.csv")
	
	if err != nil {
	  panic("Erro ao ler o arquivo")
	}
	
	tickets := strings.Split(string(file), "\n")

	totalTickets := len(tickets)

	countries := []string{}

	for i := 0; i < len(tickets); i++ {
		registersInfo := strings.Split(tickets[i], ",")
		countryInfo := registersInfo[3]

		containCountry := false

		for _, country := range countries {
			if countryInfo == country {
				containCountry = true
			}
		}

		if !containCountry {
			countries = append(countries, countryInfo)
		}
	}

	average := totalTickets / len(countries)
	return average, nil
}
 