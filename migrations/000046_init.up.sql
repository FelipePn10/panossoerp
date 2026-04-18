ALTER TABLE public.items
DROP CONSTRAINT IF EXISTS items_health_check;

ALTER TABLE public.items
    ALTER COLUMN health DROP DEFAULT;

ALTER TABLE public.items
ALTER COLUMN health TYPE health_enum
USING (
    CASE health
        WHEN 0 THEN 'ATIVO'::health_enum
        WHEN 1 THEN 'INATIVO'::health_enum
        WHEN 2 THEN 'FANTASMA'::health_enum
        ELSE 'ATIVO'::health_enum
    END
);

ALTER TABLE public.items
    ALTER COLUMN health SET DEFAULT 'ATIVO';
