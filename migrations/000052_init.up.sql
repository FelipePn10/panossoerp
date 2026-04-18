ALTER TABLE public.item_structures
    ADD CONSTRAINT uq_item_structure_unique
        UNIQUE (parent_code, child_code, parent_mask);

CREATE INDEX idx_item_structures_parent_mask_active
    ON public.item_structures (parent_code, parent_mask)
    WHERE is_active = TRUE;

CREATE INDEX idx_item_structures_generic
    ON public.item_structures (parent_code)
    WHERE parent_mask IS NULL AND is_active = TRUE;

CREATE INDEX idx_item_structures_masked
    ON public.item_structures (parent_code, parent_mask)
    WHERE parent_mask IS NOT NULL AND is_active = TRUE;