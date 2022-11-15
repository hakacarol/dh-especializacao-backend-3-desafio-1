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

	totalTickets, err := tickets.GetTotalTickets("Jamaica")
	
	if err != nil {
		panic(err)
	}

	fmt.Printf("total de tickets %d \n", totalTickets)
}
