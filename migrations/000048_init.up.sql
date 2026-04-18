DO $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM public.item_masks
        WHERE item_code !~ '^[0-9]+$'
    ) THEN
        RAISE EXCEPTION 'item_code contains non-numeric values';
END IF;
END;
$$;

ALTER TABLE public.item_masks
DROP CONSTRAINT item_masks_item_id_mask_hash_key;

ALTER TABLE public.item_masks
DROP CONSTRAINT item_masks_item_id_fkey;

ALTER TABLE public.item_masks
ALTER COLUMN item_code TYPE BIGINT USING item_code::BIGINT;

ALTER TABLE public.item_masks
DROP COLUMN item_id;

ALTER TABLE public.item_masks
    ADD CONSTRAINT item_masks_item_code_mask_hash_key UNIQUE (item_code, mask_hash);