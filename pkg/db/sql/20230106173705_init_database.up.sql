BEGIN;
  CREATE TABLE InviteCard(
    Id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    BusinessId UUID,

    SlugCode TEXT NOT NULL UNIQUE,
    PassCode TEXT,

    Business JSONB,
    Social JSONB,
    Content JSONB
  )
COMMIT;
