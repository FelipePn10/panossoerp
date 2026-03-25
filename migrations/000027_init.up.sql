CREATE TABLE enterprise (
    id BIGSERIAL PRIMARY KEY,
    code INT NOT NULL UNIQUE,
    name VARCHAR(35) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE groups (
    id BIGSERIAL PRIMARY KEY,
    code INT NOT NULL UNIQUE,
    description VARCHAR(180) NOT NULL,
    enterprise_id BIGINT NOT NULL REFERENCES enterprise(id),
    created_by UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
