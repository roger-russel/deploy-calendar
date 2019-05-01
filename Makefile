
.PHONY: compile-all
compile-all: compile-api-calendar compile-api-queue compile-consumer-queue
	@echo "Compiling all packages"

.PHONY: compile-api-calendar
compile-api-calendar:
	@go build -o ./build/api-calendar ./cmd/api/calendar/main.go

.PHONY: compile-api-queue
compile-api-queue:
	@go build -o ./build/api-queue ./cmd/api/queue/main.go

.PHONY: compile-consumer-queue
compile-consumer-queue:
	@go build -o ./build/consumer-queue ./cmd/consumer/queue/main.go