DROP TRIGGER IF EXISTS trg_generate_mask ON product_mask_answers;

DROP FUNCTION IF EXISTS generate_mask();

DROP TABLE IF EXISTS product_mask_answers;

DROP TABLE IF EXISTS generated_masks;
