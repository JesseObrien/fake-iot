CREATE DATABASE fakeiot;

-- Typically i would not give this user ALL privileges on the database
-- It would be set up to only use what it needs. So for instance having
-- a write account vs read account, different databases or tables, etc
GRANT ALL PRIVILEGES ON DATABASE fakeiot TO testuser;

\c fakeiot

CREATE TABLE IF NOT EXISTS account_logins (
  user_id varchar(36) NOT NULL,
  account_id varchar(36) NOT NULL,
  timestamp timestamp
);

-- Need an index on account_id for lookups by account_id
CREATE INDEX idx_account_logins_account_id ON account_logins(account_id);

CREATE TYPE plan_type AS ENUM ('standard', 'enterprise');

CREATE TABLE IF NOT EXISTS accounts (
  id varchar(36) NOT NULL,
  user_id varchar(36) NOT NULL,
  plan_type plan_type NOT NULL
);

-- Insert a default user into the accounts table
INSERT INTO accounts VALUES ('47f3c307-6344-49e7-961c-ea200e950a89', 'de7169a0-1ca1-4f18-8fb8-29d3a7cafd30', 'standard')