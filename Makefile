.PHONY: migrate-up migrate-down migrate-force

migrate-up:
	@echo "Applying all up migrations"
	go run cmd/migrate/main.go -up

migrate-down:
	@echo "Applying all down migrations"
	go run cmd/migrate/main.go -down

migrate-force:
	@echo "Force setting version"
	@go run cmd/migrate/main.go -force $(filter-out $@,$(MAKECMDGOALS))

%:
	@:
