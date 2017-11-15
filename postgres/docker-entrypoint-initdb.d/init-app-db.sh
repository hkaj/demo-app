#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 -U "$POSTGRES_USER" -d app <<-EOSQL
CREATE TABLE person (
    first_name text,
    last_name text
);

INSERT INTO person (first_name, last_name) VALUES ('Haissam', 'Kaj');
INSERT INTO person (first_name, last_name) VALUES ('Data', 'Dog');
INSERT INTO person (first_name, last_name) VALUES ('Leeroy', 'Jenkins');
EOSQL
