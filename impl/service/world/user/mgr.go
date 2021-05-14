package user

import (
	"github.com/75912001/xr/impl/service/common/type_define"
	"github.com/75912001/xr/lib/tcp"
)

var GUserMgr UserMgr

type USER_MAP map[*tcp.Remote]*User
type USER_ID_MAP map[type_define.USER_ID]*User

type UserMgr struct {
	userMap   USER_MAP
	userIdMap USER_ID_MAP
}

func init() {
	GUserMgr.Init()
}
func (p *UserMgr) Init() {
	p.userMap = make(USER_MAP)
	p.userIdMap = make(USER_ID_MAP)
}

func (p *UserMgr) AddUser(remote *tcp.Remote) (user *User) {
	user = new(User)

	user.remote = remote
	p.userMap[user.remote] = user
	return
}

func (p *UserMgr) DelUser(remote *tcp.Remote) {
	delete(p.userMap, remote)
}

func (p *UserMgr) Find(remote *tcp.Remote) (user *User) {
	user, _ = p.userMap[remote]
	return
}

func (p *UserMgr) AddUserId(uid type_define.USER_ID, user *User) {
	p.userIdMap[uid] = user
}

func (p *UserMgr) DelUserId(uid type_define.USER_ID) {
	delete(p.userIdMap, uid)
}

func (this *UserMgr) FindById(uid type_define.USER_ID) (user *User) {
	user, _ = this.userIdMap[uid]
	return
}
