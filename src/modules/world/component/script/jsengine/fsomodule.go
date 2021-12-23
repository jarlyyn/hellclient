package jsengine

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"hellclient/modules/world/bus"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/dop251/goja"
	"github.com/herb-go/herbplugin"
	"github.com/herb-go/herbplugin/jsplugin"
	"github.com/herb-go/uniqueid"
	"github.com/herb-go/util"
)

type TextStream struct {
	File   *os.File
	Closer func()
}

func (t *TextStream) Close(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	t.Closer()
	t.File = nil
	return nil
}
func (t *TextStream) Read(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	if t.File == nil {
		panic(errors.New("textstream closed"))
	}
	characters := int(call.Argument(0).ToInteger())
	buf := make([]byte, characters)
	_, err := t.File.Read(buf)
	if err != nil {
		panic(err)
	}
	return r.ToValue(string(buf))
}
func (t *TextStream) Skip(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	t.Read(call, r)
	return nil
}
func (t *TextStream) ReadLine(call goja.FunctionCall, r *goja.Runtime) goja.Value {

	if t.File == nil {
		panic(errors.New("textstream closed"))
	}
	reader := bufio.NewReader(t.File)
	result := ""
	for {
		buf, prefix, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		result = result + string(buf)
		if !prefix {
			break
		}
	}
	return r.ToValue(result)
}
func (t *TextStream) SkipLine(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	t.ReadLine(call, r)
	return nil
}
func (t *TextStream) ReadAll(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	if t.File == nil {
		panic(errors.New("textstream closed"))
	}
	buf, err := io.ReadAll(t.File)
	if err != nil {
		panic(err)
	}
	return r.ToValue(string(buf))
}
func (t *TextStream) Write(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	if t.File == nil {
		panic(errors.New("textstream closed"))
	}
	_, err := t.File.WriteString(call.Argument(0).String())
	if err != nil {
		panic(err)
	}
	return nil
}
func (t *TextStream) WriteBlankLines(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	if t.File == nil {
		panic(errors.New("textstream closed"))
	}
	lines := int(call.Argument(0).ToInteger())
	if lines < 0 {
		return nil
	}
	_, err := t.File.WriteString(strings.Repeat("\n", lines))
	if err != nil {
		panic(err)
	}
	return nil
}
func (t *TextStream) WriteLine(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	if t.File == nil {
		panic(errors.New("textstream closed"))
	}
	_, err := t.File.WriteString(call.Argument(0).String() + "\n")
	if err != nil {
		panic(err)
	}
	return nil
}
func (t *TextStream) Convert(r *goja.Runtime) goja.Value {
	obj := r.NewObject()
	obj.Set("Close", t.Close)
	obj.Set("Read", t.Read)
	obj.Set("ReadAll", t.ReadAll)
	obj.Set("ReadLine", t.ReadLine)
	obj.Set("Skip", t.Skip)
	obj.Set("SkipLine", t.SkipLine)
	obj.Set("Write", t.Write)
	obj.Set("WriteBlankLines", t.WriteBlankLines)
	obj.Set("WriteLine", t.WriteLine)
	return obj
}

type FileSystemObject struct {
	plugin herbplugin.Plugin
	Locker sync.Mutex
	opened map[string]*TextStream
}

func NewFileSystemObject() *FileSystemObject {
	return &FileSystemObject{
		opened: map[string]*TextStream{},
	}
}

func (o *FileSystemObject) closeAll() {
	o.Locker.Lock()
	defer o.Locker.Unlock()
	for _, v := range o.opened {
		v.File.Close()
	}
	o.opened = map[string]*TextStream{}
}
func (o *FileSystemObject) close(id string) {
	o.Locker.Lock()
	defer o.Locker.Unlock()
	f := o.opened[id]
	if f != nil {
		f.File.Close()
		delete(o.opened, id)
	}
}
func (o *FileSystemObject) open(fn string, mode int, create bool, r *goja.Runtime) goja.Value {
	o.Locker.Lock()
	defer o.Locker.Unlock()
	fn = strings.Replace(fn, "\\", "//", -1)
	var flag int
	fmode := util.DefaultFileMode
	switch mode {
	case 1:
		flag = os.O_RDONLY
	case 2:
		flag = os.O_WRONLY
	case 8:
		fmode = fmode | os.ModeAppend
	}
	if create {
		flag = flag | os.O_CREATE
	}
	f, err := os.OpenFile(fn, flag, fmode)
	if err != nil {
		panic(err)
	}
	id := uniqueid.MustGenerateID()
	ts := &TextStream{
		File: f,
		Closer: func() {
			o.close(id)
		},
	}
	o.opened[id] = ts
	return ts.Convert(r)
}
func (o *FileSystemObject) validatePath(p string) string {
	fn := o.plugin.PluginOptions().GetLocation().MustCleanInsidePath(p)
	if fn == "" {
		panic(herbplugin.NewUnauthorizePathError(p))
	}
	return fn
}
func (o *FileSystemObject) BuildPath(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(filepath.Join(call.Argument(0).String(), call.Argument(1).String()))
}
func (o *FileSystemObject) CopyFile(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	panic(errors.New("FileSystemObject.CopyFile not supported"))
}
func (o *FileSystemObject) CopyFolder(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	panic(errors.New("FileSystemObject.CopyFolder not supported"))
}
func (o *FileSystemObject) CreateFolder(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	panic(errors.New("FileSystemObject.CreateFolder not supported"))
}
func (o *FileSystemObject) CreateTextFile(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	fn := o.validatePath(call.Argument(0).String())
	create := call.Argument(1).ToBoolean()
	return o.open(fn, 0, create, r)
}
func (o *FileSystemObject) DeleteFile(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	panic(errors.New("FileSystemObject.DeleteFile not supported"))
}
func (o *FileSystemObject) DeleteFolder(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	panic(errors.New("FileSystemObject.DeleteFolder not supported"))
}
func (o *FileSystemObject) DriveExists(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(false)
}
func (o *FileSystemObject) FileExists(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	fn := o.validatePath(call.Argument(0).String())
	_, err := os.Stat(fn)
	if err != nil {
		if os.IsNotExist(err) {
			return r.ToValue(false)
		}
		panic(err)
	}
	return r.ToValue(true)
}
func (o *FileSystemObject) FolderExists(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	fn := o.validatePath(call.Argument(0).String())
	st, err := os.Stat(fn)
	if err != nil {
		if os.IsNotExist(err) {
			return r.ToValue(false)
		}
		panic(err)
	}
	return r.ToValue(st.IsDir())

}
func (o *FileSystemObject) GetAbsolutePathName(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	panic(errors.New("FileSystemObject.GetAbsolutePathName not supported"))
}
func (o *FileSystemObject) GetBaseName(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(filepath.Base(call.Argument(0).String()))
}
func (o *FileSystemObject) GetDrive(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	panic(errors.New("FileSystemObject.GetDrive not supported"))
}
func (o *FileSystemObject) GetDriveName(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(filepath.VolumeName(call.Argument(0).String()))
}
func (o *FileSystemObject) GetExtensionName(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(filepath.Ext(call.Argument(0).String()))
}
func (o *FileSystemObject) GetFile(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	panic(errors.New("FileSystemObject.GetFile not supported"))
}
func (o *FileSystemObject) GetFileName(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	list := filepath.SplitList(call.Argument(0).String())
	return r.ToValue(list[len(list)-1])

}
func (o *FileSystemObject) GetFolder(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	panic(errors.New("FileSystemObject.GetFolder not supported"))
}
func (o *FileSystemObject) GetParentFolderName(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	return r.ToValue(filepath.Dir(call.Argument(0).String()))
}
func (o *FileSystemObject) GetSpecialFolder(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	panic(errors.New("FileSystemObject.GetSpecialFolder not supported"))
}
func (o *FileSystemObject) GetTempName(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	panic(errors.New("FileSystemObject.GetTempName not supported"))

}
func (o *FileSystemObject) Move(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	panic(errors.New("FileSystemObject.Move not supported"))
}
func (o *FileSystemObject) MoveFile(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	panic(errors.New("FileSystemObject.MoveFile not supported"))
}
func (o *FileSystemObject) MoveFolder(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	panic(errors.New("FileSystemObject.MoveFolder not supported"))
}
func (o *FileSystemObject) OpenTextFile(call goja.FunctionCall, r *goja.Runtime) goja.Value {
	fn := o.validatePath(call.Argument(0).String())
	create := call.Argument(2).ToBoolean()
	return o.open(fn, int(call.Argument(1).ToInteger()), create, r)
}

func (o *FileSystemObject) Create(call goja.ConstructorCall) *goja.Object {
	name := call.Argument(0).String()
	if name != "Scripting.FileSystemObject" {
		panic(fmt.Errorf("unsupported object %s", name))
	}
	obj := call.This
	o.ConvertTo(obj)
	return obj
}
func (o *FileSystemObject) ConvertTo(obj *goja.Object) {
	obj.Set("BuildPath", o.BuildPath)
	obj.Set("CopyFile", o.CopyFile)
	obj.Set("CopyFolder", o.CopyFolder)
	obj.Set("CreateFolder", o.CreateFolder)
	obj.Set("CreateTextFile", o.CreateTextFile)
	obj.Set("DeleteFile", o.DeleteFile)
	obj.Set("DeleteFolder", o.DeleteFolder)
	obj.Set("DriveExists", o.DriveExists)
	obj.Set("FileExists", o.FileExists)
	obj.Set("FolderExists", o.FolderExists)
	obj.Set("GetAbsolutePathName", o.GetAbsolutePathName)
	obj.Set("GetBaseName", o.GetBaseName)
	obj.Set("GetDrive", o.GetDrive)
	obj.Set("GetDriveName", o.GetDriveName)
	obj.Set("GetExtensionName", o.GetExtensionName)
	obj.Set("GetFile", o.GetFile)
	obj.Set("GetFileName", o.GetFileName)
	obj.Set("GetFolder", o.GetFolder)
	obj.Set("GetParentFolderName", o.GetParentFolderName)
	obj.Set("GetSpecialFolder", o.GetSpecialFolder)
	obj.Set("GetTempName", o.GetTempName)
	obj.Set("Move", o.Move)
	obj.Set("MoveFile", o.MoveFile)
	obj.Set("MoveFolder", o.MoveFolder)
	obj.Set("OpenTextFile", o.OpenTextFile)
}
func NewFileSystemObjectModule(b *bus.Bus) *herbplugin.Module {
	fso := NewFileSystemObject()
	return herbplugin.CreateModule("jscript",
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			jsp := plugin.(*jsplugin.Plugin).LoadJsPlugin()
			fso.plugin = plugin
			r := jsp.Runtime
			r.Set("ActiveXObject", fso.Create)
			r.Set("CreateObject", fso.Create)
			next(ctx, plugin)
		},
		nil,
		func(ctx context.Context, plugin herbplugin.Plugin, next func(ctx context.Context, plugin herbplugin.Plugin)) {
			fso.closeAll()
		})
}
