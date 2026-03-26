ALTER TABLE enterprise
ADD COLUMN created_by UUID,
ADD CONSTRAINT fk_enterprise_created_by
FOREIGN KEY (created_by)
REFERENCES users(id)
ON DELETE SET NULL;
