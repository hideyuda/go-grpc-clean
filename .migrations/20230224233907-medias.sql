-- メディアを管理するテーブル
-- +migrate Up
CREATE TABLE IF NOT EXISTS medias (
  id INT NOT NULL AUTO_INCREMENT UNIQUE,
  uuid VARCHAR(36) NOT NULL UNIQUE,
  user_id INT NOT NULL,
  name VARCHAR(100) NOT NULL,
  type INT NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  PRIMARY KEY (id),
  INDEX idx_medias_user_id (user_id)
);

ALTER TABLE medias 
  ADD CONSTRAINT fk_medias_user_id
  FOREIGN KEY (user_id)
  REFERENCES users(id)
  ON DELETE CASCADE
  ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE medias DROP FOREIGN KEY fk_medias_user_id;

DROP TABLE IF EXISTS medias;
