-- Write your migrate up statements here

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE summoner_records (
	record_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	record_date TIMESTAMP,

	account_id TEXT NOT NULL,
	profile_icon_id INT NOT NULL,
	revision_date BIGINT NOT NULL,
	name TEXT NOT NULL,
	id TEXT NOT NULL,
	puuid TEXT NOT NULL,
	summoner_level BIGINT NOT NULL
);

---- create above / drop below ----

DROP TABLE summoner_records;

DROP EXTENSION IF EXISTS "uuid-ossp";

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
