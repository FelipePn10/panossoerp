BEGIN;

ALTER TABLE public.item_structures
DROP CONSTRAINT IF EXISTS fk_structure_parent_code;

ALTER TABLE public.item_structures
DROP CONSTRAINT IF EXISTS fk_structure_child_code;

ALTER TABLE public.item_structures
    ALTER COLUMN parent_code DROP DEFAULT;

ALTER TABLE public.item_structures
    ALTER COLUMN child_code DROP DEFAULT;

ALTER TABLE public.item_structures
ALTER COLUMN parent_code TYPE BIGINT USING parent_code::BIGINT;

ALTER TABLE public.item_structures
ALTER COLUMN child_code TYPE BIGINT USING child_code::BIGINT;

ALTER TABLE public.item_structures
    ADD CONSTRAINT fk_structure_parent_code
        FOREIGN KEY (parent_code) REFERENCES public.items(code);

ALTER TABLE public.item_structures
    ADD CONSTRAINT fk_structure_child_code
        FOREIGN KEY (child_code) REFERENCES public.items(code);

COMMIT;