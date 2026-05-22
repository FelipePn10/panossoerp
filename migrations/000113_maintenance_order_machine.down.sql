DROP INDEX IF EXISTS idx_maintenance_orders_machine;
ALTER TABLE maintenance_orders DROP COLUMN IF EXISTS machine_id;
