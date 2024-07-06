#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username admin --dbname admin <<-EOSQL
	CREATE DATABASE common;
  \c common;

  CREATE TABLE "flags" (
    "name" character varying NOT NULL,
    "status" boolean NOT NULL,
    "id" uuid NOT NULL
  );
EOSQL
