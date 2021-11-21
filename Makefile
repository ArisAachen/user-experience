PREFIX=/usr
PREFIXVAR=/var
PREFIXLIB=/lib
pwd := $(shell pwd)
GoPath := GOPATH=$(pwd):$(GOPATH)
GOBUILD = go build $(GO_BUILD_FLAGS)
export GO111MODULE=off

all: prepare build

build:
	$(GoPath) $(GOBUILD) -o DSServices ./

prepare:
	mkdir -p $(GOPATH)/src/github.com/ArisAachen
	rm $(GOPATH)/src/github.com/ArisAachen/experience
	ln -s $(pwd) $(GOPATH)/src/github.com/ArisAachen/experience	

test-coverage:
	env $(GoPath) go test -cover -v ./src/... | awk '$$1 ~ "^(ok|\\?)" {print $$2","$$5}' | sed "s:${CURDIR}::g" | sed 's/files\]/0\.0%/g' > coverage.csv

install:
	
	install -v -D -m +x -t $(DESTDIR)$(PREFIX)/bin DSServices
	install -v -D -m644 -t $(DESTDIR)$(PREFIX)/share/dbus-1/system.d misc/com.deepin.userexperience.Daemon.conf
	install -v -D -m644 -t $(DESTDIR)$(PREFIXLIB)/systemd/system misc/lib/systemd/system/com.deepin.userexperience.Daemon.service
test:
	$(GoPath) go test daemon

print_gopath:
	GOPATH="${CURDIR}/${GOPATH_DIR}:${GOPATH}"

clean:
	-rm -rf bin

.PHONY: all build install clean
