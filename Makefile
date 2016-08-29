include Makefile.variable

print-%: ; @echo $*=$($*)

COMMANDS=gipfeligw gipfeliauth
# Project binaries.
BINARIES=$(addprefix bin/,$(COMMANDS))

# Project images
IMAGES=$(addprefix sctlee/,$(COMMANDS))

all: build images

dist:
	@echo "🐳 $@"
	@docker build -t $(HUB_PREFIX)/$(GIPFELI_DIST) -f DockerfileBuild .

build: clean dist
	@echo "🐳 $@"
	@docker run --name $(GIPFELI_DIST) $(HUB_PREFIX)/$(GIPFELI_DIST)
	@docker cp $(GIPFELI_DIST):/dist dist

clean:
	@echo "🐳 $@ ${BINARIES}"
	@docker rm -vf $(GIPFELI_DIST)

FORCE:

# Build a binary from a cmd.
bin/%: cmd/% FORCE
	@echo "🐳 $@"
	@go build -i -tags "${DOCKER_BUILDTAGS}" -o $@ ${GO_LDFLAGS}  ${GO_GCFLAGS} ./$<

binaries: $(BINARIES) ## build binaries
	@echo "🐳 $@"

sctlee/%: dist/% FORCE
	@echo "🐳 $@"
	@docker build -t $@ .

images: $(IMAGES) ## build images
	@echo "🐳 $@"

release:
	@echo "developing..."
