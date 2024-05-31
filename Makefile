

migrations_status:
	~/go/bin/goose -dir migrations postgres "user=finance password=ecnanif dbname=finance sslmode=disable host=localhost" status

migrations_up:
	~/go/bin/goose -dir migrations postgres "user=finance password=ecnanif dbname=finance sslmode=disable host=localhost" up

migrations_down:
	~/go/bin/goose -dir migrations postgres "user=finance password=ecnanif dbname=finance sslmode=disable host=localhost" down