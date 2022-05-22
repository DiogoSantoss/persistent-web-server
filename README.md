# Persistent Web Server

Small web server written in golang with persistent data in postgresql database.


## Table of content

* [Install](#install)
* [Configuration file](#configuration-file)
* [Run](#run)
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

## Run with Docker
Create a network
```
docker network create go-postgres-network
```

Run the postgres container
```
docker run -d \
-p 5432:5432 \
-e POSTGRES_PASSWORD=password \
--name postgresdb \
--net go-postgres-network \
postgres
```

Create the database inside the container
```
docker ps
docker exec -it 8c67db24bfcf bash
psql -U postgres
CREATE DATABASE car_data;
```

Build the go image
```
docker build . \
-t go-postgres-server
```

Run the container
```
docker run -it \
-p 8080:5000 \
--network go-postgres-network \
go-web-server
```

## Run with Docker Compose

Run the docker compose
```
docker-compose -f docker-compose.yaml up
```
To turn off
```
docker-compose -f docker-compose.yaml down
```

## Additional Docker commands
List running containers
```
docker ps
```
List stored images
```
docker images
```
List networks
```
docker network ls
```

## API Testing
Use [request.py](requests.py) to stress test the API


## Build with
- [gorrila/mux](github.com/gorilla/mux) - HTTP router and URL matcher
- [gorrila/schema](github.com/gorilla/schema) - converts structs to and from form values
- [gorm](https://github.com/go-gorm/gorm) - object-relational mapping
- [godotenv](github.com/joho/godotenv) - load env variables from a .env file