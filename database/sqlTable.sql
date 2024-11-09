-- sql-01
SELECT COUNT(*) AS user_count
FROM user
WHERE username = @username;

-- sql-02
INSERT INTO user (username, password, first_name, last_name)
VALUES
    (@username, @password, @first, @last);

-- sql-03
INSERT INTO tenant (username)
VALUES
    (@username);

-- sql-04
SELECT  username, password
FROM user
WHERE username = @tenant;

-- sql-05
SELECT  username
FROM employee
WHERE username = @username;

-- sql-06
SELECT employee_role.role_name
FROM employee_role
WHERE username = 'employee3';

-- sql-07
INSERT INTO reservation(move_in_date_time, reservation_room_number, tenant_username, manager_username, bill_id)
VALUES
    (@date, @room_number, @username,NULL,NULL);

-- sql-08
SELECT bill_id, payment_term, create_date_time, bill_status
FROM bill
WHERE tenant_username = @username
ORDER BY create_date_time DESC;

-- sql-09
SELECT bill_id, payment_term, create_date_time, bill_status
FROM bill
WHERE tenant_username = @username
ORDER BY create_date_time DESC;

-- sql-10
SELECT bill_id, bill_item_number, unit, unit_price
FROM bill_item
WHERE bill_id = @bill_id;

# -- sql-11
# INSERT INTO transaction (file_number, file_path, bill_id)
# VALUES
#     (@file_number, @file_path , @bill_id);
#
-- sql-12
UPDATE bill
SET bill_status = 1
WHERE bill_id = @bill_id;

-- sql-13
SELECT reservation_id,reservation_room_number,reservation_status,bill_id
FROM reservation
WHERE bill_id = bill_id;

-- sql-14
UPDATE reservation
SET reservation_status = 3
WHERE bill_id = @bill_id;

-- sql-15
SELECT reservation_id, reservation_room_number, reservation_status, bill_id
FROM reservation
WHERE tenant_username = @username;

-- sql-16
SELECT reservation_id, reservation_room_number, reservation_status, bill_id,move_in_date_time
FROM reservation
WHERE tenant_username = @username and reservation_id = @reservation_id;

-- sql-17
SELECT contract_number, contract_year, contract_room_number
FROM contract
WHERE username = @username and contract_status = 0;

-- sql-18
SELECT contract_number, contract_year, contract_room_number,rental_price,water_rate,electricity_rate,internet_service_fee
FROM contract
WHERE contract_number = @contract_number and contract_year = @year;

-- sql-19
SELECT reservation_id, tenant_username, reservation_status,bill_id
FROM reservation;

-- sql-20
SELECT reservation_id, tenant_username, reservation_status,bill_id
FROM reservation
WHERE reservation_id = @reservation_id;

-- sql-21
SELECT reservation_id, tenant_username, reservation_status,bill_id
FROM reservation
WHERE reservation_id = @reservation_id;

-- sql-22
INSERT INTO bill (payment_term, tenant_username, cashier_username)
VALUES
    (5,@username,@username);

-- sql-23
INSERT INTO bill_item (bill_id, bill_item_number, bill_item_name,unit, unit_price)
VALUES
        (@bill_id,@bill_item_number,'ค่ามัดจำ',1,3000);

-- sql-24
UPDATE reservation
SET reservation_status = 1 , bill_id = @bill_id
WHERE reservation_id = @reservation_id;

-- sql-25
UPDATE reservation
SET reservation_status = 2
WHERE reservation_id = @reservation_id;

-- sql-26
SELECT bill_id, create_date_time, bill_status
FROM bill;

-- sql-27
SELECT bill_id, payment_term, create_date_time, bill_status, tenant_username
FROM bill
WHERE bill_id = @bill_id;

-- sql-28
SELECT bill_id, bill_item_number, bill_item_name, unit, unit_price
FROM bill_item
WHERE bill_id = @bill_id;

-- sql-29
SELECT transaction_id, payment_date_time, transaction_status, username
FROM transaction
WHERE bill_id = @bill_id;

-- sql-30
SELECT transaction_id, payment_date_time, transaction_status, username
FROM transaction
WHERE transaction_id = @transaction_id;

-- sql-31
UPDATE transaction
SET transaction_status = 1
WHERE transaction_id = @transaction_id;

-- sql-32
UPDATE bill
SET bill_status = 2
WHERE bill_id = @bill_id;

-- sql-33
SELECT reservation_id, tenant_username, reservation_status,bill_id
FROM reservation
WHERE bill_id = @bill_id;

-- sql-34
SELECT tenant_username
FROM bill
WHERE bill_id = @bill_id;

-- sql-35
INSERT INTO bill(payment_term, tenant_username, cashier_username)
VALUES
    (-1,@username,@username);

-- sql-36
INSERT INTO bill_item (bill_id, bill_item_number, bill_item_name,unit, unit_price)
VALUES
    (@bill_id,@bill_item_number,'ค่าปรับ',@days,100);

-- sql-37
UPDATE reservation
SET reservation_status = 4
WHERE bill_id = @bill_id;

-- sql-38
UPDATE transaction
SET transaction_status = 2
WHERE transaction_id = @transaction_id;

-- sql-39
SELECT user.username,last_name,first_name
FROM tenant
JOIN user on tenant.username = user.username;

-- sql-40
SELECT COUNT(*) AS contract_count
FROM contract
WHERE contract_room_number = @room_number AND contract_status = 0;

-- sql-41
INSERT INTO contract (contract_year, contract_room_number, rental_price, water_rate, electricity_rate, internet_service_fee, username)
VALUES
    (@year,@contract_room_number,@rental_price,@water_rate,@electricity_rate,@internet_service_fee,@username);

-- sql-42
UPDATE contract
SET contract_status = 1
WHERE contract_room_number = @contract_room_number;

-- sql-43
SELECT ledger_month, ledger_year
FROM ledger;

-- sql-44
SELECT ledger_item_room_number, ledger_month, ledger_year, water_unit, electricity_unit, ledger_item_status
FROM ledger_item
WHERE ledger_month = @month AND ledger_year = @year;

-- sql-45
SELECT ledger_item_room_number, ledger_month, ledger_year, water_unit, electricity_unit, ledger_item_status
FROM ledger_item
WHERE ledger_month = @month AND ledger_year = @year AND ledger_item_room_number = @room_number;

-- sql-46
UPDATE ledger_item
SET water_unit = @water_unit , electricity_unit = @electricity_unit
WHERE ledger_month = @month AND ledger_year = @year AND ledger_item_room_number = @room_number;

-- sql-47
SELECT ledger_month, ledger_year
FROM ledger
WHERE ledger_month= @month AND ledger_year=@year;

-- sql-48
INSERT INTO ledger (ledger_month, ledger_year, username)
VALUES
    (@month,@year,@username);

-- sql-49
INSERT INTO ledger_item (ledger_item_room_number, ledger_month, ledger_year)
VALUES (101, @month, @year),
       (102, @month, @year),
       (103, @month, @year),
       (104, @month, @year),
       (105, @month, @year);

-- Insert ledger_item for room numbers 201 - 206
INSERT INTO ledger_item (ledger_item_room_number, ledger_month, ledger_year)
VALUES (201, @month, @year),
       (202, @month, @year),
       (203, @month, @year),
       (204, @month, @year),
       (205, @month, @year),
       (206, @month, @year);

-- Insert ledger_item for room numbers 301 - 306
INSERT INTO ledger_item (ledger_item_room_number, ledger_month, ledger_year)
VALUES (301, @month, @year),
       (302, @month, @year),
       (303, @month, @year),
       (304, @month, @year),
       (305, @month, @year),
       (306, @month, @year);

-- Insert ledger_item for room numbers 401 - 406
INSERT INTO ledger_item (ledger_item_room_number, ledger_month, ledger_year)
VALUES (401, @month, @year),
       (402, @month, @year),
       (403, @month, @year),
       (404, @month, @year),
       (405, @month, @year),
       (406, @month, @year);

-- Insert ledger_item for room numbers 501 - 506
INSERT INTO ledger_item (ledger_item_room_number, ledger_month, ledger_year)
VALUES (501, @month, @year),
       (502, @month, @year),
       (503, @month, @year),
       (504, @month, @year),
       (505, @month, @year),
       (506, @month, @year);

-- Insert ledger_item for room numbers 601 - 603
INSERT INTO ledger_item (ledger_item_room_number, ledger_month, ledger_year)
VALUES (601, @month, @year),
       (602, @month, @year),
       (603, @month, @year);

-- sql-50
SELECT ledger_month, ledger_year
FROM ledger;

-- sql-51
SELECT ledger_item_room_number, ledger_month, ledger_year, water_unit, electricity_unit, ledger_item_status
FROM ledger_item
WHERE ledger_month = @month AND ledger_year = @year;

-- sql-52
SELECT contract_number, contract_year, contract_room_number, rental_price, water_rate, electricity_rate, internet_service_fee, username
FROM contract
WHERE contract_status = @status
  AND contract_room_number = @contract_room_number
ORDER BY contract_number DESC
LIMIT 1;

-- sql-53
SELECT first_name,last_name
FROM user
WHERE username = @username;

-- sql-54
SELECT ledger_month, ledger_year, ledger_item_room_number,water_unit,electricity_unit
FROM ledger_item
WHERE ledger_month = @month AND ledger_year = @year AND ledger_item_room_number = @room_number;

-- sql-55
UPDATE ledger_item
SET ledger_item_status = 1
WHERE ledger_month = @month AND ledger_year = @year AND ledger_item_room_number = @room_number;

-- sql-56
INSERT INTO bill(payment_term, tenant_username, cashier_username)
VALUES
    (-1,@username,@username);

-- sql-57
INSERT INTO bill_item (bill_id, bill_item_number, bill_item_name,unit, unit_price)
VALUES
    (@bill_id,@bill_item_number,'ค่าน้ำ',@unit,@water_rate),
    (@bill_id,@bill_item_number,'ค่าไฟ',@unit,@electricity_rate),
    (@bill_id,@bill_item_number,'ค่าห้อง',1,@rental_price),
    (@bill_id,@bill_item_number,'ค่าอินเตอร์เน็ต',1,@internet_service_fee);
