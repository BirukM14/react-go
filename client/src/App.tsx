import { Box, Flex } from "@chakra-ui/react";
import { useColorModeValue } from "./components/ui/color-mode";
import Navbar from "./components/ui/Navbar";
import TodoForm from "./components/ui/TodoForm";
import TodoList from "./components/ui/TodoList";

function App() {
  return (
    // Outer Flex container to center the Box
    <Flex
      justify="center"
      align="center"
      minH="100vh" // Ensures full screen height
      bg={useColorModeValue("gray.100", "gray.800")} // Background color
    >
      {/* Centered Box */}
      <Box
        bg={useColorModeValue("white", "gray.900")}
        minH="90vh" // Minimum height
        w="50vw" // Width of the box
        display="flex"
        flexDirection="column"
        alignItems="center" // Centers children horizontally
        boxShadow="xl" // Adds a shadow effect
        p={8} // Adds padding
        borderRadius="lg" // Rounded corners
      >
        {/* Navbar should always be at the top */}
        <Box w="100%" mb={4}>
          <Navbar />
        </Box>

        {/* Remaining content will take full space */}
        <Flex flexGrow={1} flexDirection="column" w="100%" align="center">
          <TodoForm />
          <TodoList />
        </Flex>
      </Box>
    </Flex>
  );
}

export default App;
