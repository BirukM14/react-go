import { Box } from '@chakra-ui/react'
import { useColorModeValue } from './components/ui/color-mode'
import './App.css'
import Navbar from './components/ui/Navbar'

function App() {
  return (
    // Full-screen box with dynamic background
    <Box 
      bg={useColorModeValue("white", "gray.900")} 
      minH="100vh" // Ensures full height
      w="100vw"    // Ensures full width
      display="flex"
      flexDirection="column"
    >
      <Navbar />
      {/* Other components go here */}
    </Box>
  );
}

export default App;