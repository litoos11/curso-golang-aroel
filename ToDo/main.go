package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name        string
	description string
	complete    bool
}

type Tasks struct {
	tasks []Task
}

func (l *Tasks) addTask(t Task) {
	l.tasks = append(l.tasks, t)
}
func (l *Tasks) completeTask(index int) {
	l.tasks[index].complete = true
}

func (l *Tasks) editTask(index int, t Task) error {
	if index < 0 || index >= len(l.tasks) {
		return fmt.Errorf("El índice %d de la tarea no existe...", index)
	}
	l.tasks[index] = t
	return nil
}

func (l *Tasks) deleteTask(index int) {
	l.tasks = append(l.tasks[:index], l.tasks[index+1:]...)
}

func main() {
	tasks := Tasks{}
	reader := bufio.NewReader(os.Stdin)

	for {
		var option int
		fmt.Println("Seleccione una opción:\n",
			"1. Agregar\n",
			"2. Completar\n",
			"3. Editar\n",
			"4. Elminar\n",
			"5. Salir")

		fmt.Print("Ingrese la opción: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			var task Task
			fmt.Print("Ingrese el nombre de la tarea: ")
			task.name, _ = reader.ReadString('\n')
			fmt.Print("Ingrese la descripción de la tarea: ")
			task.description, _ = reader.ReadString('\n')
			tasks.addTask(task)
			fmt.Println("Tarea agregada correctamente.")
		case 2:
			var index int
			fmt.Print("Ingrese el índice de la tarea que desea marcar como completada: ")
			fmt.Scanln(&index)
			tasks.completeTask(index)
			fmt.Println("Tarea completada correctamente.")
		case 3:
			var index int
			var task Task
			fmt.Print("Ingrese el índice de la tarea que desea editar: ")
			fmt.Scanln(&index)
			fmt.Print("Ingrese el nombre de la tarea: ")
			task.name, _ = reader.ReadString('\n')
			fmt.Print("Ingrese la descripción de la tarea: ")
			task.description, _ = reader.ReadString('\n')
			if err := tasks.editTask(index, task); err != nil {
				fmt.Println("Error: ", err)
			} else {
				fmt.Println("Tarea editada correctamente.")
			}
		case 4:
			var index int
			fmt.Print("Ingrese el índice de la tarea que desea eliminar: ")
			fmt.Scanln(&index)
			tasks.deleteTask(index)
			fmt.Println("Tarea eliminada correctamente.")
		case 5:
			fmt.Println("Saliendo de las tareas...")
			return
		default:
			fmt.Println("Ingrese una opción valida...")
		}

		fmt.Println("Lista de tareas:")
		fmt.Println("===================================")
		for i, t := range tasks.tasks {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.name, t.description, t.complete)
		}
		fmt.Println("===================================")
	}

}
