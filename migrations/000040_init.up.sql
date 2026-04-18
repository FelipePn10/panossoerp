CREATE TYPE unit_of_measurement_enum AS ENUM (
    'MM',
    'CM',
    'M',
    'IN',
    'KG',
    'M2',
    'M3',
    'UN',
    'MICROMETRO',
    'TONELADA'
);

CREATE TYPE health_enum AS ENUM (
    'ATIVO',
    'INATIVO',
    'FANTASMA'
);

ALTER TABLE item_structures
    ALTER COLUMN unit_of_measurement DROP DEFAULT;

ALTER TABLE item_structures
ALTER COLUMN unit_of_measurement TYPE unit_of_measurement_enum
USING unit_of_measurement::unit_of_measurement_enum;

ALTER TABLE item_structures
    ALTER COLUMN unit_of_measurement SET DEFAULT 'UN';

ALTER TABLE item_structures
    ADD COLUMN health health_enum NOT NULL DEFAULT 'ATIVO';