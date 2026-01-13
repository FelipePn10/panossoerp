ALTER TABLE boms
ADD COLUMN mask BIGINT NOT NULL REFERENCES product_masks(id);

ALTER TABLE bom_items
ADD COLUMN mask_component BIGINT NOT NULL REFERENCES component_masks(id);