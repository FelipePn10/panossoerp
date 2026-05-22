BEGIN;
DROP TABLE IF EXISTS production_sequences CASCADE;
ALTER TABLE production_appointments DROP COLUMN IF EXISTS operation_id;
DROP TABLE IF EXISTS production_order_operations CASCADE;
ALTER TABLE production_orders DROP COLUMN IF EXISTS route_id;
COMMIT;
