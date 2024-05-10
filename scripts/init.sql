CREATE TABLE IF NOT EXISTS "rescues" (
  "id"           SERIAL PRIMARY KEY NOT NULL, 
  "rescue_count" integer,
  "rescue_type"  smallint NOT NULL,
  "urgency"      smallint NOT NULL,
  "description"  text,
  "street"       varchar(100),
	"number"       integer,
	"neighborhood" varchar(100),
	"zip_code"     varchar(20),
	"complement"   varchar(100),
  "created_at"   timestamp with time zone DEFAULT now() NOT NULL,
  "updated_at"   timestamp with time zone DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS "find_animals" (
  "id"           SERIAL PRIMARY KEY NOT NULL, 
  "specie"       varchar(50) NOT NULL,
  "race"         varchar(50),
  "gender"       smallint NOT NULL, 
  "description"  text,
  "street"       varchar(100),
	"number"       integer,
	"neighborhood" varchar(100),
	"zip_code"     varchar(20),
	"complement"   varchar(100),
  "created_at"   timestamp with time zone DEFAULT now() NOT NULL,
  "updated_at"   timestamp with time zone DEFAULT now() NOT NULL
);