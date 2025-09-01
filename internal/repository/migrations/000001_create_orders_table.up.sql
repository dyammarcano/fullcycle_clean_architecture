CREATE TABLE IF NOT EXISTS orders
(
    id
    SERIAL
    PRIMARY
    KEY,
    item
    VARCHAR
    NOT
    NULL,
    amount
    INTEGER
    NOT
    NULL
);
