CREATE TABLE IF NOT EXISTS item_structures (
    id                   BIGSERIAL      PRIMARY KEY,
    parent_item_id       BIGINT         NOT NULL,
    child_item_id        BIGINT         NOT NULL,
    parent_mask          VARCHAR(500)   NULL,
    quantity             NUMERIC(15, 4) NOT NULL DEFAULT 1,
    unit_of_measurement  VARCHAR(20)    NOT NULL DEFAULT 'UN',

    -- Percentual de perda no processo produtivo (ex.: 5.00 = 5%)
    loss_percentage      NUMERIC(5, 2)  NOT NULL DEFAULT 0
        CONSTRAINT chk_loss_percentage CHECK (loss_percentage >= 0 AND loss_percentage <= 100),

    -- Ordem de exibição dentro do pai
    position             INT            NOT NULL DEFAULT 1
        CONSTRAINT chk_position CHECK (position >= 1),

    notes                TEXT           NULL,
    is_active            BOOLEAN        NOT NULL DEFAULT TRUE,
    created_by           UUID           NOT NULL,
    created_at           TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    updated_at           TIMESTAMPTZ    NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_structure_parent_item
        FOREIGN KEY (parent_item_id) REFERENCES items(id) ON DELETE RESTRICT,

    CONSTRAINT fk_structure_child_item
        FOREIGN KEY (child_item_id) REFERENCES items(id) ON DELETE RESTRICT,

    CONSTRAINT chk_no_self_reference
        CHECK (parent_item_id <> child_item_id)
);

-- ÍNDICES DE UNICIDADE PARCIAIS
-- Componentes genéricos: apenas um por par pai+filho ATIVO
CREATE UNIQUE INDEX IF NOT EXISTS uq_item_structure_generic
    ON item_structures (parent_item_id, child_item_id)
    WHERE parent_mask IS NULL
      AND is_active = TRUE;

-- Componentes mascarados: apenas um por par pai+filho+máscara ATIVO
CREATE UNIQUE INDEX IF NOT EXISTS uq_item_structure_masked
    ON item_structures (parent_item_id, child_item_id, parent_mask)
    WHERE parent_mask IS NOT NULL
      AND is_active = TRUE;

CREATE INDEX IF NOT EXISTS idx_item_structures_parent
    ON item_structures (parent_item_id)
    WHERE is_active = TRUE;

CREATE INDEX IF NOT EXISTS idx_item_structures_child
    ON item_structures (child_item_id)
    WHERE is_active = TRUE;

CREATE INDEX IF NOT EXISTS idx_item_structures_mask
    ON item_structures (parent_item_id, parent_mask)
    WHERE is_active = TRUE;
