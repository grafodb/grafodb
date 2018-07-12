.PHONY: clean clean-all deps gazelle goimports golint govet package protos test update

ANTLR:=$(shell which antlr4)
BAZEL:=$(shell which bazel)
DEP:=$(shell which dep)
GAZELLE:=$(shell which gazelle)
GO:=$(shell which go)
GOIMPORTS:=$(shell which goimports)
GOLINT:=$(shell which golint)
FLATC:=$(shell which flatc)
FLATBUFFERS_SOURCES:=$(shell find . -type f -name '*.fbs')
FLATBUFFERS_FILES:=$(patsubst %.fbs,%_generated.go,$(FLATBUFFERS_SOURCES))
PROTOC:=$(shell which protoc)
PROTO_SOURCES:=$(shell find . -type f -name '*.proto')
PROTO_FILES:=$(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))

all: package

codegen: antlr flatbuffers protos

flatbuffers: $(FLATBUFFERS_FILES)

$(FLATBUFFERS_FILES): %_generated.go: %.fbs

%_generated.go:
	@echo Compiling $<
	@$(FLATC) --go --gen-onefile --go-namespace rpc -o $(dir $<) $<
	@$(GO) tool fix $(dir $<)

protos: $(PROTO_FILES)

$(PROTO_FILES): %.pb.go: %.proto

%.pb.go:
	@echo Compiling $<
	@$(PROTOC) --go_out=plugins=grpc:. $<

# TODO: Fix makefile rule
antlr:
	@$(ANTLR) -Dlanguage=Go -package grafoql internal/graph/grafoql/GrafoQL.g4
	@$(ANTLR) -Dlanguage=Go -package schema internal/graph/schema/Schema.g4

clean:
	@$(BAZEL) clean

clean-all:
	@$(BAZEL) clean --expunge

deps:
	$(DEP) ensure

goimports:
	@$(GOIMPORTS) -w $(shell find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./bazel*")

gazelle:
	@$(BAZEL) run //:gazelle

golint:
	@$(GOLINT) -set_exit_status $(shell go list ./... | grep -v /vendor/)

govet:
	@$(GO) vet $(shell go list ./... | grep -v /vendor/)

package: update
	@$(BAZEL) run --cpu=k8 //:package

test: codegen
	@$(GO) test $(shell go list ./... | grep -v /vendor/)

update: codegen goimports gazelle


# get docker image targets
# bazel query 'filter("docker-image", kind(app_layer, //...))'
# bazel query 'kind(".*_layer", //...)'
# bazel run --cpu=k8 $(bazel query 'kind("app_layer", //...)')