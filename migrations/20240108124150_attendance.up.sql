CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS events (
	"id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
	"name" varchar NOT NULL,
	"start_time" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	"end_time" TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
	"created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	"updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS personnels (
	"id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
	"name" varchar NOT NULL,
	"created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	"updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS attendees (
	"id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
	"name" varchar NOT NULL,
	"program" varchar,
	"year_level" varchar,
	"type" varchar NOT NULL,
	"created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	"updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS attendance (
	"id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
	"event" uuid NOT NULL,
	"attendee" uuid,
	"time_in" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	"scanned_by" uuid NOT NULL,
	"created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	"updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

ALTER TABLE attendance ADD FOREIGN KEY("event") REFERENCES events("id");

ALTER TABLE attendance ADD FOREIGN KEY("scanned_by") REFERENCES personnels("id");

ALTER TABLE attendance ADD FOREIGN KEY("attendee") REFERENCES attendees("id") ON DELETE SET NULL;
