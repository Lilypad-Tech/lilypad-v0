.PHONY: all
all: build

include go.env
go.env:
	go env | zsh -c export > $@

export CGO_ENABLED=0

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

artifacts/%.json artifacts/%.dbg.json: %.sol ${HARDHAT}
	cd hardhat && npx hardhat compile

%.go: %.json ${ABIGEN}
	${ABIGEN} \
		--abi <(cat $< | jq -rc '.abi') \
		--bin <(cat $< | jq -rc '.bytecode') \
		--pkg $(shell dirname $@ | xargs basename -s .sol) \
		> $@

.PHONY: prepare
prepare: ${ABIGEN} ${HARDHAT}

CONTRACTS := $(shell find hardhat/contracts -name '*.sol')
PACKAGES  := $(shell echo '${CONTRACTS}' | xargs -n1 basename -s .sol | tr ' ' "\n" | awk '{ print "hardhat/artifacts/contracts/" $$1 ".sol/" $$1 ".go" }')
BINARY    := ./$(shell pwd | xargs basename)

${BINARY}: ${PACKAGES} $(shell find pkg -name '*.go') main.go
	go build .

.PHONY: build
build: ${PACKAGES} ${BINARY}

.PHONY: run
run: ${BINARY}
	env $$(cat hardhat/.env) ${BINARY}

.PHONY: clean
clean:
	-$(RM) ${PACKAGES}
	-${RM} ${BINARY}
