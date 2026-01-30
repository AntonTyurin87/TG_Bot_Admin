-- +goose Up
-- +goose StatementBegin


CREATE TABLE IF NOT EXISTS sources (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL UNIQUE,
    step INTEGER NOT NULL,
    type TEXT NOT NULL,
    name_ru TEXT NOT NULL,
    name_eng TEXT,
    author_ru TEXT NOT NULL,
    year TEXT,
    regions TEXT,
    time_periods TEXT,
    description TEXT,
    file_format TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    isSent BOOL NOT NULL
);

CREATE TABLE IF NOT EXISTS files (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    source_id INTEGER,
    file_data BLOB NOT NULL,
    FOREIGN KEY (source_id) REFERENCES sources (id) ON DELETE CASCADE);

-- Создаем индексы отдельно (опционально)
CREATE INDEX IF NOT EXISTS idx_created_at ON sources(created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS files;
DROP TABLE IF EXISTS sources;
-- +goose StatementEnd
