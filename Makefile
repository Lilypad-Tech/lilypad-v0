include go.env
go.env:
	go env | zsh -c export > $@

export CGO_ENABLED=0

HOMEBREW_ROOT := $(shell brew config | grep HOMEBREW_PREFIX | cut -d: -f2)
HOMEBREW_TAPS := ${HOMEBREW_ROOT}/Library/Taps

ETHEREUM_TAP := ${HOMEBREW_TAPS}/ethereum/homebrew-ethereum
${ETHEREUM_TAP}:
	brew tap ethereum/ethereum

SOLIDITY_BREW := ${HOMEBREW_ROOT}/Cellar/solidity
${SOLIDITY_BREW}: ${ETHEREUM_TAP}
	brew install solidity

PROTOC_BREW := ${HOMEBREW_ROOT}/Cellar/protoc-gen-go
${PROTOC_BREW}: ${SOLIDITY_BREW}
	brew install protoc-gen-go

GO_ETHEREUM := ${GOPATH}/pkg/mod/github.com/ethereum/go-ethereum@v1.10.26
${GO_ETHEREUM}:
	go get

ABIGEN := ${GOPATH}/bin/abigen
${ABIGEN}: ${PROTOC_BREW} ${GO_ETHEREUM}
	cd ${GO_ETHEREUM} && make devtools

%.abi: %.sol ${SOLIDITY_BREW}
	solc --output-dir $$(dirname $@) --abi $<

%.bin: %.sol ${SOLIDITY_BREW}
	solc --output-dir $$(dirname $@) --bin $<

%.go: %.abi %.bin ${ABIGEN}
	${ABIGEN} \
		--abi $(filter %.abi,$^) \
		--bin $(filter %.bin,$^) \
		--pkg $$(dirname $@ | xargs basename) \
		> $@

.PHONY: prepare
prepare: ${SOLIDITY_BREW} ${ABIGEN}

CONTRACTS := $(shell find hardhat/contracts -name '*.sol')
PACKAGES  := $(patsubst %.sol,%.go,${CONTRACTS})
BINARY    := $(shell pwd | xargs basename)

${BINARY}: ${PACKAGES} $(shell find pkg -name *.go)
	go build .

.PHONY: build
build: ${PACKAGES} ${BINARY}

.PHONY: clean
clean:
	-$(RM) contracts/*.abi
	-$(RM) contracts/*.bin
	-$(RM) contracts/*.go
	-${RM} ${BINARY}
