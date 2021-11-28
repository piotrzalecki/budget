#!/bin/bash

export DB_NAME=budget_dev
export DB_USER=postgres
export DB_PASSWORD=password



go build -o bookings ./cmd/web/*.go && ./bookings -dbname=bookings -dbuser=postgres -dbpass=mysecretpassword