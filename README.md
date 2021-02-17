# fake-iot

This project is an exercise in design, code quality and showing the ability to think through approaches carefully. It is meant to be an example of how I would design a system that consumes data from a client that's sending metrics and in turn show that to a user on the front-end in real time.

This project uses the (golang-standards project-layout)[https://github.com/golang-standards/project-layout] for folder structure consistency.

## Requirements

`docker`

`make`

`nodejs`

## Install

`make install`

## Build

`make build`

## Development

**Note for Logging In:** Hard coded username/password is `test@example.com` and `p@ssw0rd`.

Using docker-compose, you'll need to be running a postgres database.

`docker-compose up -d`

Using a little utility called (air)[https://github.com/cosmtrek/air] you can rebuild all files on watch.

`make watch`

**Note:** Development API token (For the fakeiot CLI):

`--token=882e8f9b-76a3-46fb-9f7e-bd536bdf5795`
