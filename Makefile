
.PHONY: compile-all
compile-all: packages compile-api-calendar compile-api-queue compile-consumer-queue
	@echo "Compiling all packages"

.PHONY: compile-api-calendar
compile-api-calendar:
	@go build -o ./dist/api-calendar ./cmd/api/calendar/main.go

.PHONY: compile-api-queue
compile-api-queue:
	@go build -o ./dist/api-queue ./cmd/api/queue/main.go

.PHONY: compile-consumer-queue
compile-consumer-queue:
	@go build -o ./dist/consumer-queue ./cmd/consumer/queue/main.go

.PHONY: packages
packages:
	@dep ensure

.PHONY: dev
dev:
	@realize start
