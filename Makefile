include Makefile.variable

print-%: ; @echo $*=$($*)

COMMANDS=gipfeligw gipfeliauth gipfelid gipfelis
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
$(APIS): FORCE
	@echo "🥐 $@"
	@docker run --rm -v `pwd`/$@:/protoc_dir $(HUB_PREFIX)/protoc-gateway-tool >> /dev/nul

generate: $(APIS)
	@echo "🥐 $@ finished"

release:
	@echo "developing..."
