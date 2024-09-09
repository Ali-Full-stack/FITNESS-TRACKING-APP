-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT id, username, email, profile 
FROM users
ORDER BY username;

-- name: CreateUser :one
INSERT INTO users (username, password_hash, email, profile) 
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
  set 
    username = $2,
    email = $3,
    profile = $4
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: VerifyUserLogin :one
SELECT password_hash FROM users 
WHERE id = $1;

-- name: CreateWorkout :one
INSERT INTO workouts (user_id, name, description)
VALUES ($1, $2, $3)
RETURNING * ;

-- name: GetWorkoutByUserID :many
SELECT id, user_id, name, description, date, created_at, updated_at 
FROM workouts 
WHERE user_id = $1 ;

-- name: GetWorkoutByID :one
SELECT id, user_id, name, description, date, created_at, updated_at 
FROM workouts 
WHERE id = $1 and user_id = $2 ;

-- name: UpdateWorkoutByUserID :exec
UPDATE workouts 
SET name = $3, description = $4  , updated_at = now()
WHERE id = $1 and user_id = $2;

-- name: DeleteWorkout :exec
DELETE FROM workouts 
WHERE id = $1 and user_id = $2 ;