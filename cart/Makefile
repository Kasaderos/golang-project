CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.51.1
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
PACKAGE=route256/cart/cmd/app

all: format build test lint

build: bindir
	go build -o ${BINDIR}/app ${PACKAGE}

test:
	go test ./...

run:
	go run ${PACKAGE}

lint: install-lint
	${LINTBIN} run

precommit: format build test lint
	echo "OK"

bindir:
	mkdir -p ${BINDIR}

format: install-smartimports
	${SMARTIMPORTS} -exclude internal/mocks

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})

# Используем bin в текущей директории для установки плагинов protoc
LOCAL_BIN:=$(CURDIR)/bin

# Добавляем bin в текущей директории в PATH при запуске protoc
PROTOC = PATH="$$PATH:$(LOCAL_BIN)" protoc

# Установка всех необходимых зависимостей
.PHONY: .bin-deps
.bin-deps:
	$(info Installing binary dependencies...)

	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest && \
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@latest && \
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/buf/cmd/buf@v1.21.0


# Вендоринг внешних proto файлов
.vendor-proto: vendor-proto/google/api vendor-proto/google/protobuf vendor-proto/protoc-gen-openapiv2/options vendor-proto/validate

.PHONY: .vendor-rm
.vendor-rm:
	rm -rf vendor-proto

# Устанавливаем proto описания google/googleapis
vendor-proto/google/api:
	git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/googleapis/googleapis vendor-proto/googleapis && \
 	cd vendor-proto/googleapis && \
	git sparse-checkout set --no-cone google/api && \
	git checkout
	mkdir -p  vendor-proto/google
	mv vendor-proto/googleapis/google/api vendor-proto/google
	rm -rf vendor-proto/googleapis

# Устанавливаем proto описания protoc-gen-openapiv2/options
vendor-proto/protoc-gen-openapiv2/options:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/grpc-ecosystem/grpc-gateway vendor-proto/grpc-ecosystem && \
 	cd vendor-proto/grpc-ecosystem && \
	git sparse-checkout set --no-cone protoc-gen-openapiv2/options && \
	git checkout
	mkdir -p vendor-proto/protoc-gen-openapiv2
	mv vendor-proto/grpc-ecosystem/protoc-gen-openapiv2/options vendor-proto/protoc-gen-openapiv2
	rm -rf vendor-proto/grpc-ecosystem


# Устанавливаем proto описания google/protobuf
vendor-proto/google/protobuf:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/protocolbuffers/protobuf vendor-proto/protobuf &&\
	cd vendor-proto/protobuf &&\
	git sparse-checkout set --no-cone src/google/protobuf &&\
	git checkout
	mkdir -p vendor-proto/google
	mv vendor-proto/protobuf/src/google/protobuf vendor-proto/google
	rm -rf vendor-proto/protobuf

# Устанавливаем proto описания validate
vendor-proto/validate:
	git clone -b main --single-branch --depth=2 --filter=tree:0 \
		https://github.com/bufbuild/protoc-gen-validate vendor-proto/tmp && \
		cd vendor-proto/tmp && \
		git sparse-checkout set --no-cone validate &&\
		git checkout
		mkdir -p vendor-proto/validate
		mv vendor-proto/tmp/validate vendor-proto/
		rm -rf vendor-proto/tmp


CARTS_PROTO_PATH:="api/carts/v1"
PRODUCTS_PROTO_PATH:="api/products/v1"

# Генерация протофайлов с использованием protoc
PHONY: .protoc-generate
.protoc-generate:
	mkdir -p pkg/${CARTS_PROTO_PATH}
	$(PROTOC) -I . -I vendor-proto \
    --plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go --go_out=paths=source_relative:./pkg \
    --plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc --go-grpc_out=paths=source_relative:pkg \
	--plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)/protoc-gen-grpc-gateway --grpc-gateway_out=paths=source_relative:pkg --grpc-gateway_opt logtostderr=true --grpc-gateway_opt generate_unbound_methods=true \
	--plugin=protoc-gen-validate=$(LOCAL_BIN)/protoc-gen-validate --validate_out="lang=go,paths=source_relative:pkg" \
	./$(CARTS_PROTO_PATH)/service.proto
	mkdir -p pkg/${PRODUCTS_PROTO_PATH}
	$(PROTOC) -I ${PRODUCTS_PROTO_PATH} -I vendor-proto \
	--plugin=protoc-gen-go=$(LOCAL_BIN)/protoc-gen-go --go_out pkg/${PRODUCTS_PROTO_PATH} --go_opt paths=source_relative \
	--plugin=protoc-gen-go-grpc=$(LOCAL_BIN)/protoc-gen-go-grpc --go-grpc_out pkg/${PRODUCTS_PROTO_PATH} --go-grpc_opt paths=source_relative \
	--plugin=protoc-gen-openapiv2=$(LOCAL_BIN)/protoc-gen-openapiv2 --openapiv2_out api/openapiv2/ --openapiv2_opt logtostderr=true \
	$(PRODUCTS_PROTO_PATH)/service.proto
	go mod tidy

# Генерация протофайлов с использованием protoc
PHONY: generate
generate: .bin-deps .vendor-proto .protoc-generate

.PHONY: fast-generate
fast-generate: .protoc-generate

# Генерация протофайлов с использованием buf
PHONY: generate-buf
generate-buf: .bin-deps .vendor-rm .buf-generate

# Генерация протофайлов с использованием buf
.PHONY: .buf-generate
.buf-generate:
	PATH=$(LOCAL_BIN):$(PATH) $(LOCAL_BIN)/buf mod update
	PATH=$(LOCAL_BIN):$(PATH) $(LOCAL_BIN)/buf generate

MIGRATE:=$(LOCAL_BIN)/goose
MIGRATE_DATABASE_DRIVER=postgres
MIGRATE_DATABASE_HOST=localhost
MIGRATE_DATABASE_PORT=7432
MIGRATE_DATABASE_URL="postgresql://postgres:passwd@$(MIGRATE_DATABASE_HOST):$(MIGRATE_DATABASE_PORT)/postgres?sslmode=disable"

.PHONY: migrate
migrate:
	$(MIGRATE) -dir migrations $(MIGRATE_DATABASE_DRIVER) $(MIGRATE_DATABASE_URL) up

.PHONY: reset_migrations
reset_migrations:
	$(MIGRATE) -dir migrations $(MIGRATE_DATABASE_DRIVER) $(MIGRATE_DATABASE_URL) reset

.PHONY: go_generate
go_generate:
	go generate ./...

.PHONY: test_coverage
test_coverage:
	go test ./... -count=1 -cover -coverprofile cover.out
	go tool cover -func cover.out

.PHONY: test_coverage_html
test_coverage_html:
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out