# Interactive release Makefile
# -----------------------------------------------
# Usage:
#   make release TAG=v1.1.0        # non-interactive tag
#   make release                   # prompts for tag
# -----------------------------------------------

# ---- configuration ----------------------------------------------------------
IMAGE ?= mayckol/brfiscalfaker

ifndef TAG
# Prompt once for TAG if it wasn't supplied on the CLI
TAG := $(shell read -p "Enter release tag (e.g. v1.1.0): " t; echo $${t})
endif

# Die early if user just pressed <Enter>
ifeq ($(strip $(TAG)),)
$(error TAG is required – rerun with TAG=vX.Y.Z)
endif

# ---- helper macro -----------------------------------------------------------
define CONFIRM_RUN
	@cmd='$(1)'; \
	echo "➜ $$cmd"; \
	read -p "Run this command? [y/N] " ans; \
	if [ "$$ans" = "y" ] || [ "$$ans" = "Y" ]; then \
		eval $$cmd; \
	else \
		echo "✗ skipped"; \
	fi
endef

# ---- targets ---------------------------------------------------------------
.PHONY: release tag build push push-latest

release: tag build push push-latest
	@echo "✔ Release $(TAG) complete."

tag:
	$(call CONFIRM_RUN,git tag -a $(TAG) -m "Release version $(TAG)")
	$(call CONFIRM_RUN,git push origin $(TAG))

build:
	$(call CONFIRM_RUN,docker build -t $(IMAGE):$(TAG) .)

push:
	$(call CONFIRM_RUN,docker push $(IMAGE):$(TAG))

push-latest:
	$(call CONFIRM_RUN,docker tag $(IMAGE):$(TAG) $(IMAGE):latest)
	$(call CONFIRM_RUN,docker push $(IMAGE):latest)
