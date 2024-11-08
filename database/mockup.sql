-- Insert mockup data for User
INSERT INTO user (username, password, first_name, last_name)
VALUES ('tenant1', '$2a$10$Q44RIGQ295q9Y37f1bc.G.lcV.RBx49MjrvOwWsadRMcYdRuSFaZi', 'Alice', 'Smith'),
       ('tenant2', '$2a$10$Q44RIGQ295q9Y37f1bc.G.lcV.RBx49MjrvOwWsadRMcYdRuSFaZi', 'Bob', 'Johnson'),
       ('employee1', '$2a$10$Q44RIGQ295q9Y37f1bc.G.lcV.RBx49MjrvOwWsadRMcYdRuSFaZi', 'Charlie', 'Brown'),
       ('employee2', '$2a$10$Q44RIGQ295q9Y37f1bc.G.lcV.RBx49MjrvOwWsadRMcYdRuSFaZi', 'Dana', 'White');

-- Insert mockup data for Tenant
INSERT INTO tenant (username)
VALUES ('tenant1'),
       ('tenant2');

-- Insert mockup data for Employee
INSERT INTO employee (username)
VALUES ('employee1'),
       ('employee2');

-- Insert mockup data for Role
INSERT INTO role (role_name)
VALUES ('cashier'),
       ('manager'),
       ('accountant');

-- Insert mockup data for Employee_Role
INSERT INTO employee_role (username, role_name)
VALUES ('employee1', 'cashier'),
       ('employee2', 'manager');

-- Insert mockup data for Contract
INSERT INTO contract (contract_year, contract_room_number, rental_price, water_rate, electricity_rate,
                      internet_service_fee, contract_status, username)
VALUES (2024, 101, 1200, 15, 10, 25, 0, 'tenant1'),
       (2024, 102, 1100, 14, 9, 22, 0, 'tenant2');

-- Insert mockup data for Bill
INSERT INTO bill (payment_term, bill_status, tenant_username, cashier_username)
VALUES (1, 0, 'tenant1', 'employee1'),
       (2, 1, 'tenant2', 'employee1');

-- Insert mockup data for Reservation
INSERT INTO reservation (move_in_date_time, reservation_room_number, reservation_status, tenant_username,
                         manager_username, bill_id)
VALUES ('2024-08-01 14:00:00', 101, 1, 'tenant1', 'employee2', 1),
       ('2024-08-05 10:30:00', 102, 2, 'tenant2', 'employee2', 2);

-- Insert mockup data for Transaction
INSERT INTO transaction (transaction_status, username, bill_id)
VALUES (1, 'employee1', 1),
       (2, 'employee2', 2);

-- Insert mockup data for Bill_Item
INSERT INTO bill_item (bill_id, bill_item_number, bill_item_name, unit, unit_price)
VALUES (1, 1, 'b1', 100, 1.5),
       (1, 2, 'b2', 50, 2.0),
       (2, 1, 'b3', 80, 1.6),
       (2, 2, 'b4', 40, 2.1);

-- Insert mockup data for Ledger
INSERT INTO ledger (ledger_month, ledger_year, username)
VALUES (7, 2024, 'employee1'),
       (8, 2024, 'employee2');

-- Insert mockup data for Ledger_Item
INSERT INTO ledger_item (ledger_item_room_number, ledger_month, ledger_year, water_unit, electricity_unit,
                         ledger_item_status)
VALUES (101, 7, 2024, 15, 25, 1),
       (102, 7, 2024, 20, 30, 1),
       (101, 8, 2024, 18, 28, 0),
       (102, 8, 2024, 17, 26, 0);
