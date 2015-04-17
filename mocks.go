package curator

import (
	"errors"
	"time"

	"github.com/samuel/go-zookeeper/zk"
	"github.com/stretchr/testify/mock"
)

type mockCloseable struct {
	mock.Mock

	crash bool
}

func (c *mockCloseable) Close() error {
	if c.crash {
		panic(errors.New("panic"))
	}

	return c.Called().Error(0)
}

type mockTracerDriver struct {
	mock.Mock
}

func (t *mockTracerDriver) AddTime(name string, d time.Duration) {
	t.Called(name, d)
}

func (t *mockTracerDriver) AddCount(name string, increment int) {
	t.Called(name, increment)
}

type mockRetrySleeper struct {
	mock.Mock
}

func (s *mockRetrySleeper) SleepFor(time time.Duration) error {
	return s.Called(time).Error(0)
}

type mockConn struct {
	mock.Mock

	ZookeeperConnection
}

func (c *mockConn) AddAuth(scheme string, auth []byte) error {
	args := c.Called(scheme, auth)

	return args.Error(0)
}

func (c *mockConn) Close() {
	c.Called()
}

func (c *mockConn) Create(path string, data []byte, flags int32, acl []zk.ACL) (string, error) {
	args := c.Called(path, data, flags, acl)

	return args.String(0), args.Error(1)
}

func (c *mockConn) Exists(path string) (bool, *zk.Stat, error) {
	args := c.Called(path)

	stat, _ := args.Get(1).(*zk.Stat)

	return args.Bool(0), stat, args.Error(2)
}

func (c *mockConn) ExistsW(path string) (bool, *zk.Stat, <-chan zk.Event, error) {
	args := c.Called(path)

	stat, _ := args.Get(1).(*zk.Stat)
	events, _ := args.Get(2).(chan zk.Event)

	return args.Bool(0), stat, events, args.Error(3)
}

func (c *mockConn) Delete(path string, version int32) error {
	args := c.Called(path, version)

	return args.Error(0)
}

func (c *mockConn) Get(path string) ([]byte, *zk.Stat, error) {
	args := c.Called(path)

	data, _ := args.Get(0).([]byte)
	stat, _ := args.Get(1).(*zk.Stat)

	return data, stat, args.Error(2)
}

func (c *mockConn) GetW(path string) ([]byte, *zk.Stat, <-chan zk.Event, error) {
	args := c.Called(path)

	data, _ := args.Get(0).([]byte)
	stat, _ := args.Get(1).(*zk.Stat)
	events, _ := args.Get(2).(chan zk.Event)

	return data, stat, events, args.Error(3)
}

func (c *mockConn) Set(path string, data []byte, version int32) (*zk.Stat, error) {
	args := c.Called(path, data, version)

	stat, _ := args.Get(0).(*zk.Stat)

	return stat, args.Error(1)
}

func (c *mockConn) Children(path string) ([]string, *zk.Stat, error) {
	args := c.Called(path)

	children, _ := args.Get(0).([]string)
	stat, _ := args.Get(1).(*zk.Stat)

	return children, stat, args.Error(2)
}

func (c *mockConn) ChildrenW(path string) ([]string, *zk.Stat, <-chan zk.Event, error) {
	args := c.Called(path)

	children, _ := args.Get(0).([]string)
	stat, _ := args.Get(1).(*zk.Stat)
	events, _ := args.Get(2).(chan zk.Event)

	return children, stat, events, args.Error(3)
}

func (c *mockConn) GetACL(path string) ([]zk.ACL, *zk.Stat, error) {
	args := c.Called(path)

	acls, _ := args.Get(0).([]zk.ACL)
	stat, _ := args.Get(1).(*zk.Stat)

	return acls, stat, args.Error(2)
}

func (c *mockConn) SetACL(path string, acls []zk.ACL, version int32) (*zk.Stat, error) {
	args := c.Called(path, acls, version)

	stat, _ := args.Get(0).(*zk.Stat)

	return stat, args.Error(1)
}

func (c *mockConn) Multi(ops ...interface{}) ([]zk.MultiResponse, error) {
	args := c.Called(ops)

	res, _ := args.Get(0).([]zk.MultiResponse)

	return res, args.Error(1)
}

func (c *mockConn) Sync(path string) (string, error) {
	args := c.Called(path)

	return args.String(0), args.Error(1)
}

type mockZookeeperDialer struct {
	mock.Mock
}

func (d *mockZookeeperDialer) Dial(connString string, sessionTimeout time.Duration, canBeReadOnly bool) (ZookeeperConnection, <-chan zk.Event, error) {
	args := d.Called(connString, sessionTimeout, canBeReadOnly)

	conn, _ := args.Get(0).(ZookeeperConnection)
	events, _ := args.Get(1).(chan zk.Event)

	return conn, events, args.Error(2)
}

type mockCompressionProvider struct {
	mock.Mock
}

func (p *mockCompressionProvider) Compress(path string, data []byte) ([]byte, error) {
	args := p.Called(path, data)

	compressedData, _ := args.Get(0).([]byte)

	return compressedData, args.Error(1)
}

func (p *mockCompressionProvider) Decompress(path string, compressedData []byte) ([]byte, error) {
	args := p.Called(path, compressedData)

	data, _ := args.Get(0).([]byte)

	return data, args.Error(1)
}

type mockACLProvider struct {
	mock.Mock
}

func (p *mockACLProvider) GetDefaultAcl() []zk.ACL {
	args := p.Called()

	return args.Get(0).([]zk.ACL)
}

func (p *mockACLProvider) GetAclForPath(path string) []zk.ACL {
	args := p.Called(path)

	return args.Get(0).([]zk.ACL)
}