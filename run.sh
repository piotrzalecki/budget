#!/bin/bash

export DB_NAME=budget_dev
export DB_USER=postgres
export DB_PASSWORD=password



go build -o budget ./cmd/web/*.go && ./budget -dbname=bookings -dbuser=postgres -dbpass=mysecretpassword