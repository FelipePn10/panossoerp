BEGIN;

DROP TABLE IF EXISTS route_operation_network CASCADE;
DROP TABLE IF EXISTS route_operations CASCADE;
DROP TABLE IF EXISTS manufacturing_routes CASCADE;
DROP TABLE IF EXISTS operations CASCADE;

DROP TYPE IF EXISTS route_op_situation_enum;
DROP TYPE IF EXISTS route_situation_enum;
DROP TYPE IF EXISTS operation_situation_enum;
DROP TYPE IF EXISTS operation_origin_enum;

COMMIT;
