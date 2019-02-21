var handlers={}
var onButton={}

var ws=new WebSocket("ws:"+location.host+"/ws");
        
var convertmsg=function(data){
    var sep=data.indexOf(" ")
    var msg
    if (sep>0){
        msg={
            type:data.substr(0,sep),
        }
        var msgdata =data.substr(sep+1)
        if (msgdata){
        msg.data=JSON.parse(msgdata)
        }
    }else{
        msg={
            type:data
        }
    }

    return msg
}
ws.onclose=function(){
    vm.$alert('你与程序被断开了，可能是通过别的页面打开/程序发生错误/程序死机，需要重连才能继续操作', '连接断开', {
  confirmButtonText: '重新连接',
  showClose:false,
  callback: function(){
    location.reload()
  }
});
}
ws.onmessage = function(event) {
    data=convertmsg(event.data);
    var f=handlers[data.type];
    if (f){
        f(data.data)
    }
}
var send=function(cmd,data){
    if (ws.readyState){
        var msg=JSON.stringify(data)
        console.log(msg)
        ws.send(cmd+" "+JSON.stringify(data))
    }
}