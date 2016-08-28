include Makefile.variable

print-%: ; @echo $*=$($*)

all: gipfeligw gipfelid

gipfeligw:
	docker build -t $(HUB_PREFIX)/$(GIPFELI_GATEWAY):$(GIPFELI_VERSION) ./cmd/gipfeligw

gipfelid:
	docker build -t $(HUB_PREFIX)/$(GIPFELI_DAEMON):$(DCE_GIPFELI_VERSIONVERSION) ./cmd/gipfelid

gipfeliauth:
	docker build -t $(HUB_PREFIX)/$(GIPFELI_AUTH):$(DCE_GIPFELI_VERSIONVERSION) ./cmd/gipfeliauth

release:
	@echo "developing..."
