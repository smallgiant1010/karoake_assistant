import { useContext } from "react";
import { AuthContext, type AuthContextType } from "../context/authContext";

export function useAuthContext(): AuthContextType {
	const authContext = useContext(AuthContext);
	if (!authContext) throw new Error("useAuthContext must be used inside AuthContextProvider");
	return authContext;
}