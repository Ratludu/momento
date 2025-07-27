-- name: GetAllProfiles :many
SELECT * FROM profiles;

-- name: AddProfile :one
INSERT INTO profiles (profile_name)
VALUES (
	?
	)
RETURNING *;

-- name: ResetProfiles :exec
DELETE FROM profiles;

-- name: ResetCurrentProfile :exec
UPDATE profiles
SET current_profile = 0;


-- name: SetCurrentProfile :one
UPDATE profiles
SET current_profile = 1
WHERE profile_name = ?
RETURNING *;

-- name: GetCurrentProfile :one
SELECT * FROM profiles
WHERE current_profile = 1;


