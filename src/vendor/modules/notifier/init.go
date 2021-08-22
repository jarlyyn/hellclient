package notifier

import (
	"fmt"
	"html"
	"modules/app"

	"github.com/herb-go/notification"
	"github.com/herb-go/notification-drivers/delivery/emaildelivery"
	"github.com/herb-go/util"
)

//ModuleName module name
const ModuleName = "900notifier"

type Notifier struct {
	URL  string
	SMTP *emaildelivery.SMTP
}

func (n *Notifier) WorldNotify(world string, title string, body string) {
	if n.SMTP == nil || n.SMTP.Host == "" || n.SMTP.To == "" {
		return
	}
	c := notification.NewContent()
	m, err := n.SMTP.NewEmail(c)
	if err != nil {
		util.LogError(err)
		return
	}
	m.Subject = title
	content := ""
	if n.URL != "" {
		content = fmt.Sprintf("<a href='%s#%s'>%s:</a>", n.URL, world, html.EscapeString(world))
	} else {
		content = fmt.Sprintf("%s:", world)
	}
	content += "<p>" + html.EscapeString(body) + "</p>"
	m.HTML = []byte(content)
	go func() {
		err := n.SMTP.Send(m)
		if err != nil {
			util.LogError(err)
		}
	}()
}
func New() *Notifier {
	return &Notifier{}
}

var DefaultNotifier = New()

func init() {
	util.RegisterModule(ModuleName, func() {
		//Init registered initator which registered by RegisterInitiator
		//util.RegisterInitiator(ModuleName, "func", func(){})
		util.InitOrderByName(ModuleName)
		DefaultNotifier.URL = app.System.URL
		DefaultNotifier.SMTP = app.System.SMTP
	})
}
