package phaser

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"
	"honnef.co/go/js/dom"
)

type Uint8ClampedArray struct{ *js.Object }
type Uint16Array struct{ *js.Object }
type Uint32Array struct{ *js.Object }
type Float32Array struct{ *js.Object }
type ImageData struct{ *js.Object }
type BitmapFont struct{ *js.Object }
type WebGLContext webgl.Context
type CreatureAnimation struct{ *js.Object }
type CreatureManager struct{ *js.Object }
type PhaserGame Game
type Emitter struct{ *js.Object }
type Weapon struct{ *js.Object }
type XMLDocument struct{ *js.Object }
type DOMElement dom.Element
type AudioContext struct{ *js.Object }
type TouchEvent struct{ *js.Object }
type MediaStream struct{ *js.Object }
type ArrayBuffer struct{ *js.Object }
type RenderSession struct{ *js.Object }
type XMLHttpRequest struct{ *js.Object }
type KeyboardEvent struct{ *js.Object }
type PointerEvent struct{ *js.Object }
type MouseEvent struct{ *js.Object }
type Blob struct{ *js.Object }

//PIXI
type PIXIPoint Point
type PIXIDisplayObjectContainer DisplayObjectContainer
type PIXIStage Stage
type PIXIMatrix Matrix
type PIXIGraphics Graphics
type PIXICanvasBuffer CanvasBuffer
type PIXITexture Texture

//Physics
type PhysicsBox2D struct{ *js.Object }
type PhysicsChipmunk struct{ *js.Object }
type PhysicsMatter struct{ *js.Object }
type P2World struct{ *js.Object }
type P2ContactMaterial struct{ *js.Object }
type PhysicsP2Constraint struct{ *js.Object }
type P2Body struct{ *js.Object }
type P2Shape struct{ *js.Object }
type P2Circle struct{ *js.Object }
type P2Box struct{ *js.Object }
type P2Plane struct{ *js.Object }
type P2Particle struct{ *js.Object }
type P2Line struct{ *js.Object }
type P2Capsule struct{ *js.Object }
type P2Rectangle struct{ *js.Object }
type P2RotationalSpring struct{ *js.Object }
type P2LinearSpring struct{ *js.Object }
type PhysicsCollisionGroup struct{ *js.Object }
