# get and compile roadrunner
FROM golang:1.10.4 AS roadrunner
RUN go get -d -v github.com/spiral/roadrunner/cmd/rr
RUN go install github.com/spiral/roadrunner/cmd/rr
RUN CGO_ENABLED=0 go install -ldflags "$LDFLAGS -extldflags '-static'" github.com/spiral/roadrunner/cmd/rr

# php 7
FROM php:7.2.10-alpine3.8

# composer
#RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');"
#RUN php -r "if (hash_file('SHA384', 'composer-setup.php') === '93b54496392c062774670ac18b134c3b3a95e5a5e5c8f1a9f115f203b75bf9a129d5daa8ba6a13e2cc8a1da0806388a8') { echo 'Installer verified'; } else { echo 'Installer corrupt'; unlink('composer-setup.php'); } echo PHP_EOL;"
#RUN php composer-setup.php
#RUN php -r "unlink('composer-setup.php');"
#RUN mv composer.phar /usr/local/bin/composer

# get compiled roadrunner
COPY --from=roadrunner /go/bin/rr /usr/local/bin/rr

# define how to start
ENTRYPOINT ["rr"]
CMD ["serve", "--Verbose" ,"--debug"]

# add application code
COPY application /application
WORKDIR /application
#RUN composer require spiral/roadrunner
