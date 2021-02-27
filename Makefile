MIGRATION_DIR := migrations

# this requires you have exported the env variable DATABASE_URL
apply-migrations:
	migrate -source file://${MIGRATION_DIR} -database ${DATABASE_URL} up

# this requires you have exported the env variable DATABASE_URL
apply-x-migrations:
	migrate -source file://${MIGRATION_DIR} -database ${DATABASE_URL} up "$x"

# this requires you have exported the env variable DATABASE_URL
down-migrations:
	migrate -source file://${MIGRATION_DIR} -database ${DATABASE_URL} down

# this requires you have exported the env variable DATABASE_URL
down-x-migrations:
	migrate -source file://${MIGRATION_DIR} -database ${DATABASE_URL} down "$x"

# this requires you have exported the env variable DATABASE_URL
drop-database:
	migrate -source file://${MIGRATION_DIR} -database ${DATABASE_URL} drop

create-migration:
	migrate create -ext sql -dir ${MIGRATION_DIR} -seq "$name"

run:
	./go.sh
