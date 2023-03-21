-- create the comments table
CREATE TABLE comments (
    comment_id INT PRIMARY KEY,
    text VARCHAR(255),
    parent_comment_id INT REFERENCES comments(comment_id)
);

-- insert some comments
INSERT INTO comments VALUES (1, 'First comment', NULL);
INSERT INTO comments VALUES (2, 'Reply to first comment', 1);
INSERT INTO comments VALUES (3, 'Reply to first comment', 1);
INSERT INTO comments VALUES (4, 'Reply to second comment', 2);
INSERT INTO comments VALUES (5, 'Reply to fourth comment', 4);
INSERT INTO comments VALUES (6, 'Second comment', NULL);
INSERT INTO comments VALUES (7, 'Reply to sixth comment', 6);
