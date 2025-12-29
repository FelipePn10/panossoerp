CREATE TABLE product_question_answers (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_id BIGINT NOT NULL REFERENCES products(id),
    question_id BIGINT NOT NULL REFERENCES questions(id),
    answer TEXT NOT NULL,
    created_by UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (product_id, question_id)
);

CREATE OR REPLACE FUNCTION generate_product_mask()
RETURNS TRIGGER AS $$
DECLARE
    concatenated_mask TEXT;
BEGIN
    SELECT string_agg(answer, '#' ORDER BY id)
    INTO concatenated_mask
    FROM product_question_answers
    WHERE product_id = NEW.product_id;

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
        'default_business',
        NEW.created_by,
        now()
    );

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_generate_product_mask
AFTER INSERT ON product_question_answers
FOR EACH ROW
EXECUTE FUNCTION generate_product_mask();
