ifneq (,$(filter $(OS),Windows_NT Windows))
	EXEEXT=.exe
endif

next_version := $(shell cat build_version.txt)
tag := $(shell git describe --exact-match --tags 2>/dev/null)

branch := $(shell git rev-parse --abbrev-ref HEAD)
commit := $(shell git rev-parse --short=8 HEAD)

ifdef NIGHTLY
	version := $(next_version)
	rpm_version := nightly
	rpm_iteration := 0
	deb_version := nightly
	deb_iteration := 0
	tar_version := nightly
else ifeq ($(tag),)
	version := $(next_version)
	rpm_version := $(version)~$(commit)-0
	rpm_iteration := 0
	deb_version := $(version)~$(commit)-0
	deb_iteration := 0
	tar_version := $(version)~$(commit)
else ifneq ($(findstring -rc,$(tag)),)
	version := $(word 1,$(subst -, ,$(tag)))
	version := $(version:v%=%)
	rc := $(word 2,$(subst -, ,$(tag)))
	rpm_version := $(version)-0.$(rc)
	rpm_iteration := 0.$(subst rc,,$(rc))
	deb_version := $(version)~$(rc)-1
	deb_iteration := 0
	tar_version := $(version)~$(rc)
else
	version := $(tag:v%=%)
	rpm_version := $(version)-1
	rpm_iteration := 1
	deb_version := $(version)-1
	deb_iteration := 1
	tar_version := $(version)
endif

MAKEFLAGS += --no-print-directory
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

pkgdir ?= build/dist
homedir ?= ${HOME}/.opsAgent


.PHONY: deps
deps:
	go mod download -x

.PHONY: version
version:
	@echo $(version)-$(commit)

.PHONY: build
build:
	go build -tags "$(BUILDTAGS)" ./cmd/opsAgent

.PHONY: opsAgent
opsAgent: build

# Used by dockerfile builds
.PHONY: go-install
go-install:
	go install -mod=mod ./cmd/opsAgent


.PHONY: tidy
tidy:
	go mod verify
	go mod tidy
	@if ! git diff --quiet go.mod go.sum; then \
		echo "please run go mod tidy and check in changes, you might have to use the Go version 1.19+"; \
		exit 1; \
	fi

.PHONY: install
install: $(buildbin)
	@mkdir -pv $(homedir)
	@mkdir -pv $(DESTDIR)/conf
	@if [ $(GOOS) = "linux" ]; then cp -fv conf/opsAgent.conf $(homedir)/.conf; fi
	@if [ $(GOOS) = "linux" ]; then cp -fv conf/opsAgent.conf $(DESTDIR)/conf/opsAgent.conf.sample; fi
	@cp -fv $(buildbin) $(DESTDIR)$(bindir)
	@if [ $(GOOS) = "linux" ]; then mkdir -pv $(DESTDIR)/lib/scripts; fi
	@if [ $(GOOS) = "linux" ]; then cp -af scripts/opsAgent.service $(DESTDIR)/lib/scripts; fi
	@if [ $(GOOS) = "linux" ]; then cp -af scripts/*.sh $(DESTDIR)/lib/scripts/; fi

$(buildbin):
	echo $(GOOS)
	@mkdir -pv $(dir $@)
	go build -o $(dir $@) ./cmd/opsAgent

.PHONY: clean
clean:
	rm -f opsAgent
	rm -f opsAgent.exe
	rm -rf build

# Define packages opsAgent supports, organized by architecture with a rule to echo the list to limit include_packages
# e.g. make package include_packages="$(make amd64)"
mips += linux_mips.tar.gz mips.deb
.PHONY: mips
mips:
	@ echo $(mips)
mipsel += mipsel.deb linux_mipsel.tar.gz
.PHONY: mipsel
mipsel:
	@ echo $(mipsel)
arm64 += linux_arm64.tar.gz arm64.deb aarch64.rpm
.PHONY: arm64
arm64:
	@ echo $(arm64)
amd64 += freebsd_amd64.tar.gz linux_amd64.tar.gz amd64.deb x86_64.rpm
.PHONY: amd64
amd64:
	@ echo $(amd64)
static += static_linux_amd64.tar.gz
.PHONY: static
static:
	@ echo $(static)
armel += linux_armel.tar.gz armel.rpm armel.deb
.PHONY: armel
armel:
	@ echo $(armel)
armhf += linux_armhf.tar.gz freebsd_armv7.tar.gz armhf.deb armv6hl.rpm
.PHONY: armhf
armhf:
	@ echo $(armhf)
s390x += linux_s390x.tar.gz s390x.deb s390x.rpm
.PHONY: riscv64
riscv64:
	@ echo $(riscv64)
riscv64 += linux_riscv64.tar.gz riscv64.rpm riscv64.deb
.PHONY: s390x
s390x:
	@ echo $(s390x)
ppc64le += linux_ppc64le.tar.gz ppc64le.rpm ppc64el.deb
.PHONY: ppc64le
ppc64le:
	@ echo $(ppc64le)
i386 += freebsd_i386.tar.gz i386.deb linux_i386.tar.gz i386.rpm
.PHONY: i386
i386:
	@ echo $(i386)
windows += windows_i386.zip windows_amd64.zip
.PHONY: windows
windows:
	@ echo $(windows)
darwin-amd64 += darwin_amd64.tar.gz
.PHONY: darwin-amd64
darwin-amd64:
	@ echo $(darwin-amd64)

darwin-arm64 += darwin_arm64.tar.gz
.PHONY: darwin-arm64
darwin-arm64:
	@ echo $(darwin-arm64)

include_packages := $(mips) $(mipsel) $(arm64) $(amd64) $(static) $(armel) $(armhf) $(riscv64) $(s390x) $(ppc64le) $(i386) $(windows) $(darwin-amd64) $(darwin-arm64)

.PHONY: $(include_packages)
$(include_packages):
	@$(MAKE) install
	@mkdir -p $(pkgdir)

	@if [ "$(suffix $@)" = ".rpm" ]; then \
		fpm --force \
			--log info \
			--architecture $(basename $@) \
			--input-type dir \
			--output-type rpm \
			--vendor arctic \
			--url https://github.com/ArcticClint/opsAgent \
			--maintainer arctic \
			--config-files conf/opsAgent.conf \
			--after-install scripts/rpm/post-install.sh \
			--before-install scripts/rpm/pre-install.sh \
			--after-remove scripts/rpm/post-remove.sh \
			--description "opsAgent." \
			--depends coreutils \
			--depends shadow-utils \
			--rpm-digest sha256 \
			--rpm-posttrans scripts/rpm/post-install.sh \
			--name opsAgent \
			--version $(version) \
			--iteration $(rpm_iteration) \
			--chdir $(DESTDIR) \
			--package $(pkgdir)/opsAgent-$(rpm_version).$@ \
			--prefix /opt/opsAgent; \
	elif [ "$(suffix $@)" = ".deb" ]; then \
		fpm --force \
			--log info \
			--architecture $(basename $@) \
			--input-type dir \
			--output-type deb \
			--vendor arctic \
			--url https://github.com/ArcticClint/opsAgent \
			--maintainer arctic \
			--config-files /etc/opsAgent/opsAgent.conf.sample \
			--description "opsAgent." \
			--name opsAgent \
			--version $(version) \
			--iteration $(deb_iteration) \
			--chdir $(DESTDIR) \
			--package $(pkgdir)/opsAgent_$(deb_version)_$@	;\
	elif [ "$(suffix $@)" = ".zip" ]; then \
		(cd $(dir $(DESTDIR)) && zip -r - ./*) > $(pkgdir)/opsAgent-$(tar_version)_$@ ;\
	elif [ "$(suffix $@)" = ".gz" ]; then \
		tar --owner 0 --group 0 -czvf $(pkgdir)/opsAgent-$(tar_version)_$@ -C $(dir $(DESTDIR)) . ;\
	fi

amd64.deb x86_64.rpm linux_amd64.tar.gz: export GOOS := linux
amd64.deb x86_64.rpm linux_amd64.tar.gz: export GOARCH := amd64

static_linux_amd64.tar.gz: export cgo := -nocgo
static_linux_amd64.tar.gz: export CGO_ENABLED := 0

i386.deb i386.rpm linux_i386.tar.gz: export GOOS := linux
i386.deb i386.rpm linux_i386.tar.gz: export GOARCH := 386

armel.deb armel.rpm linux_armel.tar.gz: export GOOS := linux
armel.deb armel.rpm linux_armel.tar.gz: export GOARCH := arm
armel.deb armel.rpm linux_armel.tar.gz: export GOARM := 5

armhf.deb armv6hl.rpm linux_armhf.tar.gz: export GOOS := linux
armhf.deb armv6hl.rpm linux_armhf.tar.gz: export GOARCH := arm
armhf.deb armv6hl.rpm linux_armhf.tar.gz: export GOARM := 6

arm64.deb aarch64.rpm linux_arm64.tar.gz: export GOOS := linux
arm64.deb aarch64.rpm linux_arm64.tar.gz: export GOARCH := arm64
arm64.deb aarch64.rpm linux_arm64.tar.gz: export GOARM := 7

mips.deb linux_mips.tar.gz: export GOOS := linux
mips.deb linux_mips.tar.gz: export GOARCH := mips

mipsel.deb linux_mipsel.tar.gz: export GOOS := linux
mipsel.deb linux_mipsel.tar.gz: export GOARCH := mipsle

riscv64.deb riscv64.rpm linux_riscv64.tar.gz: export GOOS := linux
riscv64.deb riscv64.rpm linux_riscv64.tar.gz: export GOARCH := riscv64

s390x.deb s390x.rpm linux_s390x.tar.gz: export GOOS := linux
s390x.deb s390x.rpm linux_s390x.tar.gz: export GOARCH := s390x

ppc64el.deb ppc64le.rpm linux_ppc64le.tar.gz: export GOOS := linux
ppc64el.deb ppc64le.rpm linux_ppc64le.tar.gz: export GOARCH := ppc64le

freebsd_amd64.tar.gz: export GOOS := freebsd
freebsd_amd64.tar.gz: export GOARCH := amd64

freebsd_i386.tar.gz: export GOOS := freebsd
freebsd_i386.tar.gz: export GOARCH := 386

freebsd_armv7.tar.gz: export GOOS := freebsd
freebsd_armv7.tar.gz: export GOARCH := arm
freebsd_armv7.tar.gz: export GOARM := 7

windows_amd64.zip: export GOOS := windows
windows_amd64.zip: export GOARCH := amd64

darwin_amd64.tar.gz: export GOOS := darwin
darwin_amd64.tar.gz: export GOARCH := amd64

darwin_arm64.tar.gz: export GOOS := darwin
darwin_arm64.tar.gz: export GOARCH := arm64

windows_i386.zip: export GOOS := windows
windows_i386.zip: export GOARCH := 386

windows_i386.zip windows_amd64.zip: export prefix =
windows_i386.zip windows_amd64.zip: export bindir = $(prefix)
windows_i386.zip windows_amd64.zip: export sysconfdir = $(prefix)
windows_i386.zip windows_amd64.zip: export localstatedir = $(prefix)
windows_i386.zip windows_amd64.zip: export EXEEXT := .exe

%.deb: export pkg := deb
%.deb: export prefix := /usr
%.deb: export conf_suffix := .sample
%.deb: export sysconfdir := /etc
%.deb: export localstatedir := /var
%.rpm: export pkg := rpm
%.rpm: export prefix := /usr
%.rpm: export sysconfdir := /etc
%.rpm: export localstatedir := /var
%.tar.gz: export pkg := tar
%.tar.gz: export prefix := /usr
%.tar.gz: export sysconfdir := /etc
%.tar.gz: export localstatedir := /var
%.zip: export pkg := zip
%.zip: export prefix := /

%.deb %.rpm %.tar.gz %.zip: export DESTDIR = build/$(GOOS)-$(GOARCH)$(GOARM)$(cgo)-$(pkg)/opsAgent-$(version)
%.deb %.rpm %.tar.gz %.zip: export buildbin = build/$(GOOS)-$(GOARCH)$(GOARM)$(cgo)/opsAgent$(EXEEXT)
