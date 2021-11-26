package models

type Item struct {
	Name, Quantity string
}

type NewOrder struct {
	ReceiverFirstName, ReceiverLastName, ReceiverPassport, ReceiverPhone, ReceiverAddress, ReceiverCity,
	SenderFirstName, SenderLastName, SenderPassport, SenderPhone, SenderEmail, SenderAddress, SenderCity, SenderPostcode,
	Comment string
	Value uint
	Items []Item
}

const POST_SENDER = `
		insert into senders (firstname, lastname, passport, phone, email, address, city, postcode) values ($1, $2, $3, $4, $5, $6, $7, $8)	
`

const POST_RECEIVER = `
		insert into receivers (firstname, lastname, passport, phone, address, city) values ($1, $2, $3, $4, $5, $6)
`

const POST_ORDER = `
		insert into orders (sender_id, receiver_id, comment, value) values ((select max(sender_id) from senders), (select max(receiver_id) from receivers), $1, $2)
`

const POST_MAIN = `
		insert into items(name,quantity,order_id) values ($1, $2 ,(select max(order_id) from orders))
`