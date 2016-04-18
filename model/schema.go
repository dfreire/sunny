package model

const SCHEMA = `
    CREATE TABLE IF NOT EXISTS Customer (
    	id        TEXT PRIMARY KEY,
    	email     TEXT NOT NULL,
    	role      TEXT NOT NULL,
    	createdAt TEXT NOT NULL
    );

    CREATE TABLE IF NOT EXISTS WineComment (
    	id         TEXT PRIMARY KEY,
    	customerId TEXT NOT NULL,
    	wineId     TEXT NOT NULL,
    	wineYear   NUMBER NOT NULL,
    	comment    TEXT NOT NULL,
        createdAt  TEXT NOT NULL,
		updatedAt  TEXT NOT NULL,

        FOREIGN KEY(customerId) REFERENCES Customer(id),
    	UNIQUE(customerId, wineId, wineYear)
    );
`
