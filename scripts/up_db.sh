DB_USER="${DB_USER-qrincard}"
DB_NAME="${DB_NAME-qrincard}"
DB_PASS="${DB_PASS-supersecret}"

DEFAULT_DATABASE="postgres://$DB_USER:$DB_PASS@0.0.0.0/$DB_NAME?sslmode=disable"
DATABASE_URL="${DATABASE_URL-$DEFAULT_DATABASE}"

migrate -path ./pkg/db/sql/ -database "$DATABASE_URL" up
