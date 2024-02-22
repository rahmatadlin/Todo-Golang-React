import { useEffect } from 'react';
import { Container, List, ThemeIcon } from '@mantine/core';
import './App.css';
import useSWR from 'swr';
import AddTodo from './components/AddTodo';
import { CheckCircleFillIcon, TrashIcon } from '@primer/octicons-react';

export const ENDPOINT = import.meta.env.VITE_BACKEND_URL;
export const JSON_HEADERS = { 'Content-Type': 'application/json; charset=UTF-8' };
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
  const { data, mutate, error } = useSWR<ITodo[]>('/todos', fetcher);

  useEffect(() => {
    if (error) {
      console.warn('Error: ', error.message);
    }
  }, [error]);

  const addTodo: AddTodoFunc = (newTodo) => {
    if (data) {
      mutate([...data, newTodo]);
    }
  };

  const markTodoAsDone = async (todo: ITodo) => {
    const resp = await fetch(`${ENDPOINT}/todos/${todo.id}/done`, {
      method: 'PATCH',
      headers: JSON_HEADERS,
    });
    const doneTodo: ITodo = await resp.json();
    if (data && data.length) {
      todo.done = !todo.done;
      mutate(data);
    }
  };

  const doDelTodo = async (todo: ITodo) => {
    await fetch(`${ENDPOINT}/todos/${todo.id}`, {
      method: 'DELETE',
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
      <List spacing="xs" size="sm" mb={12} center>
        {data?.map((todo) => (
          <List.Item
            key={`todo_list__${todo.id}`}
            icon={
              <ThemeIcon
                color={todo.done ? 'teal' : 'gray'}
                size={22}
                radius="xl"
                onClick={() => markTodoAsDone(todo)}
                style={{ cursor: 'pointer' }}
              >
                <CheckCircleFillIcon />
              </ThemeIcon>
            }
            title={todo.body}
          >
            {todo.title}
            <ThemeIcon
              color="gray"
              size={24}
              onClick={() => deleteTodo(todo)}
              style={{ marginLeft: '10px', cursor: 'pointer' }}
            >
              <TrashIcon />
            </ThemeIcon>
          </List.Item>
        ))}
      </List>
      <AddTodo addTodo={addTodo} />
    </Container>
  );
}

export default App;
