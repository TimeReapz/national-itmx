#!/bin/bash
DB_FILE="customers.db"

CREATE_TABLE_QUERY="CREATE TABLE IF NOT EXISTS customers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    age INTEGER NOT NULL
);"
INSERT_CUSTOMER_QUERY='INSERT INTO customers (name, age) VALUES ("Tony Stark","30"),("Black Widow","25"),("Scarlet Witch","35");'

if [ ! -f $DB_FILE ]; then
    touch $DB_FILE
fi

sqlite3 $DB_FILE "$CREATE_TABLE_QUERY"
sqlite3 $DB_FILE "$INSERT_CUSTOMER_QUERY"

if [ $? -eq 0 ]; then
    echo "Table 'customers' created successfully."
else
    echo "Error creating table 'customers'."
fi
