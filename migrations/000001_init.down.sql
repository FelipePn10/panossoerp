-- 1. Drop indexes
DROP INDEX IF EXISTS idx_component_masks_created_by;
DROP INDEX IF EXISTS idx_components_created_by;
DROP INDEX IF EXISTS idx_products_created_by;
DROP INDEX IF EXISTS idx_mask_composition_parent;
DROP INDEX IF EXISTS idx_component_masks_business_id;
DROP INDEX IF EXISTS idx_components_type;
DROP INDEX IF EXISTS idx_components_code_type;
DROP INDEX IF EXISTS idx_product_masks_business_id;
DROP INDEX IF EXISTS idx_products_code;

-- 2. Drop Questions and Complements (dependem apenas de themselves)
DROP TABLE IF EXISTS questions;
DROP TABLE IF EXISTS complement_a;
DROP TABLE IF EXISTS complement_b;

-- 3. Drop Material Consumption
DROP TABLE IF EXISTS material_consumption;

-- 4. Drop Mask Composition
DROP TABLE IF EXISTS mask_composition;

-- 5. Drop Component Masks
DROP TABLE IF EXISTS component_masks;

-- 6. Drop Components
DROP TABLE IF EXISTS components;

-- 7. Drop Product Masks
DROP TABLE IF EXISTS product_masks;

-- 8. Drop Products
DROP TABLE IF EXISTS products;

-- 9. Drop Users
DROP TABLE IF EXISTS users;

-- 10. Drop ENUM
DO $$
BEGIN
    DROP TYPE IF EXISTS component_type;
EXCEPTION
    WHEN undefined_object THEN null;
END
$$;
