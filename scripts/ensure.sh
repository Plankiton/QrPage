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

swag init -g api.go --dir pkg/api --parseDependency
