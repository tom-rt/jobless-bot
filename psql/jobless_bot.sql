DROP TABLE IF EXISTS chan_user;

CREATE TABLE chan_user
(
    id BIGSERIAL PRIMARY KEY,
    name text UNIQUE NOT NULL,
    sent_messages_count INTEGER DEFAULT 1
);
