DROP TRIGGER IF EXISTS trigger_generate_product_mask ON product_question_answers;

DROP FUNCTION IF EXISTS generate_product_mask();

CREATE TABLE IF NOT EXISTS product_questions (
    product_id  BIGINT NOT NULL REFERENCES products(id),
    question_id BIGINT NOT NULL REFERENCES questions(id),
    position    INT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (product_id, question_id),
    UNIQUE (product_id, position)
);
