CREATE TYPE warehouse_location AS ENUM (
    'LINHA_DE_PRODUCAO',
    'NORMAL'
);

CREATE TYPE warehouse_type AS ENUM (
    'INTERNO',
    'EXTERNO',
    'ASSISTENCIA',
    'REJEICAO',
    'INSPECAO',
    'RESERVA',
    'TRANSITO',
    'ESPECIAL'
);

ALTER TABLE warehouse
DROP COLUMN IF EXISTS name,
ADD COLUMN location warehouse_location,
ADD COLUMN type warehouse_type,
ADD COLUMN disposition BOOLEAN,
ADD COLUMN reservation_allowed BOOLEAN;
