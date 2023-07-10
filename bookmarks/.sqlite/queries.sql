-- name: CreateBookmark :exec
INSERT
    or IGNORE INTO bookmarks (path)
VALUES (?);
-- name: GetBookmarks :many
SELECT path
FROM bookmarks;
-- name: DeleteBookmark :exec
DELETE FROM bookmarks
WHERE path = ?;