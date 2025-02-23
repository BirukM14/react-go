import { Container, Stack } from '@chakra-ui/react'
import './App.css'
import Navbar from './components/ui/Navbar'
function App() {

  return (
    <Stack h= "100vh">

     <Navbar/>
     <Container>
      {/* <TodoForm/>
      <TodoForm/> */}

     </Container>
    </Stack>

    
  )
}

export default App
