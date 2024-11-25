CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Add a tsvector column for FTS
ALTER TABLE articles ADD COLUMN tsv tsvector;

-- Create a function to update the tsvector column
CREATE FUNCTION articles_tsv_trigger() RETURNS TRIGGER AS $$
BEGIN
  NEW.tsv :=
    setweight(to_tsvector('indonesian', coalesce(NEW.title, '')), 'A') ||
    setweight(to_tsvector('indonesian', coalesce(NEW.content, '')), 'B');
  RETURN NEW;
END
$$ LANGUAGE plpgsql;

-- Create a trigger to update the tsvector column on insert or update
CREATE TRIGGER tsvectorupdate BEFORE INSERT OR UPDATE
ON articles FOR EACH ROW EXECUTE FUNCTION articles_tsv_trigger();

-- Create an index on the tsvector column
CREATE INDEX articles_tsv_idx ON articles USING GIN(tsv);