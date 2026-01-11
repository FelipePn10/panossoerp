ALTER TABLE bom_items
DROP CONSTRAINT bom_items_bom_id_component_id_operation_sequence_key;

ALTER TABLE bom_items
DROP COLUMN operation_sequence;
