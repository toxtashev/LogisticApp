package models

import (
	"github.com/toxtashev/LogisticApp/utils"
)

func Create (body NewOrder) {

	utils.DB.Exec(POST_SENDER,
					body.SenderFirstName,
					body.SenderLastName,
					body.SenderPassport,
					body.SenderPhone, 
					body.SenderEmail, 
					body.SenderAddress, 
					body.SenderCity, 
					body.SenderPostcode,
				)


	utils.DB.Exec(POST_RECEIVER,
					body.ReceiverFirstName,
					body.ReceiverLastName,
					body.ReceiverPassport,
					body.ReceiverPhone,
					body.ReceiverAddress,
					body.ReceiverCity,
				)
							
	utils.DB.Exec(POST_ORDER,
					body.Comment,
					body.Value,
				)

	for _, data := range body.Items {

		utils.DB.Exec(POST_MAIN,
					data.Name,
					data.Quantity,
				)
	}

	defer utils.DB.Close()
}
