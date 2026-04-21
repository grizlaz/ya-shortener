-- +goose Up
alter table shortening add column if not exists user_id uuid;
create index if not exists idx_shortening_user_id on shortening(user_id);

-- +goose Down
alter table shortening drop column if exists user_id;
drop index if exists idx_shortening_user_id;
