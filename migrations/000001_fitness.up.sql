CREATE TABLE "users" (
  "id" serial PRIMARY KEY ,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "password_hash" varchar NOT NULL,
  "profile" jsonb DEFAULT '{}'
);