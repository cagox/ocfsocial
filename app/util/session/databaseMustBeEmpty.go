package session

import (
	"net/http"

	"github.com/cagox/fluxspells/app/user"
)

//DatabaseMustBeEmpty is a wrapper funtion to make sure the database is empty.
func DatabaseMustBeEmpty(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := GetSession(w, r)
		sessionData := GetSessionData(session)

		//fmt.Println("Checking if the Database is Empty")

		if user.AreThereAnyUsers() {
			//The database is not empty.
			sessionData.AddFlash("error", "A user already exists in the database.")
			session.Values["sessiondata"] = sessionData
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		//The database is empty
		handler(w, r)
	}

}
