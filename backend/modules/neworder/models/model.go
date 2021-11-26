package models

import(
	"time"
)

type GetOrders struct {
	ByOrder uint
	OrderId uint
	SenderFullName, ReceiverFullName, ReceiverCity string
}

type GetOrderById struct {
	Name, ReceiverFullName, ReceiverPasspost, ReceiverPhone, ReceiverAddress, ReceiverCity,
	SenderFullName, SenderPassport, SenderPhone, SenderEmail, SenderAddress, SenderCity, SenderPostcode,
	Comment,Quantity string
	Value, OrderId uint
	Ordered_at time.Time 
}

const GET = `select * from (
					select 
						row_number() over (
							order by o.order_id desc
						) as rn,
						o.order_id,
						s.firstname || ' ' || s.lastname as SenderFullname,
						r.firstname || ' ' || r.lastname as ReceiverFullname,
						r.city
					from orders as o
 					join senders as s on s.sender_id = o.sender_id					
 					join receivers as r on r.receiver_id = o.receiver_id
 					where isActivated = false And isDeleted = false
					order by o.order_id desc
					) smth
`

const GET_BY_ID = `
					select 
						i.name, i.quantity,
						r.firstname || ' ' || r.lastname as ReceiverFullname, r.passport, r.phone, r.address, r.city,
						s.firstname || ' ' || s.lastname as SenderFullname, s.passport, s.phone, s.email, s.address, s.city, s.postcode,
 						o.order_id, o.comment, o.value, o.ordered_at
 					from items as i
 					join orders as o using(order_id)
 					join receivers as r on r.receiver_id = o.receiver_id
 					join senders as s on s.sender_id = o.sender_id
 					where i.order_id = $1 AND isActivated = false AND isDeleted = false

`

type OnlyId struct {
	Id int
}

const COMPLETED = `
					update orders set (isActivated, activated_at) = (true, now()) where order_id = $1
`

const DELETED = `
					update orders set (isDeleted, deleted_at) = (true, now()) where order_id = $1
`