package Db

var CREATEDATABASE = `SELECT 'CREATE DATABASE TICKET_BOOKING'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'TICKET_BOOKING');`

var CREATETABLE = `

	CREATE TABLE IF NOT EXISTS MOVIES(
		id UUID DEFAULT uuid_generate_v4(),
		movie_name VARCHAR NOT NULL UNIQUE
);

	CREATE TABLE IF NOT EXISTS MOVIES_AVAILABILITY(
		movie_id UUID,
		tickets_available integer,
		price float8 NOT NULL,
		dates DATE NOT NULL
);

	CREATE TABLE IF NOT EXISTS BOOKING(
		id UUID DEFAULT uuid_generate_v4(),
		name VARCHAR NOT NULL,
		tickets_unit integer,
		total_price float,
		date DATE NOT NULL
);

`
