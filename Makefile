default: dev

ensure:
	bash scripts/ensure.sh

up_db:
	bash scripts/up_db.sh

down_db:
	bash scripts/down_db.sh

dev: ensure up_db
	bash scripts/dev.sh

prod: ensure up_db

.PHONY: ensure build test lint dev swagger clean resetdb
