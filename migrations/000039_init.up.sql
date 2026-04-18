DELETE FROM public.items
WHERE id NOT IN (
    SELECT MIN(id)
    FROM public.items
    GROUP BY code
);

DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_constraint
        WHERE conname = 'uq_items_code'
    ) THEN
ALTER TABLE public.items
    ADD CONSTRAINT uq_items_code UNIQUE (code);
END IF;
END $$;

UPDATE public.item_structures s
SET parent_item_id = i.id
    FROM public.items i
WHERE s.parent_code IS NOT NULL
  AND i.code = s.parent_code;

UPDATE public.item_structures s
SET child_item_id = i.id
    FROM public.items i
WHERE s.child_code IS NOT NULL
  AND i.code = s.child_code;

DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM public.item_structures
        WHERE parent_item_id IS NULL OR child_item_id IS NULL
    ) THEN
        RAISE EXCEPTION 'Migration abortada: existem registros inválidos';
END IF;
END $$;

ALTER TABLE public.item_structures
DROP COLUMN IF EXISTS parent_code,
DROP COLUMN IF EXISTS child_code;


ALTER TABLE public.item_structures
    ADD COLUMN parent_code varchar(255),
ADD COLUMN child_code varchar(255);

UPDATE public.item_structures s
SET parent_code = i.code
    FROM public.items i
WHERE i.id = s.parent_item_id;

UPDATE public.item_structures s
SET child_code = i.code
    FROM public.items i
WHERE i.id = s.child_item_id;

