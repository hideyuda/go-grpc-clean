-- キーワードを管理するテーブル
-- +migrate Up
CREATE TABLE IF NOT EXISTS keywords (
  id INT NOT NULL AUTO_INCREMENT UNIQUE,
  uuid VARCHAR(36) NOT NULL UNIQUE,
  media_id INT NOT NULL,
  keyword VARCHAR(255) NOT NULL,
  bolume INT NOT NULL DEFAULT 0,
  rank INT NOT NULL DEFAULT 0,
  seo_difficulty INT NOT NULL DEFAULT 0,
  url VARCHAR(255) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  PRIMARY KEY (id),
  INDEX idx_keywords_media_id (media_id)
);

ALTER TABLE keywords 
  ADD CONSTRAINT fk_keywords_media_id
  FOREIGN KEY (media_id)
  REFERENCES medias(id)
  ON DELETE CASCADE
  ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE keywords DROP FOREIGN KEY fk_keywords_media_id;

DROP TABLE IF EXISTS keywords;
