include Makefile.variable

print-%: ; @echo $*=$($*)

COMMANDS=gipfeligw gipfeliauth
# Project binaries.
BINARIES=$(addprefix bin/,$(COMMANDS))
DISTS=$(addprefix dist/,$(COMMANDS))

# Project images
IMAGES=$(addprefix sctlee/,$(COMMANDS))

all: build images

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

binaries: $(BINARIES) ## build binaries
	@echo "🐳 $@"

$(IMAGES): FORCE
	@echo "🐳 $@"
	@docker build --build-arg component=$(notdir $@) -t $@ .

images: $(IMAGES) ## build images
	@echo "🐳 $@"

release:
	@echo "developing..."
