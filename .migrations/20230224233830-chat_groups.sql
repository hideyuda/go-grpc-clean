-- チャットグループを管理するテーブル
-- +migrate Up
CREATE TABLE IF NOT EXISTS chat_groups (
  id INT NOT NULL AUTO_INCREMENT UNIQUE,
  uuid VARCHAR(36) NOT NULL UNIQUE,
  name VARCHAR(100) NOT NULL,
  desc VARCHAR(255) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS chat_groups;
