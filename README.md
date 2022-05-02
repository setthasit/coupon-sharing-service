# Coupon Sharing Service
this project will contain REST API for research purpose, This project build with `Gin/Gonic` and the use of `Gorm` and `Wire`. and is create to learn how wire work and how to implement it.

## This Project Create With
Service: gin/gonic
ORM: gorm
Database: Postgresql

## to run this project
1. run `go mod tidy` to tidying up the dependencies and install needed dependencies as well
2. copy `.env.example` to `.env` to create new `.env` file
3. change the `.env` file values
4. run `docker-compose up -d` to run postgres container, it will load the config from `.env` file
5. run `make dev` to start the app
