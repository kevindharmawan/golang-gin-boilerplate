POSTGRES_ADMIN_DB="postgres"
POSTGRES_ADMIN_USER="user"
POSTGRES_ADMIN_PASS="password"

docker run -d \
    --name golang-postgres \
    -p 5432:5432 \
    -e POSTGRES_DB=$POSTGRES_ADMIN_DB \
    -e POSTGRES_USER=$POSTGRES_ADMIN_USER \
    -e POSTGRES_PASSWORD=$POSTGRES_ADMIN_PASS \
    postgres