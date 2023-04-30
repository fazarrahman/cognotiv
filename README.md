# cognotiv

# MySQL create tables query
CREATE TABLE customers (
	id bigint NOT NULL AUTO_INCREMENT,
	username VARCHAR(100) NOT NULL,
	password VARCHAR(255) NOT NULL,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NULL,
	address VARCHAR(255) NULL,
	phone VARCHAR(255) NULL,
	email VARCHAR(255) NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NULL,
	PRIMARY KEY(`id`)
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
	customer_id bigint NOT NULL,
	order_date timestamp NOT NULL,
	status int NOT NULL,
	created_at timestamp NOT NULL,
	updated_at timestamp NULL,
	PRIMARY KEY(`id`),
	KEY `FK_customers` (`customer_id`),
  	CONSTRAINT `FK_customers` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`)
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

# Endpoints
## Post Order
### Curl :
curl --location 'localhost:8080/api/order' \
--header 'Content-Type: application/json' \
--data '{
    "customerId":1,
    "details": [
        {
            "productId": 1,
            "quantity": 1
        },
        {
            "productId": 2,
            "quantity": 1
        }
    ]
}'