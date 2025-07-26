-- +migrate Up
CREATE TABLE items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name TEXT NOT NULL,
    code TEXT DEFAULT NULL,
    item_type_id INT,
    status varchar(2) DEFAULT 1  -- Thêm cột status
);

CREATE TABLE resources (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(500) NOT NULL,
    code TEXT DEFAULT NULL,
    resource_type_id INT,
    status varchar(2) DEFAULT 1  -- Thêm cột status
);

CREATE TABLE item_type (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name TEXT NOT NULL,
    status varchar(2) DEFAULT 1  -- Thêm cột status
);

CREATE TABLE resources_type (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name TEXT NOT NULL,
    status varchar(2) DEFAULT 1  -- Thêm cột status
);

CREATE TABLE item_resources (
    id INT AUTO_INCREMENT PRIMARY KEY,
    item_id INT,
    resource_id INT,
    quantity INT,
    status varchar(2) DEFAULT 1 -- Thêm cột status
);

ALTER TABLE items
ADD CONSTRAINT fk_item_type FOREIGN KEY (item_type_id) REFERENCES item_type(id);

ALTER TABLE resources
ADD CONSTRAINT fk_resource_type FOREIGN KEY (resource_type_id) REFERENCES resources_type(id);

ALTER TABLE item_resources
ADD CONSTRAINT fk_item FOREIGN KEY (item_id) REFERENCES items(id);

ALTER TABLE item_resources
ADD CONSTRAINT fk_resource FOREIGN KEY (resource_id) REFERENCES resources(id);
-- +migrate Down
ALTER TABLE item_resources
DROP FOREIGN KEY fk_item;

ALTER TABLE item_resources
DROP FOREIGN KEY fk_resource;

ALTER TABLE resources
DROP FOREIGN KEY fk_resource_type;

ALTER TABLE items
DROP FOREIGN KEY fk_item_type;

DROP TABLE IF EXISTS item_resources;
DROP TABLE IF EXISTS resources;
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS item_type;
DROP TABLE IF EXISTS resources_type;