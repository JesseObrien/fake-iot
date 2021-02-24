.ONESHELL:
.PHONY: install
install:
	go get -u github.com/rakyll/statik
	go get -u github.com/cosmtrek/air
	go mod vendor
	cd web/ && npm install

.PHONY: build
build:
	cd web
	npm run build
	rm -rf ./statik
	statik -f -src=./public
	cd ..
	go build -ldflags="-s -w" -o ./tmp/fake-iot ./cmd/fake-iot/main.go

.PHONY: test
test:
	go test -race ./... 

.PHONY: watch
watch:
	export FAKEIOT_API_TOKEN=882e8f9b-76a3-46fb-9f7e-bd536bdf5795
	export DATABASE_URL=postgresql://testuser:abcd1234@localhost:5432/fakeiot?sslmode=disable
	air

.PHONY: clean 
clean:
	docker-compose down

.PHONY: certs
certs:
	cd certs

	openssl req -x509 -nodes -new -sha256 -days 1024 -newkey rsa:2048 -keyout RootCA.key -out RootCA.pem -subj "/C=US/CN=localhost"
	openssl x509 -outform pem -in RootCA.pem -out RootCA.crt
	openssl req -new -nodes -newkey rsa:2048 -keyout server.key -out server.csr -subj "/C=CA/ST=London/L=London/O=localhost/CN=localhost.local"
	openssl x509 -req -sha256 -days 1024 -in server.csr -CA RootCA.pem -CAkey RootCA.key -CAcreateserial -extfile domains.ext -out server.crt

.PHONY: psql
psql:
	docker-compose exec postgres /bin/sh -c "psql -U testuser -d fakeiot"