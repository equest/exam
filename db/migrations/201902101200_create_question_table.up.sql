-- TABLE DEFINITION --
CREATE TABLE "question" ( 
	"id" serial PRIMARY KEY,
	"owner" varchar(255) NOT NULL, 
	"type" varchar(255) NOT NULL, 
	"kind" varchar(255) NOT NULL,
	"text" varchar NOT NULL, 
	"metadata" jsonb NULL,
	"data" jsonb NULL,
	"createdAt" timestamptz NOT NULL,
	"modifiedAt" timestamptz NOT NULL,
	CONSTRAINT "bpjs_pkey" PRIMARY KEY ("nik")
)
WITH (
	OIDS=FALSE
) ;
CREATE INDEX "nik-checker_bpjs_nik" ON "bpjs" USING btree ("nik") ;
-- Indices --
