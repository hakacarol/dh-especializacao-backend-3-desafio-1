package main

import (
	"fmt"
	"github.com/hakacarol/dh-especializacao-backend-3-desafio-1/internal/tickets"
)



func main() {

	// file, err := os.Open("tickets.csv")
	// if err != nil {
	//   panic(err)
	// }
	// defer file.Close()

	// registers, err := csv.NewReader(file).ReadAll()

	// if err != nil {
    //     panic(err)
    // }

	// ticketRegister := make([]Ticket, len(registers))
    // for k, register := range registers {
	// 	price, _ := strconv.ParseFloat(register[4], 64) 
    //     ticket := Ticket{
	// 		Name: register[0],
	// 		Email: register[1],
	// 		DestinationCountry: register[2],
	// 		BoardingTime: register[3],
	// 		Price: price,
    //     }

    //     ticketRegister[k] = ticket
    // }
	
	// fmt.Println(ticketRegister)

	totalTickets, ticketsResults, err := tickets.GetTotalTickets("Brazil")
	
	if err != nil {
		panic(err)
	}

	fmt.Printf("total de tickets %d \n", totalTickets)
	fmt.Printf("tickets %s \n", ticketsResults)

	totalTicketsByPeriod, ticketsResults, err := tickets.GetPeriods("8:00")
	
	if err != nil {
		panic(err)
	}

	fmt.Printf("total de tickets %d \n", totalTicketsByPeriod)
	fmt.Printf("tickets %s \n", ticketsResults)

	averageDestination, err := tickets.AverageDestination()
	
	if err != nil {
		panic(err)
	}

	fmt.Println(averageDestination)
}
