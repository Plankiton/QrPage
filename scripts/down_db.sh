DB_USER="${DB_USER-hubla}"
DB_NAME="${DB_NAME-hubla}"
DB_PASS="${DB_PASS-supersecret}"
DATABASE_URL="${DATABASE_URL-postgres://$DB_USER:$DB_PASS@0.0.0.0/$DB_NAME?sslmode=disable}"

migrate -path ./pkg/db/sql/ -database "$DATABASE_URL" down
