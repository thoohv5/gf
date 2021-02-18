package impl

import (
	"context"
	"fmt"
	"time"

	sql "github.com/jmoiron/sqlx"

	"github.com/thoohv5/gf/db/standard"
)

type sDb struct {
	opts options
	*sql.DB
}

func NewSDb() standard.Connecter {
	return &sDb{}
}

func CopySDb(sdb *sql.DB, sos ...ServerOption) standard.IBuilder {
	var opts options
	for _, so := range sos {
		so(&opts)
	}
	return &sDb{
		opts: opts,
		DB:   sdb,
	}
}

func (s *sDb) Connect(config *standard.Config) (*standard.Wrapper, error) {
	sDb, err := sql.Open(config.Driver, config.Source)
	if err != nil {
		return nil, fmt.Errorf("sql open err: %w", err)
	}
	if err := sDb.Ping(); nil != err {
		return nil, fmt.Errorf("sql ping err: %w", err)
	}
	sDb.SetConnMaxLifetime(time.Duration(config.ConnMaxLifeTime) * time.Second)
	sDb.SetMaxIdleConns(config.MaxIdleConns)
	sDb.SetMaxOpenConns(config.MaxOpenConns)
	dsn := CopySDb(sDb, WithIsWrite(true))

	slave := make([]standard.IBuilder, 0, len(config.Slave))
	for _, slaveConfig := range config.Slave {
		sDb, err := sql.Open(config.Driver, slaveConfig.Source)
		if err != nil {
			return nil, fmt.Errorf("sql open err: %w", err)
		}
		if err := sDb.Ping(); nil != err {
			return nil, fmt.Errorf("sql ping err: %w", err)
		}
		sDb.SetConnMaxLifetime(time.Duration(config.ConnMaxLifeTime) * time.Second)
		sDb.SetMaxIdleConns(config.MaxIdleConns)
		sDb.SetMaxOpenConns(config.MaxOpenConns)
		slave = append(slave, CopySDb(sDb))
	}
	return standard.NewWrap(dsn, slave...), nil
}

func (s *sDb) Add(ctx context.Context, value interface{}) error {
	panic("implement me")
}

func (s *sDb) Update(ctx context.Context, attrs ...interface{}) error {
	panic("implement me")
}

func (s *sDb) Delete(ctx context.Context, value interface{}, where ...interface{}) error {
	panic("implement me")
}

func (s *sDb) Find(ctx context.Context, out interface{}, where ...interface{}) error {
	panic("implement me")
}

func (s *sDb) Get(ctx context.Context, out interface{}, where ...interface{}) error {
	panic("implement me")
}

func (s *sDb) Count(ctx context.Context, value interface{}) error {
	panic("implement me")
}

func (s *sDb) Where(query interface{}, args ...interface{}) standard.IBuilder {
	panic("implement me")
}

func (s *sDb) Or(query interface{}, args ...interface{}) standard.IBuilder {
	panic("implement me")
}

func (s *sDb) Offset(offset interface{}) standard.IBuilder {
	panic("implement me")
}

func (s *sDb) Limit(limit interface{}) standard.IBuilder {
	panic("implement me")
}

func (s *sDb) Order(value interface{}, reorder ...bool) standard.IBuilder {
	panic("implement me")
}

func (s *sDb) Begin(ctx context.Context) (standard.IBuilder, error) {
	panic("implement me")
}

func (s *sDb) Rollback() error {
	panic("implement me")
}

func (s *sDb) Commit() error {
	panic("implement me")
}

func (s *sDb) Model(value interface{}) standard.IBuilder {
	panic("implement me")
}

func (s *sDb) IsEmpty(e error) bool {
	panic("implement me")
}

func (s *sDb) IsWrite() bool {
	panic("implement me")
}

func (s *sDb) IsStartTx() bool {
	panic("implement me")
}

func (s *sDb) SetLogger(log standard.Logger) standard.IBuilder {
	panic("implement me")
}

func (s *sDb) Exec(ctx context.Context, sql string, values ...interface{}) error {
	_, err := s.DB.ExecContext(ctx, sql, values...)
	return err
}

func (s *sDb) Query(ctx context.Context, ret interface{}, sql string, values ...interface{}) error {
	return s.DB.SelectContext(ctx, ret, sql, values...)
}
