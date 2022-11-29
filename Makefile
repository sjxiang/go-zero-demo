


open-container:
	docker-compose up -d
	
close-container:
	docker-compose -f docker-compose.yml down

login_mysql:
	docker-compose -f ./docker-compose.yml exec db sh -c 'mysql -uroot -p${MYSQL_ROOT_PASSWORD}'

container_net:
	docker inspect db | grep IPAddress

