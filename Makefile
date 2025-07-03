# Строка подключения к БД
DB_URL=postgres://postgres:postgres@localhost:5432/newsdb?sslmode=disable

# Создание новой миграции 
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение всех миграций
migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" up

# Откат всех миграций
migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" down
