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

-- name: GetSessionsWithProfile :many
SELECT 
	s.id AS session_id,
	s.start AS session_start,
	s.end AS session_end,
	p.id AS profile_id,
	p.profile_name
FROM
	sessions AS s
LEFT JOIN 
	profiles AS p ON s.profile_id = p.id;



