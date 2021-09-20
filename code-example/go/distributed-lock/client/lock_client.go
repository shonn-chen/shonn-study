package client

type DistributedLockClient interface {
	Lock() error
	UnLock() error
}

type BlockedDistributedLockClient interface {
	BlockedLock()
	BlockedUnLock()
}
