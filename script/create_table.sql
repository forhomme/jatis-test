CREATE TABLE IF NOT EXIST customers
(
    customer_id bigserial primary key,
    company_name varchar(50),
    first_name varchar(30),
    last_name varchar(50),
    billing_address varchar(255),
    city varchar(50),
    state_or_province varchar(20),
    zip_code varchar(20),
    email varchar(75),
    company_website varchar(200),
    phone_number varchar(30),
    fax_number varchar(30),
    ship_address varchar(255),
    ship_city varchar(50)
    ship_state_or_province varchar(50),
    ship_zip_code varchar(20),
    ship_phone_number varchar(30)
);

CREATE TABLE IF NOT EXIST employees
(
    employee_id bigserial primary key,
    first_name varchar(50),
    last_name varchar(50),
    title varchar(50),
    work_phone varchar(30)
);

CREATE TABLE IF NOT EXIST products
(
    product_id bigserial primary key,
    product_name varchar(50),
    unit_price double precision,
    in_stock int
);

CREATE TABLE IF NOT EXIST shipping_methods
(
    shipping_method_id bigserial primary key,
    shipping_method varchar(20)
);

CREATE TABLE IF NOT EXIST orders
(
    order_id bigserial primary key,
    customer_id bigint not null,
    employee_id bigint not null,
    order_date timestamp with time zone default CURRENT_TIMESTAMP,
    purchase_order_number varchar(30) not null,
    ship_date timestamp with time zone,
    shipping_method_id bigint,
    freight_charge double precision,
    taxes double precision,
    payment_received int,
    comment varchar(150),
    CONSTRAINT fk_order_customer
        FOREIGN KEY(customer_id)
            REFERENCES customers(customer_id),
    CONSTRAINT fk_order_employee
        FOREIGN KEY(employee_id)
            REFERENCES employees(employee_id),
    CONSTRAINT fk_order_shipping_method
        FOREIGN KEY(shipping_method_id)
            REFERENCES shipping_methods(shipping_method_id)

);

CREATE TABLE IF NOT EXIST order_details
(
    order_detail_id bigserial primary key,
    order_id bigint,
    product_id bigint,
    quantity bigint,
    unit_price double precision,
    discount integer,
    CONSTRAINT fk_order_detail_orders
        FOREIGN KEY(order_id)
            REFERENCES orders(order_id),
    CONSTRAINT fk_order_detail_products
        FOREIGN KEY(product_id)
            REFERENCES products(product_id)
);

-- List of customers located in Irvine city.
SELECT * from customers where city = 'Irvine';

-- List of customers whose order is handled by an employee named Adam Barr.
SELECT * from customers c
    JOIN orders o on c.customer_id = o.customer_id
    JOIN employees e on o.employee_id = e.employee_id
    WHERE e.first_name = 'Adam' and e.last_name = 'Barr';

-- List of products which are ordered by "Contonso, Ltd" Company" Company.
SELECT * from products p
    JOIN order_detail od on p.product_id = od.product_id
    JOIN orders o on od.order_id = o.order_id
    JOIN customer c on o.customer_id = c.customer_id
    WHERE c.company_name = 'Contonso, Ltd" Company';

-- List of transactions (orders) which has "UPS Ground" as shipping method.
SELECT * from orders o
    JOIN order_detail od on o.order_id = od.order_id
    JOIN shipping_methods sm on o.shipping_method_id = sm.shipping_method_id
    WHERE sm.shipping_method = 'UPS Ground'

-- List of total cost (including tax and freight charge) for every order sorted by ship date.
SELECT (COALESCE(od.unit_price, 0) * COALESCE(od.quantity, 0) * (100 - COALESCE(od.discount,0)) + COALESCE(o.freight_charge,0)
            + COALESCE(o.taxes,0)) as total_cost from orders o
    JOIN order_detail od on o.order_id = od.order_id
    ORDER BY o.ship_date desc