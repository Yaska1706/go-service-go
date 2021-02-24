MIGRATION_DIR := migrations

migration:
	migrate create -ext sql -dir ${MIGRATION_DIR} -seq "$t"
