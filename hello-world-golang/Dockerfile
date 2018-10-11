# get and compile roadrunner
FROM golang:1.10.4 AS roadrunner
COPY application /application
WORKDIR /application
RUN go build .

# define how to start
ENTRYPOINT ["/application/application"]
