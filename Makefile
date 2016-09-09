include Makefile.variable

print-%: ; @echo $*=$($*)

COMMANDS=gipfeligw gipfeliauth
PROTOS=auth daemon
# Project binaries.
BINARIES=$(addprefix bin/,$(COMMANDS))
DISTS=$(addprefix dist/,$(COMMANDS))
APIS=$(addsuffix /api,$(PROTOS))

# Project images
IMAGES=$(addprefix sctlee/,$(COMMANDS))

all: images

build: clean
	@echo "🐳 $@"
	@docker build -t $(HUB_PREFIX)/$(GIPFELI_DIST) -f DockerfileBuild .
	@mkdir dist && docker run --rm -v `pwd`/dist:/dist $(HUB_PREFIX)/$(GIPFELI_DIST)

clean:
	@echo "🐳 $@"
	@rm -rf dist bin

FORCE:

# Build a binary from a cmd.
bin/%: cmd/% FORCE
	@echo "🐳 $@"
	@go build -i -tags "${DOCKER_BUILDTAGS}" -o $@ ${GO_LDFLAGS}  ${GO_GCFLAGS} ./$<

## build binaries
binaries: $(BINARIES)
	@echo "🐳 $@"

$(IMAGES): FORCE
	@echo "🐳 $@"
	@docker build --build-arg component=$(notdir $@) -t $@ .

## build images
images: build $(IMAGES)
	@echo "🐳 $@"

$(APIS): FORCE
	@echo "🐳 $@"
	@protoc --proto_path=$@ --go_out=plugins=grpc:$@  $@/*.proto

generate: $(APIS)
	@echo "🐳 $@"

release:
	@echo "developing..."
