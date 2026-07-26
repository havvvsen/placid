-- name: AddTrack :exec
INSERT INTO tracks (name, mood, audio_url, bg_url)
    VALUES ($1, $2, $3, $4);

-- name: GetTracks :many
SELECT
    *
FROM
    tracks;

-- name: GetTrack :one
SELECT
    *
FROM
    tracks
WHERE
    id = $1;

-- name: DeleteTrack :exec
DELETE FROM tracks
WHERE id = $1;

-- name: UpdateTrackAudioUrl :exec
UPDATE
    tracks
SET
    audio_url = $2
WHERE
    id = $1;

-- name: UpdateTrackBgUrl :exec
UPDATE
    tracks
SET
    bg_url = $2
WHERE
    id = $1;

-- name: UpdateTrackMood :exec
UPDATE
    tracks
SET
    mood = $2
WHERE
    id = $1;
