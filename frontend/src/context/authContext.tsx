import React, {
	createContext,
	type Context,
} from "react";

export interface AuthContextType {
	token: string;
	setToken: React.Dispatch<React.SetStateAction<string>>;
}

export const AuthContext: Context<AuthContextType | null> = createContext<AuthContextType | null>(null);