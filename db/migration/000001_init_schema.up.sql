CREATE TABLE IF NOT EXISTS users(
  id text not null,
  name text not null unique,
  phone_number text not null unique,
  otp text ,
  otp_expiration_time timestamptz,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now(),

  PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_users_name ON users (name);
CREATE INDEX IF NOT EXISTS idx_users_email ON users (phone_number);
