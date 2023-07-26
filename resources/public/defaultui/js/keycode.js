define(function (require) {
    var app = require("/public/defaultui/js/app.js")
    var vm = require("/public/defaultui/js/vm.js")
    document.addEventListener("keyup", function (event) {
        switch (event.code) {
            case "Escape":
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
                if (vm.current==""){
                    let client=vm.clients[event.code.slice(5)-1]
                    if (client){
                        app.send("change",client.ID)
                    }
                }
                break
            case "Backquote":
                if (vm.current==""){
                }
                break
        }
        console.log(event.code)
        console.log(event.ctrlKey)
    })
})