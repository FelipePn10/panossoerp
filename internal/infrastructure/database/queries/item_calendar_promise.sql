-- name: UpsertItemCalendarDay :one
INSERT INTO item_calendar_promises (item_code, mask, year, month, day, is_workday, description)
VALUES ($1, $2, $3, $4, $5, $6, $7)
    ON CONFLICT (item_code, mask, year, month, day) DO UPDATE SET is_workday = EXCLUDED.is_workday, description = EXCLUDED.description
                                                           RETURNING *;

-- name: GetItemCalendarDay :one
SELECT * FROM item_calendar_promises WHERE item_code = $1 AND mask = $2 AND year = $3 AND month = $4 AND day = $5;

-- name: GetItemWorkdaysInMonth :many
SELECT * FROM item_calendar_promises WHERE item_code = $1 AND mask = $2 AND year = $3 AND month = $4 AND is_workday = TRUE ORDER BY day;

-- name: ListItemCalendarMonth :many
SELECT * FROM item_calendar_promises WHERE item_code = $1 AND mask = $2 AND year = $3 AND month = $4 ORDER BY day;

-- name: DeleteItemCalendarDay :exec
DELETE FROM item_calendar_promises WHERE item_code = $1 AND mask = $2 AND year = $3 AND month = $4 AND day = $5;
