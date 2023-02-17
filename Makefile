SHELL = bash

.PHONY: all
all: build

export CGO_ENABLED=0
export GOOS ?= $(shell go env GOOS)
export GOARCH ?= $(shell go env GOARCH)
export GOPATH ?= $(shell go env GOPATH)

HOMEBREW_ROOT := $(shell brew config | grep HOMEBREW_PREFIX | cut -d: -f2)
PROTOC_BREW := ${HOMEBREW_ROOT}/Cellar/protoc-gen-go
${PROTOC_BREW}:
	brew install protoc-gen-go

GO_ETHEREUM := ${GOPATH}/pkg/mod/github.com/ethereum/go-ethereum@v1.10.26
${GO_ETHEREUM}:
	go get

ABIGEN ?= ${GOPATH}/bin/abigen
${ABIGEN}: ${PROTOC_BREW} ${GO_ETHEREUM}
	cd ${GO_ETHEREUM} && make devtools

HARDHAT ?= hardhat/node_modules/.bin/hardhat
${HARDHAT}:
	cd hardhat && npm install

.PHONY: prepare
prepare: ${ABIGEN} ${HARDHAT}

CONTRACTS := $(shell find hardhat/contracts -name '*.sol')

include hardhat/contracts/.deps.mk
hardhat/contracts/.deps.mk: ${CONTRACTS} ${MAKEFILES}
	echo ${CONTRACTS} | \
		xargs -n1 basename -s .sol | \
		tr ' ' "\n" | \
		awk '{ \
			print "hardhat/artifacts/contracts/" $$1 ".sol/" $$1 ".json: " "hardhat/contracts/" $$1 ".sol"; \
			print "ABIJSONS += hardhat/artifacts/contracts/" $$1 ".sol/" $$1 ".json"; \
			print "ABIJSONS += hardhat/artifacts/contracts/" $$1 ".sol/" $$1 ".dbg.json"; \
			print "PACKAGES += hardhat/artifacts/contracts/" $$1 ".sol/" $$1 ".go"; \
		}' > $@

${ABIJSONS}: | ${HARDHAT}
	cd hardhat && npx hardhat compile

%.go: %.json | ${ABIGEN}
	${ABIGEN} \
		--abi <(cat $< | jq -rc '.abi') \
		--bin <(cat $< | jq -rc '.bytecode') \
		--pkg $(shell dirname $@ | xargs basename -s .sol) \
		> $@


BASENAME  := $(shell pwd | xargs basename)
BINARY    := bin/${BASENAME}-${GOOS}-${GOARCH}

OSES     := darwin linux
ARCHES   := arm64 amd64
BINARIES := $(foreach OS,${OSES}, $(foreach ARCH,${ARCHES},bin/${BASENAME}-${OS}-${ARCH}))

bin/:
	mkdir -p $@

bin/${BASENAME}-%: ${PACKAGES} $(shell find pkg -name '*.go') main.go | bin/
	GOOS=$(shell echo $@ | cut -f2 -d'-') GOARCH=$(shell echo $@ | cut -f3 -d'-') go build -o $@ .

.PHONY: build
build: ${PACKAGES} ${BINARIES}

ENV_FILE ?= hardhat/.env
ifeq ($(shell cat ${ENV_FILE} | grep WALLET_PRIVATE_KEY),)
$(warning No WALLET_PRIVATE_KEY in ${ENV_FILE})
endif
ifeq ($(shell cat ${ENV_FILE} | grep CONTRACT_ADDRESS),)
$(warning No CONTRACT_ADDRESS in ${ENV_FILE})
endif

.PHONY: deploy
deploy: ${BINARIES} | ops/deploy.sh ${ENV_FILE}
	cd ops && ./deploy.sh

.PHONY: run
run: ${BINARY} | ${ENV_FILE}
	env $$(cat ${ENV_FILE}) ${BINARY}

.PHONY: clean
clean:
	-$(RM) hardhat/contracts/.deps.mk
	-$(RM) ${PACKAGES}
	-$(RM) ${ABIJSONS}
	-${RM} -r bin/
