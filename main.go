package main

import (
	"fmt"
	"myapp/jsonconfig"
	"myapp/myhandlers"
	"net/http"
	"os"

	"github.com/justinas/alice"
)

func main() {
	var cfg jsonconfig.Configuration
	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", &cfg)

	appC := myhandlers.LoadAppContext(cfg.Database)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	commonHandlers := alice.New(myhandlers.LoggingHandler, myhandlers.RecoverHandler)
	http.Handle("/favicon.ico", commonHandlers.ThenFunc(myhandlers.Favicon))

	http.Handle("/signin", commonHandlers.ThenFunc(appC.Signin))
	http.Handle("/signout", commonHandlers.ThenFunc(appC.Signout))

	http.Handle("/incoming/answer", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.IncomingAnswer))
	http.Handle("/incoming/instrID", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.IncomingInstr))
	http.Handle("/incoming", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.Incoming))

	http.Handle("/outgoing/cancel", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.OutgoingCancel))
	http.Handle("/outgoing/rework", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.OutgoingRework))
	http.Handle("/outgoing/accept", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.OutgoingAccept))
	http.Handle("/outgoing/add", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.OutgoingAdd))
	http.Handle("/outgoing/instrID", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.OutgoingInstr))
	http.Handle("/outgoing", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.Outgoing))

	http.Handle("/", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.Index))

	// http.Handle("/reports", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.Reports))

	http.Handle("/admin", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.Admin))
	http.Handle("/admin/adduser", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.AdminAddUser))
	http.Handle("/admin/departments", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.AdminDepartments))
	http.Handle("/admin/adddept", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.AdminAddDept))
	http.Handle("/admin/roles", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.AdminRoles))
	http.Handle("/admin/addrole", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.AdminAddRole))
	http.Handle("/admin/statuses", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.AdminStatuses))
	http.Handle("/admin/addstatus", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.AdminAddStatus))

	http.Handle("/user", commonHandlers.Append(myhandlers.AuthHandler).ThenFunc(appC.User))

	fmt.Println("Listen port :8000 ...")
	http.ListenAndServe(":8000", nil)
}
