-- +goose Up
-- +goose StatementBegin
CREATE TABLE albums(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    artist_id UUID REFERENCES users(id),
    title TEXT NOT NULL,
    description TEXT,
    genre VARCHAR(255),
    type VARCHAR(20) NOT NULL CHECK (type IN ('single', 'ep', 'album')),
    cover_image_url TEXT,
    release_date DATE,
    total_tracks INT DEFAULT 0,
    total_duration_seconds BIGINT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE albums;
-- +goose StatementEnd
