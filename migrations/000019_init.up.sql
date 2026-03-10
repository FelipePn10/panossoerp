ALTER TABLE items
ALTER COLUMN type TYPE SMALLINT USING type::smallint;

ALTER TABLE items
ALTER COLUMN status TYPE SMALLINT USING status::smallint;

ALTER TABLE items
ALTER COLUMN health TYPE SMALLINT USING health::smallint;

ALTER TABLE items
ADD CONSTRAINT items_type_check
CHECK (type BETWEEN 0 AND 5);

ALTER TABLE items
ADD CONSTRAINT items_status_check
CHECK (status BETWEEN 0 AND 3);

ALTER TABLE items
ADD CONSTRAINT items_health_check
CHECK (health BETWEEN 0 AND 3);