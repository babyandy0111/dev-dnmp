mysql:
	docker-compose -f docker-compose.yml up -d mysql

php:
	docker-compose -f docker-php-compose.yml up -d php

nginx:
	docker-compose -f docker-php-compose.yml up -d nginx

redis:
	docker-compose -f docker-compose.yml up -d redis

pgsql:
	docker-compose -f docker-pgsql-compose.yml up -d pgsql pgadmin

es:
	docker-compose -f docker-elk-compose.yml up -d elasticsearch