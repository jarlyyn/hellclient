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

const EOK = 0

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
func (a *API) DiscardQueue() int {
	return a.Bus.DoDiscardQueue()
}
func (a *API) SpeedWalkDelay() int {
	return a.Bus.GetQueueDelay()
}

func (a *API) DeleteCommandHistory() {
	a.Bus.FlushHistories()
}
