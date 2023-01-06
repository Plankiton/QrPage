all: dev

ensure:
	bash scripts/ensure.sh

dev: ensure
	bash scripts/dev.sh
