// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package modulev1

import (
	context "context"
	ormlist "cosmossdk.io/orm/model/ormlist"
	ormtable "cosmossdk.io/orm/model/ormtable"
	ormerrors "cosmossdk.io/orm/types/ormerrors"
)

type ServiceRecordTable interface {
	Insert(ctx context.Context, serviceRecord *ServiceRecord) error
	InsertReturningId(ctx context.Context, serviceRecord *ServiceRecord) (uint64, error)
	LastInsertedSequence(ctx context.Context) (uint64, error)
	Update(ctx context.Context, serviceRecord *ServiceRecord) error
	Save(ctx context.Context, serviceRecord *ServiceRecord) error
	Delete(ctx context.Context, serviceRecord *ServiceRecord) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64) (*ServiceRecord, error)
	List(ctx context.Context, prefixKey ServiceRecordIndexKey, opts ...ormlist.Option) (ServiceRecordIterator, error)
	ListRange(ctx context.Context, from, to ServiceRecordIndexKey, opts ...ormlist.Option) (ServiceRecordIterator, error)
	DeleteBy(ctx context.Context, prefixKey ServiceRecordIndexKey) error
	DeleteRange(ctx context.Context, from, to ServiceRecordIndexKey) error

	doNotImplement()
}

type ServiceRecordIterator struct {
	ormtable.Iterator
}

func (i ServiceRecordIterator) Value() (*ServiceRecord, error) {
	var serviceRecord ServiceRecord
	err := i.UnmarshalMessage(&serviceRecord)
	return &serviceRecord, err
}

type ServiceRecordIndexKey interface {
	id() uint32
	values() []interface{}
	serviceRecordIndexKey()
}

// primary key starting index..
type ServiceRecordPrimaryKey = ServiceRecordIdIndexKey

type ServiceRecordIdIndexKey struct {
	vs []interface{}
}

func (x ServiceRecordIdIndexKey) id() uint32             { return 0 }
func (x ServiceRecordIdIndexKey) values() []interface{}  { return x.vs }
func (x ServiceRecordIdIndexKey) serviceRecordIndexKey() {}

func (this ServiceRecordIdIndexKey) WithId(id uint64) ServiceRecordIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type ServiceRecordOriginIndexKey struct {
	vs []interface{}
}

func (x ServiceRecordOriginIndexKey) id() uint32             { return 1 }
func (x ServiceRecordOriginIndexKey) values() []interface{}  { return x.vs }
func (x ServiceRecordOriginIndexKey) serviceRecordIndexKey() {}

func (this ServiceRecordOriginIndexKey) WithOrigin(origin string) ServiceRecordOriginIndexKey {
	this.vs = []interface{}{origin}
	return this
}

type ServiceRecordControllerIndexKey struct {
	vs []interface{}
}

func (x ServiceRecordControllerIndexKey) id() uint32             { return 2 }
func (x ServiceRecordControllerIndexKey) values() []interface{}  { return x.vs }
func (x ServiceRecordControllerIndexKey) serviceRecordIndexKey() {}

func (this ServiceRecordControllerIndexKey) WithController(controller string) ServiceRecordControllerIndexKey {
	this.vs = []interface{}{controller}
	return this
}

type serviceRecordTable struct {
	table ormtable.AutoIncrementTable
}

func (this serviceRecordTable) Insert(ctx context.Context, serviceRecord *ServiceRecord) error {
	return this.table.Insert(ctx, serviceRecord)
}

func (this serviceRecordTable) Update(ctx context.Context, serviceRecord *ServiceRecord) error {
	return this.table.Update(ctx, serviceRecord)
}

func (this serviceRecordTable) Save(ctx context.Context, serviceRecord *ServiceRecord) error {
	return this.table.Save(ctx, serviceRecord)
}

func (this serviceRecordTable) Delete(ctx context.Context, serviceRecord *ServiceRecord) error {
	return this.table.Delete(ctx, serviceRecord)
}

func (this serviceRecordTable) InsertReturningId(ctx context.Context, serviceRecord *ServiceRecord) (uint64, error) {
	return this.table.InsertReturningPKey(ctx, serviceRecord)
}

func (this serviceRecordTable) LastInsertedSequence(ctx context.Context) (uint64, error) {
	return this.table.LastInsertedSequence(ctx)
}

func (this serviceRecordTable) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this serviceRecordTable) Get(ctx context.Context, id uint64) (*ServiceRecord, error) {
	var serviceRecord ServiceRecord
	found, err := this.table.PrimaryKey().Get(ctx, &serviceRecord, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &serviceRecord, nil
}

func (this serviceRecordTable) List(ctx context.Context, prefixKey ServiceRecordIndexKey, opts ...ormlist.Option) (ServiceRecordIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return ServiceRecordIterator{it}, err
}

func (this serviceRecordTable) ListRange(ctx context.Context, from, to ServiceRecordIndexKey, opts ...ormlist.Option) (ServiceRecordIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return ServiceRecordIterator{it}, err
}

func (this serviceRecordTable) DeleteBy(ctx context.Context, prefixKey ServiceRecordIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this serviceRecordTable) DeleteRange(ctx context.Context, from, to ServiceRecordIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this serviceRecordTable) doNotImplement() {}

var _ ServiceRecordTable = serviceRecordTable{}

func NewServiceRecordTable(db ormtable.Schema) (ServiceRecordTable, error) {
	table := db.GetTable(&ServiceRecord{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&ServiceRecord{}).ProtoReflect().Descriptor().FullName()))
	}
	return serviceRecordTable{table.(ormtable.AutoIncrementTable)}, nil
}

type CredentialTable interface {
	Insert(ctx context.Context, credential *Credential) error
	InsertReturningId(ctx context.Context, credential *Credential) (uint64, error)
	LastInsertedSequence(ctx context.Context) (uint64, error)
	Update(ctx context.Context, credential *Credential) error
	Save(ctx context.Context, credential *Credential) error
	Delete(ctx context.Context, credential *Credential) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64) (*Credential, error)
	HasByOriginHandle(ctx context.Context, origin string, handle string) (found bool, err error)
	// GetByOriginHandle returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByOriginHandle(ctx context.Context, origin string, handle string) (*Credential, error)
	HasByCredentialId(ctx context.Context, credential_id []byte) (found bool, err error)
	// GetByCredentialId returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByCredentialId(ctx context.Context, credential_id []byte) (*Credential, error)
	HasByPublicKey(ctx context.Context, public_key []byte) (found bool, err error)
	// GetByPublicKey returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByPublicKey(ctx context.Context, public_key []byte) (*Credential, error)
	List(ctx context.Context, prefixKey CredentialIndexKey, opts ...ormlist.Option) (CredentialIterator, error)
	ListRange(ctx context.Context, from, to CredentialIndexKey, opts ...ormlist.Option) (CredentialIterator, error)
	DeleteBy(ctx context.Context, prefixKey CredentialIndexKey) error
	DeleteRange(ctx context.Context, from, to CredentialIndexKey) error

	doNotImplement()
}

type CredentialIterator struct {
	ormtable.Iterator
}

func (i CredentialIterator) Value() (*Credential, error) {
	var credential Credential
	err := i.UnmarshalMessage(&credential)
	return &credential, err
}

type CredentialIndexKey interface {
	id() uint32
	values() []interface{}
	credentialIndexKey()
}

// primary key starting index..
type CredentialPrimaryKey = CredentialIdIndexKey

type CredentialIdIndexKey struct {
	vs []interface{}
}

func (x CredentialIdIndexKey) id() uint32            { return 0 }
func (x CredentialIdIndexKey) values() []interface{} { return x.vs }
func (x CredentialIdIndexKey) credentialIndexKey()   {}

func (this CredentialIdIndexKey) WithId(id uint64) CredentialIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type CredentialHandleIndexKey struct {
	vs []interface{}
}

func (x CredentialHandleIndexKey) id() uint32            { return 1 }
func (x CredentialHandleIndexKey) values() []interface{} { return x.vs }
func (x CredentialHandleIndexKey) credentialIndexKey()   {}

func (this CredentialHandleIndexKey) WithHandle(handle string) CredentialHandleIndexKey {
	this.vs = []interface{}{handle}
	return this
}

type CredentialOriginHandleIndexKey struct {
	vs []interface{}
}

func (x CredentialOriginHandleIndexKey) id() uint32            { return 2 }
func (x CredentialOriginHandleIndexKey) values() []interface{} { return x.vs }
func (x CredentialOriginHandleIndexKey) credentialIndexKey()   {}

func (this CredentialOriginHandleIndexKey) WithOrigin(origin string) CredentialOriginHandleIndexKey {
	this.vs = []interface{}{origin}
	return this
}

func (this CredentialOriginHandleIndexKey) WithOriginHandle(origin string, handle string) CredentialOriginHandleIndexKey {
	this.vs = []interface{}{origin, handle}
	return this
}

type CredentialCredentialIdIndexKey struct {
	vs []interface{}
}

func (x CredentialCredentialIdIndexKey) id() uint32            { return 3 }
func (x CredentialCredentialIdIndexKey) values() []interface{} { return x.vs }
func (x CredentialCredentialIdIndexKey) credentialIndexKey()   {}

func (this CredentialCredentialIdIndexKey) WithCredentialId(credential_id []byte) CredentialCredentialIdIndexKey {
	this.vs = []interface{}{credential_id}
	return this
}

type CredentialPublicKeyIndexKey struct {
	vs []interface{}
}

func (x CredentialPublicKeyIndexKey) id() uint32            { return 4 }
func (x CredentialPublicKeyIndexKey) values() []interface{} { return x.vs }
func (x CredentialPublicKeyIndexKey) credentialIndexKey()   {}

func (this CredentialPublicKeyIndexKey) WithPublicKey(public_key []byte) CredentialPublicKeyIndexKey {
	this.vs = []interface{}{public_key}
	return this
}

type credentialTable struct {
	table ormtable.AutoIncrementTable
}

func (this credentialTable) Insert(ctx context.Context, credential *Credential) error {
	return this.table.Insert(ctx, credential)
}

func (this credentialTable) Update(ctx context.Context, credential *Credential) error {
	return this.table.Update(ctx, credential)
}

func (this credentialTable) Save(ctx context.Context, credential *Credential) error {
	return this.table.Save(ctx, credential)
}

func (this credentialTable) Delete(ctx context.Context, credential *Credential) error {
	return this.table.Delete(ctx, credential)
}

func (this credentialTable) InsertReturningId(ctx context.Context, credential *Credential) (uint64, error) {
	return this.table.InsertReturningPKey(ctx, credential)
}

func (this credentialTable) LastInsertedSequence(ctx context.Context) (uint64, error) {
	return this.table.LastInsertedSequence(ctx)
}

func (this credentialTable) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this credentialTable) Get(ctx context.Context, id uint64) (*Credential, error) {
	var credential Credential
	found, err := this.table.PrimaryKey().Get(ctx, &credential, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &credential, nil
}

func (this credentialTable) HasByOriginHandle(ctx context.Context, origin string, handle string) (found bool, err error) {
	return this.table.GetIndexByID(2).(ormtable.UniqueIndex).Has(ctx,
		origin,
		handle,
	)
}

func (this credentialTable) GetByOriginHandle(ctx context.Context, origin string, handle string) (*Credential, error) {
	var credential Credential
	found, err := this.table.GetIndexByID(2).(ormtable.UniqueIndex).Get(ctx, &credential,
		origin,
		handle,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &credential, nil
}

func (this credentialTable) HasByCredentialId(ctx context.Context, credential_id []byte) (found bool, err error) {
	return this.table.GetIndexByID(3).(ormtable.UniqueIndex).Has(ctx,
		credential_id,
	)
}

func (this credentialTable) GetByCredentialId(ctx context.Context, credential_id []byte) (*Credential, error) {
	var credential Credential
	found, err := this.table.GetIndexByID(3).(ormtable.UniqueIndex).Get(ctx, &credential,
		credential_id,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &credential, nil
}

func (this credentialTable) HasByPublicKey(ctx context.Context, public_key []byte) (found bool, err error) {
	return this.table.GetIndexByID(4).(ormtable.UniqueIndex).Has(ctx,
		public_key,
	)
}

func (this credentialTable) GetByPublicKey(ctx context.Context, public_key []byte) (*Credential, error) {
	var credential Credential
	found, err := this.table.GetIndexByID(4).(ormtable.UniqueIndex).Get(ctx, &credential,
		public_key,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &credential, nil
}

func (this credentialTable) List(ctx context.Context, prefixKey CredentialIndexKey, opts ...ormlist.Option) (CredentialIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return CredentialIterator{it}, err
}

func (this credentialTable) ListRange(ctx context.Context, from, to CredentialIndexKey, opts ...ormlist.Option) (CredentialIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return CredentialIterator{it}, err
}

func (this credentialTable) DeleteBy(ctx context.Context, prefixKey CredentialIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this credentialTable) DeleteRange(ctx context.Context, from, to CredentialIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this credentialTable) doNotImplement() {}

var _ CredentialTable = credentialTable{}

func NewCredentialTable(db ormtable.Schema) (CredentialTable, error) {
	table := db.GetTable(&Credential{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Credential{}).ProtoReflect().Descriptor().FullName()))
	}
	return credentialTable{table.(ormtable.AutoIncrementTable)}, nil
}

type WitnessTable interface {
	Insert(ctx context.Context, witness *Witness) error
	InsertReturningIndex(ctx context.Context, witness *Witness) (uint64, error)
	LastInsertedSequence(ctx context.Context) (uint64, error)
	Update(ctx context.Context, witness *Witness) error
	Save(ctx context.Context, witness *Witness) error
	Delete(ctx context.Context, witness *Witness) error
	Has(ctx context.Context, index uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, index uint64) (*Witness, error)
	HasByOriginKey(ctx context.Context, origin string, key string) (found bool, err error)
	// GetByOriginKey returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	GetByOriginKey(ctx context.Context, origin string, key string) (*Witness, error)
	List(ctx context.Context, prefixKey WitnessIndexKey, opts ...ormlist.Option) (WitnessIterator, error)
	ListRange(ctx context.Context, from, to WitnessIndexKey, opts ...ormlist.Option) (WitnessIterator, error)
	DeleteBy(ctx context.Context, prefixKey WitnessIndexKey) error
	DeleteRange(ctx context.Context, from, to WitnessIndexKey) error

	doNotImplement()
}

type WitnessIterator struct {
	ormtable.Iterator
}

func (i WitnessIterator) Value() (*Witness, error) {
	var witness Witness
	err := i.UnmarshalMessage(&witness)
	return &witness, err
}

type WitnessIndexKey interface {
	id() uint32
	values() []interface{}
	witnessIndexKey()
}

// primary key starting index..
type WitnessPrimaryKey = WitnessIndexIndexKey

type WitnessIndexIndexKey struct {
	vs []interface{}
}

func (x WitnessIndexIndexKey) id() uint32            { return 0 }
func (x WitnessIndexIndexKey) values() []interface{} { return x.vs }
func (x WitnessIndexIndexKey) witnessIndexKey()      {}

func (this WitnessIndexIndexKey) WithIndex(index uint64) WitnessIndexIndexKey {
	this.vs = []interface{}{index}
	return this
}

type WitnessOriginKeyIndexKey struct {
	vs []interface{}
}

func (x WitnessOriginKeyIndexKey) id() uint32            { return 1 }
func (x WitnessOriginKeyIndexKey) values() []interface{} { return x.vs }
func (x WitnessOriginKeyIndexKey) witnessIndexKey()      {}

func (this WitnessOriginKeyIndexKey) WithOrigin(origin string) WitnessOriginKeyIndexKey {
	this.vs = []interface{}{origin}
	return this
}

func (this WitnessOriginKeyIndexKey) WithOriginKey(origin string, key string) WitnessOriginKeyIndexKey {
	this.vs = []interface{}{origin, key}
	return this
}

type witnessTable struct {
	table ormtable.AutoIncrementTable
}

func (this witnessTable) Insert(ctx context.Context, witness *Witness) error {
	return this.table.Insert(ctx, witness)
}

func (this witnessTable) Update(ctx context.Context, witness *Witness) error {
	return this.table.Update(ctx, witness)
}

func (this witnessTable) Save(ctx context.Context, witness *Witness) error {
	return this.table.Save(ctx, witness)
}

func (this witnessTable) Delete(ctx context.Context, witness *Witness) error {
	return this.table.Delete(ctx, witness)
}

func (this witnessTable) InsertReturningIndex(ctx context.Context, witness *Witness) (uint64, error) {
	return this.table.InsertReturningPKey(ctx, witness)
}

func (this witnessTable) LastInsertedSequence(ctx context.Context) (uint64, error) {
	return this.table.LastInsertedSequence(ctx)
}

func (this witnessTable) Has(ctx context.Context, index uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, index)
}

func (this witnessTable) Get(ctx context.Context, index uint64) (*Witness, error) {
	var witness Witness
	found, err := this.table.PrimaryKey().Get(ctx, &witness, index)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &witness, nil
}

func (this witnessTable) HasByOriginKey(ctx context.Context, origin string, key string) (found bool, err error) {
	return this.table.GetIndexByID(1).(ormtable.UniqueIndex).Has(ctx,
		origin,
		key,
	)
}

func (this witnessTable) GetByOriginKey(ctx context.Context, origin string, key string) (*Witness, error) {
	var witness Witness
	found, err := this.table.GetIndexByID(1).(ormtable.UniqueIndex).Get(ctx, &witness,
		origin,
		key,
	)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &witness, nil
}

func (this witnessTable) List(ctx context.Context, prefixKey WitnessIndexKey, opts ...ormlist.Option) (WitnessIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return WitnessIterator{it}, err
}

func (this witnessTable) ListRange(ctx context.Context, from, to WitnessIndexKey, opts ...ormlist.Option) (WitnessIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return WitnessIterator{it}, err
}

func (this witnessTable) DeleteBy(ctx context.Context, prefixKey WitnessIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this witnessTable) DeleteRange(ctx context.Context, from, to WitnessIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this witnessTable) doNotImplement() {}

var _ WitnessTable = witnessTable{}

func NewWitnessTable(db ormtable.Schema) (WitnessTable, error) {
	table := db.GetTable(&Witness{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Witness{}).ProtoReflect().Descriptor().FullName()))
	}
	return witnessTable{table.(ormtable.AutoIncrementTable)}, nil
}

// singleton store
type BaseParamsTable interface {
	Get(ctx context.Context) (*BaseParams, error)
	Save(ctx context.Context, baseParams *BaseParams) error
}

type baseParamsTable struct {
	table ormtable.Table
}

var _ BaseParamsTable = baseParamsTable{}

func (x baseParamsTable) Get(ctx context.Context) (*BaseParams, error) {
	baseParams := &BaseParams{}
	_, err := x.table.Get(ctx, baseParams)
	return baseParams, err
}

func (x baseParamsTable) Save(ctx context.Context, baseParams *BaseParams) error {
	return x.table.Save(ctx, baseParams)
}

func NewBaseParamsTable(db ormtable.Schema) (BaseParamsTable, error) {
	table := db.GetTable(&BaseParams{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&BaseParams{}).ProtoReflect().Descriptor().FullName()))
	}
	return &baseParamsTable{table}, nil
}

// singleton store
type ReadParamsTable interface {
	Get(ctx context.Context) (*ReadParams, error)
	Save(ctx context.Context, readParams *ReadParams) error
}

type readParamsTable struct {
	table ormtable.Table
}

var _ ReadParamsTable = readParamsTable{}

func (x readParamsTable) Get(ctx context.Context) (*ReadParams, error) {
	readParams := &ReadParams{}
	_, err := x.table.Get(ctx, readParams)
	return readParams, err
}

func (x readParamsTable) Save(ctx context.Context, readParams *ReadParams) error {
	return x.table.Save(ctx, readParams)
}

func NewReadParamsTable(db ormtable.Schema) (ReadParamsTable, error) {
	table := db.GetTable(&ReadParams{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&ReadParams{}).ProtoReflect().Descriptor().FullName()))
	}
	return &readParamsTable{table}, nil
}

// singleton store
type WriteParamsTable interface {
	Get(ctx context.Context) (*WriteParams, error)
	Save(ctx context.Context, writeParams *WriteParams) error
}

type writeParamsTable struct {
	table ormtable.Table
}

var _ WriteParamsTable = writeParamsTable{}

func (x writeParamsTable) Get(ctx context.Context) (*WriteParams, error) {
	writeParams := &WriteParams{}
	_, err := x.table.Get(ctx, writeParams)
	return writeParams, err
}

func (x writeParamsTable) Save(ctx context.Context, writeParams *WriteParams) error {
	return x.table.Save(ctx, writeParams)
}

func NewWriteParamsTable(db ormtable.Schema) (WriteParamsTable, error) {
	table := db.GetTable(&WriteParams{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&WriteParams{}).ProtoReflect().Descriptor().FullName()))
	}
	return &writeParamsTable{table}, nil
}

// singleton store
type OwnParamsTable interface {
	Get(ctx context.Context) (*OwnParams, error)
	Save(ctx context.Context, ownParams *OwnParams) error
}

type ownParamsTable struct {
	table ormtable.Table
}

var _ OwnParamsTable = ownParamsTable{}

func (x ownParamsTable) Get(ctx context.Context) (*OwnParams, error) {
	ownParams := &OwnParams{}
	_, err := x.table.Get(ctx, ownParams)
	return ownParams, err
}

func (x ownParamsTable) Save(ctx context.Context, ownParams *OwnParams) error {
	return x.table.Save(ctx, ownParams)
}

func NewOwnParamsTable(db ormtable.Schema) (OwnParamsTable, error) {
	table := db.GetTable(&OwnParams{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&OwnParams{}).ProtoReflect().Descriptor().FullName()))
	}
	return &ownParamsTable{table}, nil
}

type StateStore interface {
	ServiceRecordTable() ServiceRecordTable
	CredentialTable() CredentialTable
	WitnessTable() WitnessTable
	BaseParamsTable() BaseParamsTable
	ReadParamsTable() ReadParamsTable
	WriteParamsTable() WriteParamsTable
	OwnParamsTable() OwnParamsTable

	doNotImplement()
}

type stateStore struct {
	serviceRecord ServiceRecordTable
	credential    CredentialTable
	witness       WitnessTable
	baseParams    BaseParamsTable
	readParams    ReadParamsTable
	writeParams   WriteParamsTable
	ownParams     OwnParamsTable
}

func (x stateStore) ServiceRecordTable() ServiceRecordTable {
	return x.serviceRecord
}

func (x stateStore) CredentialTable() CredentialTable {
	return x.credential
}

func (x stateStore) WitnessTable() WitnessTable {
	return x.witness
}

func (x stateStore) BaseParamsTable() BaseParamsTable {
	return x.baseParams
}

func (x stateStore) ReadParamsTable() ReadParamsTable {
	return x.readParams
}

func (x stateStore) WriteParamsTable() WriteParamsTable {
	return x.writeParams
}

func (x stateStore) OwnParamsTable() OwnParamsTable {
	return x.ownParams
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormtable.Schema) (StateStore, error) {
	serviceRecordTable, err := NewServiceRecordTable(db)
	if err != nil {
		return nil, err
	}

	credentialTable, err := NewCredentialTable(db)
	if err != nil {
		return nil, err
	}

	witnessTable, err := NewWitnessTable(db)
	if err != nil {
		return nil, err
	}

	baseParamsTable, err := NewBaseParamsTable(db)
	if err != nil {
		return nil, err
	}

	readParamsTable, err := NewReadParamsTable(db)
	if err != nil {
		return nil, err
	}

	writeParamsTable, err := NewWriteParamsTable(db)
	if err != nil {
		return nil, err
	}

	ownParamsTable, err := NewOwnParamsTable(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		serviceRecordTable,
		credentialTable,
		witnessTable,
		baseParamsTable,
		readParamsTable,
		writeParamsTable,
		ownParamsTable,
	}, nil
}
