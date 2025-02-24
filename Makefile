


.PHONY: build
build:
	CGO_ENABLED=0 go build -o ./bin/swagger cmd/swagger/swagger.go


.PHONY: install
install:
	CGO_ENABLED=0 go install cmd/swagger/swagger.go
