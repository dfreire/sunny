#!/bin/bash
DATABASE="development.db"

#echo -e "\n# schema"
#sqlite3 $DATABASE ".schema"

echo -e "\n# customer_role"
sqlite3 $DATABASE "select * from customer_role"

echo -e "\n# customer"
sqlite3 $DATABASE "select * from customer"

echo -e "\n# wine_comment"
sqlite3 $DATABASE "select * from wine_comment"
