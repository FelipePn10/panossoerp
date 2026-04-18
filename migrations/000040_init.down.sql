ALTER TABLE item_structures
DROP COLUMN health;

ALTER TABLE item_structures
    ALTER COLUMN unit_of_measurement DROP DEFAULT;

ALTER TABLE item_structures
ALTER COLUMN unit_of_measurement TYPE varchar(20)
USING unit_of_measurement::text;

ALTER TABLE item_structures
    ALTER COLUMN unit_of_measurement SET DEFAULT 'UN';

DROP TYPE health_enum;
DROP TYPE unit_of_measurement_enum;