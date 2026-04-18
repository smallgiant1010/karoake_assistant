import { useState, type ReactNode } from "react";
import { AuthContext } from "./authContext";

export function AuthContextProvider({ children }: { children: ReactNode }) {
	const [token, setToken] = useState<string>("");

	return (
		<AuthContext.Provider value={{ token, setToken }}>
			{children}
		</AuthContext.Provider>
	);
}