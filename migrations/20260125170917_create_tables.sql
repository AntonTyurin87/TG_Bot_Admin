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
    year INTEGER NOT NULL,
    description TEXT,
    download_url TEXT,
    created_at TEXT,
    isSent BOOL NOT NULL
);

-- Создаем индексы отдельно (опционально)
CREATE INDEX IF NOT EXISTS idx_user_id ON sources(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sources;
-- +goose StatementEnd
