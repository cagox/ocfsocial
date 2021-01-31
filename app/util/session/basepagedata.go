package session

import (
	"github.com/cagox/fluxspells/app/user"
)

//BasePageData is the data that most pages will need. This can be used to build the data struct for templates.
type BasePageData struct {
	Page          string
	Flashes       []Flash
	Authenticated bool
	IsAdmin       bool
}

//BasicData fills in the BasePageData struct from the provided session.
func (data *BasePageData) BasicData(session Data) {
	data.Authenticated = session.Authenticated

	if data.Authenticated {
		user := user.GetUserByEmail(session.Email)
		//config.Config.Database.Where("id = ?", session.UserID).First(&user)
		data.IsAdmin = user.IsAdmin
	}

}
