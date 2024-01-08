CREATE TABLE IF NOT EXISTS instructors (
	"id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
	"name" varchar NOT NULl,
	"created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	"updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
