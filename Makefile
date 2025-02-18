migrate:
	docker compose run --rm api go run . migrate up

# make migration name=<some new migration name>;
migration:
	docker compose run --rm api go run . migrate create -n $(name)
