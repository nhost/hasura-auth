ifdef VER
VERSION=$(shell echo $(VER) | sed -e 's/^v//g' -e 's/\//_/g')
else
VERSION=$(shell grep -oP 'version\s*=\s*"\K[^"]+' flake.nix | head -n 1)
endif

ifeq ($(shell uname -m),x86_64)
  ARCH?=x86_64
else ifeq ($(shell uname -m),arm64)
  ARCH?=aarch64
else ifeq ($(shell uname -m),aarch64)
  ARCH?=aarch64
else
  ARCH?=FIXME-$(shell uname -m)
endif

ifeq ($(shell uname -o),Darwin)
  OS?=darwin
else
  OS?=linux
endif


.PHONY: help
help: ## Show this help.
	@IFS=$$'\n' ; \
	lines=(`fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//'`); \
	for line in $${lines[@]}; do \
		IFS=$$'#' ; \
		split=($$line) ; \
		command=`echo $${split[0]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		info=`echo $${split[2]} | sed -e 's/^ *//' -e 's/ *$$//'` ; \
		printf "%-38s %s\n" $$command $$info ; \
	done


.PHONY: get-version
get-version:  ## Return version
	@sed -i '/^\s*version = "0.0.0-dev";/s//version = "${VERSION}";/' flake.nix
	@sed -i '/^\s*created = "1970-.*";/s//created = "${shell date --utc '+%Y-%m-%dT%H:%M:%SZ'}";/' flake.nix
	@echo $(VERSION)


.PHONY: check
check:   ## Run nix flake check
	./build/nix.sh flake check --print-build-logs


.PHONY: check-dry-run-node
check-dry-run-node:  ## Returns the derivation of the check
	@nix build \
		--dry-run \
		--json \
		--print-build-logs \
		.\#checks.$(ARCH)-$(OS).node-checks | jq -r '.[].outputs.out'


.PHONY: check-dry-run-go
check-dry-run-go:  ## Run nix flake check
	@nix build \
		--dry-run \
		--json \
		--print-build-logs \
		.\#checks.$(ARCH)-$(OS).go-checks | jq -r '.[].outputs.out'


.PHONY: dev-env-up
dev-env-up: build-docker-image  ## Starts development environment
	cd build/dev/docker && docker compose \
		--project-name auth-dev \
		up \
		--wait --wait-timeout 120 \


.PHONY: dev-env-up-short
dev-env-up-short:  ## Starts development environment without ai service
	cd build/dev/docker && docker compose \
		--project-name auth-dev \
		up \
		--wait --wait-timeout 120 \
		postgres graphql mailhog memcached


.PHONY: dev-env-down
dev-env-down:  ## Stops development environment
	cd build/dev/docker && docker compose --project-name auth-dev down --volumes


.PHONY: build
build:  ## Build application and places the binary under ./result/bin
	./build/nix.sh build --print-build-logs


.PHONY: build-dry-run
build-dry-run:  ## Run nix flake check
	@nix path-info \
		--derivation \
		.\#packages.$(ARCH)-$(OS).default


.PHONY: build-docker-image
build-docker-image:  ## Build docker container for native architecture
	./build/nix-docker-image.sh
	skopeo copy --insecure-policy dir:./result docker-daemon:hasura-auth:$(VERSION)
	docker tag hasura-auth:$(VERSION) nhost/hasura-auth:$(VERSION)


.PHONY: migrations-add
migrations-add:  ## add a migration with NAME in the migrations folder
	migrate create -dir ./migrations/postgres -ext sql -seq $(MIGRATION_NAME)
