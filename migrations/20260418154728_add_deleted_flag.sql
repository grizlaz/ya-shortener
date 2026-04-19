-- +goose Up
alter table shortening add column if not exists is_deleted bool default false;

-- +goose Down
alter table shortening drop column if exists is_deleted;
