CREATE TABLE IF NOT EXISTS employee (
    id BIGSERIAL PRIMARY KEY,
    enterprise_id INT NOT NULL REFERENCES enterprise(id) ON DELETE CASCADE,
    code INT NOT NULL,
    description VARCHAR(180),
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
