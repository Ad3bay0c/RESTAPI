package Api_Jwt

import (
	"encoding/json"
	"fmt"
	"github.com/Ad3bay0c/WebTesting/db2"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Contact struct {
	ID			int64			`json:"id,omitempty"`
	Phone		string			`json:"phone,omitempty"`
	Name		string			`json:"name,omitempty"`
	CreatedAt	int64			`json:"created_at,omitempty"`
	UpdatedAt	int64			`json:"updated_at,omitempty"`
}

func GetAllContacts(w http.ResponseWriter, r *http.Request) {
	var contacts []Contact

	id := r.Context().Value("userId")
	userId, ok := id.(float64)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error Converting")
		json.NewEncoder(w).Encode(Message{Message: "Server Error!!!"})
		return
	}
	rows, err := db2.DB.Query("SELECT id, name, phone FROM contact WHERE user_id = $1", userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf(err.Error())
		json.NewEncoder(w).Encode(Message{Message: "Server Error!!!"})
		return
	}
	for rows.Next() {
		var contact Contact

		err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf(err.Error())
			json.NewEncoder(w).Encode(Message{Message: "Server Error!!!"})
			return
		}
		contacts = append(contacts, contact)
	}
	//id := r.Context().Value("userId")
	json.NewEncoder(w).Encode(contacts)
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	var contact Contact

	id := r.Context().Value("userId")

	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf(err.Error())
		json.NewEncoder(w).Encode(Message{Message: "Server Error!!!"})
		return
	}
	if contact.Phone == "" && contact.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Message: "Contact Name and Phone Number are Required"})
		return
	}
	if contact.Phone == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Message: "Phone Number is Required"})
		return
	}
	if contact.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Message{Message: "Contact Name is Required"})
		return
	}

	uId, ok := id.(float64)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error Converting Id")
		json.NewEncoder(w).Encode(Message{Message: "Server Error!!!"})
		return
	}
	userId := int64(uId)

	contact.ID 			= time.Now().Unix()
	contact.CreatedAt 	= time.Now().Unix()
	contact.UpdatedAt 	= time.Now().Unix()

	stmt, err := db2.DB.Prepare("INSERT INTO contact (id, phone, name, created_at, updated_at, user_id) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf(err.Error())
		json.NewEncoder(w).Encode(Message{Message: "Server Error!!!"})
		return
	}

	_, err = stmt.Exec(contact.ID, contact.Phone, contact.Name, contact.CreatedAt, contact.UpdatedAt, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf(err.Error())
		json.NewEncoder(w).Encode(Message{Message: "Server Error!!!"})
		return
	}
	//cID, _ := res.LastInsertId()

	json.NewEncoder(w).Encode(Message{Message: fmt.Sprintf("Contact Created Successfully")})
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	contactId := request["id"]

	id := r.Context().Value("userId")
	uId, ok := id.(float64)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error Converting Id")
		json.NewEncoder(w).Encode(Message{Message: "Server Error!!!"})
		return
	}
	userId := int64(uId)

	stmt, err := db2.DB.Prepare("DELETE FROM contact WHERE id = $1 AND user_id = $2")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf(err.Error())
		json.NewEncoder(w).Encode(Message{Message: "Server Error!!!"})
		return
	}

	res, err := stmt.Exec(contactId, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf(err.Error())
		json.NewEncoder(w).Encode(Message{Message: "Server Error!!!"})
		return
	}
	affected, _ := res.RowsAffected()
	json.NewEncoder(w).Encode(Message{Message: fmt.Sprintf("Contact Deleted SUccessfully; Rows Affected : %v", affected)})

}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	var contact Contact
	request := mux.Vars(r)
	contactId := request["id"]

	userId := r.Context().Value("userId")

	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf(err.Error())
		json.NewEncoder(w).Encode(Message{Message: "Server Error!!!"})
		return
	}

	stmt, err := db2.DB.Prepare("UPDATE contact SET name = $1, phone = $2 WHERE id = $3 AND user_id = $4")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf(err.Error())
		json.NewEncoder(w).Encode(Message{Message: "Server Error!!!"})
		return
	}

	_, err = stmt.Exec(contact.Name, contact.Phone, contactId, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf(err.Error())
		json.NewEncoder(w).Encode(Message{Message: "Server Error!!!"})
		return
	}

	json.NewEncoder(w).Encode(Message{Message: "Contact Updated Successfully"})
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Message{Message: "Update Contact Called"})

}
