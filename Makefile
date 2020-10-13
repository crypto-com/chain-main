PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags 2>/dev/null ) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
NETWORK ?= mainnet
COVERAGE ?= coverage.txt
BUILDDIR ?= $(CURDIR)/build
LEDGER_ENABLED ?= true

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=crypto-com-chain \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=chain-maind \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'
TESTNET_FLAGS ?=

ledger ?= HID
ifeq ($(LEDGER_ENABLED),true)
	BUILD_TAGS := -tags cgo,ledger,!test_ledger_mock,!ledger_mock
	ifeq ($(ledger), ZEMU)
		BUILD_TAGS := $(BUILD_TAGS),ledger_zemu
	else
		BUILD_TAGS := $(BUILD_TAGS),!ledger_zemu
	endif
endif

ifeq ($(NETWORK),testnet)
	BUILD_TAGS := $(BUILD_TAGS),testnet
	TEST_TAGS := "--tags=testnet"
endif

SIMAPP = github.com/crypto-com/chain-main/app
BINDIR ?= ~/go/bin

OS := $(shell uname)

all: install

install: check-network go.sum
		go install -mod=readonly $(BUILD_FLAGS) $(BUILD_TAGS) ./cmd/chain-maind

build: check-network go.sum
		go build -mod=readonly $(BUILD_FLAGS) $(BUILD_TAGS) -o $(BUILDDIR)/chain-maind ./cmd/chain-maind
.PHONY: build

build-linux: go.sum
ifeq ($(OS), Linux)
		GOOS=linux GOARCH=amd64 $(MAKE) build
else
		LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build
endif

build-mac: go.sum
ifeq ($(OS), Darwin)
		GOOS=darwin GOARCH=amd64 $(MAKE) build
else
		LEDGER_ENABLED=false GOOS=darwin GOARCH=amd64 $(MAKE) build
endif

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

test: check-network
	@go test $(TEST_TAGS) -v -mod=readonly $(PACKAGES) -coverprofile=$(COVERAGE) -covermode=atomic
.PHONY: test

# look into .golangci.yml for enabling / disabling linters
lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify

lintpy:
	@flake8
	@black --check --diff .
	@isort -c --diff -rc .

test-sim-nondeterminism: check-network
	@echo "Running non-determinism test..."
	@go test $(TEST_TAGS) -mod=readonly $(SIMAPP) -run TestAppStateDeterminism -Enabled=true \
		-NumBlocks=100 -BlockSize=200 -Commit=true -Period=0 -v -timeout 24h

test-sim-custom-genesis-fast: check-network
	@echo "Running custom genesis simulation..."
	@echo "By default, ${HOME}/.chain-maind/config/genesis.json will be used."
	@go test $(TEST_TAGS) -mod=readonly $(SIMAPP) -run TestFullAppSimulation -Genesis=${HOME}/.gaiad/config/genesis.json \
		-Enabled=true -NumBlocks=100 -BlockSize=200 -Commit=true -Seed=99 -Period=5 -v -timeout 24h

test-sim-import-export:
	@echo "Running Chain import/export simulation. This may take several minutes..."
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) 25 5 TestAppImportExport

test-sim-after-import:
	@echo "Running application simulation-after-import. This may take several minutes..."
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(SIMAPP) 50 5 TestAppSimulationAfterImport

###############################################################################
###                                Localnet                                 ###
###############################################################################

build-docker-chainmaindnode:
	$(MAKE) -C check-networks/local

# Run a 4-node testnet locally
localnet-start: build-linux build-docker-chainmaindnode localnet-stop
	@if ! [ -f $(BUILDDIR)/node0/.chain-maind/config/genesis.json ]; \
	then docker run --rm -v $(BUILDDIR):/chain-maind:Z cryptocom/chainmaindnode testnet --v 4 -o . --starting-ip-address 192.168.10.2 $(TESTNET_FLAGS); \
	fi
	BUILDDIR=$(BUILDDIR) docker-compose up -d

# Stop testnet
localnet-stop:
	docker-compose down
	docker check-network prune -f

# local build pystarport
build-pystarport:
	pip install ./pystarport

# Run a local testnet by pystarport
localnet-pystartport: build-pystarport
	pystarport serve

clean:
	rm -rf $(BUILDDIR)/

clean-docker-compose: localnet-stop
	rm -rf $(BUILDDIR)/node* $(BUILDDIR)/gentxs

###############################################################################
###                                Nix                                      ###
###############################################################################
# nix installation: https://nixos.org/download.html
nix-integration-test: check-network
	nix-shell integration_tests/shell.nix --run "pytest -v -n 3 --dist loadscope"

nix-build-%: check-network check-os
	@if [ -e ~/.nix/remote-build-env ]; then \
		. ~/.nix/remote-build-env; \
	fi && \
	nix-build -o $* -A $* docker.nix;
	docker load < $*;

chaindImage: nix-build-chaindImage
pystarportImage: nix-build-pystarportImage

check-network:
ifeq ($(NETWORK),mainnet)	
else ifeq ($(NETWORK),testnet)
else
	@echo "Unrecognized network: ${NETWORK}"
endif
	@echo "building network: ${NETWORK}"

check-os:
ifeq ($(OS), Darwin)
ifneq ("$(wildcard ~/.nix/remote-build-env))","")
	@echo "installed nix-remote-builder before" && \
	docker run --restart always --name nix-docker -d -p 3022:22 lnl7/nix:ssh 2> /dev/null || echo "nix-docker is already running"
else
	@echo install nix-remote-builder
	git clone https://github.com/holidaycheck/nix-remote-builder.git
	cd nix-remote-builder && ./init.sh
	rm -rf nix-remote-builder
endif
endif
