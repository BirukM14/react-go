import { Badge, Box, Flex, Text } from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import { useMutation } from "@tanstack/react-query";
import { BASE_URL } from "../../App";

export type Todo = {
    _id: number;
    body: string;
    completed: boolean;
};



const TodoItem = ({ todo }: { todo: Todo }) => {
    const { mutate: updateTodo, isPending: isUpdating } = useMutation({
        mutationKey: ["updateTodo"],
        mutationFn: async () => {
            if (todo.completed) return alert("Todo is already completed");

            try {
                const res = await fetch(`${BASE_URL}todos/${todo._id}`, {
                    method: "PATCH"
                });
                const data = await res.json();
                if (!res.ok) throw new Error(data.error || "Something went wrong");
            } catch (error) {
                console.error(error);
            }
        }
    });

    return (
        <Flex gap={2} alignItems={"center"}>
            <Flex
                flex={1}
                alignItems={"center"}
                border={"1px"}
                borderColor={"gray.600"}
                p={2}
                borderRadius={"lg"}
                justifyContent={"space-between"}
            >
                <Text
                    color={todo.completed ? "green.200" : "yellow.100"}
                    textDecoration={todo.completed ? "line-through" : "none"}
                >
                    {todo.body}
                </Text>
                <Badge ml="1" colorScheme={todo.completed ? "green" : "yellow"}>
                    {todo.completed ? "Done" : "In Progress"}
                </Badge>
            </Flex>
            <Flex gap={2} alignItems={"center"}>
                <Box color={"green.500"} cursor={"pointer"} onClick={() => updateTodo()}>
                    <FaCheckCircle size={20} />
                </Box>
                <Box color={"red.500"} cursor={"pointer"}>
                    <MdDelete size={25} />
                </Box>
            </Flex>
        </Flex>
    );
};

export default TodoItem;
