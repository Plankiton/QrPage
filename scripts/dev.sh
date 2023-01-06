if ! command -v docker &> /dev/null
then
    echo "MISSING DEPENDENCY: docker is not installed"
    exit 1
fi

if ! command -v go &> /dev/null
then
    echo "MISSING DEPENDENCY: go is not installed"
    exit 1
fi

echo "- Waiting for postgres to be ready..."
RETRIES=10
until docker run -it --rm --link qrincard_db:pg postgres:13.3 psql "postgres://$DB_USER:$DB_PASS@pg:5432/$DB_NAME" -c "select 1" > /dev/null || [ $RETRIES -eq 0 ]; do
    echo "    $((RETRIES--)) remaining attempts..."
    sleep 1;
done

# echo "- Seeding dev database"
# docker exec -i maagizo_db psql -U postgres -d maagizo < ./pkg/http/rest/testapi/temp_seed.sql
# echo ""

if ! air -h > /dev/null 2>&1
then
  go build -o ./bin/server ./cmd/server && ./bin/server
else
  air -build.cmd "go build -o ./bin/server ./cmd/server" -build.bin "./bin/server"
fi
