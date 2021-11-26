package models

import(
	"github.com/toxtashev/LogisticApp/utils"
)

func Get()(output []GetOrders) {

	rows, _ := utils.DB.Query(GET)
	
	defer utils.DB.Close()
	
	for rows.Next() {

		var outputItem GetOrders

		rows.Scan(
				  &outputItem.ByOrder,
				  &outputItem.OrderId,
				  &outputItem.SenderFullName,
				  &outputItem.ReceiverFullName,
				  &outputItem.ReceiverCity,
				  )

		output = append(output, outputItem)
	}

	return
}