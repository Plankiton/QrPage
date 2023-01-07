default: run

ensure:
	sh scripts/ensure.sh

up_db:
	sh scripts/up_db.sh

down_db:
	sh scripts/down_db.sh

dev: 
	sh scripts/dev.sh

run: ensure dev up_db
	sh scripts/run.sh

prod: ensure up_db

.PHONY: ensure build test lint dev swagger clean resetdb
