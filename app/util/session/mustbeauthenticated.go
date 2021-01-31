package session

import (
	"net/http"
)

//MustBeAuthenticated is a wrapper function that makes sure a user is authenticated before letting them access a page.
func MustBeAuthenticated(handler http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		session := GetSession(w, r)
		sessionData := GetSessionData(session)

		if !sessionData.Authenticated {
			sessionData.AddFlash("error", "You must be logged in to access that page.")
			session.Values["sessiondata"] = sessionData
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		handler(w, r)
	}

}
