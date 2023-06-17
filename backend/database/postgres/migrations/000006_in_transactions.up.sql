CREATE TABLE IF NOT EXISTS transactions
(
    id                  SERIAL,
    code                VARCHAR(100)  NOT NULL UNIQUE,
    product_quality_id  INT           NOT NULL,
    supplier_code       VARCHAR(100),
    customer_code       VARCHAR(100),
    description         TEXT,
    quantity            DECIMAL(10,3) NOT NULL,
    type                VARCHAR(20)  NOT NULL,
    created_at          TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (product_quality_id) REFERENCES product_qualities(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (supplier_code) REFERENCES suppliers(code) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (customer_code) REFERENCES customers(code) ON UPDATE CASCADE ON DELETE CASCADE
)