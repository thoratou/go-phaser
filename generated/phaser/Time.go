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
func NewTime(game *Game) *Time {
    return &Time{js.Global.Get("Phaser").Get("Time").New(game)}
}

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
func NewTimeI(args ...interface{}) *Time {
    return &Time{js.Global.Get("Phaser").Get("Time").New(args)}
}



// Local reference to game.
func (self *Time) Game() *Game{
    return &Game{self.Object.Get("game")}
}

// Local reference to game.
func (self *Time) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// The `Date.now()` value when the time was last updated.
func (self *Time) Time() int{
    return self.Object.Get("time").Int()
}

// The `Date.now()` value when the time was last updated.
func (self *Time) SetTimeA(member int) {
    self.Object.Set("time", member)
}

// The `now` when the previous update occurred.
func (self *Time) PrevTime() int{
    return self.Object.Get("prevTime").Int()
}

// The `now` when the previous update occurred.
func (self *Time) SetPrevTimeA(member int) {
    self.Object.Set("prevTime", member)
}

// An increasing value representing cumulative milliseconds since an undisclosed epoch.
// 
// While this value is in milliseconds and can be used to compute time deltas,
// it must must _not_ be used with `Date.now()` as it may not use the same epoch / starting reference.
// 
// The source may either be from a high-res source (eg. if RAF is available) or the standard Date.now;
// the value can only be relied upon within a particular game instance.
func (self *Time) Now() int{
    return self.Object.Get("now").Int()
}

// An increasing value representing cumulative milliseconds since an undisclosed epoch.
// 
// While this value is in milliseconds and can be used to compute time deltas,
// it must must _not_ be used with `Date.now()` as it may not use the same epoch / starting reference.
// 
// The source may either be from a high-res source (eg. if RAF is available) or the standard Date.now;
// the value can only be relied upon within a particular game instance.
func (self *Time) SetNowA(member int) {
    self.Object.Set("now", member)
}

// Elapsed time since the last time update, in milliseconds, based on `now`.
// 
// This value _may_ include time that the game is paused/inactive.
// 
// _Note:_ This is updated only once per game loop - even if multiple logic update steps are done.
// Use {@link Phaser.Timer#physicsTime physicsTime} as a basis of game/logic calculations instead.
func (self *Time) Elapsed() int{
    return self.Object.Get("elapsed").Int()
}

// Elapsed time since the last time update, in milliseconds, based on `now`.
// 
// This value _may_ include time that the game is paused/inactive.
// 
// _Note:_ This is updated only once per game loop - even if multiple logic update steps are done.
// Use {@link Phaser.Timer#physicsTime physicsTime} as a basis of game/logic calculations instead.
func (self *Time) SetElapsedA(member int) {
    self.Object.Set("elapsed", member)
}

// The time in ms since the last time update, in milliseconds, based on `time`.
// 
// This value is corrected for game pauses and will be "about zero" after a game is resumed.
// 
// _Note:_ This is updated once per game loop - even if multiple logic update steps are done.
// Use {@link Phaser.Timer#physicsTime physicsTime} as a basis of game/logic calculations instead.
func (self *Time) ElapsedMS() int{
    return self.Object.Get("elapsedMS").Int()
}

// The time in ms since the last time update, in milliseconds, based on `time`.
// 
// This value is corrected for game pauses and will be "about zero" after a game is resumed.
// 
// _Note:_ This is updated once per game loop - even if multiple logic update steps are done.
// Use {@link Phaser.Timer#physicsTime physicsTime} as a basis of game/logic calculations instead.
func (self *Time) SetElapsedMSA(member int) {
    self.Object.Set("elapsedMS", member)
}

// The physics update delta, in fractional seconds.
// 
// This should be used as an applicable multiplier by all logic update steps (eg. `preUpdate/postUpdate/update`)
// to ensure consistent game timing. Game/logic timing can drift from real-world time if the system
// is unable to consistently maintain the desired FPS.
// 
// With fixed-step updates this is normally equivalent to `1.0 / desiredFps`.
func (self *Time) PhysicsElapsed() int{
    return self.Object.Get("physicsElapsed").Int()
}

// The physics update delta, in fractional seconds.
// 
// This should be used as an applicable multiplier by all logic update steps (eg. `preUpdate/postUpdate/update`)
// to ensure consistent game timing. Game/logic timing can drift from real-world time if the system
// is unable to consistently maintain the desired FPS.
// 
// With fixed-step updates this is normally equivalent to `1.0 / desiredFps`.
func (self *Time) SetPhysicsElapsedA(member int) {
    self.Object.Set("physicsElapsed", member)
}

// The physics update delta, in milliseconds - equivalent to `physicsElapsed * 1000`.
func (self *Time) PhysicsElapsedMS() int{
    return self.Object.Get("physicsElapsedMS").Int()
}

// The physics update delta, in milliseconds - equivalent to `physicsElapsed * 1000`.
func (self *Time) SetPhysicsElapsedMSA(member int) {
    self.Object.Set("physicsElapsedMS", member)
}

// The desiredFps multiplier as used by Game.update.
func (self *Time) DesiredFpsMult() int{
    return self.Object.Get("desiredFpsMult").Int()
}

// The desiredFps multiplier as used by Game.update.
func (self *Time) SetDesiredFpsMultA(member int) {
    self.Object.Set("desiredFpsMult", member)
}

// The suggested frame rate for your game, based on an averaged real frame rate.
// This value is only populated if `Time.advancedTiming` is enabled.
// 
// _Note:_ This is not available until after a few frames have passed; until then
// it's set to the same value as desiredFps.
func (self *Time) SuggestedFps() int{
    return self.Object.Get("suggestedFps").Int()
}

// The suggested frame rate for your game, based on an averaged real frame rate.
// This value is only populated if `Time.advancedTiming` is enabled.
// 
// _Note:_ This is not available until after a few frames have passed; until then
// it's set to the same value as desiredFps.
func (self *Time) SetSuggestedFpsA(member int) {
    self.Object.Set("suggestedFps", member)
}

// Scaling factor to make the game move smoothly in slow motion
// - 1.0 = normal speed
// - 2.0 = half speed
func (self *Time) SlowMotion() int{
    return self.Object.Get("slowMotion").Int()
}

// Scaling factor to make the game move smoothly in slow motion
// - 1.0 = normal speed
// - 2.0 = half speed
func (self *Time) SetSlowMotionA(member int) {
    self.Object.Set("slowMotion", member)
}

// If true then advanced profiling, including the fps rate, fps min/max, suggestedFps and msMin/msMax are updated.
func (self *Time) AdvancedTiming() bool{
    return self.Object.Get("advancedTiming").Bool()
}

// If true then advanced profiling, including the fps rate, fps min/max, suggestedFps and msMin/msMax are updated.
func (self *Time) SetAdvancedTimingA(member bool) {
    self.Object.Set("advancedTiming", member)
}

// Advanced timing result: The number of render frames record in the last second.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
func (self *Time) Frames() int{
    return self.Object.Get("frames").Int()
}

// Advanced timing result: The number of render frames record in the last second.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
func (self *Time) SetFramesA(member int) {
    self.Object.Set("frames", member)
}

// Advanced timing result: Frames per second.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
func (self *Time) Fps() int{
    return self.Object.Get("fps").Int()
}

// Advanced timing result: Frames per second.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
func (self *Time) SetFpsA(member int) {
    self.Object.Set("fps", member)
}

// Advanced timing result: The lowest rate the fps has dropped to.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
// This value can be manually reset.
func (self *Time) FpsMin() int{
    return self.Object.Get("fpsMin").Int()
}

// Advanced timing result: The lowest rate the fps has dropped to.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
// This value can be manually reset.
func (self *Time) SetFpsMinA(member int) {
    self.Object.Set("fpsMin", member)
}

// Advanced timing result: The highest rate the fps has reached (usually no higher than 60fps).
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
// This value can be manually reset.
func (self *Time) FpsMax() int{
    return self.Object.Get("fpsMax").Int()
}

// Advanced timing result: The highest rate the fps has reached (usually no higher than 60fps).
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
// This value can be manually reset.
func (self *Time) SetFpsMaxA(member int) {
    self.Object.Set("fpsMax", member)
}

// Advanced timing result: The minimum amount of time the game has taken between consecutive frames.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
// This value can be manually reset.
func (self *Time) MsMin() int{
    return self.Object.Get("msMin").Int()
}

// Advanced timing result: The minimum amount of time the game has taken between consecutive frames.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
// This value can be manually reset.
func (self *Time) SetMsMinA(member int) {
    self.Object.Set("msMin", member)
}

// Advanced timing result: The maximum amount of time the game has taken between consecutive frames.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
// This value can be manually reset.
func (self *Time) MsMax() int{
    return self.Object.Get("msMax").Int()
}

// Advanced timing result: The maximum amount of time the game has taken between consecutive frames.
// 
// Only calculated if {@link Phaser.Time#advancedTiming advancedTiming} is enabled.
// This value can be manually reset.
func (self *Time) SetMsMaxA(member int) {
    self.Object.Set("msMax", member)
}

// Records how long the game was last paused, in milliseconds.
// (This is not updated until the game is resumed.)
func (self *Time) PauseDuration() int{
    return self.Object.Get("pauseDuration").Int()
}

// Records how long the game was last paused, in milliseconds.
// (This is not updated until the game is resumed.)
func (self *Time) SetPauseDurationA(member int) {
    self.Object.Set("pauseDuration", member)
}

// The value that setTimeout needs to work out when to next update
func (self *Time) TimeToCall() int{
    return self.Object.Get("timeToCall").Int()
}

// The value that setTimeout needs to work out when to next update
func (self *Time) SetTimeToCallA(member int) {
    self.Object.Set("timeToCall", member)
}

// The time when the next call is expected when using setTimer to control the update loop
func (self *Time) TimeExpected() int{
    return self.Object.Get("timeExpected").Int()
}

// The time when the next call is expected when using setTimer to control the update loop
func (self *Time) SetTimeExpectedA(member int) {
    self.Object.Set("timeExpected", member)
}

// A {@link Phaser.Timer} object bound to the master clock (this Time object) which events can be added to.
func (self *Time) Events() *Timer{
    return &Timer{self.Object.Get("events")}
}

// A {@link Phaser.Timer} object bound to the master clock (this Time object) which events can be added to.
func (self *Time) SetEventsA(member *Timer) {
    self.Object.Set("events", member)
}

// The desired frame rate of the game.
// 
// This is used is used to calculate the physic / logic multiplier and how to apply catch-up logic updates. The desired frame rate of the game. Defaults to 60.
func (self *Time) DesiredFps() int{
    return self.Object.Get("desiredFps").Int()
}

// The desired frame rate of the game.
// 
// This is used is used to calculate the physic / logic multiplier and how to apply catch-up logic updates. The desired frame rate of the game. Defaults to 60.
func (self *Time) SetDesiredFpsA(member int) {
    self.Object.Set("desiredFps", member)
}



// Called automatically by Phaser.Game after boot. Should not be called directly.
func (self *Time) Boot() {
    self.Object.Call("boot")
}

// Called automatically by Phaser.Game after boot. Should not be called directly.
func (self *Time) BootI(args ...interface{}) {
    self.Object.Call("boot", args)
}

// Adds an existing Phaser.Timer object to the Timer pool.
func (self *Time) Add(timer *Timer) *Timer{
    return &Timer{self.Object.Call("add", timer)}
}

// Adds an existing Phaser.Timer object to the Timer pool.
func (self *Time) AddI(args ...interface{}) *Timer{
    return &Timer{self.Object.Call("add", args)}
}

// Creates a new stand-alone Phaser.Timer object.
func (self *Time) Create() *Timer{
    return &Timer{self.Object.Call("create")}
}

// Creates a new stand-alone Phaser.Timer object.
func (self *Time) Create1O(autoDestroy bool) *Timer{
    return &Timer{self.Object.Call("create", autoDestroy)}
}

// Creates a new stand-alone Phaser.Timer object.
func (self *Time) CreateI(args ...interface{}) *Timer{
    return &Timer{self.Object.Call("create", args)}
}

// Remove all Timer objects, regardless of their state and clears all Timers from the {@link Phaser.Time#events events} timer.
func (self *Time) RemoveAll() {
    self.Object.Call("removeAll")
}

// Remove all Timer objects, regardless of their state and clears all Timers from the {@link Phaser.Time#events events} timer.
func (self *Time) RemoveAllI(args ...interface{}) {
    self.Object.Call("removeAll", args)
}

// Refreshes the Time.time and Time.elapsedMS properties from the system clock.
func (self *Time) Refresh() {
    self.Object.Call("refresh")
}

// Refreshes the Time.time and Time.elapsedMS properties from the system clock.
func (self *Time) RefreshI(args ...interface{}) {
    self.Object.Call("refresh", args)
}

// Updates the game clock and if enabled the advanced timing data. This is called automatically by Phaser.Game.
func (self *Time) Update(time int) {
    self.Object.Call("update", time)
}

// Updates the game clock and if enabled the advanced timing data. This is called automatically by Phaser.Game.
func (self *Time) UpdateI(args ...interface{}) {
    self.Object.Call("update", args)
}

// Handles the updating of the Phaser.Timers (if any)
// Called automatically by Time.update.
func (self *Time) UpdateTimers() {
    self.Object.Call("updateTimers")
}

// Handles the updating of the Phaser.Timers (if any)
// Called automatically by Time.update.
func (self *Time) UpdateTimersI(args ...interface{}) {
    self.Object.Call("updateTimers", args)
}

// Handles the updating of the advanced timing values (if enabled)
// Called automatically by Time.update.
func (self *Time) UpdateAdvancedTiming() {
    self.Object.Call("updateAdvancedTiming")
}

// Handles the updating of the advanced timing values (if enabled)
// Called automatically by Time.update.
func (self *Time) UpdateAdvancedTimingI(args ...interface{}) {
    self.Object.Call("updateAdvancedTiming", args)
}

// Called when the game enters a paused state.
func (self *Time) GamePaused() {
    self.Object.Call("gamePaused")
}

// Called when the game enters a paused state.
func (self *Time) GamePausedI(args ...interface{}) {
    self.Object.Call("gamePaused", args)
}

// Called when the game resumes from a paused state.
func (self *Time) GameResumed() {
    self.Object.Call("gameResumed")
}

// Called when the game resumes from a paused state.
func (self *Time) GameResumedI(args ...interface{}) {
    self.Object.Call("gameResumed", args)
}

// The number of seconds that have elapsed since the game was started.
func (self *Time) TotalElapsedSeconds() int{
    return self.Object.Call("totalElapsedSeconds").Int()
}

// The number of seconds that have elapsed since the game was started.
func (self *Time) TotalElapsedSecondsI(args ...interface{}) int{
    return self.Object.Call("totalElapsedSeconds", args).Int()
}

// How long has passed since the given time.
func (self *Time) ElapsedSince(since int) int{
    return self.Object.Call("elapsedSince", since).Int()
}

// How long has passed since the given time.
func (self *Time) ElapsedSinceI(args ...interface{}) int{
    return self.Object.Call("elapsedSince", args).Int()
}

// How long has passed since the given time (in seconds).
func (self *Time) ElapsedSecondsSince(since int) int{
    return self.Object.Call("elapsedSecondsSince", since).Int()
}

// How long has passed since the given time (in seconds).
func (self *Time) ElapsedSecondsSinceI(args ...interface{}) int{
    return self.Object.Call("elapsedSecondsSince", args).Int()
}

// Resets the private _started value to now and removes all currently running Timers.
func (self *Time) Reset() {
    self.Object.Call("reset")
}

// Resets the private _started value to now and removes all currently running Timers.
func (self *Time) ResetI(args ...interface{}) {
    self.Object.Call("reset", args)
}
