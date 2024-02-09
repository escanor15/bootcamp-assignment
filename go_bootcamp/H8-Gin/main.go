package main

import (
	// "github.com/docker/docker/libnetwork/drivers/host"
	// "go_bootcamp/routes"

	"go_bootcamp/H8-Gin/routes"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "04091997"
// 	dbname   = "bootcamp"
// )

// var (
// 	db  *sql.DB
// 	err error
// )

func main() {
	var PORT = ":8080"

	routes.StartServer().Run(PORT)

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Sucessfully Connected to a Database")

	// var employee = api.Employee{}

	// sqlStatement := `INSERT INTO employee (full_name, email, age, division)
	// VALUES ($1, $2, $3, $4)
	// Returning *
	// `

	// err = db.QueryRow(sqlStatement, "Airell Jordan", "airell@mail.com", 23, "IT").Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("New Employee Data : %+v", employee)

	// http.HandleFunc("/employees", getEmployees)
	// http.HandleFunc("/employee", createEmployees)
	// fmt.Println("Application is listening on port", PORT)
	// http.ListenAndServe(PORT, nil)

}

// func getEmployees(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json")

// 	if r.Method == "GET" {
// 		json.NewEncoder(w).Encode(api.Employees)
// 		return
// 	}

// 	http.Error(w, "Invalid method	", http.StatusBadRequest)
// }

// func getEmployees(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json")

// 	if r.Method == "GET" {
// 		tpl, err := template.ParseFiles("api/template.html")

// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		tpl.Execute(w, api.Employees)
// 		return
// 	}

// 	http.Error(w, "Invalid method	", http.StatusBadRequest)
// }

// func createEmployees(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json")

// 	if r.Method == "POST" {
// 		name := r.FormValue("name")
// 		age := r.FormValue("age")
// 		division := r.FormValue("division")

// 		convertAge, err := strconv.Atoi(age)

// 		if err != nil {
// 			http.Error(w, "Invalid Age", http.StatusBadRequest)
// 			return
// 		}

// 		newEmployee := api.Employee{
// 			ID:       len(api.Employees) + 1,
// 			Name:     name,
// 			Age:      convertAge,
// 			Division: division,
// 		}

// 		api.Employees = append(api.Employees, newEmployee)

// 		json.NewEncoder(w).Encode(newEmployee)
// 		return
// 	}

// 	http.Error(w, "Invalid method:", http.StatusBadRequest)
// }
