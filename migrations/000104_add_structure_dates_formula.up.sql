BEGIN;

ALTER TABLE item_structures
    ADD COLUMN IF NOT EXISTS start_date   DATE,
    ADD COLUMN IF NOT EXISTS end_date     DATE,
    ADD COLUMN IF NOT EXISTS loss_formula TEXT;

COMMIT;
