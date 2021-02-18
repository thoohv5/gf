package impl

import (
	"context"
	"reflect"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/thoohv5/gf/db/standard"
)

// gDb
type gDb struct {
	opts options
	*gorm.DB
}

func NewGDb() standard.Connecter {
	return &gDb{}
}

func CopyGDb(gdb *gorm.DB, sos ...ServerOption) standard.IBuilder {
	var opts options
	for _, so := range sos {
		so(&opts)
	}
	return &gDb{
		opts: opts,
		DB:   gdb,
	}
}

// 连接
func (g *gDb) Connect(config *standard.Config) (*standard.Wrapper, error) {
	var (
		err error
		gdb *gorm.DB
	)
	gdb, err = gorm.Open(config.Driver, config.Source)
	if err != nil {
		return nil, err
	}
	gdb.DB().SetConnMaxLifetime(time.Duration(config.ConnMaxLifeTime) * time.Second)
	gdb.DB().SetMaxIdleConns(config.MaxIdleConns)
	gdb.DB().SetMaxOpenConns(config.MaxOpenConns)
	if config.Log.Mode > 0 {
		gdb.LogMode(true)
	}
	// todo callback
	dsn := CopyGDb(gdb, WithIsWrite(true))
	slave := make([]standard.IBuilder, 0, len(config.Slave))
	for _, s := range config.Slave {
		var gDb *gorm.DB

		gDb, err = gorm.Open(config.Driver, s.Source)
		if err != nil {
			return nil, err
		}
		gDb.DB().SetConnMaxLifetime(time.Duration(config.ConnMaxLifeTime) * time.Second)
		gDb.DB().SetMaxIdleConns(config.MaxIdleConns)
		gDb.DB().SetMaxOpenConns(config.MaxOpenConns)
		if config.Log.Mode > 1 {
			gDb.LogMode(true)
		}
		slave = append(slave, CopyGDb(gDb))
	}
	return standard.NewWrap(dsn, slave...), err
}

// Add
func (g *gDb) Add(ctx context.Context, value interface{}) error {
	return g.DB.Create(value).Error
}

// Update
func (g *gDb) Update(ctx context.Context, attrs ...interface{}) error {
	return g.DB.Update(attrs).Error
}

// delete
func (g *gDb) Delete(ctx context.Context, value interface{}, where ...interface{}) error {
	return g.DB.Delete(value, where...).Error
}

// find
func (g *gDb) Find(ctx context.Context, out interface{}, where ...interface{}) error {
	return g.First(out).Error
}

// get
func (g *gDb) Get(ctx context.Context, out interface{}, where ...interface{}) error {
	return g.DB.Find(out).Error
}

// txCount
func (g *gDb) Count(ctx context.Context, value interface{}) error {
	return g.DB.Count(value).Error
}

// where
func (g *gDb) Where(query interface{}, args ...interface{}) standard.IBuilder {
	return CopyGDb(g.DB.Where(query, args...), WithIsWrite(g.IsWrite()), WithTxCount(g.TxCount()))
}

// or
func (g *gDb) Or(query interface{}, args ...interface{}) standard.IBuilder {
	return CopyGDb(g.DB.Or(query, args...), WithIsWrite(g.IsWrite()), WithTxCount(g.TxCount()))
}

// offset
func (g *gDb) Offset(offset interface{}) standard.IBuilder {
	return CopyGDb(g.DB.Offset(offset), WithIsWrite(g.IsWrite()), WithTxCount(g.TxCount()))
}

// limit
func (g *gDb) Limit(limit interface{}) standard.IBuilder {
	return CopyGDb(g.DB.Limit(limit), WithIsWrite(g.IsWrite()), WithTxCount(g.TxCount()))
}

// order
func (g *gDb) Order(value interface{}, reorder ...bool) standard.IBuilder {
	var db = g.DB
	if reflect.ValueOf(value).Kind() == reflect.Slice {
		if sort, ok := value.([]string); ok {
			for _, s := range sort {
				db = db.Order(s, reorder...)
			}
		}
	} else {
		db = db.Order(value, reorder...)
	}
	return CopyGDb(db, WithIsWrite(g.IsWrite()), WithTxCount(g.TxCount()))
}

// begin
func (g *gDb) Begin(ctx context.Context) (standard.IBuilder, error) {
	var (
		err     error
		db      = g.DB
		txCount = g.TxCount()
	)
	if txCount == 0 {
		db = db.Begin()
		err = db.Error
	}
	if nil == err {
		txCount++
	}
	return CopyGDb(db, WithIsWrite(g.IsWrite()), WithTxCount(txCount)), err
}

// rollback
func (g *gDb) Rollback() error {
	g.txClear()
	return g.DB.Rollback().Error
}

// commit
func (g *gDb) Commit() error {
	var (
		err     error
		txCount = g.txRelief()
	)
	if txCount == 0 {
		err = g.DB.Commit().Error
	}
	return err
}

// Model
func (g *gDb) Model(value interface{}) standard.IBuilder {
	return CopyGDb(g.DB.Model(value), WithIsWrite(g.IsWrite()), WithTxCount(g.TxCount()))
}

// IsEmpty
func (g *gDb) IsEmpty(err error) bool {
	var flag bool
	if err == gorm.ErrRecordNotFound {
		flag = true
	}
	return flag
}

// IsWrite
func (g *gDb) IsWrite() bool {
	return g.opts.isWrite
}

// IsStartTx
func (g *gDb) IsStartTx() bool {
	return g.opts.txCount > 0
}

// txRelief
func (g *gDb) txRelief() uint32 {
	g.opts.txCount -= 1
	return g.opts.txCount
}

// txClear
func (g *gDb) txClear() {
	g.opts.txCount = 0
}

// TxGagarin
func (g *gDb) TxGagarin() {
	g.opts.txCount += 1
}

// TxCount
func (g *gDb) TxCount() uint32 {
	return g.opts.txCount
}

// SetLogger
func (g *gDb) SetLogger(log standard.Logger) standard.IBuilder {
	g.DB.SetLogger(log)
	return CopyGDb(g.DB, WithIsWrite(g.IsWrite()), WithTxCount(g.TxCount()))
}

// Exec
func (g *gDb) Exec(ctx context.Context, sql string, values ...interface{}) error {
	return g.DB.New().Exec(sql, values...).Error
}

// Query
func (g *gDb) Query(ctx context.Context, dest interface{}, sql string, values ...interface{}) error {
	return g.DB.Raw(sql, values...).Scan(dest).Error
}
