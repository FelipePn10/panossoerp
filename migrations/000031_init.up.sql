ALTER TABLE warehouse
ADD COLUMN IF NOT EXISTS created_by UUID,
ADD CONSTRAINT fk_warehouse_created_by
FOREIGN KEY (created_by)
REFERENCES users(id)
ON DELETE SET NULL;
