include Makefile.variable

print-%: ; @echo $*=$($*)

COMMANDS=gipfeligw gipfeliauth gipfelid gipfeliss
PROTOS=auth daemon

# Project binaries.
BINARIES=$(addprefix bin/,$(COMMANDS))
DISTS=$(addprefix dist/,$(COMMANDS))

# Project protobufs
APIS=$(addsuffix /api,$(PROTOS))

# Project images
IMAGES=$(addprefix sctlee/,$(COMMANDS))

all: images

build: clean
	@echo "🥐 $@"
	@docker build -t $(HUB_PREFIX)/$(GIPFELI_DIST) -f DockerfileBuild .
	@mkdir dist && docker run --rm -v `pwd`/dist:/dist $(HUB_PREFIX)/$(GIPFELI_DIST)

clean:
	@echo "🥐 $@"
	@rm -rf dist bin

FORCE:

# Build a binary from a cmd.
bin/%: cmd/% FORCE
	@echo "🥐 $@"
	@go build -i -tags "${DOCKER_BUILDTAGS}" -o $@ ${GO_LDFLAGS}  ${GO_GCFLAGS} ./$<

## build binaries
binaries: $(BINARIES)
	@echo "🥐 $@"

$(IMAGES): FORCE
	@echo "🥐 $@"
	@docker build --build-arg component=$(notdir $@) -t $@ .

## build images
images: build $(IMAGES)
	@echo "🥐 $@"

# generate protobuf
# setup_protobuf:
# 	@echo "🥐 $@"
# 	@protoc --proto_path=protobuf --go_out=plugins=grpc:protobuf  protobuf/plugin/*.proto

GPRC_PROTO_PATH=${HOME}/Codes/Tools/protoc-3/include
GRPC_GATEWAY_PROTO_PATH=${HOME}/Codes/GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

$(APIS): FORCE
	@echo "🥐 $@"
	@protoc --proto_path=$@ --proto_path=${GRPC_GATEWAY_PROTO_PATH} --proto_path=${GPRC_PROTO_PATH} \
	--go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:$@  $@/*.proto
	@protoc --proto_path=$@ --proto_path=${GRPC_GATEWAY_PROTO_PATH} --proto_path=${GPRC_PROTO_PATH} \
	--grpc-gateway_out=logtostderr=true:$@  $@/*.proto

generate: $(APIS)
	@echo "🥐 $@"

release:
	@echo "developing..."
