package main

import (
	"net/http"
)

func main() {
	/*
			c := db.C("siafiles")
			err = c.Insert(&Siafile{SiafilePath: "sia/1231231", Title: "First File"})

			if err != nil {
				log.Fatal(err)
			}

		result := new(Siafile)
		err = c.Find(bson.M{"title": "First File"}).One(&result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("SiaFiles:", result.SiafilePath)
	*/

	sf := New()
	defer sf.Close()
	http.ListenAndServe(":10181", sf.Router)
}
