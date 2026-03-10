ALTER TABLE boms
ADD COLUMN mask BIGINT NOT NULL REFERENCES product_masks(id);

ALTER TABLE bom_items
ADD COLUMN mask_component BIGINT NOT NULL REFERENCES component_masks(id);

CREATE TABLE IF NOT EXISTS question_options(
  question_id BIGINT NOT NULL REFERENCES questions(id),
  created_by UUID NOT NULL REFERENCES users(id),
  value TEXT NOT NULL
);

ALTER TABLE question_options
ADD COLUMN created_by UUID NOT NULL;
