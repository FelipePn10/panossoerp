ALTER TABLE product_masks
DROP CONSTRAINT IF EXISTS product_masks_business_id_product_id_key,
DROP COLUMN IF EXISTS business_id,
DROP COLUMN IF EXISTS product_id;