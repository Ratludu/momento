-- name: CreateSession :one
INSERT INTO sessions (profile_id, note, start)
VALUES (
	?,
	?,
	?
	)
RETURNING *;
