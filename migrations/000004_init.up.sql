CREATE TABLE votes (
    id SERIAL PRIMARY KEY,
    linkid TEXT NOT NULL,
    userid INT NOT NULL,
    votetype INT NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (userid, linkid)
);
