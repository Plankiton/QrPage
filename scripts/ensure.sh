if ! command -v migrate &> /dev/null
then
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.1
fi

go install -tags 'swag' github.com/swaggo/swag/cmd/swag@v1.7.8
# go install github.com/vektra/mockery/v2@v2.10.0

if ! command -v air &> /dev/null
then
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
fi

swag init -g api.go --dir pkg/api

DB_USER="${DB_USER-qrincard}"
DB_NAME="${DB_NAME-qrincard}"
DB_PASS="${DB_PASS-supersecret}"

if ! docker container inspect qrincard_db > /dev/null 2>&1
then
    echo "- Starting a postgres on port 5432..."
    docker run --name qrincard_db -d -p 5432:5432 \
        -e POSTGRES_USER="$DB_USER" \
        -e POSTGRES_PASSWORD="$DB_PASS" \
        -e POSTGRES_DB="$DB_NAME" \
        postgres:13.3
    echo ""
fi

docker start qrincard_db > /dev/null

