FROM php:8.0.0
RUN  --mount=type=cache,target=/var/cache/apt --mount=type=cache,target=/var/lib/apt apt-get update -y && apt-get -y install gcc \
    curl \
    make \
    autoconf \
    libc-dev \
    pkg-config \
    libzip-dev \
    libfreetype6-dev \
    libpng-dev \
    libjpeg62-turbo-dev \
    libssl-dev \
    libwebp-dev \
    libmcrypt-dev \
    vim

# Install PHP extensions
RUN docker-php-ext-install pdo_mysql

RUN curl -sSL https://getcomposer.org/installer | php \
    && mv composer.phar /usr/local/bin/composer

RUN groupmod -o -g 1000 www-data && \
    usermod -o -u 1000 -g www-data www-data