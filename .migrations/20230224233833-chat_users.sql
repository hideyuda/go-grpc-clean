-- チャットを管理するテーブル
-- +migrate Up
CREATE TABLE IF NOT EXISTS chat_users (
  id INT NOT NULL AUTO_INCREMENT UNIQUE,
  uuid VARCHAR(36) NOT NULL UNIQUE,
  group_id INT NOT NULL,
  user_id INT NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  PRIMARY KEY (id),
  INDEX idx_chat_users_group_id (group_id),
  INDEX idx_chat_users_user_id (user_id)
);

ALTER TABLE chat_users 
  ADD CONSTRAINT fk_chat_users_group_id
  FOREIGN KEY (group_id) 
  REFERENCES chat_groups(id) 
  ON DELETE CASCADE
  ON UPDATE CASCADE;

ALTER TABLE chat_users
  ADD CONSTRAINT fk_chat_users_user_id
  FOREIGN KEY (user_id)
  REFERENCES users(id)
  ON DELETE CASCADE
  ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE chat_users DROP FOREIGN KEY fk_chat_users_group_id;
ALTER TABLE chat_users DROP FOREIGN KEY fk_chat_users_user_id;

DROP TABLE IF EXISTS chat_users;
