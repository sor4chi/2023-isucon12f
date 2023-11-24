ALTER TABLE user_present_all_received_history ADD INDEX idx_user_present_all_received_history_user_id (user_id, present_all_id);

ALTER TABLE user_presents ADD INDEX idx_user_presents_user_id (user_id, deleted_at, created_at DESC, id);
