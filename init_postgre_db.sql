DROP TABLE IF EXISTS product;

CREATE TABLE product (
  id          bigserial PRIMARY KEY,
  name        varchar(100) NOT NULL,
  description varchar(1000) NOT NULL,
  price       integer NOT NULL,
  count       integer NOT NULL
);

INSERT INTO product (name, description, price, count) VALUES (
  'docker in action',
  'a book for docker',
  100,
  50
) RETURNING id;

INSERT INTO product (name, description, price, count) VALUES (
  'golang in action',
  'a book for golang',
  100,
  30
) RETURNING id;
