package model

const SCHEMA = `
	PRAGMA foreign_keys=ON;
	
	CREATE TABLE IF NOT EXISTS CustomerRoleEnum (
		id TEXT PRIMARY KEY
	);
	INSERT OR IGNORE INTO CustomerRoleEnum(id) VALUES
		('sommelier'), ('restaurant'), ('wine_distribution'),
		('wine_shop'), ('wine_lover'), ('other');
		
	CREATE TABLE IF NOT EXISTS SignupOriginEnum (
		id TEXT PRIMARY KEY
	);
	INSERT OR IGNORE INTO SignupOriginEnum(id) VALUES
		('wine_comment'), ('newsletter');
	
    CREATE TABLE IF NOT EXISTS Customer (
    	id             TEXT PRIMARY KEY,
    	email          TEXT NOT NULL,
    	roleId         TEXT NOT NULL,
    	createdAt      TEXT NOT NULL,
		signupOriginId TEXT NOT NULL,
		
		FOREIGN KEY (roleId) REFERENCES CustomerRoleEnum(id),
		FOREIGN KEY (signupOriginId) REFERENCES SignupOriginEnum(id)
    );

    CREATE TABLE IF NOT EXISTS WineComment (
    	id         TEXT PRIMARY KEY,
    	customerId TEXT NOT NULL,
    	wineId     TEXT NOT NULL,
    	wineYear   NUMBER NOT NULL,
    	comment    TEXT NOT NULL,
        createdAt  TEXT NOT NULL,
		updatedAt  TEXT NOT NULL,

        FOREIGN KEY (customerId) REFERENCES Customer(id),
    	UNIQUE (customerId, wineId, wineYear)
    );
`
