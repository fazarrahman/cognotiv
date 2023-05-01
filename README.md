# cognotiv

# Create new MySQL database 
#### Create new MySQL database. You can name it anything, for example : cognotivdb

# Fill .env file with these config
You can change the config to your mysql server settings

APP_PORT=8080
DB_HOST=localhost
DB_PORT=3306
DB_NAME=cognotivdb
DB_USERNAME=root
DB_PASSWORD=
AUTH_CLIENT_ID=000000
AUTH_SECRET=999999
APP_DOMAIN=http://localhost

# MySQL create tables queries
CREATE TABLE roles (
	id bigint NOT NULL AUTO_INCREMENT,
	code VARCHAR(255) NOT NULL,
	name varchar(255) NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NULL,
	PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

CREATE TABLE users (
	id bigint NOT NULL AUTO_INCREMENT,
	username VARCHAR(100) NOT NULL,
	`password` VARCHAR(255) NOT NULL,
	role_id bigint NOT NULL,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NULL,
	address VARCHAR(255) NULL,
	phone VARCHAR(255) NULL,
	email VARCHAR(255) NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NULL,
	PRIMARY KEY(`id`),
	KEY `FK_users_roles` (`role_id`),
  	CONSTRAINT `FK_users_roles` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

CREATE TABLE products (
	id bigint NOT NULL AUTO_INCREMENT,
	name varchar(255) NOT NULL,
	price decimal NOT NULL,
	description varchar(255) NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NULL,
	PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;


CREATE TABLE orders (
	id bigint NOT NULL AUTO_INCREMENT,
	user_id bigint NOT NULL,
	order_date timestamp NOT NULL,
	status int NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NULL,
	PRIMARY KEY(`id`),
	KEY `FK_users` (`user_id`),
  	CONSTRAINT `FK_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

CREATE TABLE order_details (
	id bigint NOT NULL AUTO_INCREMENT,
	order_id bigint NOT NULL,
	product_id bigint NOT NULL,
	quantity bigint NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NULL,
	PRIMARY KEY(`id`),
	KEY `FK_order_detail_order` (`order_id`),
  	CONSTRAINT `FK_order_detail_order` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
  	KEY `FK_order_detail_product` (`product_id`),
  	CONSTRAINT `FK_order_detail_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

# Insert initial data To MySQL DB
insert into roles (code, name, created_at)
values ('ADMIN', 'admin', NOW())

insert into roles (code, name, created_at)
values ('CUSTOMER', 'customer', NOW())

insert into products (name, price, created_at)
values ('PS5', 7000000, NOW())

insert into products (name, price, created_at)
values ('TWS', 1000000, NOW())

# API Endpoints
## Sign Up New User
Role id 1 is admin
Role id 2 is customer

Login with admin user, you can get all active orders
Login with customer, you can only get your own orders

### Curl :
curl --location 'http://localhost:8080/api/auth/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "fazar2",
    "password": "fazar123",
    "firstName": "Fazar",
    "lastName": "Rahman",
    "email": "fazar.rahman@gmail.com",
    "roleId": 2
}'

## User login
The response of user login will contain access token. Copy it

### Curl : 
curl --location 'http://localhost:8080/api/auth/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "fazar2",
    "password": "fazar123"
}'

## Token Info
### Curl : 
Paste the access token to authorization header

curl --location 'http://localhost:8080/api/auth/tokeninfo' \
--header 'Authorization: Bearer {Access token}'

## Post Order
### Curl :
curl --location 'localhost:8080/api/order' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {Access token}' \
--data '{
    "userId":1,
    "details": [
        {
            "productId": 1,
            "quantity": 2
        },
        {
            "productId": 2,
            "quantity": 3
        }
    ]
}'

## Get Order detail with product list
Paste the access token to authorization header

### Curl :
curl --location 'localhost:8080/api/order' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {Access token}' \
--data ''