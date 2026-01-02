DROP TRIGGER IF EXISTS trigger_generate_product_mask ON product_question_answers;

DROP FUNCTION IF EXISTS generate_product_mask();

CREATE OR REPLACE FUNCTION generate_product_mask()
RETURNS TRIGGER AS $$
DECLARE
    concatenated_mask TEXT;
BEGIN
    SELECT string_agg(
        pq.position || ':' || pqa.answer,
        '#' ORDER BY pq.position
    )
    INTO concatenated_mask
    FROM product_question_answers pqa
    JOIN product_questions pq
      ON pq.product_id = pqa.product_id
     AND pq.question_id = pqa.question_id
    WHERE pqa.product_id = NEW.product_id;

    INSERT INTO product_masks (
        product_id,
        mask,
        mask_hash,
        business_id,
        created_by,
        created_at
    )
    VALUES (
        NEW.product_id,
        concatenated_mask,
        substr(md5(concatenated_mask), 1, 8),
        'default_business', -- você pode mudar para parametrizável se quiser
        NEW.created_by,
        now()
    );

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 4️⃣ Cria novo trigger
CREATE TRIGGER trigger_generate_product_mask
AFTER INSERT ON product_question_answers
FOR EACH ROW
EXECUTE FUNCTION generate_product_mask();
