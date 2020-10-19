MAINTAINER=fabiotavarespr
PROJECT=api-backend
VERSION=1.0
DATABASE-IMAGE=mysql:8.0.21

all: build

build:
	@docker build --tag=${MAINTAINER}/${PROJECT}:${VERSION} .
	@docker build --tag=${MAINTAINER}/${PROJECT}:latest .

push:
	@docker push ${MAINTAINER}/${PROJECT}:${VERSION}
	@docker push ${MAINTAINER}/${PROJECT}:latest

database-up:
	@docker run --rm --name db-backend -d -p 3306:3306 -v ${PWD}/.sql-scripts/:/docker-entrypoint-initdb.d/ -e MYSQL_DATABASE=backend -e MYSQL_ROOT_PASSWORD=passwd123 ${DATABASE-IMAGE}

database-down:
	@docker stop db-backend

docker-up: build
	@docker-compose up -d

docker-down:
	@docker-compose down

docker-log:
	@docker-compose logs -f