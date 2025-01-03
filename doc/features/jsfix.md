# Javascript 修正

## eval

脚本里的eval已经不是标准的eval函数。而是调用Goja库中的RunScript函数，将一段代码以指定文件名的名义运行。

所以可以指定第二个字符串函数，方便进行Debug

## VBArray 对象

对[VBArray对象](https://documentation.help/MS-Office-JScript/jsobjVBArray.htm)做了简单模拟，提供了

* dimensions
* getitem
* lbound
* ubound
* toArray
方法

为部分老旧代码提供了兼容性支持

## Enumerator 对象

对[Enumerator对象](https://learn.microsoft.com/en-us/dotnet/api/microsoft.jscript.enumeratorobject?view=netframework-4.8.1)做了简单模拟，提供了

* atEnd
* moveFirst
* moveNext
* item
方法

为部分老旧代码提供了兼容性支持

## FileSystemObject 对象

对[FileSystemObject对象](https://documentation.help/MS-Office-JScript/jsobjFileSystem.htm)进行了模拟，实现了读取文本文件的功能，提供了

* CreateTextFile
* FileExists
* FolderExists
* GetDriveName
* GetExtensionName
* GetFileName
* GetParentFolderName
* OpenTextFile
* Create("Scripting.FileSystemObject")
的支持

实现了 TextStream 以及 TextStream的

* WriteLine
* WriteBlankLines
* Write
* SkipLine
* Skip
* ReadLine
* ReadAll
* Read
* Close
方法

为部分老旧代码提供了兼容性支持
