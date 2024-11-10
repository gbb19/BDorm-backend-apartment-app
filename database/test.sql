INSERT INTO user (username, password, first_name, last_name)
VALUES ('employee3', 'password', 'first', 'last');

INSERT INTO employee (username)
VALUES ('employee3');

INSERT INTO employee_role (username, role_name)
VALUES ('employee3', 'cashier'),
       ('employee3', 'manager');

SELECT reservation_id
FROM reservation;

SELECT *
FROM contract
WHERE contract_room_number = 101
  AND contract_status = 0;


SELECT internet_service_fee,water_rate,electricity_rate,username
FROM contract
WHERE contract_room_number = 104
  AND contract_status = 0

