package model

const SCHEMA = `
	PRAGMA foreign_keys=ON;
	
	CREATE TABLE IF NOT EXISTS customer_role (
		id TEXT PRIMARY KEY
	);
	INSERT OR IGNORE INTO customer_role VALUES
		('sommelier'), ('restaurant'), ('wine_distribution'),
		('wine_shop'), ('wine_lover'), ('other');
			
    CREATE TABLE IF NOT EXISTS customer (
    	id               TEXT PRIMARY KEY,
        created_at       DATETIME NOT NULL,
        updated_at       DATETIME NOT NULL,
		name             TEXT,
    	email            TEXT NOT NULL UNIQUE,
    	role_id          TEXT NOT NULL,
        wants_newsletter BOOL,
		in_newsletter    BOOL,
		
		FOREIGN KEY (role_id) REFERENCES customer_role(id)
    );

    CREATE TABLE IF NOT EXISTS wine_comment (
    	id          TEXT PRIMARY KEY,
        created_at  DATETIME NOT NULL,
        updated_at  DATETIME NOT NULL,
    	customer_id TEXT NOT NULL,
    	wine_id     TEXT NOT NULL,
    	wine_year   NUMBER NOT NULL,
    	comment     TEXT NOT NULL,

        FOREIGN KEY (customer_id) REFERENCES customer(id),
    	UNIQUE (customer_id, wine_id, wine_year)
    );
`
