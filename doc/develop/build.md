# 关于编译本程序

本程序需要使用cgo链接libprre(实现mush lua的rex部分,如不需要可以使用nopcre标签关闭该功能，无需cgo)。

理论上可以在绝大部分标准golang环境编译。

目前作者使用Linux环境进行编译/交叉编译并发布。具体环境为

* Debian testing(目前为12)
* Linux发布采用 musl (apt get install musl)进行静态编译，避免glibc版本冲突
* Windows发布采用 mingw-w64(apt get install mingw-w64) 发布

自行编译发布可参考 /src/build/ 下的各种编译脚本。