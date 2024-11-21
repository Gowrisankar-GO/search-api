package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"search_api/model"
	"strconv"
	"time"
)

func (d *Dependency) SearchUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet{

		w.WriteHeader(404)

		return
	}

	startTime := time.Now()

	name := r.URL.Query().Get("name")

	var offset int

	if r.URL.Query().Get("offset")==""{

		offset = Offset

	}else{

		offset,_ = strconv.Atoi(r.URL.Query().Get("offset"))
	}

	if name == "" {

		msg := "name query parameter is required"

		http.Error(w,msg,400)

		return
	}

	rows, err := d.DB.Query("select users.*,similarity(soundex(name),soundex($1)) as score from users where lower(name) ilike lower($2) order by score desc limit $3 offset $4",name, `%`+name+`%`, Limit, offset)

	if err != nil {

		http.Error(w,fmt.Sprint(err),500)

		return
	}

	users := []model.Users{}

	for rows.Next() {

		var user model.Users

		err = rows.Scan(&user.Id, &user.Name, &user.PhoneNumber, &user.Country,&user.Score)

		if err != nil{

			http.Error(w,fmt.Sprint(err),500)

			return
		}

		users = append(users, user)
	}

	var count int

	err = d.DB.QueryRow("select count(*) from users where lower(name) ilike lower($1)", `%`+name+`%`).Scan(&count)

	if err != nil {

		http.Error(w,fmt.Sprint(err),500)

		return
	}

	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(200)

	json.NewEncoder(w).Encode(Resp{Total: count, Results: users, TimeTakenMS: time.Since(startTime).Milliseconds()})

}
