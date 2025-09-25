-- +gooseUp
CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(32) UNIQUE NOT NULL,
    password VARCHAR(64) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);


CREATE TABLE rooms (
    id UUID PRIMARY KEY,
    name VARCHAR(32) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    creator_id UUID NOT NULL REFERENCES users(id),
    owner_id UUID NOT NULL REFERENCES users(id)
);

CREATE TABLE messages (
    id UUID PRIMARY KEY,
    data VARCHAR(1000) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    sender_id UUID NOT NULL REFERENCES users(id),
    room_id UUID NOT NULL REFERENCES rooms(id)
);

-- +goose Down
DROP TABLE users;
DROP TABLE rooms;
DROP TABLE messages;