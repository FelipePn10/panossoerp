DROP TRIGGER IF EXISTS trigger_generate_item_mask ON item_question_answers;
DROP FUNCTION IF EXISTS generate_mask();
DROP TABLE IF EXISTS item_mask_answers;
DROP TABLE IF EXISTS generated_masks;
