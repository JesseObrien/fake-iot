# fake-iot

This project is an exercise in design, code quality and showing the ability to think through approaches carefully. It is meant to be an example of how I would design a system that consumes data from a client that's sending metrics and in turn show that to a user on the front-end in real time.

This project uses the (golang-standards project-layout)[https://github.com/golang-standards/project-layout] for folder structure consistency.

## Requirements

- `docker`
- `docker-compose`
- `make`
- `nodejs`

## Install

`make install`

## Build

`make build`

## Development

Using docker-compose, you'll need to be running a postgres database.

`docker-compose up -d`

Using a little utility called (air)[https://github.com/cosmtrek/air] you can rebuild all files on watch.

`make watch`

### Development Notes

The user id and account ID for the user are hard coded as follows:

- user_id: de7169a0-1ca1-4f18-8fb8-29d3a7cafd30
- account_id: 47f3c307-6344-49e7-961c-ea200e950a89

Run the fakeiot client with the hard coded token and account ID:

`fakeiot --token=882e8f9b-76a3-46fb-9f7e-bd536bdf5795 --url="https://127.0.0.1:8080" --ca-cert=./certs/server.crt run --account-id=47f3c307-6344-49e7-961c-ea200e950a89 --users=100`

The database file is seeded via [scripts/init-pg.sql](scripts/init-pg.sql)`.
