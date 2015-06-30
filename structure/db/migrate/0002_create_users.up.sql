CREATE TABLE users (
  id uuid primary key default uuid_generate_v4(),
  name character varying(255),
  email character varying(255),
  created_at timestamp with time zone default current_timestamp,
  updated_at timestamp with time zone default current_timestamp
);
