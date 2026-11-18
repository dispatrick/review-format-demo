package store

func Backup(data []byte) {
	Save("/var/backups/app/data", data)
}
