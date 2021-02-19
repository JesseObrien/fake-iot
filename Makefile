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
	go test ./...

.PHONY: watch
watch:
	export FAKEIOT_API_TOKEN=882e8f9b-76a3-46fb-9f7e-bd536bdf5795
	export DATABASE_URL=postgresql://testuser:abcd1234@localhost:5432/fakeiot
	air

.PHONY: clean 
clean:
	docker-compose down

.PHONY: certs
certs:
	cd certs
	openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes \
		-keyout server.key -out server.crt -extensions san -config \
		<(echo "[req]"; 
			echo distinguished_name=req; 
			echo "[san]"; 
			echo subjectAltName=DNS:example.com,DNS:www.example.net,IP:127.0.0.1
			) \
		-subj "/CN=example.com"

.PHONY: psql
psql:
	docker-compose exec postgres /bin/sh -c "psql -U testuser -d fakeiot"