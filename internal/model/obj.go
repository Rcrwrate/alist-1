package model

import (
	"io"
	"sort"
	"time"

	"github.com/maruel/natural"
)

type UnwarpObj interface {
	Unwarp() Obj
}

type Obj interface {
	GetSize() int64
	GetName() string
	ModTime() time.Time
	IsDir() bool

	// The internal information of the driver.
	// If you want to use it, please understand what it means
	GetID() string
	GetPath() string
}

type FileStreamer interface {
	io.ReadCloser
	Obj
	GetMimetype() string
	SetReadCloser(io.ReadCloser)
	NeedStore() bool
	GetReadCloser() io.ReadCloser
}

type URL interface {
	URL() string
}

type Thumb interface {
	Thumb() string
}

type SetPath interface {
	SetPath(path string)
}

func SortFiles(objs []Obj, orderBy, orderDirection string) {
	if orderBy == "" {
		return
	}
	sort.Slice(objs, func(i, j int) bool {
		switch orderBy {
		case "name":
			{
				c := natural.Less(objs[i].GetName(), objs[j].GetName())
				if orderDirection == "desc" {
					return !c
				}
				return c
			}
		case "size":
			{
				if orderDirection == "desc" {
					return objs[i].GetSize() >= objs[j].GetSize()
				}
				return objs[i].GetSize() <= objs[j].GetSize()
			}
		case "modified":
			if orderDirection == "desc" {
				return objs[i].ModTime().After(objs[j].ModTime())
			}
			return objs[i].ModTime().Before(objs[j].ModTime())
		}
		return false
	})
}

func ExtractFolder(objs []Obj, extractFolder string) {
	if extractFolder == "" {
		return
	}
	front := extractFolder == "front"
	sort.SliceStable(objs, func(i, j int) bool {
		if objs[i].IsDir() || objs[j].IsDir() {
			if !objs[i].IsDir() {
				return !front
			}
			if !objs[j].IsDir() {
				return front
			}
		}
		return false
	})
}

func WarpObjsName(objs []Obj) {
	for i := 0; i < len(objs); i++ {
		objs[i] = &ObjWarpName{Obj: objs[i]}
	}
}
