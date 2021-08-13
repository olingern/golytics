

createdb:
	sqlite3 ./db/analytics.db "VACUUM;"

migrate:
	dbmate migrate

up:
	DATABASE_URL="sqlite:db/database.sqlite3" dbmate up

down:
	DATABASE_URL="sqlite:db/database.sqlite3" dbmate down
	