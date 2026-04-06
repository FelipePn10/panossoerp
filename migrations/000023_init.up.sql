DO $$
BEGIN
    CREATE TYPE warehouse_location AS ENUM (
        'LINHA_DE_PRODUCAO',
        'NORMAL'
    );
EXCEPTION
    WHEN duplicate_object THEN NULL;
END
$$;

DO $$
BEGIN
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
EXCEPTION
    WHEN duplicate_object THEN NULL;
END
$$;


ALTER TABLE warehouse
DROP COLUMN IF EXISTS name,
DROP COLUMN IF EXISTS types,
ADD COLUMN IF NOT EXISTS location warehouse_location,
ADD COLUMN IF NOT EXISTS type warehouse_type,
ADD COLUMN IF NOT EXISTS disposition BOOLEAN,
ADD COLUMN IF NOT EXISTS reservation_allowed BOOLEAN;
