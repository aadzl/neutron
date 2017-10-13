all: build
start:
	node mailserver/server.js --config=./mail_config.toml && \
	go run neutron.go
build:
	go build -o neutron neutron.go
build-client:
	cd public && \
	npm install && \
	node_modules/.bin/grunt ngconstant:dev build
build-mail:
	cd mailserver && \
	npm install
build-docker:
	CGO_ENABLED=0 GOOS=linux go build -o neutron-docker -a -installsuffix cgo -ldflags '-s'
clean-client-dist:
	rm -rf public/node_modules public/test
