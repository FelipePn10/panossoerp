BEGIN;

-- 1. Remover constraints que dependem das colunas
ALTER TABLE public.item_structures
DROP CONSTRAINT IF EXISTS fk_structure_parent_item;

ALTER TABLE public.item_structures
DROP CONSTRAINT IF EXISTS fk_structure_child_item;

ALTER TABLE public.item_structures
DROP CONSTRAINT IF EXISTS chk_no_self_reference;

-- 2. Remover índices relacionados (se existirem)
DROP INDEX IF EXISTS idx_item_structures_parent;
DROP INDEX IF EXISTS idx_item_structures_child;
DROP INDEX IF EXISTS idx_item_structures_mask;
DROP INDEX IF EXISTS uq_item_structure_generic;
DROP INDEX IF EXISTS uq_item_structure_masked;

-- 3. Dropar colunas
ALTER TABLE public.item_structures
DROP COLUMN IF EXISTS parent_item_id;

ALTER TABLE public.item_structures
DROP COLUMN IF EXISTS child_item_id;

COMMIT;