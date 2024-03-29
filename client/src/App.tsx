import { useEffect, useState } from "react";
import { Container, Card, Text, Button } from "@mantine/core";
import "./App.css";
import useSWR from "swr";
import AddTodo from "./components/AddTodo";
import { TrashIcon, CheckIcon } from "@primer/octicons-react";

export const ENDPOINT = "http://localhost:4000/api";
export const JSON_HEADERS = {
  "Content-Type": "application/json; charset=UTF-8",
};
export interface ITodo {
  id: number;
  title: string;
  done: boolean;
  body: string;
}

export type AddTodoFunc = (newTodo: ITodo) => void;

const fetcher = async (url: string) => {
  const resp = await fetch(`${ENDPOINT}${url}`);
  return await resp.json();
};

function App() {
  const { data, mutate, error } = useSWR<ITodo[]>("/todos", fetcher);

  useEffect(() => {
    if (error) {
      console.warn("Error: ", error.message);
    }
  }, [error]);

  const addTodo: AddTodoFunc = (newTodo) => {
    if (data) {
      mutate([...data, newTodo]);
    }
  };

  const markTodoAsDone = async (todo: ITodo) => {
    const resp = await fetch(`${ENDPOINT}/todos/${todo.id}/done`, {
      method: "PATCH",
      headers: JSON_HEADERS,
    });
    const doneTodo: ITodo = await resp.json();
    if (data && data.length) {
      todo.done = !todo.done;
      mutate([...data]); // Fix: Should pass a new array reference
    }
  };

  const doDelTodo = async (todo: ITodo) => {
    await fetch(`${ENDPOINT}/todos/${todo.id}`, {
      method: "DELETE",
      headers: JSON_HEADERS,
    });
  };

  const deleteTodo = async (todo: ITodo) => {
    await doDelTodo(todo);
    if (data && data.length) {
      mutate(data.filter((tdo) => tdo.id !== todo.id));
    }
  };

  return (
    <Container size="xs" className="app-container">
      <h2>Todo list:</h2>
      {data?.map((todo) => (
        <Card
          key={todo.id} // Fixed: Used todo.id as the key
          shadow="sm"
          padding="md"
          className="todo-card"
          style={{ marginBottom: "20px", border: "1px solid blue" }}
        >
          <div
            style={{
              display: "flex",
              justifyContent: "space-between",
              alignItems: "center",
            }}
          >
            <Text size="xl">{todo.title}</Text>
            <Button
              variant="filled"
              color="red"
              size="xs"
              onClick={() => deleteTodo(todo)}
              style={{ marginLeft: "10px" }}
            >
              <TrashIcon />
            </Button>
          </div>
          <Text size="sm" color={todo.done ? "teal" : "gray"}>
            {todo.body}
          </Text>
          <div style={{ marginTop: "10px", display: "flex", alignItems: "center" }}>
            <Button
              variant="outline"
              color="teal"
              size="xs"
              onClick={() => markTodoAsDone(todo)}
              style={{ marginRight: "10px" }}
            >
              {todo.done ? "Undo" : "Done"}
            </Button>
            {todo.done && <CheckIcon size={20} />}
          </div>
        </Card>
      ))}
      <AddTodo addTodo={addTodo} />
    </Container>
  );
}

export default App;
