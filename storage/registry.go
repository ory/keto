package storage

type Registry interface {
	StorageManager() Manager
}
