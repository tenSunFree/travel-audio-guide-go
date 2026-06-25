-- name: GetProfileByID :one
select id, email, display_name, avatar_url, preferred_language, created_at, updated_at
from profiles
where id = $1;

-- name: CreateProfile :one
insert into profiles (id, email, preferred_language)
values ($1, $2, 'zh-TW')
returning id, email, display_name, avatar_url, preferred_language, created_at, updated_at;

-- name: UpdateProfile :one
update profiles
set display_name       = coalesce(sqlc.narg(display_name), display_name),
    avatar_url         = coalesce(sqlc.narg(avatar_url), avatar_url),
    preferred_language = coalesce(sqlc.narg(preferred_language), preferred_language),
    updated_at         = now()
where id = $1
returning id, email, display_name, avatar_url, preferred_language, created_at, updated_at;
