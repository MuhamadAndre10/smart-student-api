DSN=postgres://postgres:root@localhost:5432/users?sslmode=disable

migrate_up:
	@echo "Migrating database"
	migrate -path db/schema -database ${DSN} -verbose up
	@echo "Done..."

migrate_down:
	@echo "Migrating database"
	migrate -path db/schema -database ${DSN} -verbose down
	@echo "Done..."