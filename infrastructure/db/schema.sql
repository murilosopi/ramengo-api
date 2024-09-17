CREATE TABLE address (
	id integer auto_increment primary key,
	street varchar(255),
	number integer,
	zipcode varchar(10)
);


CREATE TABLE users (
	id integer auto_increment primary key,
	name varchar(255),
	email varchar(320),
	password varchar(72),
	address_id integer,
	
	foreign key (address_id) references address(id)
);


CREATE TABLE kitchens (
	id integer auto_increment primary key,
	address_id integer,
	
	foreign key (address_id) references address(id)
);

CREATE TABLE users_kitchens (
	user_id integer,
	kitchen_id integer,
	
	foreign key (user_id) references users(id),
	foreign key (kitchen_id) references kitchens(id)
);

CREATE TABLE status (
	id integer auto_increment primary key,
	description varchar(50)
);
INSERT INTO status (description) VALUES ("Confirmed"), ("Preparing"), ("Ready"), ("Canceled");

CREATE TABLE orders (
	id integer auto_increment primary key,
	kitchen_id integer,
	status_id integer,
	user_id integer,
	date date default current_date,
	
	foreign key (kitchen_id) references kitchens(id),
	foreign key (status_id) references status(id),
	foreign key (user_id) references users(id)
);