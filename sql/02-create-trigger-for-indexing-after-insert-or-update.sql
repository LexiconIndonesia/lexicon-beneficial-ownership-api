-- create postgres trigger after insert or update on table cases for indexing tsvector column for full text search only for affected row

CREATE TRIGGER cases_fulltext_search_index_update BEFORE INSERT OR UPDATE
ON cases FOR EACH ROW EXECUTE FUNCTION
tsvector_update_trigger('fulltext_search_index', 'pg_catalog.english', 'subject', 'summary');




