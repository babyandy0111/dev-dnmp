FROM php:7.4-fpm

#RUN apk add --update \
#		$PHPIZE_DEPS \
#		freetype-dev \
#		libjpeg-turbo-dev \
#		libpng-dev \
#		libxml2-dev \
#		libzip-dev \
#		openssh-client \
#		php7-json \
#		php7-openssl \
#		php7-pdo \
#		php7-pdo_mysql \
#		php7-session \
#		php7-simplexml \
#		php7-tokenizer \
#		php7-xml \
#		imagemagick \
#		imagemagick-libs \
#		imagemagick-dev \
#		php7-imagick \
#		php7-pcntl \
#		php7-zip \
#		sqlite \
#	&& docker-php-ext-install iconv soap sockets exif bcmath pdo_mysql pcntl

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

RUN docker-php-ext-configure gd --with-freetype --with-jpeg \
	&& docker-php-ext-install -j$(nproc) gd

RUN set -xe; \
    apt-get update -yqq && \
    pecl channel-update pecl.php.net && \
    apt-get install -yqq \
    apt-utils \
    libzip-dev zip unzip && \
    docker-php-ext-configure zip; \
    docker-php-ext-install zip && \
    php -m | grep -q 'zip'

# add mysqli
RUN printf "\n" | docker-php-ext-install mysqli pdo pdo_mysql

###########################################################################
# Human Language and Character Encoding Support:
###########################################################################
RUN apt-get install -y libicu-dev g++ --fix-missing && \
    docker-php-ext-configure intl && \
    docker-php-ext-install intl

###########################################################################
# ImageMagick:
###########################################################################
RUN apt-get install -y libmagickwand-dev imagemagick && \
    pecl install imagick && \
    docker-php-ext-enable imagick

###########################################################################
# for geoip_country_code_by_name()
###########################################################################
RUN apt-get install -y geoip-database libgeoip-dev && \
    pecl install geoip-beta && \
    docker-php-ext-enable geoip

# add composer
RUN curl -sSL https://getcomposer.org/installer | php \
    && mv composer.phar /usr/local/bin/composer

###########################################################################
# PHP REDIS EXTENSION
###########################################################################
RUN pecl install -o -f redis-3.1.6 && \
    docker-php-ext-enable redis

###########################################################################
# Set Timezone
###########################################################################

ARG TZ=UTC
ENV TZ ${TZ}

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN docker-php-ext-install bcmath


# change www-data's uid and gid for laravel folder permisstion
RUN groupmod -o -g 1000 www-data && \
    usermod -o -u 1000 -g www-data www-data
#EOF