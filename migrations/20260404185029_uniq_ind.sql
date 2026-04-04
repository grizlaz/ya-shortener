-- +goose Up
create unique index if not exists idx_shortening_original_url on shortening(original_url);

-- +goose Down
drop index if exists idx_shortening_original_url;
