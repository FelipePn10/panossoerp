BEGIN;
DROP TABLE IF EXISTS maintenance_orders;
DROP TABLE IF EXISTS maintenance_plans;
DROP TYPE IF EXISTS maintenance_order_status_enum;
DROP TYPE IF EXISTS maintenance_frequency_enum;
COMMIT;
