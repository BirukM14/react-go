import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import { ChakraProvider, createSystem, defaultConfig } from "@chakra-ui/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const system = createSystem(defaultConfig);
// ✅ Properly define the theme

const queryClient = new QueryClient();

ReactDOM.createRoot(document.getElementById("root")!).render(
	<React.StrictMode>
		<QueryClientProvider client={queryClient}>
			<ChakraProvider value={system}>
				<App />
			</ChakraProvider>
		</QueryClientProvider>
	</React.StrictMode>
);
