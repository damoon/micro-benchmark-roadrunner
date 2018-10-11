FROM alpine:3.8

RUN apk add --update --no-cache nginx

COPY nginx/nginx.conf /etc/nginx/
COPY nginx/symfony.conf /etc/nginx/conf.d/default.conf

RUN echo "upstream php-upstream { server php:9001; }" > /etc/nginx/conf.d/upstream.conf

RUN adduser -D -g '' -G www-data www-data

COPY --chown=www-data:www-data application /var/www/symfony

CMD ["nginx"]
