CREATE TABLE IF NOT EXISTS suppliers
(
    id           SERIAL,
    code         VARCHAR(100) NOT NULL UNIQUE,
    name         VARCHAR(100) NOT NULL,
    address      VARCHAR(256) NOT NULL,
    phone        VARCHAR(20)  NOT NULL,
    created_at   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
)