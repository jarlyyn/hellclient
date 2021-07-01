package api

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"modules/version"
	"modules/world"
	"modules/world/bus"
	"strconv"
	"strings"
	"sync/atomic"

	uuid "github.com/satori/go.uuid"

	"github.com/herb-go/uniqueid"
)

var uniqueNumber = int32(0)

type API struct {
	Bus *bus.Bus
}

func (a *API) Version() string {
	return version.Version
}
func (a *API) Note(cmd string) {
	a.Bus.DoPrint(cmd)
}
func (a *API) SendImmediate(message string) int {
	cmd := world.CreateCommand(message)
	a.Bus.DoSend(cmd)
	return EOK
}
func (a *API) Send(message string) int {
	cmd := world.CreateCommand(message)
	a.Bus.DoSendToQueue(cmd)
	return EOK
}
func (a *API) SendNoEcho(message string) int {
	cmd := world.CreateCommand(message)
	cmd.Echo = false
	a.Bus.DoSend(cmd)
	return EOK
}
func (a *API) SendPush(message string) int {
	cmd := world.CreateCommand(message)
	cmd.History = true
	a.Bus.DoSend(cmd)
	return EOK
}
func (a *API) SendSpecial(message string, echo bool, queue bool, log bool, history bool) int {
	cmd := world.CreateCommand(message)
	cmd.Echo = echo
	cmd.Log = log
	cmd.History = history
	if queue {
		a.Bus.DoSendToQueue(cmd)
	} else {
		a.Bus.DoSend(cmd)
	}
	return EOK
}
func (a *API) LogSend(message string) int {
	cmd := world.CreateCommand(message)
	cmd.Log = true
	a.Bus.DoSend(cmd)
	return EOK
}
func (a *API) Execute(message string) int {
	a.Bus.DoExecute(message)
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
func (a *API) GetVariable(text string) *string {
	allvar := a.Bus.GetParams()
	val, ok := allvar[text]
	if !ok {
		return nil
	}
	return &val
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
	id, err := uuid.NewV1()
	if err != nil {
		panic(err)
	}
	return id.String()
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
func (a *API) Queue(message string) int {
	cmd := world.CreateCommand(message)
	a.Bus.DoSendToQueue(cmd)
	return EOK
}
func (a *API) DiscardQueue() int {
	return a.Bus.DoDiscardQueue()
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
	t.Second = int(seconds)
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
	t.Second = int(seconds)
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
	t.Second = int(seconds)
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
	t.Second = int(seconds)
	t.SendTo = sendto
	t.Send = sendtext
	t.Temporary = true
	a.Bus.AddTimer(t, false)
	return EOK
}

func (a *API) AddTimer(timerName string, hour int, minute int, second float64, responseText string, flags int, scriptName string) int {
	t := world.CreateTimer()
	t.Hour = hour
	t.Minute = minute
	t.Second = int(second)
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
	return a.Bus.DoDeleteTimerGroup(group)
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
	return a.Bus.DoListTimerNames()
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
	ok := a.Bus.DoDeleteAlias(name)
	if !ok {
		return EAliasNotFound
	}
	return EOK
}
func (a *API) DeleteAliasGroup(group string) int {
	return a.Bus.DoDeleteAliasGroup(group)
}
func (a *API) EnableAlias(name string, enabled bool) int {
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
	return a.Bus.DoListAliasNames()
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
	ok := a.Bus.DoDeleteTrigger(name)
	if !ok {
		return ETriggerNotFound
	}
	return EOK
}
func (a *API) DeleteTriggerGroup(group string) int {
	return a.Bus.DoDeleteTriggerGroup(group)
}
func (a *API) EnableTrigger(name string, enabled bool) int {
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
	return a.Bus.DoListTriggerNames()
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

func (a *API) ColourNameToRGB(v string) string {
	return v
}
