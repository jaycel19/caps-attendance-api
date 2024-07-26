CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create events table
CREATE TABLE IF NOT EXISTS events (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "name" varchar NOT NULL,
    "start_time" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    "end_time" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create personnels table
CREATE TABLE IF NOT EXISTS personnels (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "name" varchar NOT NULL,
    "username" varchar NOT NULL,
    "password" varchar NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create attendees table
CREATE TABLE IF NOT EXISTS attendees (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "name" varchar NOT NULL,
    "program" varchar,
    "year_level" varchar,
    "type" varchar NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create attendance table
CREATE TABLE IF NOT EXISTS attendance (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "event" uuid NOT NULL,
    "attendee" uuid,
    "time_in" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    "scanned_by" uuid NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    FOREIGN KEY ("event") REFERENCES events("id") ON DELETE CASCADE,
    FOREIGN KEY ("attendee") REFERENCES attendees("id") ON DELETE CASCADE,
    FOREIGN KEY ("scanned_by") REFERENCES personnels("id") ON DELETE CASCADE
);

-- Ensure only deleting from the attendees table cascades to attendance
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_constraint WHERE conname = 'fk_attendance_attendee') THEN
        ALTER TABLE attendance DROP CONSTRAINT fk_attendance_attendee;
    END IF;
    ALTER TABLE attendance ADD CONSTRAINT fk_attendance_attendee
        FOREIGN KEY ("attendee") REFERENCES attendees("id") ON DELETE CASCADE;
END $$;
