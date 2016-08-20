// Automatic generation for Phaser.Time
// generated file Time.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

)

// This is the core internal game clock.
// 
// It manages the elapsed time and calculation of elapsed values, used for game object motion and tweens,
// and also handles the standard Timer pool.
// 
// To create a general timed event, use the master {@link Phaser.Timer} accessible through {@link Phaser.Time.events events}.
// 
// There are different *types* of time in Phaser:
// 
// - ***Game time*** always runs at the speed of time in real life.
// 
//   Unlike wall-clock time, *game time stops when Phaser is paused*.
// 
//   Game time is used for {@link Phaser.Timer timer events}.
// 
// - ***Physics time*** represents the amount of time given to physics calculations.
// 
//   *When {@link Phaser.Time#slowMotion slowMotion} is in effect physics time runs slower than game time.*
//   Like game time, physics time stops when Phaser is paused.
// 
//   Physics time is used for physics calculations and {@link Phaser.Tween tweens}.
// 
// - {@link https://en.wikipedia.org/wiki/Wall-clock_time ***Wall-clock time***} represents the duration between two events in real life time.
// 
//   This time is independent of Phaser and always progresses, regardless of if Phaser is paused.
type Time struct {
    *js.Object
}


// Local reference to game.
func (self *Time) GetGame() Game{
    return Game{self.Get("game")}
}

// Local reference to game.
func (self *Time) SetGame(member Game) {
    self.Set("game", member)
}

// The `Date.now()` value when the time was last updated.
func (self *Time) GetTime() int{
    return self.Get("time").Int()
}

// The `Date.now()` value when the time was last updated.
func (self *Time) SetTime(member int) {
    self.Set("time", member)
}

// The `now` when the previous update occurred.
func (self *Time) GetPrevTime() float64{
    return self.Get("prevTime").Float()
}

// The `now` when the previous update occurred.
func (self *Time) SetPrevTime(member float64) {
    self.Set("prevTime", member)
}

// An increasing value representing cumulative milliseconds since an undisclosed epoch.
// 
// While this value is in milliseconds and can be used to compute time deltas,
// it must must _not_ be used with `Date.now()` as it may not use the same epoch / starting reference.
// 
// The source may either be from a high-res source (eg. if RAF is available) or the standard Date.now;
// the value can only be relied upon within a particular game instance.
func (self *Time) GetNow() float64{
    return self.Get("now").Float()
}

// An increasing value representing cumulative milliseconds since an undisclosed epoch.
// 
// While this value is in milliseconds and can be used to compute time deltas,
// it must must _not_ be used with `Date.now()` as it may not use the same epoch / starting reference.
// 
// The source may either be from a high-res source (eg. if RAF is available) or the standard Date.now;
// the value can only be relied upon within a particular game instance.
func (self *Time) SetNow(member float64) {
    self.Set("now", member)
}

// Elapsed time since the last time update, in milliseconds, based on `now`.
// 
// This value _may_ include time that the game is paused/inactive.
// 
// _Note:_ This is updated only once per game loop - even if multiple logic update steps are done.
// Use {@link Phaser.Timer#physicsTime physicsTime} as a basis of game/logic calculations instead.
func (self *Time) GetElapsed() float64{
    return self.Get("elapsed").Float()
}

// Elapsed time since the last time update, in milliseconds, based on `now`.
// 
// This value _may_ include time that the game is paused/inactive.
// 
// _Note:_ This is updated only once per game loop - even if multiple logic update steps are done.
// Use {@link Phaser.Timer#physicsTime physicsTime} as a basis of game/logic calculations instead.
func (self *Time) SetElapsed(member float64) {
    self.Set("elapsed", member)
}

// The time in ms since the last time update, in milliseconds, based on `time`.
// 
// This value is corrected for game pauses and will be "about zero" after a game is resumed.
// 
// _Note:_ This is updated once per game loop - even if multiple logic update steps are done.
// Use {@link Phaser.Timer#physicsTime physicsTime} as a basis of game/logic calculations instead.
func (self *Time) GetElapsedMS() int{
    return self.Get("elapsedMS").Int()
}

// The time in ms since the last time update, in milliseconds, based on `time`.
// 
// This value is corrected for game pauses and will be "about zero" after a game is resumed.
// 
// _Note:_ This is updated once per game loop - even if multiple logic update steps are done.
// Use {@link Phaser.Timer#physicsTime physicsTime} as a basis of game/logic calculations instead.
func (self *Time) SetElapsedMS(member int) {
    self.Set("elapsedMS", member)
}

// The physics update delta, in fractional seconds.
// 
// This should be used as an applicable multiplier by all logic update steps (eg. `preUpdate/postUpdate/update`)
// to ensure consistent game timing. Game/logic timing can drift from real-world time if the system
// is unable to consistently maintain the desired FPS.
// 
// With fixed-step updates this is normally equivalent to `1.0 / desiredFps`.
func (self *Time) GetPhysicsElapsed() float64{
    return self.Get("physicsElapsed").Float()
}

// The physics update delta, in fractional seconds.
// 
// This should be used as an applicable multiplier by all logic update steps (eg. `preUpdate/postUpdate/update`)
// to ensure consistent game timing. Game/logic timing can drift from real-world time if the system
// is unable to consistently maintain the desired FPS.
// 
// With fixed-step updates this is normally equivalent to `1.0 / desiredFps`.
func (self *Time) SetPhysicsElapsed(member float64) {
    self.Set("physicsElapsed", member)
}

// The physics update delta, in milliseconds - equivalent to `physicsElapsed * 1000`.
func (self *Time) GetPhysicsElapsedMS() float64{
    return self.Get("physicsElapsedMS").Float()
}

// The physics update delta, in milliseconds - equivalent to `physicsElapsed * 1000`.
func (self *Time) SetPhysicsElapsedMS(member float64) {
    self.Set("physicsElapsedMS", member)
}

// The desiredFps multiplier as used by Game.update.
func (self *Time) GetDesiredFpsMult() int{
    return self.Get("desiredFpsMult").Int()
}

// The desiredFps multiplier as used by Game.update.
func (self *Time) SetDesiredFpsMult(member int) {
    self.Set("desiredFpsMult", member)
}

// The suggested frame rate for your game, based on an averaged real frame rate.
// This value is only populated if `Time.advancedTiming` is enabled.
// 
// _Note:_ This is not available until after a few frames have passed; until then
// it's set to the same value as desiredFps.
func (self *Time) GetSuggestedFps() float64{
    return self.Get("suggestedFps").Float()
}

// The suggested frame rate for your game, based on an averaged real frame rate.
// This value is only populated if `Time.advancedTiming` is enabled.
// 
// _Note:_ This is not available until after a few frames have passed; until then
// it's set to the same value as desiredFps.
func (self *Time) SetSuggestedFps(member float64) {
    self.Set("suggestedFps", member)
}

// Scaling factor to make the game move smoothly in slow motion
// - 1.0 = normal speed
// - 2.0 = half speed
func (self *Time) GetSlowMotion() float64{
    return self.Get("slowMotion").Float()
}

// Scaling factor to make the game move smoothly in slow motion
// - 1.0 = normal speed
// - 2.0 = half speed
func (self *Time) SetSlowMotion(member float64) {
    self.Set("slowMotion", member)
}

// If true then advanced profiling, including the fps rate, fps min/max, suggestedFps and msMin/msMax are updated.
func (self *Time) GetAdvancedTiming() bool{
    return self.Get("advancedTiming").Bool()
}

// If true then advanced profiling, including the fps rate, fps min/max, suggestedFps and msMin/msMax are updated.
func (self *Time) SetAdvancedTiming(member bool) {
    self.Set("advancedTiming", member)
}

// Advanced timing result: The number of render frames record in the last second.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
func (self *Time) GetFrames() int{
    return self.Get("frames").Int()
}

// Advanced timing result: The number of render frames record in the last second.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
func (self *Time) SetFrames(member int) {
    self.Set("frames", member)
}

// Advanced timing result: Frames per second.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
func (self *Time) GetFps() float64{
    return self.Get("fps").Float()
}

// Advanced timing result: Frames per second.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
func (self *Time) SetFps(member float64) {
    self.Set("fps", member)
}

// Advanced timing result: The lowest rate the fps has dropped to.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
// This value can be manually reset.
func (self *Time) GetFpsMin() float64{
    return self.Get("fpsMin").Float()
}

// Advanced timing result: The lowest rate the fps has dropped to.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
// This value can be manually reset.
func (self *Time) SetFpsMin(member float64) {
    self.Set("fpsMin", member)
}

// Advanced timing result: The highest rate the fps has reached (usually no higher than 60fps).
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
// This value can be manually reset.
func (self *Time) GetFpsMax() float64{
    return self.Get("fpsMax").Float()
}

// Advanced timing result: The highest rate the fps has reached (usually no higher than 60fps).
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
// This value can be manually reset.
func (self *Time) SetFpsMax(member float64) {
    self.Set("fpsMax", member)
}

// Advanced timing result: The minimum amount of time the game has taken between consecutive frames.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
// This value can be manually reset.
func (self *Time) GetMsMin() float64{
    return self.Get("msMin").Float()
}

// Advanced timing result: The minimum amount of time the game has taken between consecutive frames.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
// This value can be manually reset.
func (self *Time) SetMsMin(member float64) {
    self.Set("msMin", member)
}

// Advanced timing result: The maximum amount of time the game has taken between consecutive frames.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
// This value can be manually reset.
func (self *Time) GetMsMax() float64{
    return self.Get("msMax").Float()
}

// Advanced timing result: The maximum amount of time the game has taken between consecutive frames.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
// This value can be manually reset.
func (self *Time) SetMsMax(member float64) {
    self.Set("msMax", member)
}

// Records how long the game was last paused, in milliseconds.
// (This is not updated until the game is resumed.)
func (self *Time) GetPauseDuration() float64{
    return self.Get("pauseDuration").Float()
}

// Records how long the game was last paused, in milliseconds.
// (This is not updated until the game is resumed.)
func (self *Time) SetPauseDuration(member float64) {
    self.Set("pauseDuration", member)
}

// The value that setTimeout needs to work out when to next update
func (self *Time) GetTimeToCall() float64{
    return self.Get("timeToCall").Float()
}

// The value that setTimeout needs to work out when to next update
func (self *Time) SetTimeToCall(member float64) {
    self.Set("timeToCall", member)
}

// The time when the next call is expected when using setTimer to control the update loop
func (self *Time) GetTimeExpected() float64{
    return self.Get("timeExpected").Float()
}

// The time when the next call is expected when using setTimer to control the update loop
func (self *Time) SetTimeExpected(member float64) {
    self.Set("timeExpected", member)
}

// A {@link Phaser.Timer} object bound to the master clock (this Time object) which events can be added to.
func (self *Time) GetEvents() Timer{
    return Timer{self.Get("events")}
}

// A {@link Phaser.Timer} object bound to the master clock (this Time object) which events can be added to.
func (self *Time) SetEvents(member Timer) {
    self.Set("events", member)
}

// The desired frame rate of the game.
// 
// This is used is used to calculate the physic / logic multiplier and how to apply catch-up logic updates. The desired frame rate of the game. Defaults to 60.
func (self *Time) GetDesiredFps() int{
    return self.Get("desiredFps").Int()
}

// The desired frame rate of the game.
// 
// This is used is used to calculate the physic / logic multiplier and how to apply catch-up logic updates. The desired frame rate of the game. Defaults to 60.
func (self *Time) SetDesiredFps(member int) {
    self.Set("desiredFps", member)
}



// Called automatically by Phaser.Game after boot. Should not be called directly.
func (self *Time) BootI(args ...interface{}) {
    self.Call("boot", args)
}

// Adds an existing Phaser.Timer object to the Timer pool.
func (self *Time) AddI(args ...interface{}) Timer{
    return Timer{self.Call("add", args)}
}

// Creates a new stand-alone Phaser.Timer object.
func (self *Time) CreateI(args ...interface{}) Timer{
    return Timer{self.Call("create", args)}
}

// Remove all Timer objects, regardless of their state and clears all Timers from the {@link Phaser.Time#events events} timer.
func (self *Time) RemoveAllI(args ...interface{}) {
    self.Call("removeAll", args)
}

// Refreshes the Time.time and Time.elapsedMS properties from the system clock.
func (self *Time) RefreshI(args ...interface{}) {
    self.Call("refresh", args)
}

// Updates the game clock and if enabled the advanced timing data. This is called automatically by Phaser.Game.
func (self *Time) UpdateI(args ...interface{}) {
    self.Call("update", args)
}

// Handles the updating of the Phaser.Timers (if any)
// Called automatically by Time.update.
func (self *Time) UpdateTimersI(args ...interface{}) {
    self.Call("updateTimers", args)
}

// Handles the updating of the advanced timing values (if enabled)
// Called automatically by Time.update.
func (self *Time) UpdateAdvancedTimingI(args ...interface{}) {
    self.Call("updateAdvancedTiming", args)
}

// Called when the game enters a paused state.
func (self *Time) GamePausedI(args ...interface{}) {
    self.Call("gamePaused", args)
}

// Called when the game resumes from a paused state.
func (self *Time) GameResumedI(args ...interface{}) {
    self.Call("gameResumed", args)
}

// The number of seconds that have elapsed since the game was started.
func (self *Time) TotalElapsedSecondsI(args ...interface{}) float64{
    return self.Call("totalElapsedSeconds", args).Float()
}

// How long has passed since the given time.
func (self *Time) ElapsedSinceI(args ...interface{}) float64{
    return self.Call("elapsedSince", args).Float()
}

// How long has passed since the given time (in seconds).
func (self *Time) ElapsedSecondsSinceI(args ...interface{}) float64{
    return self.Call("elapsedSecondsSince", args).Float()
}

// Resets the private _started value to now and removes all currently running Timers.
func (self *Time) ResetI(args ...interface{}) {
    self.Call("reset", args)
}
