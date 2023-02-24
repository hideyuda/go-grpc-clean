-- 競合メディアを管理するテーブル
-- +migrate Up
CREATE TABLE IF NOT EXISTS competitors (
  id INT NOT NULL AUTO_INCREMENT UNIQUE,
  uuid VARCHAR(36) NOT NULL UNIQUE,
  media_id INT NOT NULL,
  url VARCHAR(255) NOT NULL,
  monthly_traffic INT NOT NULL DEFAULT 0,
  monthly_ad_revenue INT NOT NULL DEFAULT 0,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  PRIMARY KEY (id),
  INDEX idx_competitors_media_id (media_id)
);

ALTER TABLE competitors 
  ADD CONSTRAINT fk_competitors_media_id
  FOREIGN KEY (media_id)
  REFERENCES medias(id)
  ON DELETE CASCADE
  ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE competitors DROP FOREIGN KEY fk_competitors_media_id;

DROP TABLE IF EXISTS competitors;
