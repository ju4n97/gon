-- Create "todo" table
CREATE TABLE "todo" (
  "id" bigserial NOT NULL,
  "title" text NOT NULL,
  "is_completed" boolean NOT NULL DEFAULT false,
  PRIMARY KEY ("id")
);
