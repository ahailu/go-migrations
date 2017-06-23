CREATE TABLE jobs (
	id SERIAL,
	name text,
	webhook text,
	created_at timestamptz,
	removed_at timestamptz
);
