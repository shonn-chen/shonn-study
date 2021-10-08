package client

import (
	"database/sql"

	"github.com/shonn-study/code-example/go/distributed-lock/model"
)

type MysqlInsertDeleteLockClient struct {
	db   *sql.DB
	lock *model.MysqlInsterDeleteLock
}

func NewMysqlInsertDeleteLockClient(db *sql.DB, lock *model.MysqlInsterDeleteLock) *MysqlInsertDeleteLockClient {
	return &MysqlInsertDeleteLockClient{
		db:   db,
		lock: lock,
	}
}

func (c *MysqlInsertDeleteLockClient) Lock() error {
	const sqlStr = `
		INSERT INTO distributed_lock_tab(
			lock_name, lock_remark, create_time, update_time) 
		VALUES(
			?, ?, REPLACE(unix_timestamp(now(3)),'.',''), REPLACE(unix_timestamp(now(3)),'.',''))
	`
	_, err := c.db.Exec(sqlStr, c.lock.GetLockName(), c.lock.GetLocRemark())
	if err != nil {
		return err
	}
	return nil
}

func (c *MysqlInsertDeleteLockClient) UnLock() error {
	const sqlStr = `
		DELETE FROM distributed_lock_tab 
		WHERE
			lock_name=?
	`
	_, err := c.db.Exec(sqlStr, c.lock.GetLockName())
	if err != nil {
		return err
	}
	return nil
}

func (c *MysqlInsertDeleteLockClient) GetLockDetail() (*model.MysqlInsterDeleteLock, error) {
	const sqlStr = `
		SELECT 
			id, lock_name, lock_remark, create_time, update_time
		FROM distributed_lock_tab
		WHERE
			lock_name=?
	`
	lock := model.NewEmptyMysqlInsertDeleteLock()
	row := c.db.QueryRow(sqlStr, c.lock.GetLockName())
	err := row.Scan(&lock.ID, &lock.LockName, &lock.LockRemark, &lock.CreateTime, &lock.UpdateTime)
	if err != nil {
		return nil, err
	}
	return lock, nil
}

func (c *MysqlInsertDeleteLockClient) Close() error {
	return c.db.Close()
}
