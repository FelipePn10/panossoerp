-- 1. Remover coluna inválida/legada
ALTER TABLE warehouse
DROP COLUMN IF EXISTS types,
DROP COLUMN IF EXISTS active;

-- 2. Padronizar naming
ALTER TABLE warehouse
RENAME COLUMN reservation_allowed TO reservations_allowed;

-- 3. Garantir colunas de auditoria
ALTER TABLE warehouse
ADD COLUMN IF NOT EXISTS created_by UUID;

ALTER TABLE warehouse
ADD COLUMN IF NOT EXISTS created_at TIMESTAMPTZ NOT NULL DEFAULT now();

-- 4. Garantir constraints básicas
ALTER TABLE warehouse
ALTER COLUMN description SET NOT NULL;

ALTER TABLE warehouse
ALTER COLUMN code SET NOT NULL;
