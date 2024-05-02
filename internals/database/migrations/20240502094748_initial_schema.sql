-- Create "todos" table
CREATE TABLE "todos" (
  "id" bigserial NOT NULL,
  "title" text NOT NULL,
  "completed" boolean NOT NULL DEFAULT false,
  PRIMARY KEY ("id")
);
