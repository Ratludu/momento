-- name: CreateSession :one
INSERT INTO sessions (profile_id, note, start)
VALUES (
	?,
	?,
	?
	)
RETURNING *;

-- name: GetSessions :many
SELECT * FROM sessions WHERE end = "EMPTY";

-- name: ResetSessions :exec
DELETE FROM sessions;

-- name: CloseSession :one
UPDATE sessions
SET updated_at = DATETIME('now'),
end = ?
WHERE end = "EMPTY"
RETURNING *;
