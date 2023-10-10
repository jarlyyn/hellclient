package api

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"modules/notifier"
	"modules/version"
	"modules/world"
	"modules/world/bus"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbsecurity/secret/encrypt/aesencrypt"
	"github.com/herb-go/util"

	uuid "github.com/satori/go.uuid"

	"github.com/herb-go/uniqueid"
)

var uniqueNumber = int32(0)

type API struct {
	Bus *bus.Bus
}

func (a *API) Version() string {
	return version.Version.FullVersionCode()
}
func (a *API) Note(cmd string) {
	a.Bus.DoPrint(cmd)
}
func (a *API) SendImmediate(message string) int {
	cmd := world.CreateCommand(message)
	cmd.Creator, cmd.CreatorType = a.Bus.GetScriptCaller()
	a.Bus.DoMetronomeSend(cmd)
	return EOK
}
func (a *API) Send(message string) int {
	cmd := world.CreateCommand(message)
	cmd.Creator, cmd.CreatorType = a.Bus.GetScriptCaller()
	a.Bus.DoMetronomeSend(cmd)
	return EOK
}
func (a *API) SendNoEcho(message string) int {
	cmd := world.CreateCommand(message)
	cmd.Creator, cmd.CreatorType = a.Bus.GetScriptCaller()
	cmd.Echo = false
	a.Bus.DoMetronomeSend(cmd)
	return EOK
}
func (a *API) SendPush(message string) int {
	cmd := world.CreateCommand(message)
	cmd.History = true
	a.Bus.DoMetronomeSend(cmd)
	return EOK
}
func (a *API) SendSpecial(message string, echo bool, queue bool, log bool, history bool) int {
	cmd := world.CreateCommand(message)
	cmd.Creator, cmd.CreatorType = a.Bus.GetScriptCaller()
	cmd.Echo = echo
	cmd.Log = log
	cmd.History = history
	if queue {
		a.Bus.DoSendToQueue(cmd)
	} else {
		a.Bus.DoMetronomeSend(cmd)
	}
	return EOK
}
func (a *API) LogSend(message string) int {
	cmd := world.CreateCommand(message)
	cmd.Creator, cmd.CreatorType = a.Bus.GetScriptCaller()
	cmd.Log = true
	a.Bus.DoMetronomeSend(cmd)
	return EOK
}
func (a *API) Execute(message string) int {
	go a.Bus.DoExecute(message)
	return EOK
}
func (a *API) SendPkt(packet string) int {
	return EOK
}

func (a *API) Connect() int {
	a.Bus.HandleConnError(a.Bus.DoConnectServer())
	return EOK
}
func (a *API) IsConnected() bool {
	return a.Bus.GetConnConnected()
}
func (a *API) Disconnect() int {
	a.Bus.HandleConnError(a.Bus.DoCloseServer())
	return EOK
}
func (a *API) Hash(text string) string {
	result := sha1.Sum([]byte(text))

	return hex.EncodeToString(result[:])
}
func (a *API) Base64Encode(text string, mutliline bool) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	if !mutliline {
		return encoded
	}
	var result = ""
	for len(encoded) > 76 {
		result = result + encoded[:76] + "\n"
		encoded = encoded[76:]
	}
	result = result + encoded
	return result
}
func (a *API) Base64Decode(text string) *string {
	decoded, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return nil
	}
	result := string(decoded)
	return &result
}
func (a *API) GetVariable(text string) string {
	return a.Bus.GetParam(text)
}
func (a *API) SetVariable(name string, content string) int {
	a.Bus.SetParam(name, content)
	return EOK
}
func (a *API) DeleteVariable(name string) int {
	a.Bus.DeleteParam(name)
	return EOK
}
func (a *API) GetVariableList() map[string]string {
	allvar := a.Bus.GetParams()
	result := make(map[string]string, len(allvar))
	for k := range allvar {
		result[k] = allvar[k]
	}
	return result
}
func (a *API) GetVariableComment(text string) string {
	return a.Bus.GetParamComment(text)
}
func (a *API) SetVariableComment(name string, content string) int {
	a.Bus.SetParamComment(name, content)
	return EOK
}
func (a *API) GetUniqueNumber() int {
	v := atomic.AddInt32(&uniqueNumber, 1)
	if v < 0 {
		v = v + 2147483647
	}
	return int(v)
}
func (a *API) GetUniqueID() string {
	return uniqueid.MustGenerateID()
}
func (a *API) CreateGUID() string {
	uid := uuid.NewV1()
	return strings.ToUpper(uid.String())
}

func (a *API) SetStatus(text string) {
	a.Bus.SetStatus(text)
}

func (a *API) GetWorldById(WorldID string) interface{} {
	return nil
}

func (a *API) GetWorld(WorldName string) interface{} {
	return nil
}

func (a *API) GetWorldID() string {
	return a.Bus.ID
}
func (a *API) GetWorldIdList() []string {
	return []string{}
}
func (a *API) GetWorldList() []string {
	return []string{}
}
func (a *API) WorldName() string {
	return a.Bus.ID
}
func (a *API) WorldAddress() string {
	return a.Bus.GetHost()
}
func (a *API) WorldPort() int {
	port, err := strconv.Atoi(a.Bus.GetPort())
	if err != nil {
		return 0
	}
	return port
}
func (a *API) WorldProxy() string {
	return a.Bus.GetProxy()
}
func (a *API) Trim(source string) string {
	return strings.TrimSpace(source)
}

func (a *API) FlashIcon() {}

func (a *API) GetQueue() []string {
	cmds := a.Bus.GetQueue()
	var result = make([]string, len(cmds))
	for k := range cmds {
		result[k] = cmds[k].Mesasge
	}
	return result
}
func (a *API) Queue(message string, echo bool) int {
	cmd := world.CreateCommand(message)
	cmd.Echo = echo
	cmd.Creator, cmd.CreatorType = a.Bus.GetScriptCaller()
	a.Bus.DoSendToQueue(cmd)
	return EOK
}
func (a *API) DiscardQueue(force bool) int {
	return a.Bus.DoDiscardQueue(force)
}
func (a *API) LockQueue() {
	a.Bus.DoLockQueue()
}
func (a *API) SpeedWalkDelay() int {
	return a.Bus.GetQueueDelay()
}
func (a *API) SetSpeedWalkDelay(d int) {
	a.Bus.SetQueueDelay(d)
}
func (a *API) DeleteCommandHistory() {
	a.Bus.FlushHistories()
}

func (a *API) DoAfter(seconds float64, sendtext string) int {
	t := world.CreateTimer()
	t.SetByUser(false)
	t.Enabled = true
	t.OneShot = true
	t.Second = seconds
	t.SendTo = world.SendtoWorld
	t.Send = sendtext
	t.Temporary = true
	a.Bus.AddTimer(t, false)
	return EOK
}
func (a *API) DoAfterNote(seconds float64, sendtext string) int {
	t := world.CreateTimer()
	t.SetByUser(false)
	t.Enabled = true
	t.OneShot = true
	t.Second = seconds
	t.SendTo = world.SendtoOutput
	t.Send = sendtext
	t.Temporary = true
	a.Bus.AddTimer(t, false)
	return EOK
}
func (a *API) DoAfterSpeedWalk(seconds float64, sendtext string) int {
	t := world.CreateTimer()
	t.SetByUser(false)
	t.Enabled = true
	t.OneShot = true
	t.Second = seconds
	t.SendTo = world.SendtoSpeedwalk
	t.Send = sendtext
	t.Temporary = true
	a.Bus.AddTimer(t, false)
	return EOK
}

func (a *API) DoAfterSpecial(seconds float64, sendtext string, sendto int) int {
	t := world.CreateTimer()
	t.Enabled = true
	t.OneShot = true
	t.Second = seconds
	t.SendTo = sendto
	t.Send = sendtext
	t.Temporary = true
	a.Bus.AddTimer(t, false)
	return EOK
}

func (a *API) DeleteGroup(group string) int {
	return a.Bus.DoDeleteTriggerGroup(group, false) + a.Bus.DoDeleteTimerGroup(group, false) + a.Bus.DoDeleteAliasGroup(group, false)
}
func (a *API) AddTimer(timerName string, hour int, minute int, second float64, responseText string, flags int, scriptName string) int {
	t := world.CreateTimer()
	t.Name = timerName
	t.Hour = hour
	t.Minute = minute
	t.Second = second
	t.Send = responseText
	t.Script = scriptName
	t.Enabled = flags&world.TimerFlagEnabled != 0
	t.AtTime = flags&world.TimerFlagAtTime != 0
	t.OneShot = flags&world.TimerFlagOneShot != 0
	t.ActionWhenDisconnectd = flags&world.TimerFlagActiveWhenClosed != 0
	t.Temporary = flags&world.TimerFlagTemporary != 0
	t.SetByUser(false)
	a.Bus.AddTimer(t, flags&world.TimerFlagReplace != 0)
	return EOK
}
func (a *API) DeleteTemporaryTimers() int {
	return a.Bus.DoDeleteTemporaryTimers()
}
func (a *API) DeleteTimer(name string) int {
	name = world.PrefixedName(name, false)
	if !a.Bus.DoDeleteTimerByName(name) {
		return ETimerNotFound
	}
	return EOK
}

func (a *API) DeleteTimerGroup(group string) int {
	return a.Bus.DoDeleteTimerGroup(group, false)
}

func (a *API) EnableTimer(name string, enabled bool) int {
	name = world.PrefixedName(name, false)
	if !a.Bus.DoEnableTimerByName(name, enabled) {
		return ETimerNotFound
	}
	return EOK
}

func (a *API) EnableTimerGroup(group string, enabled bool) int {
	return a.Bus.DoEnableTimerGroup(group, enabled)
}
func (a *API) GetTimerList() []string {
	return a.Bus.DoListTimerNames(false)
}

func (a *API) IsTimer(name string) int {
	name = world.PrefixedName(name, false)
	if !a.Bus.HasNamedTimer(name) {
		return ETimerNotFound
	}
	return EOK
}
func (a *API) ResetTimer(name string) int {
	name = world.PrefixedName(name, false)
	if !a.Bus.DoResetNamedTimer(name) {
		return ETimerNotFound
	}
	return EOK
}
func (a *API) ResetTimers() {
	a.Bus.DoResetTimers()
}

func (a *API) GetTimerOption(name string, option string) (string, int) {
	name = world.PrefixedName(name, false)
	result, ofound, tfound := a.Bus.GetTimerOption(name, option)
	if !tfound {
		return "", ETimerNotFound
	}
	if !ofound {
		return "", EOptionOutOfRange
	}
	return result, EOK
}
func (a *API) SetTimerOption(name string, option string, value string) int {
	name = world.PrefixedName(name, false)
	result, ofound, tfound := a.Bus.SetTimerOption(name, option, value)
	if !tfound {
		return ETimerNotFound
	}
	if !ofound {
		return EOK
	}
	if !result {
		return ETimeInvalid
	}
	return EOK
}

func (a *API) AddAlias(aliasName string, match string, responseText string, flags int, scriptName string) int {
	if match == "" {
		return EAliasCannotBeEmpty
	}
	alias := world.CreateAlias()
	alias.Name = aliasName
	alias.Match = match
	alias.Send = responseText
	alias.Script = scriptName
	alias.Enabled = flags&world.AliasFlagEnabled != 0
	alias.KeepEvaluating = flags&world.AliasFlagKeepEvaluating != 0
	alias.IgnoreCase = flags&world.AliasFlagIgnoreAliasCase != 0
	alias.OmitFromLog = flags&world.AliasFlagOmitFromLogFile != 0
	alias.Regexp = flags&world.AliasFlagRegularExpression != 0
	alias.ExpandVariables = flags&world.AliasFlagExpandVariables != 0
	alias.Temporary = flags&world.AliasFlagTemporary != 0
	if flags&world.AliasFlagAliasSpeedWalk != 0 {
		alias.SendTo = world.SendtoSpeedwalk
	}
	if flags&world.AliasFlagAliasQueue != 0 {
		alias.SendTo = world.SendtoCommandqueue
	}
	alias.Menu = flags&world.AliasFlagAliasMenu != 0
	alias.SetByUser(false)
	ok := a.Bus.AddAlias(alias, flags&world.AliasFlagReplace != 0)
	if !ok {
		return EAliasAlreadyExists
	}
	return EOK
}

func (a *API) DeleteAlias(name string) int {
	name = world.PrefixedName(name, false)
	ok := a.Bus.DoDeleteAliasByName(name)
	if !ok {
		return EAliasNotFound
	}
	return EOK
}
func (a *API) DeleteAliasGroup(group string) int {
	return a.Bus.DoDeleteAliasGroup(group, false)
}
func (a *API) DeleteTemporaryAliases() int {
	return a.Bus.DoDeleteTemporaryAliases()
}
func (a *API) EnableAlias(name string, enabled bool) int {
	name = world.PrefixedName(name, false)
	ok := a.Bus.DoEnableAliasByName(name, enabled)
	if !ok {
		return EAliasNotFound
	}
	return EOK
}

func (a *API) EnableAliasGroup(group string, enabled bool) int {
	return a.Bus.DoEnableAliasGroup(group, enabled)
}

func (a *API) GetAliasList() []string {
	return a.Bus.DoListAliasNames(false)
}

func (a *API) GetAliasOption(name string, option string) (string, int) {
	name = world.PrefixedName(name, false)
	result, ofound, tfound := a.Bus.GetAliasOption(name, option)
	if !tfound {
		return "", ETimerNotFound
	}
	if !ofound {
		return "", EOptionOutOfRange
	}
	return result, EOK

}

func (a *API) IsAlias(name string) int {
	name = world.PrefixedName(name, false)
	if !a.Bus.HasNamedAlias(name) {
		return EAliasNotFound
	}
	return EOK
}

func (a *API) SetAliasOption(name string, option string, value string) int {
	name = world.PrefixedName(name, false)
	_, ofound, tfound := a.Bus.SetAliasOption(name, option, value)
	if !tfound {
		return EAliasNotFound
	}
	if !ofound {
		return EOK
	}
	return EOK
}

func (a *API) AddTrigger(triggerName string, match string, responseText string, flags int, colour int, wildcard int, soundFileName string, scriptName string) int {
	if match == "" {
		return ETriggerCannotBeEmpty
	}
	trigger := world.CreateTrigger()
	trigger.Name = triggerName
	trigger.Match = match
	trigger.Send = responseText
	trigger.Colour = colour
	trigger.SoundFileName = soundFileName
	trigger.Script = scriptName
	trigger.Enabled = flags&world.TriggerFlagEnabled != 0
	trigger.KeepEvaluating = flags&world.TriggerFlagKeepEvaluating != 0
	trigger.IgnoreCase = flags&world.TriggerFlagIgnoreCase != 0
	trigger.OmitFromLog = flags&world.TriggerFlagOmitFromLog != 0
	trigger.Regexp = flags&world.TriggerFlagRegularExpression != 0
	trigger.ExpandVariables = flags&world.TriggerFlagExpandVariables != 0
	trigger.Temporary = flags&world.TriggerFlagTemporary != 0
	trigger.SetByUser(false)
	ok := a.Bus.AddTrigger(trigger, flags&world.TriggerFlagReplace != 0)
	if !ok {
		return ETriggerAlreadyExists
	}
	return EOK
}

func (a *API) AddTriggerEx(triggerName string, match string, responseText string, flags int, colour int, wildcard int, soundFileName string, scriptName string, sendTo int, sequence int) int {
	if match == "" {
		return ETriggerCannotBeEmpty
	}
	trigger := world.CreateTrigger()
	trigger.Name = triggerName
	trigger.Match = match
	trigger.Send = responseText
	trigger.Colour = colour
	trigger.SoundFileName = soundFileName
	trigger.SendTo = sendTo
	trigger.Sequence = sequence
	trigger.Script = scriptName
	trigger.Enabled = flags&world.TriggerFlagEnabled != 0
	trigger.KeepEvaluating = flags&world.TriggerFlagKeepEvaluating != 0
	trigger.IgnoreCase = flags&world.TriggerFlagIgnoreCase != 0
	trigger.OmitFromLog = flags&world.TriggerFlagOmitFromLog != 0
	trigger.Regexp = flags&world.TriggerFlagRegularExpression != 0
	trigger.ExpandVariables = flags&world.TriggerFlagExpandVariables != 0
	trigger.Temporary = flags&world.TriggerFlagTemporary != 0
	trigger.SetByUser(false)
	ok := a.Bus.AddTrigger(trigger, flags&world.TriggerFlagReplace != 0)
	if !ok {
		return ETriggerAlreadyExists
	}
	return EOK
}

func (a *API) DeleteTrigger(name string) int {
	name = world.PrefixedName(name, false)
	ok := a.Bus.DoDeleteTrigger(name)
	if !ok {
		return ETriggerNotFound
	}
	return EOK
}
func (a *API) DeleteTriggerGroup(group string) int {
	return a.Bus.DoDeleteTriggerGroup(group, false)
}
func (a *API) EnableTrigger(name string, enabled bool) int {
	name = world.PrefixedName(name, false)
	ok := a.Bus.DoEnableTriggerByName(name, enabled)
	if !ok {
		return ETriggerNotFound
	}
	return EOK
}

func (a *API) EnableTriggerGroup(group string, enabled bool) int {
	return a.Bus.DoEnableTriggerGroup(group, enabled)
}

func (a *API) GetTriggerList() []string {
	return a.Bus.DoListTriggerNames(false)
}

func (a *API) GetTriggerOption(name string, option string) (string, int) {
	name = world.PrefixedName(name, false)
	result, ofound, tfound := a.Bus.GetTriggerOption(name, option)
	if !tfound {
		return "", ETimerNotFound
	}
	if !ofound {
		return "", EOptionOutOfRange
	}
	return result, EOK

}

func (a *API) IsTrigger(name string) int {
	name = world.PrefixedName(name, false)
	if !a.Bus.HasNamedTrigger(name) {
		return ETriggerNotFound
	}
	return EOK
}

func (a *API) SetTriggerOption(name string, option string, value string) int {
	name = world.PrefixedName(name, false)
	_, ofound, tfound := a.Bus.SetTriggerOption(name, option, value)
	if !tfound {
		return ETriggerNotFound
	}
	if !ofound {
		return EOK
	}
	return EOK
}
func (a *API) StopEvaluatingTriggers() {
	a.Bus.DoStopEvaluatingTriggers()
}
func (a *API) GetTriggerWildcard(triggername string, wildcard string) *string {
	triggername = world.PrefixedName(triggername, false)
	w := a.Bus.DoGetTriggerWildcard(triggername)
	if w == nil {
		return nil
	}
	var result, ok = w.Named[wildcard]
	if ok {
		return &result
	}
	var index, err = strconv.Atoi(wildcard)
	if err != nil || index < 0 || index >= len(w.List) {
		return nil

	}
	result = w.List[index]
	return &result
}
func (a *API) ColourNameToRGB(v string) int {
	color, ok := world.NamedColor[v]
	if !ok {
		return -1
	}
	return color
}
func (a *API) MustCheckHome() {
	home := a.Bus.GetScriptHome()
	if home == "" {
		return
	}
	_, err := os.Stat(home)
	if err != nil {
		id := a.Bus.GetScriptID()
		if id == "" {
			return
		}
	}
}
func (a *API) MustCleanHomeFileInsidePath(name string) string {
	home := a.Bus.GetScriptHome()
	name = herbplugin.MustCleanPath(home, name)
	if name == "" {
		return name
	}
	if !strings.HasPrefix(name, home) {
		return ""
	}
	return name
}
func (a *API) HasHomeFile(p herbplugin.Plugin, name string) bool {
	a.MustCheckHome()
	filename := a.MustCleanHomeFileInsidePath(name)
	if filename == "" {
		panic(fmt.Errorf("read %s not allowed", name))
	}
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}
func (a *API) ReadHomeFile(p herbplugin.Plugin, name string) string {
	a.MustCheckHome()
	filename := a.MustCleanHomeFileInsidePath(name)
	if filename == "" {
		panic(fmt.Errorf("read %s not allowed", name))
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}
func (a *API) ReadHomeLines(p herbplugin.Plugin, name string) []string {
	data := a.ReadHomeFile(p, name)
	return strings.Split(lineReplacer.Replace(data), "\n")
}
func (a *API) WriteHomeFile(p herbplugin.Plugin, name string, body []byte) {
	a.MustCheckHome()
	filename := a.MustCleanHomeFileInsidePath(name)
	if filename == "" {
		panic(fmt.Errorf("write %s not allowed", name))
	}
	err := os.WriteFile(filename, body, util.DefaultFileMode)
	if err != nil {
		panic(err)
	}
}

func (a *API) MakeHomeFolder(p herbplugin.Plugin, name string) bool {
	a.MustCheckHome()
	filename := a.MustCleanHomeFileInsidePath(name)
	if filename == "" {
		panic(fmt.Errorf("make folder %s not allowed", name))
	}
	err := os.MkdirAll(filename, util.DefaultFolderMode)
	if err != nil {
		if os.IsExist(err) {
			return false
		}
		panic(err)
	}
	return true
}

func (a *API) MustCleanModFileInsidePath(name string) string {
	if !a.Bus.GetModEnabled() {
		return ""
	}
	sid := a.Bus.GetScriptID()
	if sid == "" {
		return ""
	}
	modpath := filepath.Join(a.Bus.GetModPath(), sid+".mod")
	name = herbplugin.MustCleanPath(modpath, name)
	if name == "" {
		return name
	}
	if !strings.HasPrefix(name, modpath) {
		return ""
	}
	return name
}

func (a *API) HasModFile(p herbplugin.Plugin, name string) bool {
	if !a.Bus.GetModEnabled() {
		return false
	}
	filename := a.MustCleanModFileInsidePath(name)
	if filename == "" {
		panic(fmt.Errorf("read %s not allowed", name))
	}
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}
func (a *API) ReadModFile(p herbplugin.Plugin, name string) string {
	filename := a.MustCleanModFileInsidePath(name)
	if filename == "" {
		panic(fmt.Errorf("read %s not allowed", name))
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}
func (a *API) ReadModLines(p herbplugin.Plugin, name string) []string {
	data := a.ReadModFile(p, name)
	return strings.Split(lineReplacer.Replace(data), "\n")
}

func (a *API) HasFile(p herbplugin.Plugin, name string) bool {
	o := p.PluginOptions()
	filename := o.GetLocation().MustCleanInsidePath(name)
	if filename == "" {
		panic(fmt.Errorf("read %s not allowed", name))
	}
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		panic(err)
	}
	return true
}
func (a *API) ReadFile(p herbplugin.Plugin, name string) string {
	o := p.PluginOptions()
	filename := o.GetLocation().MustCleanInsidePath(name)
	if filename == "" {
		panic(fmt.Errorf("read %s not allowed", name))
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

var lineReplacer = strings.NewReplacer("\r\n", "\n", "\n\r", "\n")

func (a *API) ReadLines(p herbplugin.Plugin, name string) []string {
	data := a.ReadFile(p, name)
	return strings.Split(lineReplacer.Replace(data), "\n")
}

func (a *API) SplitN(text string, sep string, n int) []string {
	return strings.SplitN(text, sep, n)
}

func (a *API) UTF8Len(text string) int {
	v := []rune(text)
	return len(v)
}
func (a *API) UTF8Index(text string, substring string) int {
	idx := strings.Index(text, substring)
	if idx <= 0 {
		return idx
	}
	return len([]rune(text[:idx]))
}
func (a *API) UTF8Sub(text string, start int, end int) string {
	v := []rune(text)
	if start < 0 {
		start = 0
	}
	if start >= len(v) {
		return ""
	}
	if end > len(v) || end <= 0 {
		end = len(v)
	}
	return string(v[start:end])
}
func (a *API) ToUTF8(code string, text string) *string {
	result, err := world.ToUTF8(code, []byte(text))
	if err != nil {
		return nil
	}
	output := string(result)
	return &output
}
func (a *API) FromUTF8(code string, text string) *string {
	result, err := world.FromUTF8(code, []byte(text))
	if err != nil {
		return nil
	}
	output := string(result)
	return &output
}
func (a *API) Info(text string) {
	a.Bus.SetStatus(a.Bus.GetStatus() + text)
}
func (a *API) InfoClear() {
	a.Bus.SetStatus("")
}

func (a *API) GetAlphaOption(name string) string {
	switch name {
	case "name":
		return a.Bus.GetName()
	case "id":
		return a.Bus.ID
	case "command_stack_character":
		return a.Bus.GetCommandStackCharacter()
	case "script_prefix":
		return a.Bus.GetScriptPrefix()
	}
	panic(fmt.Errorf("world alpha option %s not supported", name))
}

func (a *API) SetAlphaOption(name string, value string) int {
	switch name {
	case "name":
		a.Bus.SetName(value)
	default:
		panic(fmt.Errorf("world alpha option %s not supported", name))
	}
	return EOK
}

func (a *API) GetLinesInBufferCount() int {
	return a.Bus.GetLinesInBufferCount()
}
func (a *API) DeleteOutput() {

}

func (a *API) DeleteLines(count int) {
	a.Bus.DoDeleteLines(count)
}

func (a *API) GetLineCount() int {
	return a.Bus.GetLineCount()
}

func (a *API) GetRecentLines(count int) string {
	recent := a.Bus.GetMaxRecent()
	if count > recent {
		count = recent
	}
	lines := a.Bus.GetRecentLines(count)
	var result = make([]string, 0, len(lines))
	for _, v := range lines {
		result = append(result, v.Plain())
	}
	return strings.Join(result, "\n")
}

func (a *API) GetLineInfo(linenumber int, infotype int) (string, bool) {
	line := a.Bus.GetLine(linenumber)
	if line == nil {
		return "", false
	}
	switch infotype {
	case 1:
		return line.Plain(), true
	case 2:
		return strconv.Itoa(len(line.Plain())), true
	case 3:
		return world.ToStringBool(line.IsNewline()), true
	case 4:
		return world.ToStringBool(line.Type == world.LineTypePrint), true
	case 5:
		return world.ToStringBool(line.Type == world.LineTypeEcho), true
	case 6:
		return world.ToStringBool(!line.OmitFromLog), true
	case 7:
		return world.ToStringBool(false), true
	case 8:
		return world.ToStringBool(false), true
	case 9:
		return strconv.FormatInt(line.Time, 10), true
	case 10:
		return line.ID, true
	case 11:
		return strconv.Itoa(len(line.Words)), true
	}
	return "", false
}
func (a *API) BoldColour(WhichColour int) int {
	//bold colour should equal to normalcolour
	return world.GetNormalColour(WhichColour)
}
func (a *API) NormalColour(WhichColour int) int {
	return world.GetNormalColour(WhichColour)
}
func (a *API) GetStyleInfo(linenumber int, style int, infotype int) (string, bool) {
	line := a.Bus.GetLine(linenumber)
	if line == nil {
		return "", false
	}
	if style < 1 || style > len(line.Words) {
		return "", false
	}
	word := line.Words[style-1]
	switch infotype {
	case 1:
		return word.Text, true
	case 2:
		return strconv.Itoa(len(word.Text)), true
	case 3:
		sc := line.GetWordStartColumn(style)
		return strconv.Itoa(sc), true
	case 8:
		return world.ToStringBool(word.Bold), true
	case 9:
		return world.ToStringBool(word.Underlined), true
	case 10:
		return world.ToStringBool(word.Blinking), true
	case 11:
		return world.ToStringBool(word.Inverse), true
	case 14:
		return strconv.Itoa(word.GetColorRGB()), true
	case 15:
		return strconv.Itoa(word.GetBGColorRGB()), true
	}
	return "", false

}
func (a *API) WriteLog(message string) int {
	a.Bus.DoLog(message)
	return EOK
}
func (a *API) CloseLog() int {
	return EOK
}
func (a *API) FlushLog() int {
	return EOK
}
func (a *API) OpenLog() int {
	return EOK
}
func (a *API) GetGlobalOption(optionname string) string {
	switch optionname {
	case "AllTypingToCommandWindow":
	case "AlwaysOnTop":
	case "AppendToLogFiles":
	case "AutoConnectWorlds":
	case "AutoExpandConfig":
	case "FlatToolbars":
	case "AutoLogWorld":
	case "BleedBackground":
	case "ColourGradientConfig":
	case "ConfirmBeforeClosingMXPdebug":
	case "ConfirmBeforeClosingMushclient":
	case "ConfirmBeforeClosingWorld":
	case "ConfirmBeforeSavingVariables":
	case "ConfirmLogFileClose":
	case "EnableSpellCheck":
	case "AllowLoadingDlls":
	case "F1macro":
	case "FixedFontForEditing":
	case "NotepadWordWrap":
	case "NotifyIfCannotConnect":
	case "ErrorNotificationToOutputWindow":
	case "NotifyOnDisconnect":
	case "OpenActivityWindow":
	case "OpenWorldsMaximised":
	case "WindowTabsStyle":
	case "ReconnectOnLinkFailure":
	case "RegexpMatchEmpty":
	case "ShowGridLinesInListViews":
	case "SmoothScrolling":
	case "SmootherScrolling":
	case "DisableKeyboardMenuActivation":
	case "TriggerRemoveCheck":
	case "NotepadBackColour":
	case "NotepadTextColour":
	case "ActivityButtonBarStyle":
	case "AsciiArtLayout":
	case "DefaultInputFontHeight":
	case "DefaultInputFontItalic ":
	case "DefaultInputFontWeight":
	case "DefaultOutputFontHeight":
	case "Icon Placement":
	case "Tray Icon":
	case "ActivityWindowRefreshInterval":
	case "ParenMatchFlags":
	case "PrinterFontSize":
	case "PrinterLeftMargin":
	case "PrinterLinesPerPage":
	case "PrinterTopMargin":
	case "FixedPitchFontSize":
	case "TabInsertsTabInMultiLineDialogs":
	case "AsciiArtFont":
	case "FixedPitchFont":
	case "WordDelimitersDblClick":
		return "0"
	case "TimerInterval":
		return "0"
	case "ActivityWindowRefreshType":
	case "PluginList":
	case "PluginsDirectory":
	case "StateFilesDirectory":
	case "PrinterFont":
	case "TrayIconFileName":
	case "WordDelimiters":
	case "WorldList":
	case "LuaScript":
	case "Locale":
	case "DefaultAliasesFile":
	case "DefaultColoursFile":
	case "DefaultInputFont":
	case "DefaultLogFileDirectory":
	case "DefaultMacrosFile":
	case "DefaultNameGenerationFile":
	case "DefaultOutputFont ":
	case "DefaultTimersFile ":
	case "DefaultTriggersFile":
	case "DefaultWorldFileDirectory":
	case "NotepadQuoteString":
		return ""
	}
	return ""
}
func (a *API) GetInfo(infotype int) string {
	switch infotype {
	case 1:
		return a.Bus.GetHost()
	case 2:
		return a.Bus.GetName()
	case 8:
		return ""
	case 28:
		return a.Bus.GetScriptType()
	case 35:
		return a.Bus.GetScriptID()
	case 36:
		return a.Bus.GetScriptPrefix()
	case 40:
		return a.Bus.ID + ".log"
	case 51:
		return a.Bus.ID + ".log"
	case 53:
		return a.Bus.GetStatus()
	case 54:
		return a.Bus.ID + ".toml"
	case 55:
		return a.Bus.ID
	case 56:
		return "hellclient"
	case 57:
		return "./"
	case 58:
		return "./"
	case 59:
		return "./"
	case 64:
		return "./"
	case 66:
		return "./"
	case 67:
		return "./"
	case 68:
		return "./"

	}
	panic(fmt.Errorf("unknown world.GetInfo infotype %d", infotype))
}

func (a *API) GetTimerInfo(name string, infotype int) (string, int) {
	name = world.PrefixedName(name, false)
	result, ofound, tfound := a.Bus.GetTimerInfo(name, infotype)
	if !tfound {
		return "", ETimerNotFound
	}
	if !ofound {
		panic(fmt.Errorf("unknown world.GetTimerInfo infotype %d", infotype))
	}
	return result, EOK

}

func (a *API) GetTriggerInfo(name string, infotype int) (string, int) {
	name = world.PrefixedName(name, false)
	result, ofound, tfound := a.Bus.GetTriggerInfo(name, infotype)
	if !tfound {
		return "", ETriggerNotFound
	}
	if !ofound {
		panic(fmt.Errorf("unknown world.GetTriggerInfo infotype %d", infotype))
	}
	return result, EOK
}

func (a *API) GetAliasInfo(name string, infotype int) (string, int) {
	name = world.PrefixedName(name, false)
	result, ofound, tfound := a.Bus.GetAliasInfo(name, infotype)
	if !tfound {
		return "", EAliasNotFound
	}
	if !ofound {
		panic(fmt.Errorf("unknown world.GetAliasInfo infotype %d", infotype))
	}
	return result, EOK
}
func (a *API) Broadcast(msg string, gloabl bool) {
	if msg == "" {
		return
	}
	channel := a.Bus.GetScriptData().Channel
	if channel == "" {
		return
	}
	bc := world.CreateBroadcast(channel, msg, gloabl)
	a.Bus.RaiseBroadcastEvent(bc)
	if a.Bus.GetShowBroadcast() {
		if gloabl {
			a.Bus.DoPrintGlobalBroadcastOut(msg)
		} else {
			a.Bus.DoPrintLocalBroadcastOut(msg)
		}
	}
}

func (a *API) Notify(title string, body string) {
	notifier.DefaultNotifier.WorldNotify(a.Bus.ID, title, body)
}

func (a *API) CheckPermissions(p []string) bool {
	permissions := a.Bus.GetPermissions()
NEED:
	for _, need := range p {
		for _, own := range permissions {
			if own == need {
				continue NEED
			}
		}
		return false
	}
	return true
}
func (a *API) RequestPermissions(permissions []string, reason string, script string) {
	a.Bus.RequestPermissions(world.CreateAuthorization(a.Bus.ID, permissions, reason, script))
}
func (a *API) CheckTrustedDomains(d []string) bool {
	domains := a.Bus.GetTrusted().Domains
NEED:
	for _, need := range d {
		for _, own := range domains {
			if own == need {
				continue NEED
			}
		}
		return false
	}
	return true
}

func (a *API) RequestTrustDomains(domains []string, reason string, script string) {
	a.Bus.RequestTrustDomains(world.CreateAuthorization(a.Bus.ID, domains, reason, script))
}

func (a *API) Encrypt(data, key string) *string {
	result, err := aesencrypt.AESNonceEncryptBase64([]byte(data), []byte(key))
	if err != nil {
		return nil
	}
	return &result
}

func (a *API) Decrypt(data string, key string) *string {
	result, err := aesencrypt.AESNonceDecryptBase64(data, []byte(key))
	if err != nil {
		return nil
	}
	str := string(result)
	return &str
}

func (a *API) Request(reqtype string, data string) string {
	msg := world.CreateMessage(a.Bus.ID, reqtype, data)
	a.Bus.RaiseRequestEvent(msg)
	if a.Bus.GetShowBroadcast() {
		a.Bus.DoPrintRequest(msg.Desc())
	}
	return msg.ID
}

func (a *API) DumpOutput(length int, offset int) string {
	if length < 0 {
		length = 0
	}
	if offset < 0 {
		offset = 0
	}
	line := a.Bus.GetRecentLines(offset + length)
	if length > len(line) {
		length = len(line)
	}
	var data []byte
	var err error
	data, err = json.Marshal(line[0:length])
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (a *API) ConcatOutput(output1 string, output2 string) string {
	var list1 = []*world.Line{}
	var list2 = []*world.Line{}
	var err error
	err = json.Unmarshal([]byte(output1), &list1)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(output2), &list2)
	if err != nil {
		panic(err)
	}
	list1 = append(list1, list2...)
	result, err := json.Marshal(list1)
	if err != nil {
		panic(err)
	}
	return string(result)
}
func (a *API) SliceOutput(output string, start int, end int) string {
	var list = []*world.Line{}
	err := json.Unmarshal([]byte(output), &list)
	if err != nil {
		panic(err)
	}
	if start < 0 {
		start = 0
	}
	if start >= len(list) {
		start = len(list) - 1
	}
	if end <= 0 || end > len(list) {
		end = len(list)
	}
	if end < start {
		end = start - 1
	}
	data, err := json.Marshal(list[start:end])
	if err != nil {
		panic(err)
	}
	return string(data)
}
func (a *API) OutputToText(output string) string {
	var list = []*world.Line{}
	err := json.Unmarshal([]byte(output), &list)
	if err != nil {
		panic(err)
	}
	var lines = make([]string, len(list))
	for k := range list {
		var words = make([]string, len(list[k].Words))
		for wordsindex := range words {
			words[wordsindex] = list[k].Words[wordsindex].Text
		}
		lines[k] = strings.Join(words, "")
	}
	return strings.Join(lines, "\n")
}

func (a *API) FormatOutput(output string) string {
	var list = []*world.Line{}
	err := json.Unmarshal([]byte(output), &list)
	if err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(data)

}

func (a *API) PrintOutput(text string) string {
	line := world.NewLine()
	line.Type = world.LineTypePrint
	word := world.NewWord()
	word.Text = text
	line.Words = append(line.Words, word)
	list := []*world.Line{line}
	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(data)

}

func (a *API) Simulate(text string) {
	var list = strings.Split(text, "\n")
	go func() {
		for _, t := range list {
			line := world.NewLine()
			line.Type = world.LineTypeReal
			word := world.NewWord()
			word.Text = t
			line.Words = append(line.Words, word)
			a.Bus.RaiseLineEvent(line)
		}
	}()
}

func (a *API) SimulateOutput(output string) {
	var list = []*world.Line{}
	err := json.Unmarshal([]byte(output), &list)
	if err != nil {
		panic(err)
	}
	go func() {
		for _, line := range list {
			line.ID = uniqueid.MustGenerateID()
			a.Bus.RaiseLineEvent(line)
		}
	}()
}

func (a *API) DumpTriggers(byUser bool) string {
	t := a.Bus.GetTriggersByType(byUser)
	data, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(data)
}
func (a *API) RestoreTriggers(data string, byUser bool) {
	var list = []*world.Trigger{}
	err := json.Unmarshal([]byte(data), &list)
	if err != nil {
		panic(err)
	}
	for _, v := range list {
		v.SetByUser(byUser)
	}
	a.Bus.AddTriggers(list)
}
func (a *API) DumpTimers(byUser bool) string {
	t := a.Bus.GetTimersByType(byUser)
	data, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(data)
}
func (a *API) RestoreTimers(data string, byUser bool) {
	var list = []*world.Timer{}
	err := json.Unmarshal([]byte(data), &list)
	if err != nil {
		panic(err)
	}
	for _, v := range list {
		v.SetByUser(byUser)
	}
	a.Bus.AddTimers(list)
}
func (a *API) DumpAliases(byUser bool) string {
	al := a.Bus.GetAliasesByType(byUser)
	data, err := json.Marshal(al)
	if err != nil {
		panic(err)
	}
	return string(data)
}
func (a *API) RestoreAliases(data string, byUser bool) {
	var list = []*world.Alias{}
	err := json.Unmarshal([]byte(data), &list)
	if err != nil {
		panic(err)
	}
	for _, v := range list {
		v.SetByUser(byUser)
	}
	a.Bus.AddAliases(list)
}

func (a *API) SetHUDSize(size int) {
	a.Bus.SetHUDSize(size)
}

func (a *API) GetHUDContent() string {
	var content = a.Bus.GetHUDContent()
	data, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (a *API) GetHUDSize() int {
	return a.Bus.GetHUDSize()
}
func (a *API) UpdateHUD(start int, content string) bool {
	var lines = []*world.Line{}
	err := json.Unmarshal([]byte(content), &lines)
	if err != nil {
		panic(err)
	}
	return a.Bus.UpdateHUDContent(start, lines)
}
func (a *API) NewLine() string {
	line := world.NewLine()
	line.Type = world.LineTypeReal
	data, err := json.Marshal(line)
	if err != nil {
		panic(err)
	}
	return string(data)
}
func (a *API) NewWord(value string) string {
	word := world.NewWord()
	word.Text = value
	data, err := json.Marshal(word)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (a *API) GetModInfo(herbplugin.Plugin) *world.Mod {
	mod := world.NewMod()
	if !a.Bus.GetModEnabled() {
		return mod
	}
	mod.Enabled = true
	modpath := a.MustCleanModFileInsidePath("")
	if modpath == "" {
		panic(fmt.Errorf("get mod info not allowed"))
	}
	files, err := ioutil.ReadDir(modpath)
	if err != nil {
		if os.IsNotExist(err) {
			return mod
		}
		panic(err)
	}
	mod.Exists = true
	for _, f := range files {
		name := filepath.Base(f.Name())
		if f.IsDir() {
			mod.FolderList = append(mod.FolderList, name)
		} else {
			mod.FileList = append(mod.FileList, name)
		}
	}
	sort.Strings(mod.FolderList)
	sort.Strings(mod.FileList)
	return mod
}

func (a *API) SetPriority(value int) {
	a.Bus.SetPriority(value)
}
func (a *API) GetPriority() int {
	return a.Bus.GetPriority()
}
func (a *API) SetSummary(content string) {
	var lines = []*world.Line{}
	err := json.Unmarshal([]byte(content), &lines)
	if err != nil {
		panic(err)
	}
	a.Bus.SetSummary(lines)
}
func (a *API) GetSummary() string {
	var content = a.Bus.GetSummary()
	data, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}
	return string(data)
}
