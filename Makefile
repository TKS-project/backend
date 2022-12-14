up:
	docker-compose up
connect:
	docker-compose exec db bash -c 'mysql -u root -pecc test'

down:
	docker-compose down