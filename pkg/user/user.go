package user

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	dt "github.com/sonr-io/core/internal/data"
	md "github.com/sonr-io/core/internal/models"
	"google.golang.org/protobuf/proto"
)

// @ Sonr User Struct
type User struct {
	Call    dt.NodeCallback
	Contact *md.Contact
	Device  *md.Device
	Peer    *md.Peer
	User    *md.User
}

// ^ Method Loads User Data from Disk ^ //
func (fs *FileSystem) LoadUser() (*md.User, error) {
	if fs.Initialized {
		path := filepath.Join(fs.Root, K_SONR_USER_PATH)

		// @ Check for Path
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return nil, errors.New("User File Does Not Exist")
		} else {
			// @ Read User Data File
			dat, err := ioutil.ReadFile(path)
			if err != nil {
				return nil, err
			}

			// Get User Data
			user := &md.User{}
			err = proto.Unmarshal(dat, user)
			if err != nil {
				return nil, err
			}
			return user, nil
		}
	}
	return nil, errors.New("Sonr FileSystem not Initialized")
}

// ^ Method Updates User Contact ^ //
func (fs *FileSystem) SaveContact(contact *md.Contact) error {
	if fs.Initialized {
		// Load User
		user, err := fs.LoadUser()
		if err != nil {
			return err
		}

		// Set Contact
		user.Contact = contact

		// Save User
		if err := fs.WriteUser(user); err != nil {
			return err
		}
		return nil
	}
	return errors.New("Sonr FileSystem not Initialized")
}

// ^ Method Adds Device to User ^ //
func (fs *FileSystem) SaveDevice(device *md.Device) error {
	if fs.Initialized {
		// Load User
		user, err := fs.LoadUser()
		if err != nil {
			return err
		}

		// Append Devices List
		user.Devices = append(user.Devices, device)

		// Save User
		err = fs.WriteUser(user)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("Sonr FileSystem not Initialized")
}

// ^ Write User Data at Path ^
func (sfs *FileSystem) WriteUser(user *md.User) error {
	userPath := filepath.Join(sfs.Root, K_SONR_USER_PATH)

	// Convert User to Bytes
	userData, err := proto.Marshal(user)
	if err != nil {
		return err
	}

	// Check for User File at Path
	file, err := os.OpenFile(userPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// Defer Close
	defer file.Close()

	// Write User Data to File
	_, err = file.Write(userData)
	if err != nil {
		return err
	}
	return nil
}
