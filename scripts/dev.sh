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

echo "- Waiting for postgres to be ready..."
RETRIES=10
until docker run -it --rm --link qrincard_db:pg postgres:13.3 psql "postgres://$DB_USER:$DB_PASS@pg:5432/$DB_NAME" -c "select 1" > /dev/null || [ $RETRIES -eq 0 ]; do
    echo "    $((RETRIES--)) remaining attempts..."
    sleep 1;
done

echo "- Seeding dev database"
docker exec -i maagizo_db psql -U postgres -d maagizo < ./pkg/http/rest/testapi/temp_seed.sql
echo ""

