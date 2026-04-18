ALTER TABLE item_structures
    ALTER COLUMN parent_code SET NOT NULL,
ALTER COLUMN child_code SET NOT NULL;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_constraint
        WHERE conname = 'uq_items_code'
    ) THEN
ALTER TABLE items
    ADD CONSTRAINT uq_items_code UNIQUE (code);
END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_constraint
        WHERE conname = 'fk_structure_parent_code'
    ) THEN
ALTER TABLE item_structures
    ADD CONSTRAINT fk_structure_parent_code
        FOREIGN KEY (parent_code) REFERENCES items(code)
            ON DELETE RESTRICT;
END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_constraint
        WHERE conname = 'fk_structure_child_code'
    ) THEN
ALTER TABLE item_structures
    ADD CONSTRAINT fk_structure_child_code
        FOREIGN KEY (child_code) REFERENCES items(code)
            ON DELETE RESTRICT;
END IF;
END$$;

CREATE INDEX IF NOT EXISTS idx_item_structures_parent_code
    ON item_structures(parent_code);

CREATE INDEX IF NOT EXISTS idx_item_structures_child_code
    ON item_structures(child_code);