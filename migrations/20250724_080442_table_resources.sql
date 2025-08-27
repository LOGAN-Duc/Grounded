-- +migrate Up
CREATE TABLE item_type (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name TEXT NOT NULL,
    status VARCHAR(2) DEFAULT '1'
);

CREATE TABLE resources_type (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name TEXT NOT NULL,
    status VARCHAR(2) DEFAULT '1'
);

CREATE TABLE items (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name TEXT NOT NULL,
    code TEXT DEFAULT NULL,
    url TEXT,
    item_type_id BIGINT UNSIGNED,
    status VARCHAR(2) DEFAULT '1',
    CONSTRAINT fk_item_type FOREIGN KEY (item_type_id) REFERENCES item_type(id)
);

CREATE TABLE resources (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(500) NOT NULL,
    code TEXT DEFAULT NULL,
    url TEXT,
    resource_type_id BIGINT UNSIGNED,
    status VARCHAR(2) DEFAULT '1',
    CONSTRAINT fk_resource_type FOREIGN KEY (resource_type_id) REFERENCES resources_type(id)
);

CREATE TABLE item_resources (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    item_id BIGINT UNSIGNED,
    resource_id BIGINT UNSIGNED,
    quantity INT,
    status VARCHAR(2) DEFAULT '1',
    CONSTRAINT fk_item FOREIGN KEY (item_id) REFERENCES items(id) ON DELETE CASCADE,
    CONSTRAINT fk_resource FOREIGN KEY (resource_id) REFERENCES resources(id) ON DELETE CASCADE
);

CREATE TABLE biomes (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    url TEXT,
    description TEXT,
    status VARCHAR(2) DEFAULT '1',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE landmarks (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    biome_id BIGINT UNSIGNED NOT NULL,
    name VARCHAR(255) NOT NULL,
    url TEXT,
    status VARCHAR(2) DEFAULT '1',
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_landmarks_biomes FOREIGN KEY (biome_id) REFERENCES biomes(id) ON DELETE CASCADE
);

-- Bảng quan hệ N-N giữa biomes và resources
CREATE TABLE biome_resources (
    biome_id BIGINT UNSIGNED NOT NULL,
    resource_id BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (biome_id, resource_id),
    CONSTRAINT fk_biome_resources_biome FOREIGN KEY (biome_id) REFERENCES biomes(id) ON DELETE CASCADE,
    CONSTRAINT fk_biome_resources_resource FOREIGN KEY (resource_id) REFERENCES resources(id) ON DELETE CASCADE
);

-- Thêm quan hệ 1-N giữa items và biomes
ALTER TABLE items
ADD COLUMN biome_id BIGINT UNSIGNED NULL,
ADD CONSTRAINT fk_items_biomes FOREIGN KEY (biome_id) REFERENCES biomes(id) ON DELETE SET NULL;

-- +migrate Down
ALTER TABLE items DROP FOREIGN KEY fk_items_biomes;
ALTER TABLE items DROP COLUMN biome_id;

DROP TABLE IF EXISTS biome_resources;
DROP TABLE IF EXISTS landmarks;
DROP TABLE IF EXISTS biomes;
DROP TABLE IF EXISTS item_resources;
DROP TABLE IF EXISTS resources;
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS item_type;
DROP TABLE IF EXISTS resources_type;
