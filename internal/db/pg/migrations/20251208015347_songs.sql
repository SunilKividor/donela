-- +goose Up
-- +goose StatementBegin
CREATE TABLE songs (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    artist_id UUID NOT NULL REFERENCES users(id),
    album_id UUID NOT NULL REFERENCES albums(id),
    title TEXT NOT NULL,
    genre VARCHAR(255),
    audio_file_url TEXT,
    cover_image_url TEXT,
    status VARCHAR(255) NOT NULL,
    duration_seconds BIGINT DEFAULT 0,
    bitrate BIGINT DEFAULT 0,
    file_format VARCHAR(255),
    file_size_bytes BIGINT DEFAULT 0,
    release_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE songs;
-- +goose StatementEnd
 