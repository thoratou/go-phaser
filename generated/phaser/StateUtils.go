package phaser

import "github.com/gopherjs/gopherjs/js"

type statePreload interface {
	Preload()
}

type stateCreate interface {
	Create()
}

type stateUpdate interface {
	Update()
}

type stateRender interface {
	Render()
}

type CustomState struct {
	*js.Object
	State interface{}
}

func WrapState(state interface{}) *CustomState {
	customState := &CustomState{
		Object: js.MakeWrapper(state),
		State:  state,
	}

	if typedState, match := customState.State.(statePreload); match {
		customState.Set("preload", typedState.Preload)
	}

	if typedState, match := customState.State.(stateCreate); match {
		customState.Set("create", typedState.Create)
	}

	if typedState, match := customState.State.(stateUpdate); match {
		customState.Set("update", typedState.Update)
	}

	if typedState, match := customState.State.(stateRender); match {
		customState.Set("render", typedState.Render)
	}

	return customState
}
