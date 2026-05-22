ALTER TABLE maintenance_orders ADD COLUMN IF NOT EXISTS machine_id BIGINT REFERENCES machines(id);
CREATE INDEX IF NOT EXISTS idx_maintenance_orders_machine ON maintenance_orders(machine_id);
