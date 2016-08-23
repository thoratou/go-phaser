package phaser

import (
	"github.com/gopherjs/gopherjs/js"
)

type PhysicsGlobal struct {
	*js.Object
	ARCADE   int `js:"ARCADE"`
	P2JS     int `js:"P2JS"`
	NINJA    int `js:"NINJA"`
	BOX2D    int `js:"BOX2D"`
	MATTER   int `js:"MATTER"`
	CHIPMUNK int `js:"CHIPMUNK"`
}

func GetGlobalConst(key string) int {
	return js.Global.Get("Phaser").Get(key).Int()
}

func GetKeyboardKey(key string) int {
	return js.Global.Get("Phaser").Get("Keyboard").Get(key).Int()
}
