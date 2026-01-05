CREATE TABLE product_masks (
    id BIGINT PRIMARY KEY,
    product_code VARCHAR(10) NOT NULL,
    mask TEXT NOT NULL,
    mask_hash CHAR(8) NOT NULL,
    created_by UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL,
    UNIQUE (mask_hash)
);
