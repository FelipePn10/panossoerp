ALTER TABLE public.items
    ALTER COLUMN warehouse_unit_of_measurement DROP DEFAULT;

ALTER TABLE public.items
ALTER COLUMN warehouse_unit_of_measurement TYPE TEXT;

ALTER TABLE public.items
ALTER COLUMN warehouse_unit_of_measurement TYPE unit_of_measurement_enum
USING (
    CASE warehouse_unit_of_measurement
        WHEN '0' THEN 'MM'::unit_of_measurement_enum
        WHEN '1' THEN 'CM'::unit_of_measurement_enum
        WHEN '2' THEN 'M'::unit_of_measurement_enum
        WHEN '3' THEN 'IN'::unit_of_measurement_enum
        WHEN '4' THEN 'KG'::unit_of_measurement_enum
        WHEN '5' THEN 'M2'::unit_of_measurement_enum
        WHEN '6' THEN 'M3'::unit_of_measurement_enum
        WHEN '7' THEN 'UN'::unit_of_measurement_enum
        WHEN '8' THEN 'MICROMETRO'::unit_of_measurement_enum
        WHEN '9' THEN 'TONELADA'::unit_of_measurement_enum
        ELSE 'UN'::unit_of_measurement_enum
    END
);

ALTER TABLE public.items
    ALTER COLUMN warehouse_unit_of_measurement SET DEFAULT 'UN';
