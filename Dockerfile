FROM golang:alpine AS builder
WORKDIR /app/
RUN apk add --update openssl git nodejs npm make
RUN go get github.com/rakyll/statik 
COPY . .
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY ./certs ./certs
COPY --from=builder /app/tmp/fake-iot .

## @NOTE I am hard coding these here to save time on running. Ideally they would be in some kind of secrets
## manager or provided at run time when the container is ran in whatever environment it is.
ENV FAKEIOT_API_TOKEN=882e8f9b-76a3-46fb-9f7e-bd536bdf5795
ENV DATABASE_URL=postgresql://testuser:abcd1234@fakeiot-postgres:5432/fakeiot?sslmode=disable

CMD ["./fake-iot", "-host=0.0.0.0"]