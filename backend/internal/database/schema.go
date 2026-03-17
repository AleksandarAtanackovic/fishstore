package database

/*
CREATE TABLE customers (
	id SERIAL PRIMARY KEY,
	name TEXT,
	surname TEXT,
	phone_number TEXT,
	created_at TIMESTAMP
);

CREATE TABLE orders (
	id SERIAL PRIMARY KEY,
	customer_id INT REFERENCES customers(id),
	order_time TIMESTAMP,
	completed BOOLEAN
);

CREATE TABLE order_items (
	id SERIAL PRIMARY KEY,
	order_id INT REFERENCES orders(id),
	fish_type TEXT,
	order_type TEXT,
	prepared BOOLEAN
);
*/
