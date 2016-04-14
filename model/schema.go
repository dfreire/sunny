package model

const SCHEMA = `
    CREATE TABLE IF NOT EXISTS Customer (
    	id        TEXT PRIMARY KEY,
    	email     TEXT NOT NULL,
    	role      TEXT,
    	createdAt TEXT
    );

    CREATE TABLE IF NOT EXISTS WineComment (
    	id         TEXT PRIMARY KEY,
    	customerId TEXT NOT NULL,
    	wineId     TEXT NOT NULL,
    	wineYear   NUMBER,
    	createdAt  TEXT,
    	updatedAt  TEXT,
    	comment    TEXT NOT NULL,

        FOREIGN KEY(customerId) REFERENCES Customer(id),
    	UNIQUE(customerId, wineId, wineYear)
    );
`
