
CREATE TABLE cases (
    id VARCHAR(26) PRIMARY KEY,
    "subject" VARCHAR(255) NOT NULL,
    subject_type VARCHAR(255) NOT NULL,
    person_in_charge VARCHAR(255),
    benificiary_ownership VARCHAR(255),
    "date" DATE,
    decision_number VARCHAR(255),
    source VARCHAR(255),
    link VARCHAR(255),
    nation VARCHAR(255),
    punishment_duration VARCHAR(255),
    "type" VARCHAR(255),
    "year" VARCHAR(4),
    summary TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- add index to SubjectType, Year, Type, Nation Postgresql
CREATE INDEX idx_search_filter ON cases (subject_type, year, type, nation);


-- alter table to add search index column for fast full text search
ALTER TABLE cases ADD COLUMN fulltext_search_index tsvector GENERATED ALWAYS AS (to_tsvector('english', subject || ' ' || summary)) STORED;

-- create fulltext search index on Subject and Summary Postgresql
CREATE INDEX idx_search_fulltext ON cases USING GIN (fulltext_search_index);
