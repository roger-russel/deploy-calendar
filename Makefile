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
packages: jwt
	@dep ensure

.PHONY: dev
dev:
	@realize start

.PHONY: jwt jwt-private jwt-public
jwt: jwt-private jwt-public
	@echo "JWT Keys have been created"

jwt-private:
	@openssl ecparam -name secp521r1 -genkey -noout -out ./keys/jwt_ecdsa_private_key.pem

jwt-public:
	@openssl ec -in ./keys/jwt_ecdsa_private_key.pem -pubout -out ./keys/jwt_ecdsa_public_key.pem

.PHONY: prisma
prisma:
	@prisma deploy
