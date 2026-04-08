CREATE TABLE IF NOT EXISTS item_mask_answers (
    id BIGINT PRIMARY KEY,
    question_id BIGINT NOT NULL REFERENCES questions(id),
    option_id BIGINT NOT NULL REFERENCES question_options(id),
    position INTEGER NOT NULL,
    mask_id BIGINT NOT NULL REFERENCES item_masks(id)
);
