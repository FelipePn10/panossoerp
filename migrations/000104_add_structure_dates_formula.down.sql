BEGIN;

ALTER TABLE item_structures
    DROP COLUMN IF EXISTS start_date,
    DROP COLUMN IF EXISTS end_date,
    DROP COLUMN IF EXISTS loss_formula;

COMMIT;
