// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q           = new(Query)
	Contact     *contact
	GroupMember *groupMember
	Message     *message
	Room        *room
	RoomFriend  *roomFriend
	RoomGroup   *roomGroup
	User        *user
	UserApply   *userApply
	UserFriend  *userFriend
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Contact = &Q.Contact
	GroupMember = &Q.GroupMember
	Message = &Q.Message
	Room = &Q.Room
	RoomFriend = &Q.RoomFriend
	RoomGroup = &Q.RoomGroup
	User = &Q.User
	UserApply = &Q.UserApply
	UserFriend = &Q.UserFriend
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:          db,
		Contact:     newContact(db, opts...),
		GroupMember: newGroupMember(db, opts...),
		Message:     newMessage(db, opts...),
		Room:        newRoom(db, opts...),
		RoomFriend:  newRoomFriend(db, opts...),
		RoomGroup:   newRoomGroup(db, opts...),
		User:        newUser(db, opts...),
		UserApply:   newUserApply(db, opts...),
		UserFriend:  newUserFriend(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Contact     contact
	GroupMember groupMember
	Message     message
	Room        room
	RoomFriend  roomFriend
	RoomGroup   roomGroup
	User        user
	UserApply   userApply
	UserFriend  userFriend
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:          db,
		Contact:     q.Contact.clone(db),
		GroupMember: q.GroupMember.clone(db),
		Message:     q.Message.clone(db),
		Room:        q.Room.clone(db),
		RoomFriend:  q.RoomFriend.clone(db),
		RoomGroup:   q.RoomGroup.clone(db),
		User:        q.User.clone(db),
		UserApply:   q.UserApply.clone(db),
		UserFriend:  q.UserFriend.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:          db,
		Contact:     q.Contact.replaceDB(db),
		GroupMember: q.GroupMember.replaceDB(db),
		Message:     q.Message.replaceDB(db),
		Room:        q.Room.replaceDB(db),
		RoomFriend:  q.RoomFriend.replaceDB(db),
		RoomGroup:   q.RoomGroup.replaceDB(db),
		User:        q.User.replaceDB(db),
		UserApply:   q.UserApply.replaceDB(db),
		UserFriend:  q.UserFriend.replaceDB(db),
	}
}

type queryCtx struct {
	Contact     IContactDo
	GroupMember IGroupMemberDo
	Message     IMessageDo
	Room        IRoomDo
	RoomFriend  IRoomFriendDo
	RoomGroup   IRoomGroupDo
	User        IUserDo
	UserApply   IUserApplyDo
	UserFriend  IUserFriendDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Contact:     q.Contact.WithContext(ctx),
		GroupMember: q.GroupMember.WithContext(ctx),
		Message:     q.Message.WithContext(ctx),
		Room:        q.Room.WithContext(ctx),
		RoomFriend:  q.RoomFriend.WithContext(ctx),
		RoomGroup:   q.RoomGroup.WithContext(ctx),
		User:        q.User.WithContext(ctx),
		UserApply:   q.UserApply.WithContext(ctx),
		UserFriend:  q.UserFriend.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
