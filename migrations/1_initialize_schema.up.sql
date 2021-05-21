CREATE TABLE candidate (
    id CHAR(36) NOT NULL,
    name VARCHAR(255),
    address VARCHAR(255),
    phone VARCHAR(255),
    PRIMARY KEY(id)
);

CREATE TABLE status (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    candidate_id CHAR(36) NOT NULL,
    type TINYINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    PRIMARY KEY(id),
    CONSTRAINT FK_candidate_status FOREIGN KEY (candidate_id) REFERENCES candidate(id) ON UPDATE CASCADE, NO DELETE CASCADE
);