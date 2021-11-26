package models

import(
	"github.com/toxtashev/LogisticApp/utils"
)

func GetById(id string)(output []GetOrderById) {

	rows, _ := utils.DB.Query(GET_BY_ID, id)
	
	defer utils.DB.Close()
	
	for rows.Next() {

		var outputItem GetOrderById
	
		rows.Scan(&outputItem.Name,
				  &outputItem.Quantity,
				  &outputItem.ReceiverFullName,
				  &outputItem.ReceiverPasspost,
				  &outputItem.ReceiverPhone,
				  &outputItem.ReceiverAddress,
				  &outputItem.ReceiverCity,
				  &outputItem.SenderFullName,
				  &outputItem.SenderPassport,
				  &outputItem.SenderPhone,
				  &outputItem.SenderEmail,
				  &outputItem.SenderAddress,
				  &outputItem.SenderCity,
				  &outputItem.SenderPostcode,
				  &outputItem.OrderId,
				  &outputItem.Comment,
				  &outputItem.Value,
				  &outputItem.Ordered_at,
				  &outputItem.Deleted_at,
			)

		output = append(output, outputItem)
	}

	return
}