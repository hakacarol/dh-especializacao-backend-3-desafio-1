package tickets

import (
	"os"
	"strings"
	"fmt"
)

// estrutura que receberá os dados do csv
type Ticket struct {
	Id int
	Name string 
	Email string
	DestinationCountry string
	BoardingTime string
	Price float64
}

// Uma função que calcule a quantidade de pessoas que viajam para um país determinado
func GetTotalTickets(DestinationCountry string) (int, error) {

	file, err := os.ReadFile("./tickets.csv")
	
	if err != nil {
	  panic("Erro ao ler o arquivo")
	}
	
	tickets := strings.Split(string(file), "\n")

	//fmt.Println(tickets)

	countTickets := 0

	for i := 0; i < len(tickets)-1; i++ {
		registersInfo := strings.Split(tickets[i], ",")

		//fmt.Println(registersInfo[3])

		if registersInfo[3] == DestinationCountry {
			countTickets++
		}
	}
	fmt.Println(countTickets)
	return countTickets, nil
}

// exemplo 2
// func GetMornings(time string) (int, error) {}

// exemplo 3
// func AverageDestination(destination string, total int) (int, error) {}
