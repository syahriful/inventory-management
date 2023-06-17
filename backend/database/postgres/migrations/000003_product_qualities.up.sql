CREATE TABLE IF NOT EXISTS product_qualities
(
   id               SERIAL,
   product_code     VARCHAR(100)    NOT NULL,
   quality          VARCHAR(100)    NOT NULL,
   price            BIGINT          NOT NULL DEFAULT 0,
   quantity         DECIMAL(10,3)   NOT NULL,
   type             VARCHAR(20)     NOT NULL,
   created_at       TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at       TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (id),
   FOREIGN KEY (product_code) REFERENCES products(code) ON DELETE CASCADE ON UPDATE CASCADE
)