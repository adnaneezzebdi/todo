package httpapi

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/internal/tasks"
)

func Startserver() { //funzione che avvia il server http
	http.HandleFunc("/tasks", handleTasks)             //registra un handler per il percorso /tasks e quanfo arriva una richiesta /tasks, il server chiamera la funziown handleTask
	http.HandleFunc("/tasks/complete", handleComplete) //stesas logica
	http.HandleFunc("/tasks/delete", handleDelete)     //stessa logica

	println("Server HTTP attivo su http://localhost:8080") //fa capire a noi che il server sta per partire
	http.ListenAndServe(":8080", nil)                      //mette in ascolto un server su quella porta del computer(8080)
}

func handleTasks(w http.ResponseWriter, r *http.Request) { //tutto quello che scriviamo su w va al browser o postman, r è la richeista del client
	switch r.Method { //metodi get, post

	case http.MethodGet: //client vuole leggere i dati
		ts, err := tasks.List()
		if err != nil { //controlla se la lettura del json da errore
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(ts) //trasforma ts in json e le scrive nella risposta http

	case http.MethodPost: //client vuole aggiuingere nuove task
		w.Header().Set("Content-Type", "application/json") //indica che la risposta sara in formato json
		var body struct {                                  //serve per leggere il body della richiesta post
			Title string `json:"title"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil { //crea un lettore del body, trasfroma il json in oggetto go
			http.Error(w, "json non valido", http.StatusBadRequest) //se err!=nil da errore
			return
		}

		err := tasks.Add(body.Title) //crea una task, genre id, la aggiunge alla lista e salva nel json
		if err != nil {              //gestisce l'errore
			http.Error(w, err.Error(), 500) //errore 500 come risposta
			return
		}

		w.WriteHeader(201)                           //201 indica che la risorsa è stat creata con successo
		json.NewEncoder(w).Encode(map[string]string{ //creo un json
			"message": "task aggiunta!", //contenuto del json in risposta al client
		})
	}
}

func handleComplete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id") //prende il valore del parametro id dalla richiesta del client
	id, _ := strconv.Atoi(idStr)     //converto la stringa dell'id in un int(ignoro errore con _ che verrà catturato nella fuznione Complete)

	task, err := tasks.Complete(id) //chiedo al package di completare la task con quell'id, contiene la task aggiornata e errori
	if err != nil {                 //gestione errore
		http.Error(w, err.Error(), 404)
		return
	}

	json.NewEncoder(w).Encode(task) //prende la task completata, la trasfroma in json e la manda al client
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id") //prende il valore del parametro id dalla richiesta del client
	id, _ := strconv.Atoi(idStr)     //converto la stringa dell'id in un int

	err := tasks.Delete(id) //cerca una task tramite id nel file tasks.json,la rimuove dall'array e salva il nuovo array nel json
	if err != nil {         //gestisce l'errore
		http.Error(w, err.Error(), 404)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{ //operazione andata a buon fine e rispondo con un json
		"message": "task eliminata!", //messaggio che verrà mostrato al client
	})
}
