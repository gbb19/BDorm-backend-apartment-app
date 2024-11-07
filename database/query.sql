-- User
CREATE TABLE user
(
    username   VARCHAR(256),
    password   VARCHAR(256) NOT NULL,
    first_name VARCHAR(256) NOT NULL,
    last_name  VARCHAR(256) NOT NULL,

    PRIMARY KEY (username)
);

-- Tenant
CREATE TABLE tenant
(
    username VARCHAR(256),

    PRIMARY KEY (username),
    FOREIGN KEY (username) REFERENCES user (username)
);

-- Employee
CREATE TABLE employee
(
    username VARCHAR(256),

    PRIMARY KEY (username),
    FOREIGN KEY (username) REFERENCES user (username)
);

-- Contract
CREATE TABLE contract
(
    contract_number      INT AUTO_INCREMENT,
    contract_year        YEAR,
    contract_room_number INT,
    rental_price         FLOAT        NOT NULL CHECK (rental_price >= 0),
    water_rate           FLOAT        NOT NULL CHECK (water_rate >= 0),
    electricity_rate     FLOAT        NOT NULL CHECK (electricity_rate >= 0),
    internet_service_fee FLOAT        NOT NULL CHECK (internet_service_fee >= 0),
    contract_status      INT(1)       NOT NULL CHECK (contract_status IN (0, 1)) DEFAULT 0,
    username             VARCHAR(256) NOT NULL,

    PRIMARY KEY (contract_number, contract_year),
    FOREIGN KEY (username) REFERENCES user (username)
);

-- Bill
CREATE TABLE bill
(
    bill_id          INT AUTO_INCREMENT,
    payment_term     INT          NOT NULL,
    create_date_time TIMESTAMP    NOT NULL                                    DEFAULT CURRENT_TIMESTAMP,
    bill_status      INT(1)       NOT NULL CHECK ( bill_status IN (0, 1, 2) ) DEFAULT 0,
    tenant_username  VARCHAR(256) NOT NULL,
    cashier_username VARCHAR(256) NOT NULL,

    PRIMARY KEY (bill_id),
    FOREIGN KEY (tenant_username) REFERENCES tenant (username),
    FOREIGN KEY (cashier_username) REFERENCES employee (username)
);

-- Reservation
CREATE TABLE reservation
(
    reservation_id          INT AUTO_INCREMENT,
    move_in_date_time       DATETIME     NOT NULL,
    reservation_room_number INT(3)       NOT NULL,
    reservation_status      INT(1)       NOT NULL CHECK ( reservation_status IN (0, 1, 2, 3, 4) ) DEFAULT 0,

    tenant_username         VARCHAR(256) NOT NULL,
    manager_username        VARCHAR(256) ,
    bill_id                 INT,

    PRIMARY KEY (reservation_id),
    FOREIGN KEY (tenant_username) REFERENCES tenant (username),
    FOREIGN KEY (manager_username) REFERENCES employee (username),
    FOREIGN KEY (bill_id) REFERENCES bill (bill_id)
);

-- Transaction
CREATE TABLE transaction
(
    transaction_id     INT AUTO_INCREMENT,
    payment_date_time  TIMESTAMP     NOT NULL                                           DEFAULT CURRENT_TIMESTAMP,
    transaction_status INT(1)        NOT NULL CHECK ( transaction_status IN (0, 1, 2) ) DEFAULT 0,
    username           VARCHAR(256),
    bill_id            INT           NOT NULL,

    PRIMARY KEY (transaction_id),
    FOREIGN KEY (username) REFERENCES employee (username),
    FOREIGN KEY (bill_id) REFERENCES bill (bill_id)
);

-- Bill_Item
CREATE TABLE bill_item
(
    bill_id          INT,
    bill_item_number INT,
    bill_item_name VARCHAR(256) NOT NULL,
    unit             INT   NOT NULL CHECK ( unit >= 1 ),
    unit_price       FLOAT NOT NULL check ( unit_price >= 0 ),

    PRIMARY KEY (bill_id, bill_item_number),
    FOREIGN KEY (bill_id) REFERENCES bill (bill_id)
);

-- Ledger
CREATE TABLE ledger
(
    ledger_month INT(2),
    ledger_year  YEAR,
    username     VARCHAR(256) NOT NULL,

    PRIMARY KEY (ledger_month, ledger_year),
    FOREIGN KEY (username) REFERENCES employee (username)
);

-- Ledger_Item
CREATE TABLE ledger_item
(
    ledger_item_room_number INT(3),
    ledger_month            INT(2),
    ledger_year             YEAR,
    water_unit              INT    NOT NULL CHECK ( water_unit >= 0 )              DEFAULT 0,
    electricity_unit        INT    NOT NULL CHECK ( electricity_unit >= 0 )        DEFAULT 0,
    ledger_item_status      INT(1) NOT NULL CHECK ( ledger_item_status IN (0, 1) ) DEFAULT 0,

    PRIMARY KEY (ledger_item_room_number, ledger_month, ledger_year),
    FOREIGN KEY (ledger_month, ledger_year) references ledger (ledger_month, ledger_year)
);

-- Role
CREATE TABLE role
(
    role_name VARCHAR(256),

    PRIMARY KEY (role_name)
);

-- Employee_Role
CREATE TABLE employee_role
(
    username  VARCHAR(256),
    role_name VARCHAR(256),

    PRIMARY KEY (username, role_name),
    FOREIGN KEY (username) REFERENCES employee (username),
    FOREIGN KEY (role_name) REFERENCES role (role_name)
);
