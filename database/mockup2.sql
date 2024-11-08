INSERT INTO employee(username)
VALUES ('employee3');

INSERT INTO employee_role (username, role_name)
VALUES ('employee3', 'cashier'),
       ('employee3', 'accountant');

SELECT @@global.time_zone, @@session.time_zone;

INSERT INTO bill (payment_term, bill_status, tenant_username, cashier_username)
VALUES (1, 0, 'tenant1', 'employee1');

INSERT INTO bill_item (bill_id, bill_item_number, bill_item_name, unit, unit_price)
VALUES (3, 1, 'b1', 100, 1.5),
       (3, 2, 'b2', 50, 2.0);


INSERT INTO transaction( bill_id)
VALUES (@bill_id)