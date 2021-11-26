create table senders(
	sender_id serial not null primary key,
	firstname varchar(64) not null,
	lastname varchar(64) not null, 
	passport varchar(10) not null,
	phone varchar(25) not null,
	email varchar(254) not null,
	address varchar(256) not null,
	city varchar(128) not null,
	postcode varchar(10)
);

create table receivers(
	receiver_id serial not null primary key,
	firstname varchar(64) not null,
	lastname varchar(64) not null, 
	passport varchar(10) not null,
	phone varchar(15) not null,
	address varchar(256) not null,
	city varchar(128) not null
);

create table orders(
	order_id serial not null primary key,
	sender_id int not null references senders(sender_id),
	receiver_id int not null references receivers(receiver_id),
	comment text,
	value int not null,
	ordered_at timestamptz default current_timestamp,
	isActivated bool default false,
	activated_at timestamptz,
	isFinished bool default false,
	finished_at timestamptz,
	isDeleted bool default false,
	deleted_at timestamptz
);

create table items(
	item_id serial not null primary key,
	name varchar(256) not null,
	quantity int not null,
	order_id int not null references orders(order_id)
);