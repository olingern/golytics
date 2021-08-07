

createdb:
	sqlite3 ./db/analytics.db "VACUUM;"

up:
	DATABASE_URL="sqlite:db/database.sqlite3" dbmate up

down:
	DATABASE_URL="sqlite:db/database.sqlite3" dbmate down	

run-client:
	npx http-server ./client
	