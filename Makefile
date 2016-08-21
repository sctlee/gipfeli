include Makefile.variable

print-%: ; @echo $*=$($*)

all: gipfeligw gipfelid

gipfeligw:
	docker build -t $(HUB_PREFIX)/$(GIPFELI_GATEWAY):$(GIPFELI_VERSION)

gipfelid:
	docker build -t $(HUB_PREFIX)/$(GIPFELI_DAEMON):$(DCE_GIPFELI_VERSIONVERSION)

release:
	@echo "developing..."
