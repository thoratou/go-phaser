// Automatic generation for Phaser.Video
// generated file Video.go
package phaser

import (
	"github.com/gopherjs/gopherjs/js"

	dom "honnef.co/go/js/dom"

)

// A Video object that takes a previously loaded Video from the Phaser Cache and handles playback of it.
// 
// Alternatively it takes a getUserMedia feed from an active webcam and streams the contents of that to
// the Video instead (see `startMediaStream` method)
// 
// The video can then be applied to a Sprite as a texture. If multiple Sprites share the same Video texture and playback
// changes (i.e. you pause the video, or seek to a new time) then this change will be seen across all Sprites simultaneously.
// 
// Due to a bug in IE11 you cannot play a video texture to a Sprite in WebGL. For IE11 force Canvas mode.
// 
// If you need each Sprite to be able to play a video fully independently then you will need one Video object per Sprite.
// Please understand the obvious performance implications of doing this, and the memory required to hold videos in RAM.
// 
// On some mobile browsers such as iOS Safari, you cannot play a video until the user has explicitly touched the screen.
// This works in the same way as audio unlocking. Phaser will handle the touch unlocking for you, however unlike with audio
// it's worth noting that every single Video needs to be touch unlocked, not just the first one. You can use the `changeSource`
// method to try and work around this limitation, but see the method help for details.
// 
// Small screen devices, especially iPod and iPhone will launch the video in its own native video player,
// outside of the Safari browser. There is no way to avoid this, it's a device imposed limitation.
// 
// Note: On iOS if you need to detect when the user presses the 'Done' button (before the video ends)
// then you need to add your own event listener
type Video struct {
    *js.Object
}


// A reference to the currently running game.
func (self *Video) GetGameA() *Game{
    return &Game{self.Object.Get("game")}
}

// A reference to the currently running game.
func (self *Video) SetGameA(member *Game) {
    self.Object.Set("game", member)
}

// The key of the Video in the Cache, if stored there. Will be `null` if this Video is using the webcam instead.
func (self *Video) GetKeyA() string{
    return self.Object.Get("key").String()
}

// The key of the Video in the Cache, if stored there. Will be `null` if this Video is using the webcam instead.
func (self *Video) SetKeyA(member string) {
    self.Object.Set("key", member)
}

// The width of the video in pixels.
func (self *Video) GetWidthA() int{
    return self.Object.Get("width").Int()
}

// The width of the video in pixels.
func (self *Video) SetWidthA(member int) {
    self.Object.Set("width", member)
}

// The height of the video in pixels.
func (self *Video) GetHeightA() int{
    return self.Object.Get("height").Int()
}

// The height of the video in pixels.
func (self *Video) SetHeightA(member int) {
    self.Object.Set("height", member)
}

// The const type of this object.
func (self *Video) GetTypeA() int{
    return self.Object.Get("type").Int()
}

// The const type of this object.
func (self *Video) SetTypeA(member int) {
    self.Object.Set("type", member)
}

// If true this video will never send its image data to the GPU when its dirty flag is true. This only applies in WebGL.
func (self *Video) GetDisableTextureUploadA() bool{
    return self.Object.Get("disableTextureUpload").Bool()
}

// If true this video will never send its image data to the GPU when its dirty flag is true. This only applies in WebGL.
func (self *Video) SetDisableTextureUploadA(member bool) {
    self.Object.Set("disableTextureUpload", member)
}

// true if this video is currently locked awaiting a touch event. This happens on some mobile devices, such as iOS.
func (self *Video) GetTouchLockedA() bool{
    return self.Object.Get("touchLocked").Bool()
}

// true if this video is currently locked awaiting a touch event. This happens on some mobile devices, such as iOS.
func (self *Video) SetTouchLockedA(member bool) {
    self.Object.Set("touchLocked", member)
}

// This signal is dispatched when the Video starts to play. It sends 3 parameters: a reference to the Video object, if the video is set to loop or not and the playback rate.
func (self *Video) GetOnPlayA() *Signal{
    return &Signal{self.Object.Get("onPlay")}
}

// This signal is dispatched when the Video starts to play. It sends 3 parameters: a reference to the Video object, if the video is set to loop or not and the playback rate.
func (self *Video) SetOnPlayA(member *Signal) {
    self.Object.Set("onPlay", member)
}

// This signal is dispatched if the Video source is changed. It sends 3 parameters: a reference to the Video object and the new width and height of the new video source.
func (self *Video) GetOnChangeSourceA() *Signal{
    return &Signal{self.Object.Get("onChangeSource")}
}

// This signal is dispatched if the Video source is changed. It sends 3 parameters: a reference to the Video object and the new width and height of the new video source.
func (self *Video) SetOnChangeSourceA(member *Signal) {
    self.Object.Set("onChangeSource", member)
}

// This signal is dispatched when the Video completes playback, i.e. enters an 'ended' state. On iOS specifically it also fires if the user hits the 'Done' button at any point during playback. Videos set to loop will never dispatch this signal.
func (self *Video) GetOnCompleteA() *Signal{
    return &Signal{self.Object.Get("onComplete")}
}

// This signal is dispatched when the Video completes playback, i.e. enters an 'ended' state. On iOS specifically it also fires if the user hits the 'Done' button at any point during playback. Videos set to loop will never dispatch this signal.
func (self *Video) SetOnCompleteA(member *Signal) {
    self.Object.Set("onComplete", member)
}

// This signal is dispatched if the user allows access to their webcam.
func (self *Video) GetOnAccessA() *Signal{
    return &Signal{self.Object.Get("onAccess")}
}

// This signal is dispatched if the user allows access to their webcam.
func (self *Video) SetOnAccessA(member *Signal) {
    self.Object.Set("onAccess", member)
}

// This signal is dispatched if an error occurs either getting permission to use the webcam (for a Video Stream) or when trying to play back a video file.
func (self *Video) GetOnErrorA() *Signal{
    return &Signal{self.Object.Get("onError")}
}

// This signal is dispatched if an error occurs either getting permission to use the webcam (for a Video Stream) or when trying to play back a video file.
func (self *Video) SetOnErrorA(member *Signal) {
    self.Object.Set("onError", member)
}

// This signal is dispatched if when asking for permission to use the webcam no response is given within a the Video.timeout limit.
// This may be because the user has picked `Not now` in the permissions window, or there is a delay in establishing the LocalMediaStream.
func (self *Video) GetOnTimeoutA() *Signal{
    return &Signal{self.Object.Get("onTimeout")}
}

// This signal is dispatched if when asking for permission to use the webcam no response is given within a the Video.timeout limit.
// This may be because the user has picked `Not now` in the permissions window, or there is a delay in establishing the LocalMediaStream.
func (self *Video) SetOnTimeoutA(member *Signal) {
    self.Object.Set("onTimeout", member)
}

// The amount of ms allowed to elapsed before the Video.onTimeout signal is dispatched while waiting for webcam access.
func (self *Video) GetTimeoutA() int{
    return self.Object.Get("timeout").Int()
}

// The amount of ms allowed to elapsed before the Video.onTimeout signal is dispatched while waiting for webcam access.
func (self *Video) SetTimeoutA(member int) {
    self.Object.Set("timeout", member)
}

// The HTML Video Element that is added to the document.
func (self *Video) GetVideoA() dom.HTMLVideoElement{
    return WrapHTMLVideoElement(self.Object.Get("video"))
}

// The HTML Video Element that is added to the document.
func (self *Video) SetVideoA(member dom.HTMLVideoElement) {
    self.Object.Set("video", member)
}

// The Video Stream data. Only set if this Video is streaming from the webcam via `startMediaStream`.
func (self *Video) GetVideoStreamA() *MediaStream{
    return &MediaStream{self.Object.Get("videoStream")}
}

// The Video Stream data. Only set if this Video is streaming from the webcam via `startMediaStream`.
func (self *Video) SetVideoStreamA(member *MediaStream) {
    self.Object.Set("videoStream", member)
}

// Is there a streaming video source? I.e. from a webcam.
func (self *Video) GetIsStreamingA() bool{
    return self.Object.Get("isStreaming").Bool()
}

// Is there a streaming video source? I.e. from a webcam.
func (self *Video) SetIsStreamingA(member bool) {
    self.Object.Set("isStreaming", member)
}

// When starting playback of a video Phaser will monitor its readyState using a setTimeout call.
// The setTimeout happens once every `Video.retryInterval` ms. It will carry on monitoring the video
// state in this manner until the `retryLimit` is reached and then abort.
func (self *Video) GetRetryLimitA() int{
    return self.Object.Get("retryLimit").Int()
}

// When starting playback of a video Phaser will monitor its readyState using a setTimeout call.
// The setTimeout happens once every `Video.retryInterval` ms. It will carry on monitoring the video
// state in this manner until the `retryLimit` is reached and then abort.
func (self *Video) SetRetryLimitA(member int) {
    self.Object.Set("retryLimit", member)
}

// The current retry attempt.
func (self *Video) GetRetryA() int{
    return self.Object.Get("retry").Int()
}

// The current retry attempt.
func (self *Video) SetRetryA(member int) {
    self.Object.Set("retry", member)
}

// The number of ms between each retry at monitoring the status of a downloading video.
func (self *Video) GetRetryIntervalA() int{
    return self.Object.Get("retryInterval").Int()
}

// The number of ms between each retry at monitoring the status of a downloading video.
func (self *Video) SetRetryIntervalA(member int) {
    self.Object.Set("retryInterval", member)
}

// The PIXI.Texture.
func (self *Video) GetTextureA() *Texture{
    return &Texture{self.Object.Get("texture")}
}

// The PIXI.Texture.
func (self *Video) SetTextureA(member *Texture) {
    self.Object.Set("texture", member)
}

// The Frame this video uses for rendering.
func (self *Video) GetTextureFrameA() *Frame{
    return &Frame{self.Object.Get("textureFrame")}
}

// The Frame this video uses for rendering.
func (self *Video) SetTextureFrameA(member *Frame) {
    self.Object.Set("textureFrame", member)
}

// A snapshot grabbed from the video. This is initially black. Populate it by calling Video.grab().
// When called the BitmapData is updated with a grab taken from the current video playing or active video stream.
// If Phaser has been compiled without BitmapData support this property will always be `null`.
func (self *Video) GetSnapshotA() *BitmapData{
    return &BitmapData{self.Object.Get("snapshot")}
}

// A snapshot grabbed from the video. This is initially black. Populate it by calling Video.grab().
// When called the BitmapData is updated with a grab taken from the current video playing or active video stream.
// If Phaser has been compiled without BitmapData support this property will always be `null`.
func (self *Video) SetSnapshotA(member *BitmapData) {
    self.Object.Set("snapshot", member)
}

// The current time of the video in seconds. If set the video will attempt to seek to that point in time.
func (self *Video) GetCurrentTimeA() int{
    return self.Object.Get("currentTime").Int()
}

// The current time of the video in seconds. If set the video will attempt to seek to that point in time.
func (self *Video) SetCurrentTimeA(member int) {
    self.Object.Set("currentTime", member)
}

// The duration of the video in seconds.
func (self *Video) GetDurationA() int{
    return self.Object.Get("duration").Int()
}

// The duration of the video in seconds.
func (self *Video) SetDurationA(member int) {
    self.Object.Set("duration", member)
}

// The progress of this video. This is a value between 0 and 1, where 0 is the start and 1 is the end of the video.
func (self *Video) GetProgressA() int{
    return self.Object.Get("progress").Int()
}

// The progress of this video. This is a value between 0 and 1, where 0 is the start and 1 is the end of the video.
func (self *Video) SetProgressA(member int) {
    self.Object.Set("progress", member)
}

// Gets or sets the muted state of the Video.
func (self *Video) GetMuteA() bool{
    return self.Object.Get("mute").Bool()
}

// Gets or sets the muted state of the Video.
func (self *Video) SetMuteA(member bool) {
    self.Object.Set("mute", member)
}

// Gets or sets the paused state of the Video.
// If the video is still touch locked (such as on iOS devices) this call has no effect.
func (self *Video) GetPausedA() bool{
    return self.Object.Get("paused").Bool()
}

// Gets or sets the paused state of the Video.
// If the video is still touch locked (such as on iOS devices) this call has no effect.
func (self *Video) SetPausedA(member bool) {
    self.Object.Set("paused", member)
}

// Gets or sets the volume of the Video, a value between 0 and 1. The value given is clamped to the range 0 to 1.
func (self *Video) GetVolumeA() int{
    return self.Object.Get("volume").Int()
}

// Gets or sets the volume of the Video, a value between 0 and 1. The value given is clamped to the range 0 to 1.
func (self *Video) SetVolumeA(member int) {
    self.Object.Set("volume", member)
}

// Gets or sets the playback rate of the Video. This is the speed at which the video is playing.
func (self *Video) GetPlaybackRateA() int{
    return self.Object.Get("playbackRate").Int()
}

// Gets or sets the playback rate of the Video. This is the speed at which the video is playing.
func (self *Video) SetPlaybackRateA(member int) {
    self.Object.Set("playbackRate", member)
}

// Gets or sets if the Video is set to loop.
// Please note that at present some browsers (i.e. Chrome) do not support *seamless* video looping.
// If the video isn't yet set this will always return false.
func (self *Video) GetLoopA() bool{
    return self.Object.Get("loop").Bool()
}

// Gets or sets if the Video is set to loop.
// Please note that at present some browsers (i.e. Chrome) do not support *seamless* video looping.
// If the video isn't yet set this will always return false.
func (self *Video) SetLoopA(member bool) {
    self.Object.Set("loop", member)
}

// True if the video is currently playing (and not paused or ended), otherwise false.
func (self *Video) GetPlayingA() bool{
    return self.Object.Get("playing").Bool()
}

// True if the video is currently playing (and not paused or ended), otherwise false.
func (self *Video) SetPlayingA(member bool) {
    self.Object.Set("playing", member)
}



// Connects to an external media stream for the webcam, rather than using a local one.
func (self *Video) ConnectToMediaStream(video *dom.HTMLVideoElement, stream *MediaStream) *Video{
    return &Video{self.Object.Call("connectToMediaStream", video, stream)}
}

// Connects to an external media stream for the webcam, rather than using a local one.
func (self *Video) ConnectToMediaStreamI(args ...interface{}) *Video{
    return &Video{self.Object.Call("connectToMediaStream", args)}
}

// Instead of playing a video file this method allows you to stream video data from an attached webcam.
// 
// As soon as this method is called the user will be prompted by their browser to "Allow" access to the webcam.
// If they allow it the webcam feed is directed to this Video. Call `Video.play` to start the stream.
// 
// If they block the webcam the onError signal will be dispatched containing the NavigatorUserMediaError
// or MediaStreamError event.
// 
// You can optionally set a width and height for the stream. If set the input will be cropped to these dimensions.
// If not given then as soon as the stream has enough data the video dimensions will be changed to match the webcam device.
// You can listen for this with the onChangeSource signal.
func (self *Video) StartMediaStream(captureAudio bool, width int, height int) *Video{
    return &Video{self.Object.Call("startMediaStream", captureAudio, width, height)}
}

// Instead of playing a video file this method allows you to stream video data from an attached webcam.
// 
// As soon as this method is called the user will be prompted by their browser to "Allow" access to the webcam.
// If they allow it the webcam feed is directed to this Video. Call `Video.play` to start the stream.
// 
// If they block the webcam the onError signal will be dispatched containing the NavigatorUserMediaError
// or MediaStreamError event.
// 
// You can optionally set a width and height for the stream. If set the input will be cropped to these dimensions.
// If not given then as soon as the stream has enough data the video dimensions will be changed to match the webcam device.
// You can listen for this with the onChangeSource signal.
func (self *Video) StartMediaStreamI(args ...interface{}) *Video{
    return &Video{self.Object.Call("startMediaStream", args)}
}

// 
func (self *Video) GetUserMediaTimeout() {
    self.Object.Call("getUserMediaTimeout")
}

// 
func (self *Video) GetUserMediaTimeoutI(args ...interface{}) {
    self.Object.Call("getUserMediaTimeout", args)
}

// 
func (self *Video) GetUserMediaError() {
    self.Object.Call("getUserMediaError")
}

// 
func (self *Video) GetUserMediaErrorI(args ...interface{}) {
    self.Object.Call("getUserMediaError", args)
}

// 
func (self *Video) GetUserMediaSuccess() {
    self.Object.Call("getUserMediaSuccess")
}

// 
func (self *Video) GetUserMediaSuccessI(args ...interface{}) {
    self.Object.Call("getUserMediaSuccess", args)
}

// Creates a new Video element from the given Blob. The Blob must contain the video data in the correct encoded format.
// This method is typically called by the Phaser.Loader and Phaser.Cache for you, but is exposed publicly for convenience.
func (self *Video) CreateVideoFromBlob(blob *Blob) *Video{
    return &Video{self.Object.Call("createVideoFromBlob", blob)}
}

// Creates a new Video element from the given Blob. The Blob must contain the video data in the correct encoded format.
// This method is typically called by the Phaser.Loader and Phaser.Cache for you, but is exposed publicly for convenience.
func (self *Video) CreateVideoFromBlobI(args ...interface{}) *Video{
    return &Video{self.Object.Call("createVideoFromBlob", args)}
}

// Creates a new Video element from the given URL.
func (self *Video) CreateVideoFromURL(url string, autoplay bool) *Video{
    return &Video{self.Object.Call("createVideoFromURL", url, autoplay)}
}

// Creates a new Video element from the given URL.
func (self *Video) CreateVideoFromURLI(args ...interface{}) *Video{
    return &Video{self.Object.Call("createVideoFromURL", args)}
}

// Called automatically if the video source changes and updates the internal texture dimensions.
// Then dispatches the onChangeSource signal.
func (self *Video) UpdateTexture(event interface{}, width int, height int) {
    self.Object.Call("updateTexture", event, width, height)
}

// Called automatically if the video source changes and updates the internal texture dimensions.
// Then dispatches the onChangeSource signal.
func (self *Video) UpdateTextureI(args ...interface{}) {
    self.Object.Call("updateTexture", args)
}

// Called when the video completes playback (reaches and ended state).
// Dispatches the Video.onComplete signal.
func (self *Video) Complete() {
    self.Object.Call("complete")
}

// Called when the video completes playback (reaches and ended state).
// Dispatches the Video.onComplete signal.
func (self *Video) CompleteI(args ...interface{}) {
    self.Object.Call("complete", args)
}

// Starts this video playing if it's not already doing so.
func (self *Video) Play(loop bool, playbackRate int) *Video{
    return &Video{self.Object.Call("play", loop, playbackRate)}
}

// Starts this video playing if it's not already doing so.
func (self *Video) PlayI(args ...interface{}) *Video{
    return &Video{self.Object.Call("play", args)}
}

// Called when the video starts to play. Updates the texture.
func (self *Video) PlayHandler() {
    self.Object.Call("playHandler")
}

// Called when the video starts to play. Updates the texture.
func (self *Video) PlayHandlerI(args ...interface{}) {
    self.Object.Call("playHandler", args)
}

// Stops the video playing.
// 
// This removes all locally set signals.
// 
// If you only wish to pause playback of the video, to resume at a later time, use `Video.paused = true` instead.
// If the video hasn't finished downloading calling `Video.stop` will not abort the download. To do that you need to
// call `Video.destroy` instead.
// 
// If you are using a video stream from a webcam then calling Stop will disconnect the MediaStream session and disable the webcam.
func (self *Video) Stop() *Video{
    return &Video{self.Object.Call("stop")}
}

// Stops the video playing.
// 
// This removes all locally set signals.
// 
// If you only wish to pause playback of the video, to resume at a later time, use `Video.paused = true` instead.
// If the video hasn't finished downloading calling `Video.stop` will not abort the download. To do that you need to
// call `Video.destroy` instead.
// 
// If you are using a video stream from a webcam then calling Stop will disconnect the MediaStream session and disable the webcam.
func (self *Video) StopI(args ...interface{}) *Video{
    return &Video{self.Object.Call("stop", args)}
}

// Updates the given Display Objects so they use this Video as their texture.
// This will replace any texture they will currently have set.
func (self *Video) Add(object interface{}) *Video{
    return &Video{self.Object.Call("add", object)}
}

// Updates the given Display Objects so they use this Video as their texture.
// This will replace any texture they will currently have set.
func (self *Video) AddI(args ...interface{}) *Video{
    return &Video{self.Object.Call("add", args)}
}

// Creates a new Phaser.Image object, assigns this Video to be its texture, adds it to the world then returns it.
func (self *Video) AddToWorld(x int, y int, anchorX int, anchorY int, scaleX int, scaleY int) *Image{
    return &Image{self.Object.Call("addToWorld", x, y, anchorX, anchorY, scaleX, scaleY)}
}

// Creates a new Phaser.Image object, assigns this Video to be its texture, adds it to the world then returns it.
func (self *Video) AddToWorldI(args ...interface{}) *Image{
    return &Image{self.Object.Call("addToWorld", args)}
}

// If the game is running in WebGL this will push the texture up to the GPU if it's dirty.
// This is called automatically if the Video is being used by a Sprite, otherwise you need to remember to call it in your render function.
// If you wish to suppress this functionality set Video.disableTextureUpload to `true`.
func (self *Video) Render() {
    self.Object.Call("render")
}

// If the game is running in WebGL this will push the texture up to the GPU if it's dirty.
// This is called automatically if the Video is being used by a Sprite, otherwise you need to remember to call it in your render function.
// If you wish to suppress this functionality set Video.disableTextureUpload to `true`.
func (self *Video) RenderI(args ...interface{}) {
    self.Object.Call("render", args)
}

// Internal handler called automatically by the Video.mute setter.
func (self *Video) SetMute() {
    self.Object.Call("setMute")
}

// Internal handler called automatically by the Video.mute setter.
func (self *Video) SetMuteI(args ...interface{}) {
    self.Object.Call("setMute", args)
}

// Internal handler called automatically by the Video.mute setter.
func (self *Video) UnsetMute() {
    self.Object.Call("unsetMute")
}

// Internal handler called automatically by the Video.mute setter.
func (self *Video) UnsetMuteI(args ...interface{}) {
    self.Object.Call("unsetMute", args)
}

// Internal handler called automatically by the Video.paused setter.
func (self *Video) SetPause() {
    self.Object.Call("setPause")
}

// Internal handler called automatically by the Video.paused setter.
func (self *Video) SetPauseI(args ...interface{}) {
    self.Object.Call("setPause", args)
}

// Internal handler called automatically by the Video.paused setter.
func (self *Video) SetResume() {
    self.Object.Call("setResume")
}

// Internal handler called automatically by the Video.paused setter.
func (self *Video) SetResumeI(args ...interface{}) {
    self.Object.Call("setResume", args)
}

// On some mobile browsers you cannot play a video until the user has explicitly touched the video to allow it.
// Phaser handles this via the `setTouchLock` method. However if you have 3 different videos, maybe an "Intro", "Start" and "Game Over"
// split into three different Video objects, then you will need the user to touch-unlock every single one of them.
// 
// You can avoid this by using just one Video object and simply changing the video source. Once a Video element is unlocked it remains
// unlocked, even if the source changes. So you can use this to your benefit to avoid forcing the user to 'touch' the video yet again.
// 
// As you'd expect there are limitations. So far we've found that the videos need to be in the same encoding format and bitrate.
// This method will automatically handle a change in video dimensions, but if you try swapping to a different bitrate we've found it
// cannot render the new video on iOS (desktop browsers cope better).
// 
// When the video source is changed the video file is requested over the network. Listen for the `onChangeSource` signal to know
// when the new video has downloaded enough content to be able to be played. Previous settings such as the volume and loop state
// are adopted automatically by the new video.
func (self *Video) ChangeSource(src string, autoplay bool) *Video{
    return &Video{self.Object.Call("changeSource", src, autoplay)}
}

// On some mobile browsers you cannot play a video until the user has explicitly touched the video to allow it.
// Phaser handles this via the `setTouchLock` method. However if you have 3 different videos, maybe an "Intro", "Start" and "Game Over"
// split into three different Video objects, then you will need the user to touch-unlock every single one of them.
// 
// You can avoid this by using just one Video object and simply changing the video source. Once a Video element is unlocked it remains
// unlocked, even if the source changes. So you can use this to your benefit to avoid forcing the user to 'touch' the video yet again.
// 
// As you'd expect there are limitations. So far we've found that the videos need to be in the same encoding format and bitrate.
// This method will automatically handle a change in video dimensions, but if you try swapping to a different bitrate we've found it
// cannot render the new video on iOS (desktop browsers cope better).
// 
// When the video source is changed the video file is requested over the network. Listen for the `onChangeSource` signal to know
// when the new video has downloaded enough content to be able to be played. Previous settings such as the volume and loop state
// are adopted automatically by the new video.
func (self *Video) ChangeSourceI(args ...interface{}) *Video{
    return &Video{self.Object.Call("changeSource", args)}
}

// Internal callback that monitors the download progress of a video after changing its source.
func (self *Video) CheckVideoProgress() {
    self.Object.Call("checkVideoProgress")
}

// Internal callback that monitors the download progress of a video after changing its source.
func (self *Video) CheckVideoProgressI(args ...interface{}) {
    self.Object.Call("checkVideoProgress", args)
}

// Sets the Input Manager touch callback to be Video.unlock.
// Required for mobile video unlocking. Mostly just used internally.
func (self *Video) SetTouchLock() {
    self.Object.Call("setTouchLock")
}

// Sets the Input Manager touch callback to be Video.unlock.
// Required for mobile video unlocking. Mostly just used internally.
func (self *Video) SetTouchLockI(args ...interface{}) {
    self.Object.Call("setTouchLock", args)
}

// Enables the video on mobile devices, usually after the first touch.
// If the SoundManager hasn't been unlocked then this will automatically unlock that as well.
// Only one video can be pending unlock at any one time.
func (self *Video) Unlock() {
    self.Object.Call("unlock")
}

// Enables the video on mobile devices, usually after the first touch.
// If the SoundManager hasn't been unlocked then this will automatically unlock that as well.
// Only one video can be pending unlock at any one time.
func (self *Video) UnlockI(args ...interface{}) {
    self.Object.Call("unlock", args)
}

// Grabs the current frame from the Video or Video Stream and renders it to the Video.snapshot BitmapData.
// 
// You can optionally set if the BitmapData should be cleared or not, the alpha and the blend mode of the draw.
// 
// If you need more advanced control over the grabbing them call `Video.snapshot.copy` directly with the same parameters as BitmapData.copy.
func (self *Video) Grab(clear bool, alpha int, blendMode string) *BitmapData{
    return &BitmapData{self.Object.Call("grab", clear, alpha, blendMode)}
}

// Grabs the current frame from the Video or Video Stream and renders it to the Video.snapshot BitmapData.
// 
// You can optionally set if the BitmapData should be cleared or not, the alpha and the blend mode of the draw.
// 
// If you need more advanced control over the grabbing them call `Video.snapshot.copy` directly with the same parameters as BitmapData.copy.
func (self *Video) GrabI(args ...interface{}) *BitmapData{
    return &BitmapData{self.Object.Call("grab", args)}
}

// Removes the Video element from the DOM by calling parentNode.removeChild on itself.
// Also removes the autoplay and src attributes and nulls the reference.
func (self *Video) RemoveVideoElement() {
    self.Object.Call("removeVideoElement")
}

// Removes the Video element from the DOM by calling parentNode.removeChild on itself.
// Also removes the autoplay and src attributes and nulls the reference.
func (self *Video) RemoveVideoElementI(args ...interface{}) {
    self.Object.Call("removeVideoElement", args)
}

// Destroys the Video object. This calls `Video.stop` and then `Video.removeVideoElement`.
// If any Sprites are using this Video as their texture it is up to you to manage those.
func (self *Video) Destroy() {
    self.Object.Call("destroy")
}

// Destroys the Video object. This calls `Video.stop` and then `Video.removeVideoElement`.
// If any Sprites are using this Video as their texture it is up to you to manage those.
func (self *Video) DestroyI(args ...interface{}) {
    self.Object.Call("destroy", args)
}
