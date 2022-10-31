package types

const (
	// ModuleName defines the module name
	ModuleName = "bucket"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bucket"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	BucketCountKey    = "WhereIs-count-"
	BucketKeyPrefix   = "Bucket/value/"
	DocumentKeyPrefix = "Document/value/"
	ServiceKeyPrefix     = "Service/value/"
)

func BucketKey(creator string) []byte {
	var key []byte
	creatorBytes := []byte(creator)
	key = append(key, creatorBytes...)
	key = append(key, []byte("/")...)

	return key
}
