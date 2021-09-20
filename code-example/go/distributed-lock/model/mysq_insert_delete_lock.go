package model

type MysqlInsterDeleteLock struct {
	ID         *uint64
	LockName   *string
	LockRemark *string
	CreateTime *uint64
	UpdateTime *uint64
}

func NewEmptyMysqlInsertDeleteLock() *MysqlInsterDeleteLock {
	return &MysqlInsterDeleteLock{}
}

func (l *MysqlInsterDeleteLock) GetID() uint64 {
	if l != nil && l.ID != nil {
		return *l.ID
	}
	return 0
}

func (l *MysqlInsterDeleteLock) GetLockName() string {
	if l != nil && l.LockName != nil {
		return *l.LockName
	}
	return ""
}

func (l *MysqlInsterDeleteLock) GetLocRemark() string {
	if l != nil && l.LockRemark != nil {
		return *l.LockRemark
	}
	return ""
}

func (l *MysqlInsterDeleteLock) GetCreateTime() uint64 {
	if l != nil && l.CreateTime != nil {
		return *l.CreateTime
	}
	return 0
}

func (l *MysqlInsterDeleteLock) GetUpdateTime() uint64 {
	if l != nil && l.UpdateTime != nil {
		return *l.UpdateTime
	}
	return 0
}
