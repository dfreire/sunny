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
    	wineYear   NUMBER NOT NULL,
    	comment    TEXT NOT NULL,
        createdAt  TEXT NOT NULL,
    	updatedAt  TEXT NOT NULL,

        FOREIGN KEY(customerId) REFERENCES Customer(id),
    	UNIQUE(customerId, wineId, wineYear)
    );

    INSERT INTO Customer (id, email) VALUES ("customer-1", "dario.freire@gmail.com");
    INSERT INTO WineComment (id, customerId, wineId, wineYear, comment, createdAt, updatedAt) VALUES
        ("comment-1", "customer-1", "wine-1", 2015, "great", "a", "b");
`
