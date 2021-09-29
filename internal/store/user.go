package store

import (
	"time"

	"github.com/sonr-io/core/internal/common"
	bolt "go.etcd.io/bbolt"
	"google.golang.org/protobuf/proto"
)

var (
	PROFILE_KEY = []byte("profile")
)

// GetProfile returns the profile for the user from diskDB
func (s *Store) GetProfile() (*common.Profile, error) {
	var profile common.Profile
	err := s.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b, err := tx.CreateBucketIfNotExists(USER_BUCKET)
		if err != nil {
			return err
		}

		// Get profile buffer
		buf := b.Get(PROFILE_KEY)
		if buf == nil {
			return nil
		}

		// Unmarshal profile
		err = proto.Unmarshal(buf, &profile)
		if err != nil {
			return err
		}
		return nil
	})
	return &profile, err
}

// PutProfile stores the profile for the user in diskDB
func (s *Store) PutProfile(profile *common.Profile) error {
	// Verify timestamp
	if profile.GetLastModified() == 0 {
		profile.LastModified = time.Now().Unix()
	}

	// Put in Bucket
	return s.db.Update(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b, err := tx.CreateBucketIfNotExists(USER_BUCKET)
		if err != nil {
			return err
		}

		// Marshal profile
		buf, err := proto.Marshal(profile)
		if err != nil {
			return err
		}

		// Put profile
		err = b.Put(PROFILE_KEY, buf)
		if err != nil {
			return err
		}
		return nil
	})
}
