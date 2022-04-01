CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name TEXT,
    email TEXT
);

CREATE TABLE posts (
    post_id SERIAL PRIMARY KEY,
    title TEXT,
    content TEXT,
    author_id INT,
    CONSTRAINT fk_author
        FOREIGN KEY(author_id)
            REFERENCES users(user_id)
);

CREATE TABLE comments (
    comment_id SERIAL PRIMARY KEY,
    content TEXT,
    author_id INT,
    post_id INT,
    CONSTRAINT fk_post
        FOREIGN KEY(post_id)
            REFERENCES posts(post_id),
    CONSTRAINT fk_author
        FOREIGN KEY(author_id)
            REFERENCES users(user_id)
);

CREATE TABLE votes (
    vote_id SERIAL PRIMARY KEY,
    voter_id INT,
    post_id INT,
    CONSTRAINT fk_post
        FOREIGN KEY(post_id)
            REFERENCES posts(post_id),
    CONSTRAINT fk_voter
        FOREIGN KEY(voter_id)
            REFERENCES users(user_id)
);

INSERT INTO users (name, email) VALUES ('Obi-Wan Kenobi', 'obi-wan@mail.com');

INSERT INTO posts (title, content, author_id) VALUES ('Initial post', 'Here be description.', 1);

INSERT INTO comments (content, author_id, post_id) VALUES ('Here be comment.', 1, 1);
INSERT INTO comments (content, author_id, post_id) VALUES ('Here be another comment.', 1, 1);

INSERT INTO votes (voter_id, post_id) VALUES (1, 1)
