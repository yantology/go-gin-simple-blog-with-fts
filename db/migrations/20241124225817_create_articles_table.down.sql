-- Drop the index on the tsvector column
DROP INDEX IF EXISTS articles_tsv_idx;

-- Drop the trigger
DROP TRIGGER IF EXISTS tsvectorupdate ON articles;

-- Drop the function
DROP FUNCTION IF EXISTS articles_tsv_trigger();

-- Drop the table
DROP TABLE IF EXISTS articles;