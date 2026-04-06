DO $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_name = 'item_questions'
        AND column_name = 'product_id'
    ) THEN
        ALTER TABLE item_questions
        RENAME COLUMN product_id TO item_id;
    END IF;
END $$;
