-- bookmarks
CREATE TABLE IF NOT EXISTS bookmarks (
    id integer PRIMARY KEY,
    path TEXT NOT NULL UNIQUE ON CONFLICT IGNORE
);
-- CREATE UNIQUE INDEX IF NOT EXISTS index_bookmarks_on_path_unique ON bookmarks (path);
-- set WAL mode
PRAGMA journal_mode = WAL;