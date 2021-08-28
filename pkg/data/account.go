package data

import (
	"fmt"
	"os"
	"path"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
	msg "github.com/libp2p/go-msgio"
	util "github.com/sonr-io/core/pkg/util"
	"google.golang.org/protobuf/proto"
)

// Method Initializes User Info Struct ^ //
func InitAccount(ir *InitializeRequest, d *Device) (*Account, *SonrError) {
	// Fetch Key Pair
	keychain, err := d.Initialize(ir)
	if err != nil {
		return nil, err
	}

	// Check for existing account
	if d.GetFileSystem().GetSupport().IsFile(util.ACCOUNT_FILE) {
		// Load Account
		buf, err := d.GetFileSystem().GetSupport().ReadFile(util.ACCOUNT_FILE)
		if err != nil {
			return nil, err
		}

		// Unmarshal Account
		loadedAccount := &Account{}
		serr := proto.Unmarshal(buf, loadedAccount)
		if serr != nil {
			return nil, NewError(serr, ErrorEvent_ACCOUNT_LOAD)
		}

		LogInfo(fmt.Sprintf("LoadedAccount: %s", loadedAccount.String()))

		// Set Account
		loadedAccount.KeyChain = keychain
		loadedAccount.Current = d
		loadedAccount.ApiKeys = ir.GetApiKeys()
		loadedAccount.State = ir.AccountState()
		return loadedAccount, nil
	} else {
		// Return User
		u := &Account{
			KeyChain: keychain,
			Current:  d,
			ApiKeys:  ir.GetApiKeys(),
			State:    ir.AccountState(),
			Devices:  make([]*Device, 0),
			Member: &Member{
				Reach:      Member_ONLINE,
				Associated: make([]*Peer, 0),
			},
		}
		err := u.Save()
		if err != nil {
			return nil, NewError(err, ErrorEvent_ACCOUNT_SAVE)
		}
		return u, nil
	}
}

// Set the User with ConnectionRequest
func (u *Account) SetConnection(cr *ConnectionRequest) {
	// Initialize Params
	u.PushToken = cr.GetPushToken()
	u.SName = cr.GetContact().GetProfile().GetSName()
	u.Contact = cr.GetContact()
	u.Member.PushToken = cr.GetPushToken()
	u.Save()
}

// Update Account after Device Peer set for Member
func (u *Account) HandleSetPeer(p *Peer, isPrimary bool) {
	if isPrimary {
		u.Member.Active = p
	} else {
		u.Member.Associated = append(u.Member.Associated, p)
	}
	u.Save()
}

// Update Account after LinkPacket is received
func (u *Account) HandleLinkPacket(lp *LinkPacket) {
	u.Primary = lp.Primary
	u.Devices = append(u.Devices, lp.Secondary)
	u.GetCurrent().replaceKeyChain(lp.GetKeyChain())
	u.Save()
}

// Checks Whether User is Ready to Communicate
func (u *Device) IsReady() bool {
	return u.Contact != nil && u.Location != nil && u.Status != Status_DEFAULT
}

// Return Client API Keys
func (u *Account) APIKeys() *APIKeys {
	return u.GetApiKeys()
}

// Method Returns DeviceID
func (u *Account) DeviceID() string {
	return u.GetCurrent().GetId()
}

// Method Returns Profile First Name
func (u *Account) FirstName() string {
	return u.GetContact().GetProfile().GetFirstName()
}

// Method Returns Profile Last Name
func (u *Account) LastName() string {
	return u.GetContact().GetProfile().GetLastName()
}

// Method Returns Profile
func (u *Account) Profile() *Profile {
	return u.GetContact().GetProfile()
}

// Method Returns Account KeyPair
func (u *Account) AccountKeys() *KeyPair {
	return u.GetKeyChain().GetAccount()
}

// Method Returns Device KeyPair
func (u *Account) DeviceKeys() *KeyPair {
	return u.GetKeyChain().GetDevice()
}

// Method Returns Device Link Public Key
func (u *Account) DevicePubKey() *KeyPair_Public {
	return u.GetKeyChain().GetDevice().GetPublic()
}

// Method Returns Group KeyPair
func (u *Account) GroupKeys() *KeyPair {
	return u.GetKeyChain().GetGroup()
}

// Method Returns Exportable Keychain for Linked Devices
func (u *Account) ExportKeychain() *KeyChain {
	return &KeyChain{
		Account: u.AccountKeys(),
		Device:  u.DeviceKeys(),
		Group:   u.GroupKeys(),
	}
}

func (u *Account) Save() error {
	// Marshal Account to Protobuf
	data, err := proto.Marshal(u)
	if err != nil {
		return err
	}

	// Open File at Path
	f, err := os.OpenFile(u.AccountFilePath(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	// Write Data to File
	_, err = f.Write(data)
	if err != nil {
		return err
	}

	// Close File
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

// Method Signs Data with KeyPair
func (u *Account) Sign(req *AuthRequest) *AuthResponse {
	// Create Prefix
	prefixResult := u.GetKeyChain().GetAccount().Sign(fmt.Sprintf("%s%s", req.GetSName(), u.DeviceID()))

	// Get Prefix Appended and Place
	prefix := util.Substring(prefixResult, 0, 16)

	// Get FingerPrint from Mnemonic and Place
	fingerprint := u.GetKeyChain().GetAccount().Sign(req.GetMnemonic())
	pubKey := u.GetKeyChain().GetAccount().PubKeyBase64()

	// Return Response
	return &AuthResponse{
		SignedPrefix:      prefix,
		SignedFingerprint: fingerprint,
		PublicKey:         pubKey,
		GivenSName:        req.GetSName(),
		GivenMnemonic:     req.GetMnemonic(),
	}
}

// Method Signs Packet with Keys
func (u *Account) SignLinkPacket(resp *LinkResponse) *LinkPacket {
	return &LinkPacket{
		Primary:   u.GetPrimary(),
		Secondary: resp.GetDevice(),
		KeyChain:  u.ExportKeychain(),
	}
}

// Method Returns support directory file for account
func (u *Account) AccountFilePath() string {
	return path.Join(u.GetCurrent().GetFileSystem().GetSupport().GetPath(), util.ACCOUNT_FILE)
}

// Method Updates User Contact
func (u *Account) UpdateContact(c *Contact) {
	u.Contact = c
	u.GetMember().UpdateProfile(c)
	u.Save()
}

// Method Verifies the Device Link Public Key
func (u *Account) VerifyDevicePubKey(pub crypto.PubKey) bool {
	return u.GetKeyChain().Device.VerifyPubKey(pub)
}

// Method Updates User Contact
func (u *Account) VerifyRead() *VerifyResponse {
	kp := u.GetKeyChain().GetAccount()
	return &VerifyResponse{
		PublicKey: kp.PubKeyBase64(),
		ShortID:   u.GetCurrent().ShortID(),
	}
}

// ** ─── Linker MANAGEMENT ────────────────────────────────────────────────────────

type Linker struct {
	// Inherited Properties
	pid        protocol.ID
	account    *Account
	primary    *Device
	associated *Device
	linkPacket *LinkPacket
}

// Node Device is Sending Primary Account Information
func NewOutLinker(pid protocol.ID, account *Account, lp *LinkPacket) *Linker {
	return &Linker{
		pid:        pid,
		account:    account,
		primary:    account.GetCurrent(),
		linkPacket: lp,
	}
}

// RPC Device is Receiving Primary Account Information
func NewInLinker(pid protocol.ID, account *Account) *Linker {
	return &Linker{
		pid:        pid,
		account:    account,
		associated: account.GetCurrent(),
	}
}

// Read Buffer sent from Stream to Handle Linker
func (s *Linker) ReadFromStream(stream network.Stream) {
	// Concurrent Function
	go func(rs msg.ReadCloser, stream network.Stream) {
		buf, err := rs.ReadMsg()
		if err != nil {
			LogError(err)
			return
		}

		// Unmarshal linkPacket
		lp := &LinkPacket{}
		err = proto.Unmarshal(buf, lp)
		if err != nil {
			LogError(err)
			return
		}

		// Callback on Linked
		s.account.HandleLinkPacket(lp)
		stream.Close()
	}(msg.NewReader(stream), stream)
}

// Write Buffer onto from Stream to New Associated Device
func (s *Linker) WriteToStream(stream network.Stream) {
	// Marshal linkPacket
	buf, err := proto.Marshal(s.linkPacket)
	if err != nil {
		LogError(err)
		return
	}

	// Concurrent Function
	go func(ws msg.WriteCloser) {
		err := ws.WriteMsg(buf)
		if err != nil {
			LogError(err)
		}
	}(msg.NewWriter(stream))
}
