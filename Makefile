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


.PHONY: watch
watch:
	export FAKEIOT_API_TOKEN=882e8f9b-76a3-46fb-9f7e-bd536bdf5795
	air

.PHONY: certs
certs:
	docker run --rm -w /certs --mount src="`pwd`/certs",target=/certs,type=bind golang:alpine /bin/sh -c "go run /usr/local/go/src/crypto/tls/generate_cert.go --host localhost && chmod 755 *.pem"