package main

import "fmt"

// Crear un programa para la gestión de tareas diarias
// Crear, modificar y eliminar tareas
// usuario será capaz de gestionar sus tareas

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarDeList(index int) {
	// con el primer argumento :index(exclusivo), se obtiene toda la lista excepto ese elemento
	// con el segundo argumento index+1:(inclusivo), se obtiene toda la lista excepto ese elemento
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) verListaTareas() {
	fmt.Println("*************** Lista de tareas *******************")
	for _, value := range t.tasks {
		fmt.Println(value)
	}
	fmt.Println("***************************************************")
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func main() {
	t1 := &task{
		nombre:      "Completar curso de orquídeas",
		descripcion: "Rociar orquídeas",
		completado:  false,
	}
	t2 := &task{
		nombre:      "Completar curso de rosas",
		descripcion: "Rociar rosas",
		completado:  false,
	}
	t3 := &task{
		nombre:      "Completar curso de claveles",
		descripcion: "Abonar claveles",
		completado:  false,
	}
	// lista de tareas
	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	lista.agregarALista(t3)
	fmt.Println(lista.tasks[2])
	lista.verListaTareas()

	lista.eliminarDeList(1)
	fmt.Println(len(lista.tasks))
	fmt.Println(lista.tasks[:1])

	//fmt.Printf("Nombre: %v, Descripción: %s, completado: %v\n", t1.nombre, t1.nombre, t1.completado)
	// fmt.Printf("%+v\n", t1)
	// t1.marcarCompleta()
	// t1.actualizarDescripcion("Comprar semillas")
	// fmt.Printf("%+v\n", t1)
	// implementando map con structs
	mapaTareas := make(map[string]*taskList)
	mapaTareas["Juan"] = lista
}
