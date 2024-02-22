import { Button, Group, Modal, Textarea, TextInput } from "@mantine/core";
import { useForm } from "@mantine/form";
import { useState } from "react";
import { AddTodoFunc, ENDPOINT, JSON_HEADERS } from "../App";

const AddTodo = ({ addTodo }: { addTodo: AddTodoFunc }) => {
  const [open, setOpen] = useState(false);

  const form = useForm({
    initialValues: {
      title: "",
      body: "",
    },
  });

  const fetchNewTodo: any = async (values: { title: string; body: string }) => {
    const resp = await fetch(`${ENDPOINT}/todos`, {
      method: `POST`,
      headers: JSON_HEADERS,
      body: JSON.stringify(values),
    });

    return await resp.json();
  };

  const createTodo = async (values: { title: string; body: string }) => {
    const newTodo = await fetchNewTodo(values);
    addTodo(newTodo);
    form.reset();
    setOpen(false);
  };

  return (
    <>
      <Modal opened={open} onClose={() => setOpen(false)} title="Create Todo">
        <form onSubmit={form.onSubmit(createTodo)}>
          <TextInput
            required
            mb={12}
            label="Todo"
            placeholder="What do you want todo..."
            {...form.getInputProps("title")}
          />
          <Textarea
            required
            mb={12}
            label="Body"
            placeholder="Tell me more..."
            {...form.getInputProps("body")}
          />
          <Button type="submit">Create todo</Button>
        </form>
      </Modal>
      <Group
        style={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        <Button fullWidth mb={12} onClick={() => setOpen(true)}>
          Add todo
        </Button>
      </Group>
    </>
  );
};

export default AddTodo;
