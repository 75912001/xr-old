package world_mgr

import (
	"fmt"
	"github.com/75912001/xr/lib/tcp"
	"github.com/75912001/xr/lib/util"
	"log"
	"sort"
)

type WORLD_MAP map[*tcp.Remote]*World
type WORLD_ID_SLICE []*World

type WorldMgr struct {
	worldMap     WORLD_MAP
	worldIdSlice WORLD_ID_SLICE
}

func (p *WorldMgr) Init() {
	p.worldMap = make(WORLD_MAP)
	p.worldIdSlice = make(WORLD_ID_SLICE, 0)
}

func (p *WorldMgr) Add(remote *tcp.Remote) (world *World) {
	world = &World{
		Remote: remote,
	}

	p.worldMap[remote] = world
	return
}

func (p *WorldMgr) Del(remote *tcp.Remote) {
	delete(p.worldMap, remote)
}

func (p *WorldMgr) Find(remote *tcp.Remote) (world *World) {
	world, _ = p.worldMap[remote]
	return
}

func (p *WorldMgr) AddById(id uint32, world *World) {
	w := p.Find(world.Remote)
	w.Id = id
	p.worldIdSlice = append(p.worldIdSlice, w)

	sort.Stable(p.worldIdSlice)
}

func (p *WorldMgr) DelById(id uint32) {
	if 0 == id {
		return
	}
	var idx int = -1
	for k, v := range p.worldIdSlice {
		if v.Id == id {
			idx = k
			break
		}
	}
	if idx < 0 {
		return
	}
	p.worldIdSlice = append(p.worldIdSlice[:idx], p.worldIdSlice[idx+1:]...)

	sort.Stable(p.worldIdSlice)
}

func (p *WorldMgr) GetWorldService(platform uint32, account string) (world *World) {
	defer func() {
		//world 可能在其他goroutine中已被移除.导致slice越界.
		if err := recover(); err != nil {
			log.Printf("GetWorldService panic:%v", err)
			world = nil
		}
	}()
	if len(p.worldIdSlice) == 0 {
		return
	}
	var key string
	key += "platform=" + fmt.Sprint(platform)
	key += "account=" + account
	idx := util.HASH32([]byte(key))

	world = p.worldIdSlice[idx%uint32(len(p.worldIdSlice))]
	return
}
