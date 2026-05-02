ALTER TABLE public.independent_demands
DROP CONSTRAINT IF EXISTS independent_demands_cost_center_id_fkey;

ALTER TABLE public.independent_demands
    RENAME COLUMN cost_center_id TO cost_center_code;

ALTER TABLE public.independent_demands
ALTER COLUMN cost_center_code TYPE int8
USING cost_center_code::int8;

ALTER TABLE public.independent_demands
    ADD CONSTRAINT independent_demands_cost_center_code_fkey
        FOREIGN KEY (cost_center_code)
            REFERENCES public.cost_centers(code);

CREATE INDEX IF NOT EXISTS idx_independent_demands_cost_center_code
    ON public.independent_demands (cost_center_code);
