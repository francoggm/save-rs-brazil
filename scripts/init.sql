CREATE TABLE IF NOT EXISTS "animals" (
  "id"           SERIAL PRIMARY KEY NOT NULL, 
  "specie"       varchar(50) NOT NULL,
  "race"         varchar(50),
  "gender"       smallint NOT NULL, 
  "state"        smallint NOT NULL,
  "description"  text,
  "street"       varchar(100),
	"number"       integer,
	"neighborhood" varchar(100) NOT NULL,
	"zip_code"     varchar(20),
	"complement"   varchar(100),
  "phone"        varchar(20),
  "created_at"   timestamp with time zone DEFAULT now() NOT NULL,
  "updated_at"   timestamp with time zone DEFAULT now() NOT NULL
);