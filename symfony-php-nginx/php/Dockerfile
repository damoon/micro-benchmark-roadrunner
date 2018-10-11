FROM alpine:3.8

RUN apk add --update --no-cache \
    php7-fpm \
    php7-apcu \
    php7-ctype \
    php7-curl \
    php7-dom \
    php7-gd \
    php7-iconv \
    php7-imagick \
    php7-json \
    php7-intl \
    php7-mcrypt \
    php7-mbstring \
    php7-opcache \
    php7-openssl \
    php7-pdo \
    php7-pdo_mysql \
    php7-mysqli \
    php7-xml \
    php7-zlib \
    php7-phar \
    php7-tokenizer \
    php7-session \
    php7-simplexml \
#    php7-xdebug \
    make \
    curl

RUN apk add --update --no-cache php7-sqlite3 php7-pdo_sqlite

#RUN rm -rf /var/cache/apk/* && rm -rf /tmp/* && \
#    curl --insecure https://getcomposer.org/composer.phar -o /usr/bin/composer && chmod +x /usr/bin/composer

ADD php/symfony.ini /etc/php7/php-fpm.d/
ADD php/symfony.ini /etc/php7/cli/conf.d/
#ADD php/xdebug.ini  /etc/php7/conf.d/

ADD php/symfony.pool.conf /etc/php7/php-fpm.d/

CMD ["php-fpm7", "-F"]

#COPY --chown=www-data:www-data application /var/www/symfony
COPY application /var/www/symfony
RUN chmod a+rwx /var/www/symfony/var/log

WORKDIR /var/www/symfony
EXPOSE 9000


ENV APP_ENV=prod
ENV APP_DEBUG=0
ENV APP_SECRET=67d829bf61dc5f87a73fd814e2c9f629

ENV DATABASE_URL=sqlite:///%kernel.project_dir%/data/database.sqlite
ENV MAILER_URL=null://localhost
