-- チャットを管理するテーブル
-- +migrate Up
CREATE TABLE IF NOT EXISTS chats (
  id INT NOT NULL AUTO_INCREMENT UNIQUE,
  uuid VARCHAR(36) NOT NULL UNIQUE,
  group_id INT NOT NULL,
  name VARCHAR(100) NOT NULL,
  from INT NOT NULL,
  to INT NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  PRIMARY KEY (id),
  INDEX idx_chats_group_id (group_id)
);

ALTER TABLE chats 
  ADD CONSTRAINT fk_chats_group_id
  FOREIGN KEY (group_id) 
  REFERENCES chat_groups(id) 
  ON DELETE CASCADE
  ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE chats DROP FOREIGN KEY fk_chats_group_id;

DROP TABLE IF EXISTS chats;
