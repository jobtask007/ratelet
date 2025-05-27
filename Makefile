.PHONY: ut deps lint run mocks image run-image

deps:
	go mod tidy

ut:
	go test -v -tags=unit ./...

run: deps
	go run cmd/main.go

mocks:
	mockery

image:
	docker build -t ratelet:1.0.0 .

run-image:
	docker run -it --rm -e OPEN_EXCHANGE_RATES_APP_ID=ca589a9519fa4596aef3e4ea0b84c2c0 -p 8080:8080 --name ratelet ratelet:1.0.0