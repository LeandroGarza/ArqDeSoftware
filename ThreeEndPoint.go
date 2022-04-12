package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

package main

import (
"encoding/json"
"fmt"
"log"      //si hay error quiero ver por consola que ha pasado
"net/http" //importamos el modulo http desde net

//para poder manejar entradas y salidas de los datos que llegan al sv
"strconv" //convierte el string en entero

"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json: ID`     //respondemos con un id
	Name    string `json:Name`    //respondemos con un nombre
	Content string `json:Content` //respondemos con un  content
}

type allTasks []task //una lista de tareas

var tasks = allTasks{ //definimos una lista de todas las tareas comoo variable asi despues la puedo usar
	{
		ID:      1,
		Name:    "Task One",
		Content: "some content",
	},
}

//creamos rutas de endpoint
//esta permite obtener las tareas
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //para q cuando envio las tareas lo envio en formaro json
	json.NewEncoder(w).Encode(tasks)                   //em formato json retorna las tareas
}

//ruta que nos permite crear tareas
func createTasks(w http.ResponseWriter, r *http.Request) {
	var newTask task //recibimos lo que el cliente nos manda

	reqBody, err := ioutil.readAll(r.Body) //puedo recibir request o error. El r.body es lo que recibo que envia el cliente. El read all es un metodo de ioutil
	if err != nil {                        //si encontramos error
		fmt.Fprintf(w, "insert a valid task")
	}

	//en caso que la infromacion sea correcta, lo manipulamos
	json.Unmarshal(reqBody, &newTask) //se lo asigno a newtask lo que recibo del request body

	newTask.ID = len(tasks) + 1    //creo un ID autogenerado para asignarle a la tarea. El +1 es porq al ser uno nuevo, sera 1  mayor que el existente.
	tasks = append(tasks, newTask) //lo guardo en la lista de tareas. De las tareas hago una nueva tarea

	w.Header().Set("Content-Type", "application/json") //aclaramos que tipo de dato estoy enviando en este caso json. Un header es info extra por cada peticion que el usuario hace
	w.WriteHeader(http.StatusCreated)                  //para decirle que todo ha ido bien y q lo q ha enviado a sido aniadido a la lista
	json.NewEncoder(w).Encode(newTask)                 //respondo al cliente con la lista q acabo de crear

}

func getTask(w http.ResponseWriter, r *http.Request) { //obtiene una unica tarea
	//hacemos una busqued de una tarea, y una vez q la tenemos la devolvemos
	vars := mux.Vars(r)                     //esta variable guarda las variables de la peticion. El mux lo que hace es extraer las variables del request y lo asigna a vars
	taskID, err := strconv.Atoi(vars["id"]) //recibe string y convierte en entero. desde vars le pasamos el id y lo guardamos en taskID

	if err != nil {
		fmt.Fprintf(w, "invalid id")
		return
	}

	//pero si capturo un id o pudo convertir el entero, hacemos una busqueda recorriendo la lista y vemos si el q me envian coincide con uno de la lista
	for _, task := range tasks { //puedo recibir id o tarea de la lista tasks
		if task.ID == taskID { //si el q tengo es igual al  que me envia el cliente
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task) //lo devolvemos
		}
	}
	//vars["id"]	se llama id xq en la func main lo llamamos asi en el enrutador
}

func deleteTask(w http.ResponseWriter, r *http.Request){
	//hacemos una busqueda similar a la que hicimos en get task
	vars := mux.Vars(r)
	taskID, err strconv.Atoi(vars["id"])

	if err != nil{
		fmt.Fprintf(w, "invalid ID")
		return
	}

	for i, t := range tasks{	//obtenemos el indice i para luego saber cual eliminar
		if t.ID == taskID{
			tasks = append(tasks[:i], tasks[i + 1:]...)	//el :i es porq lo de antes de i en la lista lo voy a conservar. Ademas lo siguiente es ya que esto anterior lo voy a concatenar con lo que esta despues del indice
			fmt.Fprintf(w,"la tarea de ID %v ha sido removida", taskID)
		}
	}
}

func updateTask(w http.ResponseWriter, r *http.Request){
	//primero lo buscamos por el id
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updatedTask task //almacenamos los nuevos datos

	if err != nil{
		fmt.Fprintf(w, "invalid ID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w, "inserte datos validos")
	}
	json.Unmarshal(reqBody, &updatedTask)

	for i, t := range tasks{
		if t.ID == taskID{
			tasks = append(tasks[:i], tasks[i+1:]...) //quito tarea de antes
			updatedTask.ID = taskID	//una vez eliminado, voy a asignar a lo que me mandaron(en su id) el id de la tarea que el usuario me ha enviado
			tasks = append(tasks, updatedTask)	//lo inserto de nuevo en la lista

			fmt.Fprintf(w, "el task con id %v ha sido actualizada satisfactoriamente", taskID)
		}
	}
}

//w es la rta al cliente, r la info que el usuario a enviado
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bienvenido a mi API") //le respondemos esto
}

func main() {

	router := mux.NewRouter().StrictSlash(true) //primera ruta, quiero q escriba la url correcta por eso modo stricto

	router.HandleFunc("/", indexRoute)                       //cuando visites la url principal osea "/", quiero que ejecutes indexRoute
	router.HandleFunc("/tasks", getTasks).Methods("GET")     //cuando me piden las tareas respondo con la funcion get tasks. Funciona solo con la peticion get
	router.HandleFunc("/taks", createTasks).Methods("POST")  //nuevo enrutador para el post manejado con create task pero funciona solo con la peticion post
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET") //es por si queremos (desde el postman) solicitar una unica funcion y no todo el programa. Despues del tasks va a venir un nombre o algo q es el id
	router.HandleFunc("/tasks/{id}" , deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}" , updateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))          //pasamos el http, que escuche el puerto 3000 y el enrutador es router
	//fmt.Println("hello world")
}