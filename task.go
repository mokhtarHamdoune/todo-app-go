package main

import (
	"database/sql"
)

type Task struct{
    Id int;
    Content string;
    IsDone bool; 
}

type TaskManager struct {
  database *sql.DB
}


func (tm TaskManager) save(t Task) (Task, error){
  query := `INSERT INTO TASKS (content, is_done) VALUES (?, ?)`
  result, err := tm.database.Exec(query,t.Content, t.IsDone)
	if err != nil {
		return Task{}, err
	}
	id,err :=  result.LastInsertId()
	if err != nil {
		return Task{}, err
	}
	t.Id = int(id)
  return t, nil; 
}

func (tm TaskManager) update(taskId int, upatedTask Task) error{
	query := `UPDATE TASKS`
	if upatedTask.Content != "" {
		query = query +  `SET content=?`
	} 

	if upatedTask.IsDone {
		query = query +  `SET content=?`
	} 
	// TODO: you should understand spreading in go I want to spread the args according to 
	// the passed updatedTask
   _, err := tm.database.Exec(query,upatedTask.Content, upatedTask.IsDone)
  return err
}

func (tm TaskManager) delete(taskId int) error{
	query := `DELETE FROM TASKS WHERE id=?`
	_, err := tm.database.Exec(query,taskId)
	return err
}

func(tm TaskManager) get(taskId int) (Task, error){
  query := `SELECT id, content, is_done FROM TASKS WHERE id=?`
	task := Task{}
	err := tm.database.QueryRow(query,taskId).Scan(task.Id,task.Content,task.IsDone)
	return task, err;	
}

func(tm TaskManager) getAll()([]Task,error) {
	query := `SELECT id, content, is_done FROM TASKS`
	rows, err := tm.database.Query(query)
	var tasks []Task
	if err != nil {
		return tasks , err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id,&task.Content, &task.IsDone)
		if err != nil {
			emptyList := []Task{}
			return emptyList , err
		}
		tasks = append(tasks, task)
	}

	return tasks, err
}