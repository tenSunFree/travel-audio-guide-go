create table if not exists profiles (
  id uuid primary key,
  email text,
  display_name text,
  avatar_url text,
  preferred_language text not null default 'zh-TW',
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

create index if not exists idx_profiles_email on profiles(email);
