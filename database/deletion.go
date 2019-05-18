package database

const stmtDeletePostTagsByPostId = "DELETE FROM posts_tags WHERE post_id = ?"
const stmtDeletePostById = "DELETE FROM posts WHERE id = ?"

func DeletePostTagsForPostId(postId int64) error {
	writeDB, err := readDB.Begin()
	if err != nil {
		_ = writeDB.Rollback()
		return err
	}
	_, err = writeDB.Exec(stmtDeletePostTagsByPostId, postId)
	if err != nil {
		_ = writeDB.Rollback()
		return err
	}
	return writeDB.Commit()
}

func DeletePostById(id int64) error {
	writeDB, err := readDB.Begin()
	if err != nil {
		_ = writeDB.Rollback()
		return err
	}
	_, err = writeDB.Exec(stmtDeletePostById, id)
	if err != nil {
		_ = writeDB.Rollback()
		return err
	}
	return writeDB.Commit()
}
