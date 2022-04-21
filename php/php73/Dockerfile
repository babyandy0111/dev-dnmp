FROM php:7.3-fpm-alpine
# @see https://hub.docker.com/r/jpswade/php7.4-fpm-alpine

# Install gd, iconv, mbstring, mysql, soap, sockets, zip, and zlib extensions
# $PHPIZE_DEPS include autoconf, make...etc
# see example at https://hub.docker.com/_/php/
RUN apk add --update \
		$PHPIZE_DEPS \
		freetype-dev \
		git \
		libjpeg-turbo-dev \
		libpng-dev \
		libxml2-dev \
		libzip-dev \
		openssh-client \
		php7-json \
		php7-openssl \
		php7-pdo \
		php7-pdo_mysql \
		php7-session \
		php7-simplexml \
		php7-tokenizer \
		php7-xml \
		imagemagick \
		imagemagick-libs \
		imagemagick-dev \
		php7-imagick \
		php7-pcntl \
		php7-zip \
		sqlite \
	&& docker-php-ext-install soap sockets exif bcmath pdo_mysql pcntl \
	#&& docker-php-ext-configure gd --with-jpeg --with-freetype \
	#&& docker-php-ext-install gd \
	&& docker-php-ext-install zip

# add mysqli
RUN printf "\n" | docker-php-ext-install mysqli

# add intl
RUN printf "\n" | apk add --update \
		icu-dev \
	&& docker-php-ext-configure intl \
	&& docker-php-ext-install intl

# add mcrypt
RUN printf "\n" | apk add --update \
		libmcrypt-dev \
	&& pecl install \
					mcrypt && \
	docker-php-ext-enable mcrypt

# add imagick
RUN printf "\n" | pecl install \
		imagick && \
		docker-php-ext-enable --ini-name 20-imagick.ini imagick

# add pcov
RUN printf "\n" | pecl install \
		pcov && \
		docker-php-ext-enable pcov

# add composer
RUN php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');" \
	&& php composer-setup.php \
	&& php -r "unlink('composer-setup.php');" \
	&& mv composer.phar /usr/bin/composer

# setup timezone
RUN sed -i 's/;date.timezone =/date.timezone = "Asia\/Taipei"/g' /etc/php7/php.ini

# change www-data's uid and gid for laravel folder permisstion
RUN apk --no-cache add shadow && \
    usermod -u 1000 www-data && \
    groupmod -g 1000 www-data

#EOF