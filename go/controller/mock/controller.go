// Code generated by MockGen. DO NOT EDIT.
// Source: controller.go
//
// Generated by this command:
//
//	mockgen -package mock -destination mock/controller.go --source=controller.go
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	pgtype "github.com/jackc/pgx/v5/pgtype"
	notifications "github.com/nhost/hasura-auth/go/notifications"
	sql "github.com/nhost/hasura-auth/go/sql"
	gomock "go.uber.org/mock/gomock"
)

// MockEmailer is a mock of Emailer interface.
type MockEmailer struct {
	ctrl     *gomock.Controller
	recorder *MockEmailerMockRecorder
}

// MockEmailerMockRecorder is the mock recorder for MockEmailer.
type MockEmailerMockRecorder struct {
	mock *MockEmailer
}

// NewMockEmailer creates a new mock instance.
func NewMockEmailer(ctrl *gomock.Controller) *MockEmailer {
	mock := &MockEmailer{ctrl: ctrl}
	mock.recorder = &MockEmailerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailer) EXPECT() *MockEmailerMockRecorder {
	return m.recorder
}

// SendEmail mocks base method.
func (m *MockEmailer) SendEmail(ctx context.Context, to, locale string, templateName notifications.TemplateName, data notifications.TemplateData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEmail", ctx, to, locale, templateName, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendEmail indicates an expected call of SendEmail.
func (mr *MockEmailerMockRecorder) SendEmail(ctx, to, locale, templateName, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmail", reflect.TypeOf((*MockEmailer)(nil).SendEmail), ctx, to, locale, templateName, data)
}

// MockDBClientGetUser is a mock of DBClientGetUser interface.
type MockDBClientGetUser struct {
	ctrl     *gomock.Controller
	recorder *MockDBClientGetUserMockRecorder
}

// MockDBClientGetUserMockRecorder is the mock recorder for MockDBClientGetUser.
type MockDBClientGetUserMockRecorder struct {
	mock *MockDBClientGetUser
}

// NewMockDBClientGetUser creates a new mock instance.
func NewMockDBClientGetUser(ctrl *gomock.Controller) *MockDBClientGetUser {
	mock := &MockDBClientGetUser{ctrl: ctrl}
	mock.recorder = &MockDBClientGetUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBClientGetUser) EXPECT() *MockDBClientGetUserMockRecorder {
	return m.recorder
}

// GetUser mocks base method.
func (m *MockDBClientGetUser) GetUser(ctx context.Context, id uuid.UUID) (sql.AuthUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(sql.AuthUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockDBClientGetUserMockRecorder) GetUser(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockDBClientGetUser)(nil).GetUser), ctx, id)
}

// GetUserByEmail mocks base method.
func (m *MockDBClientGetUser) GetUserByEmail(ctx context.Context, email pgtype.Text) (sql.AuthUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(sql.AuthUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockDBClientGetUserMockRecorder) GetUserByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockDBClientGetUser)(nil).GetUserByEmail), ctx, email)
}

// GetUserByRefreshTokenHash mocks base method.
func (m *MockDBClientGetUser) GetUserByRefreshTokenHash(ctx context.Context, arg sql.GetUserByRefreshTokenHashParams) (sql.AuthUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByRefreshTokenHash", ctx, arg)
	ret0, _ := ret[0].(sql.AuthUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByRefreshTokenHash indicates an expected call of GetUserByRefreshTokenHash.
func (mr *MockDBClientGetUserMockRecorder) GetUserByRefreshTokenHash(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByRefreshTokenHash", reflect.TypeOf((*MockDBClientGetUser)(nil).GetUserByRefreshTokenHash), ctx, arg)
}

// MockDBClientInsertUser is a mock of DBClientInsertUser interface.
type MockDBClientInsertUser struct {
	ctrl     *gomock.Controller
	recorder *MockDBClientInsertUserMockRecorder
}

// MockDBClientInsertUserMockRecorder is the mock recorder for MockDBClientInsertUser.
type MockDBClientInsertUserMockRecorder struct {
	mock *MockDBClientInsertUser
}

// NewMockDBClientInsertUser creates a new mock instance.
func NewMockDBClientInsertUser(ctrl *gomock.Controller) *MockDBClientInsertUser {
	mock := &MockDBClientInsertUser{ctrl: ctrl}
	mock.recorder = &MockDBClientInsertUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBClientInsertUser) EXPECT() *MockDBClientInsertUserMockRecorder {
	return m.recorder
}

// InsertUser mocks base method.
func (m *MockDBClientInsertUser) InsertUser(ctx context.Context, arg sql.InsertUserParams) (sql.InsertUserRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", ctx, arg)
	ret0, _ := ret[0].(sql.InsertUserRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockDBClientInsertUserMockRecorder) InsertUser(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockDBClientInsertUser)(nil).InsertUser), ctx, arg)
}

// InsertUserWithRefreshToken mocks base method.
func (m *MockDBClientInsertUser) InsertUserWithRefreshToken(ctx context.Context, arg sql.InsertUserWithRefreshTokenParams) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserWithRefreshToken", ctx, arg)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUserWithRefreshToken indicates an expected call of InsertUserWithRefreshToken.
func (mr *MockDBClientInsertUserMockRecorder) InsertUserWithRefreshToken(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserWithRefreshToken", reflect.TypeOf((*MockDBClientInsertUser)(nil).InsertUserWithRefreshToken), ctx, arg)
}

// InsertUserWithSecurityKeyAndRefreshToken mocks base method.
func (m *MockDBClientInsertUser) InsertUserWithSecurityKeyAndRefreshToken(ctx context.Context, arg sql.InsertUserWithSecurityKeyAndRefreshTokenParams) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserWithSecurityKeyAndRefreshToken", ctx, arg)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUserWithSecurityKeyAndRefreshToken indicates an expected call of InsertUserWithSecurityKeyAndRefreshToken.
func (mr *MockDBClientInsertUserMockRecorder) InsertUserWithSecurityKeyAndRefreshToken(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserWithSecurityKeyAndRefreshToken", reflect.TypeOf((*MockDBClientInsertUser)(nil).InsertUserWithSecurityKeyAndRefreshToken), ctx, arg)
}

// MockDBClientUpdateUser is a mock of DBClientUpdateUser interface.
type MockDBClientUpdateUser struct {
	ctrl     *gomock.Controller
	recorder *MockDBClientUpdateUserMockRecorder
}

// MockDBClientUpdateUserMockRecorder is the mock recorder for MockDBClientUpdateUser.
type MockDBClientUpdateUserMockRecorder struct {
	mock *MockDBClientUpdateUser
}

// NewMockDBClientUpdateUser creates a new mock instance.
func NewMockDBClientUpdateUser(ctrl *gomock.Controller) *MockDBClientUpdateUser {
	mock := &MockDBClientUpdateUser{ctrl: ctrl}
	mock.recorder = &MockDBClientUpdateUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBClientUpdateUser) EXPECT() *MockDBClientUpdateUserMockRecorder {
	return m.recorder
}

// InsertUserWithSecurityKey mocks base method.
func (m *MockDBClientUpdateUser) InsertUserWithSecurityKey(ctx context.Context, arg sql.InsertUserWithSecurityKeyParams) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserWithSecurityKey", ctx, arg)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUserWithSecurityKey indicates an expected call of InsertUserWithSecurityKey.
func (mr *MockDBClientUpdateUserMockRecorder) InsertUserWithSecurityKey(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserWithSecurityKey", reflect.TypeOf((*MockDBClientUpdateUser)(nil).InsertUserWithSecurityKey), ctx, arg)
}

// UpdateUserChangeEmail mocks base method.
func (m *MockDBClientUpdateUser) UpdateUserChangeEmail(ctx context.Context, arg sql.UpdateUserChangeEmailParams) (sql.AuthUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserChangeEmail", ctx, arg)
	ret0, _ := ret[0].(sql.AuthUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserChangeEmail indicates an expected call of UpdateUserChangeEmail.
func (mr *MockDBClientUpdateUserMockRecorder) UpdateUserChangeEmail(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserChangeEmail", reflect.TypeOf((*MockDBClientUpdateUser)(nil).UpdateUserChangeEmail), ctx, arg)
}

// UpdateUserDeanonymize mocks base method.
func (m *MockDBClientUpdateUser) UpdateUserDeanonymize(ctx context.Context, arg sql.UpdateUserDeanonymizeParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserDeanonymize", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserDeanonymize indicates an expected call of UpdateUserDeanonymize.
func (mr *MockDBClientUpdateUserMockRecorder) UpdateUserDeanonymize(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserDeanonymize", reflect.TypeOf((*MockDBClientUpdateUser)(nil).UpdateUserDeanonymize), ctx, arg)
}

// UpdateUserLastSeen mocks base method.
func (m *MockDBClientUpdateUser) UpdateUserLastSeen(ctx context.Context, id uuid.UUID) (pgtype.Timestamptz, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserLastSeen", ctx, id)
	ret0, _ := ret[0].(pgtype.Timestamptz)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserLastSeen indicates an expected call of UpdateUserLastSeen.
func (mr *MockDBClientUpdateUserMockRecorder) UpdateUserLastSeen(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserLastSeen", reflect.TypeOf((*MockDBClientUpdateUser)(nil).UpdateUserLastSeen), ctx, id)
}

// UpdateUserTicket mocks base method.
func (m *MockDBClientUpdateUser) UpdateUserTicket(ctx context.Context, arg sql.UpdateUserTicketParams) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserTicket", ctx, arg)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserTicket indicates an expected call of UpdateUserTicket.
func (mr *MockDBClientUpdateUserMockRecorder) UpdateUserTicket(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserTicket", reflect.TypeOf((*MockDBClientUpdateUser)(nil).UpdateUserTicket), ctx, arg)
}

// MockDBClient is a mock of DBClient interface.
type MockDBClient struct {
	ctrl     *gomock.Controller
	recorder *MockDBClientMockRecorder
}

// MockDBClientMockRecorder is the mock recorder for MockDBClient.
type MockDBClientMockRecorder struct {
	mock *MockDBClient
}

// NewMockDBClient creates a new mock instance.
func NewMockDBClient(ctrl *gomock.Controller) *MockDBClient {
	mock := &MockDBClient{ctrl: ctrl}
	mock.recorder = &MockDBClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBClient) EXPECT() *MockDBClientMockRecorder {
	return m.recorder
}

// CountSecurityKeysUser mocks base method.
func (m *MockDBClient) CountSecurityKeysUser(ctx context.Context, userID uuid.UUID) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountSecurityKeysUser", ctx, userID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountSecurityKeysUser indicates an expected call of CountSecurityKeysUser.
func (mr *MockDBClientMockRecorder) CountSecurityKeysUser(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountSecurityKeysUser", reflect.TypeOf((*MockDBClient)(nil).CountSecurityKeysUser), ctx, userID)
}

// DeleteRefreshTokens mocks base method.
func (m *MockDBClient) DeleteRefreshTokens(ctx context.Context, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRefreshTokens", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRefreshTokens indicates an expected call of DeleteRefreshTokens.
func (mr *MockDBClientMockRecorder) DeleteRefreshTokens(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRefreshTokens", reflect.TypeOf((*MockDBClient)(nil).DeleteRefreshTokens), ctx, userID)
}

// DeleteUserRoles mocks base method.
func (m *MockDBClient) DeleteUserRoles(ctx context.Context, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserRoles", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserRoles indicates an expected call of DeleteUserRoles.
func (mr *MockDBClientMockRecorder) DeleteUserRoles(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserRoles", reflect.TypeOf((*MockDBClient)(nil).DeleteUserRoles), ctx, userID)
}

// GetRefreshTokenByRefreshTokenHash mocks base method.
func (m *MockDBClient) GetRefreshTokenByRefreshTokenHash(ctx context.Context, refreshTokenHash pgtype.Text) (sql.AuthRefreshToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRefreshTokenByRefreshTokenHash", ctx, refreshTokenHash)
	ret0, _ := ret[0].(sql.AuthRefreshToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRefreshTokenByRefreshTokenHash indicates an expected call of GetRefreshTokenByRefreshTokenHash.
func (mr *MockDBClientMockRecorder) GetRefreshTokenByRefreshTokenHash(ctx, refreshTokenHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRefreshTokenByRefreshTokenHash", reflect.TypeOf((*MockDBClient)(nil).GetRefreshTokenByRefreshTokenHash), ctx, refreshTokenHash)
}

// GetUser mocks base method.
func (m *MockDBClient) GetUser(ctx context.Context, id uuid.UUID) (sql.AuthUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(sql.AuthUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockDBClientMockRecorder) GetUser(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockDBClient)(nil).GetUser), ctx, id)
}

// GetUserByEmail mocks base method.
func (m *MockDBClient) GetUserByEmail(ctx context.Context, email pgtype.Text) (sql.AuthUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(sql.AuthUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockDBClientMockRecorder) GetUserByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockDBClient)(nil).GetUserByEmail), ctx, email)
}

// GetUserByRefreshTokenHash mocks base method.
func (m *MockDBClient) GetUserByRefreshTokenHash(ctx context.Context, arg sql.GetUserByRefreshTokenHashParams) (sql.AuthUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByRefreshTokenHash", ctx, arg)
	ret0, _ := ret[0].(sql.AuthUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByRefreshTokenHash indicates an expected call of GetUserByRefreshTokenHash.
func (mr *MockDBClientMockRecorder) GetUserByRefreshTokenHash(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByRefreshTokenHash", reflect.TypeOf((*MockDBClient)(nil).GetUserByRefreshTokenHash), ctx, arg)
}

// GetUserRoles mocks base method.
func (m *MockDBClient) GetUserRoles(ctx context.Context, userID uuid.UUID) ([]sql.AuthUserRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRoles", ctx, userID)
	ret0, _ := ret[0].([]sql.AuthUserRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRoles indicates an expected call of GetUserRoles.
func (mr *MockDBClientMockRecorder) GetUserRoles(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRoles", reflect.TypeOf((*MockDBClient)(nil).GetUserRoles), ctx, userID)
}

// InsertRefreshtoken mocks base method.
func (m *MockDBClient) InsertRefreshtoken(ctx context.Context, arg sql.InsertRefreshtokenParams) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertRefreshtoken", ctx, arg)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertRefreshtoken indicates an expected call of InsertRefreshtoken.
func (mr *MockDBClientMockRecorder) InsertRefreshtoken(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertRefreshtoken", reflect.TypeOf((*MockDBClient)(nil).InsertRefreshtoken), ctx, arg)
}

// InsertUser mocks base method.
func (m *MockDBClient) InsertUser(ctx context.Context, arg sql.InsertUserParams) (sql.InsertUserRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", ctx, arg)
	ret0, _ := ret[0].(sql.InsertUserRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockDBClientMockRecorder) InsertUser(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockDBClient)(nil).InsertUser), ctx, arg)
}

// InsertUserWithRefreshToken mocks base method.
func (m *MockDBClient) InsertUserWithRefreshToken(ctx context.Context, arg sql.InsertUserWithRefreshTokenParams) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserWithRefreshToken", ctx, arg)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUserWithRefreshToken indicates an expected call of InsertUserWithRefreshToken.
func (mr *MockDBClientMockRecorder) InsertUserWithRefreshToken(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserWithRefreshToken", reflect.TypeOf((*MockDBClient)(nil).InsertUserWithRefreshToken), ctx, arg)
}

// InsertUserWithSecurityKey mocks base method.
func (m *MockDBClient) InsertUserWithSecurityKey(ctx context.Context, arg sql.InsertUserWithSecurityKeyParams) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserWithSecurityKey", ctx, arg)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUserWithSecurityKey indicates an expected call of InsertUserWithSecurityKey.
func (mr *MockDBClientMockRecorder) InsertUserWithSecurityKey(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserWithSecurityKey", reflect.TypeOf((*MockDBClient)(nil).InsertUserWithSecurityKey), ctx, arg)
}

// InsertUserWithSecurityKeyAndRefreshToken mocks base method.
func (m *MockDBClient) InsertUserWithSecurityKeyAndRefreshToken(ctx context.Context, arg sql.InsertUserWithSecurityKeyAndRefreshTokenParams) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUserWithSecurityKeyAndRefreshToken", ctx, arg)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUserWithSecurityKeyAndRefreshToken indicates an expected call of InsertUserWithSecurityKeyAndRefreshToken.
func (mr *MockDBClientMockRecorder) InsertUserWithSecurityKeyAndRefreshToken(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUserWithSecurityKeyAndRefreshToken", reflect.TypeOf((*MockDBClient)(nil).InsertUserWithSecurityKeyAndRefreshToken), ctx, arg)
}

// UpdateRefreshTokenExpiresAt mocks base method.
func (m *MockDBClient) UpdateRefreshTokenExpiresAt(ctx context.Context, arg sql.UpdateRefreshTokenExpiresAtParams) (pgtype.Timestamptz, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRefreshTokenExpiresAt", ctx, arg)
	ret0, _ := ret[0].(pgtype.Timestamptz)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRefreshTokenExpiresAt indicates an expected call of UpdateRefreshTokenExpiresAt.
func (mr *MockDBClientMockRecorder) UpdateRefreshTokenExpiresAt(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRefreshTokenExpiresAt", reflect.TypeOf((*MockDBClient)(nil).UpdateRefreshTokenExpiresAt), ctx, arg)
}

// UpdateUserChangeEmail mocks base method.
func (m *MockDBClient) UpdateUserChangeEmail(ctx context.Context, arg sql.UpdateUserChangeEmailParams) (sql.AuthUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserChangeEmail", ctx, arg)
	ret0, _ := ret[0].(sql.AuthUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserChangeEmail indicates an expected call of UpdateUserChangeEmail.
func (mr *MockDBClientMockRecorder) UpdateUserChangeEmail(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserChangeEmail", reflect.TypeOf((*MockDBClient)(nil).UpdateUserChangeEmail), ctx, arg)
}

// UpdateUserDeanonymize mocks base method.
func (m *MockDBClient) UpdateUserDeanonymize(ctx context.Context, arg sql.UpdateUserDeanonymizeParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserDeanonymize", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserDeanonymize indicates an expected call of UpdateUserDeanonymize.
func (mr *MockDBClientMockRecorder) UpdateUserDeanonymize(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserDeanonymize", reflect.TypeOf((*MockDBClient)(nil).UpdateUserDeanonymize), ctx, arg)
}

// UpdateUserLastSeen mocks base method.
func (m *MockDBClient) UpdateUserLastSeen(ctx context.Context, id uuid.UUID) (pgtype.Timestamptz, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserLastSeen", ctx, id)
	ret0, _ := ret[0].(pgtype.Timestamptz)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserLastSeen indicates an expected call of UpdateUserLastSeen.
func (mr *MockDBClientMockRecorder) UpdateUserLastSeen(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserLastSeen", reflect.TypeOf((*MockDBClient)(nil).UpdateUserLastSeen), ctx, id)
}

// UpdateUserTicket mocks base method.
func (m *MockDBClient) UpdateUserTicket(ctx context.Context, arg sql.UpdateUserTicketParams) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserTicket", ctx, arg)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserTicket indicates an expected call of UpdateUserTicket.
func (mr *MockDBClientMockRecorder) UpdateUserTicket(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserTicket", reflect.TypeOf((*MockDBClient)(nil).UpdateUserTicket), ctx, arg)
}
