INSERT INTO `accounts` (id, name, saldo, group_id, nfc_chip_uid)
VALUES (2, 'testaccount1', 120, 1, 'testchipid2');

INSERT INTO `transactions` (old_saldo, new_saldo, amount, account_id, created)
VALUES (120, 115, -5, 1, '2019-01-17 16:15:14'),
       (115, 110, -5, 1, '2019-02-17 16:15:14'),
       (110, 105, -5, 1, '2019-03-17 16:15:14'),
       (105, 100, -5, 1, '2019-04-17 16:15:14'),
       (100, 95, -5, 1, '2019-05-17 16:15:14'),
       (120, 115, -5, 2, '2019-06-17 16:15:14'),
       (115, 110, -5, 2, '2019-07-17 16:15:14'),
       (110, 105, -5, 2, '2019-08-17 16:15:14'),
       (105, 100, -5, 2, '2019-09-17 16:15:14');
