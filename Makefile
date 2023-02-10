.PHONY: all
all: build

include go.env
go.env:
	go env | zsh -c export > $@

export CGO_ENABLED=0
export GOOS ?= $(shell go env GOOS)
export GOARCH ?= $(shell go env GOARCH)

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

CONTRACTS := $(shell find hardhat/artifacts/contracts -name '*.sol')
PACKAGES  := $(shell echo '${CONTRACTS}' | xargs -n1 basename -s .sol | tr ' ' "\n" | awk '{ print "hardhat/artifacts/contracts/" $$1 ".sol/" $$1 ".go" }')
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

.PHONY: run
run: ${BINARY}
	env $$(cat hardhat/.env) ${BINARY}

.PHONY: clean
clean:
	-$(RM) ${PACKAGES}
	-${RM} -r bin/
