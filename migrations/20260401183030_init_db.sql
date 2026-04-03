-- +goose Up
create table if not exists shortening (
	id int primary key,
	short_url text not null,
	original_url text not null
);

create index if not exists idx_shortening_short_url on shortening(short_url);

-- +goose Down
drop table if exists shortening;
