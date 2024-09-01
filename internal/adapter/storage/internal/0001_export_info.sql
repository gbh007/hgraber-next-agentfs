-- +goose Up

CREATE TABLE export_infos (
    book_id TEXT NOT NULL,
    book_url TEXT,
    relative_path TEXT NOT NULL,
    exported_at BIGINT,
    PRIMARY KEY(book_id, relative_path)
);
