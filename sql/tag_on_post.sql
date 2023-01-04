CREATE TABLE tag_on_post (
    post_id integer REFERENCES post(id),
    tag_id integer REFERENCES tag(id),
    PRIMARY KEY (post_id, tag_id)
);