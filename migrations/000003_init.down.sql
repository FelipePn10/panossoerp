DROP INDEX IF EXISTS idx_component_masks_created_by;
DROP INDEX IF EXISTS idx_components_created_by;
DROP INDEX IF EXISTS idx_products_created_by;

ALTER TABLE component_masks
DROP CONSTRAINT IF EXISTS fk_component_masks_created_by;

ALTER TABLE component_masks
DROP COLUMN IF EXISTS created_by;

ALTER TABLE components
DROP CONSTRAINT IF EXISTS fk_components_created_by;

ALTER TABLE components
DROP COLUMN IF EXISTS created_by;

ALTER TABLE product_masks
DROP CONSTRAINT IF EXISTS fk_product_masks_created_by;

ALTER TABLE products
DROP CONSTRAINT IF EXISTS fk_products_created_by;

DROP TABLE IF EXISTS users;
