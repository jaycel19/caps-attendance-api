CREATE TABLE IF NOT EXISTS personnels (
	"id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
	"name" varchar NOT NULL,
	"created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	"updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
