DROP INDEX IF EXISTS idx_mask_composition_parent;
DROP INDEX IF EXISTS idx_component_masks_business_id;
DROP INDEX IF EXISTS idx_components_type;
DROP INDEX IF EXISTS idx_components_code_type;
DROP INDEX IF EXISTS idx_product_masks_business_id;
DROP INDEX IF EXISTS idx_products_code;

DROP TABLE IF EXISTS material_consumption;
DROP TABLE IF EXISTS mask_composition;
DROP TABLE IF EXISTS component_masks;
DROP TABLE IF EXISTS components;
DROP TABLE IF EXISTS product_masks;
DROP TABLE IF EXISTS products;

DROP TYPE IF EXISTS component_type;
