insert into account (document_number, credit_limit) values ("12345678900", 5000.00);

insert into operation_type (id, description) values (1, "COMPRA A VISTA");
insert into operation_type (id, description) values (2, "COMPRA PARCELADA");
insert into operation_type (id, description) values (3, "SAQUE");
insert into operation_type (id, description) values (4, "PAGAMENTO");

insert into transaction (account_id, operation_type_id, amount, event_date) values (1, 1, -50.00, '2020-01-01 10:32:07.7199222');
insert into transaction (account_id, operation_type_id, amount, event_date) values (1, 1, -23.05, '2020-01-01 10:48:12.2135875');
insert into transaction (account_id, operation_type_id, amount, event_date) values (1, 1, -18.07, '2020-01-02 19:01:23.1458543');
insert into transaction (account_id, operation_type_id, amount, event_date) values (1, 4, 60.00, '2020-01-05 09:34:18.5893223');