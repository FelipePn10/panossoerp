CREATE TABLE question_options (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    question_id BIGINT NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    value TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (question_id, value)
);

ALTER TABLE questions
DROP COLUMN complement_a_id,
DROP COLUMN complement_b_id;

DROP TABLE complement_a;
DROP TABLE complement_b;
