#!/bin/bash
set -e
export PGPASSWORD=postgres123;
psql -v ON_ERROR_STOP=1 --username "postgres" <<-EOSQL
  CREATE DATABASE bmdb;
  GRANT ALL PRIVILEGES ON DATABASE bmdb TO "postgres";
EOSQL