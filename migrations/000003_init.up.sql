CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

ALTER TABLE products
ADD CONSTRAINT fk_products_created_by
FOREIGN KEY (created_by) REFERENCES users(id);

ALTER TABLE product_masks
ADD CONSTRAINT fk_product_masks_created_by
FOREIGN KEY (created_by) REFERENCES users(id);

ALTER TABLE components
ADD COLUMN created_by UUID NOT NULL;

ALTER TABLE components
ADD CONSTRAINT fk_components_created_by
FOREIGN KEY (created_by) REFERENCES users(id);

ALTER TABLE component_masks
ADD COLUMN created_by UUID NOT NULL;

ALTER TABLE component_masks
ADD CONSTRAINT fk_component_masks_created_by
FOREIGN KEY (created_by) REFERENCES users(id);

CREATE INDEX idx_products_created_by
ON products(created_by);

CREATE INDEX idx_components_created_by
ON components(created_by);

CREATE INDEX idx_component_masks_created_by
ON component_masks(created_by);
