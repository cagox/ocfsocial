package session

import (
	"encoding/gob"
	"net/http"
	//"fmt"

	"github.com/gorilla/sessions"

	"github.com/cagox/fluxspells/common/config"
)

//Store is the session store.
var Store *sessions.CookieStore

func init() {
	//authKeyOne := securecookie.GenerateRandomKey(64)
	//encryptionKeyOne := securecookie.GenerateRandomKey(32)

	authKeyOne := []byte(config.Config.AuthKey)
	encryptionKeyOne := []byte(config.Config.EncKey)

	Store = sessions.NewCookieStore(authKeyOne, encryptionKeyOne)

	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 24 * 30, //Max Age 30 Days. This site is not exactly high risk.
		HttpOnly: true,
	}

	//Register necessary structs.
	gob.Register(Data{})
	gob.Register(Flash{})
	gob.Register(BasePageData{})
}

//Data is a the struct to move data between the session cookie and the program.
type Data struct {
	Email         string
	Authenticated bool
	Flashes       []Flash
}

/*
Flash will be used to add flash messages to the session cookie.
Class indicates the type of message, and will be used for CSS purposes.
Message is the message itself.

Classes should be of the following values:

error:   An Error messages
success: Success that warents notification.
info:    Useful information that might not be overly important.
warning: Warning the user that they should procede carefuly.

*/
type Flash struct {
	Class   string
	Message string
}

//NewSessionData returns a default SessionData struct.
func NewSessionData() Data {
	return Data{Email: "", Authenticated: false}
}

//GetSessionData grabs the SessionData struct from the cookie and returns it.
func GetSessionData(session *sessions.Session) Data {

	data := session.Values["sessiondata"]

	if data != nil {
		if page, ok := data.(Data); ok {
			//The cookie exists and is ok.
			return page
		}

		//The cookie exists but is not ok.
		return NewSessionData()
	}
	//The cookie doesn't exist.
	return NewSessionData()
}

//GetSession returns the session with the open cookie file.
func GetSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := Store.Get(r, config.Config.CookieName)

	if err != nil {
		//fmt.Println("err:", err)      //We don't actually care what the errors are, just that they exist.
		session.Values["sessiondata"] = Data{Email: "", Authenticated: false}
		session.Save(r, w)
		return session
	}
	if session == nil {
		//The session didn't exist.
		session.Values["sessiondata"] = Data{Email: "", Authenticated: false}
		session.Save(r, w)
		return session
	}
	return session
}

//AddFlash adds a flash message to the session.Data object
func (sessionData *Data) AddFlash(class string, message string) {
	flash := Flash{Class: class, Message: message}
	sessionData.Flashes = append(sessionData.Flashes, flash)
}

//GetFlashes will add the flash messages from the SessionData struct t the PageData struct.
func (sessionData *Data) GetFlashes(clearData bool) []Flash {
	flashes := sessionData.Flashes
	if clearData {
		sessionData.Flashes = make([]Flash, 0)
	}
	return flashes
}
