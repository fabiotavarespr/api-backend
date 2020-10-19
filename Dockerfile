FROM golang:1.15 as BUILDER
# create a working directory
WORKDIR /go/src/app
# add source code

COPY . .

ENV GO115MODULE "on" 
RUN go mod vendor

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/api-backend

ENV API_LOG_LEVEL "info"
ENV API_SERVER_HOSTNAME "server-hostname"
ENV API_SERVER_PORT "8080"
ENV API_DATABASE_HOSTNAME "database-hostname"
ENV API_DATABASE_PORT "3306"
ENV API_DATABASE_USERNAME "database-username"
ENV API_DATABASE_PASSWORD "database-password"
ENV API_DATABASE_NAME "database-name"

USER root

FROM alpine

COPY startup.sh /
RUN chmod -R 777 /startup.sh

COPY --from=BUILDER /go/bin/api-backend /bin/api-backend

# run compiled go app
ENTRYPOINT [ "/bin/sh" ]
CMD [ "/startup.sh" ]