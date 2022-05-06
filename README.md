# Persistent Web Server

Small web server written in golang with persistent data in postgresql database.


## Table of content

* [Install](#install)
* [Configuration file](#configuration-file)
* [Run](#run)
* [API Documentation](#api-documentation)
* [API Usage](#api-usage)
* [Built with](#build-with)


## Install

You can install Go from [Go official website](https://go.dev/) or by doing
```console
wget https://go.dev/dl/go1.18.1.linux-amd64.tar.gz
tar -xvf go1.18.1.linux-amd64.tar.gz
mv go go-1.18.1
sudo mv go-1.18.1 /usr/local
```

You can install Postgresql from [PostgreSQL official website](https://www.postgresql.org/) or by doing 
```console
sudo apt install postgresql 
```

## Configuration File
Rename `example.env` to `.env` and fill its fields  
Should look something like this
```
DIALECT="postgres"
DB_HOST="localhost"
DB_PORT="5432"
DB_USER="user"
DB_NAME="car_data"
DB_PASSWORD="password"
```


## Run

To start postgresql and create database
```
sudo service postgresql start
createdb car_data
```
Run web server
```
go run main.go
```

## API Documentation

TODO

## API Usage
You can find [demo.rest](demo.rest) in the root directoy with some request examples.  
For further testing use `curl`.  
Some examples:
1. Insert data with POST request (using Form Data)
```
curl -X POST -F 'Latitude=387365578' -F 'Longitude=-91389050' -F 'Time=12:35:13' -F 'Speed=57' http://localhost:5000/put_data
```
2. Insert data with GET request
```
curl "http://localhost:5000/put_data?Latitude=387365578&Longitude=-91389050&Time=12:35:13&Speed=57"
```
3. Get Data with GET request
```
curl "http://localhost:5000/get_data"
```

## Build with
- [gorrila/mux](github.com/gorilla/mux) - HTTP router and URL matcher
- [gorrila/schema](github.com/gorilla/schema) - converts structs to and from form values
- [gorm](https://github.com/go-gorm/gorm) - object-relational mapping
- [godotenv](github.com/joho/godotenv) - load env variables from a .env file