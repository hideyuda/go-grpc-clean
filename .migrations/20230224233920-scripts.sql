-- スクリプトを管理するテーブル
-- +migrate Up
CREATE TABLE IF NOT EXISTS scripts (
  id INT NOT NULL AUTO_INCREMENT UNIQUE,
  uuid VARCHAR(36) NOT NULL UNIQUE,
  media_id INT NOT NULL,
  title VARCHAR(255) NOT NULL,
  content VARCHAR(255) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  PRIMARY KEY (id),
  INDEX idx_scripts_media_id (media_id)
);

ALTER TABLE scripts 
  ADD CONSTRAINT fk_scripts_media_id
  FOREIGN KEY (media_id)
  REFERENCES medias(id)
  ON DELETE CASCADE
  ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE scripts DROP FOREIGN KEY fk_scripts_media_id;

DROP TABLE IF EXISTS scripts;
