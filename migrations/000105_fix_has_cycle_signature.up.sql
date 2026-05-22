BEGIN;

-- Remove the legacy varchar version created in migration 000044.
-- After migration 000045 converted parent_code / child_code to BIGINT,
-- this function causes "operator does not exist: bigint = character varying"
-- whenever the Go driver sends the arguments as text.
DROP FUNCTION IF EXISTS has_cycle(VARCHAR, VARCHAR);
DROP FUNCTION IF EXISTS has_cycle(CHARACTER VARYING, CHARACTER VARYING);

COMMIT;
