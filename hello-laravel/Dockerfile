FROM php:8-alpine3.14

COPY --from=composer /usr/bin/composer /usr/bin/composer

WORKDIR /app

COPY . .
RUN composer self-update && composer install

CMD ["php", "artisan", "serve", "--host=0.0.0.0"]
