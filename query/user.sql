
-- name: CreateNewUser :one
-- CreateNewUser will create a new user
INSERT INTO users ("name", "email", "password")
VALUES (
    $1, $2, $3
)
RETURNING id
;

-- name: CheckUserExist :one
-- CheckUserExist will check if the user exist, retuns true if the exists
SELECT Cast(count(1)::int as boolean) FROM users
WHERE email=$1
LIMIT 1;
