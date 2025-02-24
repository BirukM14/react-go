import { Box } from '@chakra-ui/react'
import { useColorModeValue } from './components/ui/color-mode'
import Navbar from './components/ui/Navbar'
import TodoForm from './components/ui/TodoForm'                                 
import TodoList from './components/ui/TodoList'

function App() {
  return (
    // Full-screen box with dynamic background
    <Box 
      bg={useColorModeValue("white", "gray.900")} 
      minH="90vh" // Ensures full height
      w="50vw"    // Ensures full width
      display="flex"
      flexDirection="column"
    >
      <Navbar />
      <TodoForm/>
      <TodoList/>
    </Box>
  );
}

export default App;