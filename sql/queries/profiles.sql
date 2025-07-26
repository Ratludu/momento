-- name: GetAllProfiles :many
SELECT * FROM profiles;

-- name: AddProfile :one
INSERT INTO profiles (profile_name)
VALUES (
	?
	)
RETURNING *;
