Insert into users (name, email, password) VALUES
('sbres', 'nope@nope.com', '$2a$10$OaIA9tJk6VH7gdex1tOsr.F.7DZX07axteBU9NcPXTQ87xm7CkPRW');
-- That should be 12345678 password

insert into "subdomains" ("prefix", "user_id", "destination_url", "destination_secret", "percentage_fee", "name") values
('wow', 1, '127.0.0.1:402', 'random_secret_key', '3', 'Example API'),
('much.subdomain', 1, '127.0.0.1', 'random_secret_key', '3', 'Example API 2');

Insert into "subdomain_description" ("subdomain_id", "description", "doc")
VALUES (1, 'This api is used to test that everything is working', 'Here we can put a more complete documentation');

insert into "endpoints" ("cost", "method", "path", "retry_on_failure", "subdomain_id", "success_code", "timeout")
values
(1000, 'GET', '/ping', false, 1, 200, 1000),
(10000, 'GET', '/sleep', false, 1, 200, 1000)
;

Insert into "endpoint_description" ("endpoint_id", "description", "swagger_doc") VALUES 
(1, 'This endpoint will respond pong when the call is received', '');

INSERT INTO "global_fees" ("base_fee", "percentage_fee") VALUES
(1000, 3.50)