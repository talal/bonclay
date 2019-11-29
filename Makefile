PREFIX  := /usr/local
PKG      = github.com/talal/bonclay
VERSION := $(shell util/find_version.sh)

GO          := GOBIN=$(CURDIR)/build go
BUILD_FLAGS :=
LD_FLAGS    := -s -w -X main.version=$(VERSION)

ifndef GOOS
	GOOS := $(word 1, $(subst /, " ", $(word 4, $(shell go version))))
endif

BINARY64  := bonclay-$(GOOS)_amd64
RELEASE64 := bonclay-$(VERSION)-$(GOOS)_amd64

################################################################################

all: build/bonclay

# This target uses the incremental rebuild capabilities of the Go compiler to speed things up.
# If no source files have changed, `go install` exits quickly without doing anything.
build/bonclay: FORCE
	$(GO) install $(BUILD_FLAGS) -ldflags '$(LD_FLAGS)' '$(PKG)'

install: FORCE all
	install -D build/bonclay "$(DESTDIR)$(PREFIX)/bin/bonclay"

ifeq ($(GOOS),windows)
release: FORCE release/$(BINARY64)
	cd release && cp -f $(BINARY64) bonclay.exe && zip $(RELEASE64).zip bonclay.exe
	cd release && rm -f bonclay.exe
else
release: FORCE release/$(BINARY64)
	cd release && cp -f $(BINARY64) bonclay && tar -czf $(RELEASE64).tar.gz bonclay
	cd release && rm -f bonclay
endif

release-all: FORCE clean
	GOOS=darwin make release
	GOOS=linux  make release

release/$(BINARY64): FORCE
	GOARCH=amd64 $(GO) build $(BUILD_FLAGS) -o $@ -ldflags '$(LD_FLAGS)' '$(PKG)'

clean: FORCE
	rm -rf -- build release

.PHONY: FORCE
