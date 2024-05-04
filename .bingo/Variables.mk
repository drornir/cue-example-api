# Auto generated binary variables helper managed by https://github.com/bwplotka/bingo v0.9. DO NOT EDIT.
# All tools are designed to be build inside $GOBIN.
BINGO_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
GOPATH ?= $(shell go env GOPATH)
GOBIN  ?= $(firstword $(subst :, ,${GOPATH}))/bin
GO     ?= $(shell which go)

# Below generated variables ensure that every time a tool under each variable is invoked, the correct version
# will be used; reinstalling only if needed.
# For example for bingo variable:
#
# In your main Makefile (for non array binaries):
#
#include .bingo/Variables.mk # Assuming -dir was set to .bingo .
#
#command: $(BINGO)
#	@echo "Running bingo"
#	@$(BINGO) <flags/args..>
#
BINGO := $(GOBIN)/bingo-v0.9.0
$(BINGO): $(BINGO_DIR)/bingo.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/bingo-v0.9.0"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=bingo.mod -o=$(GOBIN)/bingo-v0.9.0 "github.com/bwplotka/bingo"

CUE := $(GOBIN)/cue-v0.8.2
$(CUE): $(BINGO_DIR)/cue.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/cue-v0.8.2"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=cue.mod -o=$(GOBIN)/cue-v0.8.2 "cuelang.org/go/cmd/cue"

GOMODIFYTAGS := $(GOBIN)/gomodifytags-v1.16.0
$(GOMODIFYTAGS): $(BINGO_DIR)/gomodifytags.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/gomodifytags-v1.16.0"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=gomodifytags.mod -o=$(GOBIN)/gomodifytags-v1.16.0 "github.com/fatih/gomodifytags"

GUM := $(GOBIN)/gum-v0.14.0
$(GUM): $(BINGO_DIR)/gum.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/gum-v0.14.0"
	@cd $(BINGO_DIR) && GOWORK=off $(GO) build -mod=mod -modfile=gum.mod -o=$(GOBIN)/gum-v0.14.0 "github.com/charmbracelet/gum"

