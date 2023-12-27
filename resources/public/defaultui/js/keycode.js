define(function (require) {
    var app = require("/public/defaultui/js/app.js")
    var vm = require("/public/defaultui/js/vm.js")
    app.onInputKeyDown = function (event) {
        switch (event.code) {
            case "NumpadDivide":
            case "NumpadMultiply":
            case "NumpadSubtract":
            case "NumpadAdd":
            case "NumpadDecimal":
            case "Numpad0":
            case "Numpad1":
            case "Numpad2":
            case "Numpad3":
            case "Numpad4":
            case "Numpad5":
            case "Numpad6":
            case "Numpad7":
            case "Numpad8":
            case "Numpad9":
                event.preventDefault()
                return false
                break

        }
    }
    let isDialogEvent=function(event){
        if (document.querySelector("body.el-popup-parent--hidden")){
            return true
        }
        for (var i=0;i<event.path.length;i++){
            if (event.path[i].className=="el-dialog__wrapper"){
                return true
            }
        }
        return false
    }
    document.addEventListener("keyup", function (event) {
        switch (event.code) {
            case "Backspace":
                if (!event.ctrlKey) {
                    return
                }
            case "Pause":
                app.send("change", "")
                break
            case "Digit1":
            case "Digit2":
            case "Digit3":
            case "Digit4":
            case "Digit5":
            case "Digit6":
            case "Digit7":
            case "Digit8":
            case "Digit9":
                if (isDialogEvent(event)){
                    return
                }
                if (vm.current == "") {
                    let client = vm.clients[event.code.slice(5) - 1]
                    if (client) {
                        app.send("change", client.ID)
                        event.preventDefault()
                    }
                }
                break
            case "Backquote":
                if (vm.current == "" || event.ctrlKey) {
                    vm.onButton('clientquick')
                    event.preventDefault()
                    return
                }
                break
            case "PageUp":
                if (vm.$refs.visualPromptSlide) {
                    vm.$refs.visualPromptSlide.prev()
                    event.preventDefault()
                    return
                }
                break
            case "PageDown":
                if (vm.$refs.visualPromptSlide) {
                    vm.$refs.visualPromptSlide.next()
                    event.preventDefault()
                    return
                }
                break
            case "Enter":
                if (vm.visualPrompt) {
                    vm.onButton('visualPromptSubmit')
                    vm.doFocus()
                    return
                }
                break
            case "ScrollLock":
                vm.onButton('clientquick')
                break
            case "NumpadDivide":
            case "NumpadMultiply":
            case "NumpadSubtract":
            case "NumpadAdd":
            case "NumpadDecimal":
            case "Numpad0":
            case "Numpad1":
            case "Numpad2":
            case "Numpad3":
            case "Numpad4":
            case "Numpad5":
            case "Numpad6":
            case "Numpad7":
            case "Numpad8":
            case "Numpad9":
                app.send("keyup", event.code)
                event.preventDefault()
                break
        }
        // console.log(event.code)
        // console.log(event.ctrlKey)
    })
})