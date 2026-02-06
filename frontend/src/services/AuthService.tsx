import { BACKEND_API } from "../config";

export async function CreateNewAccount(acc: Account): Promise<Credentials> {
	const response = await fetch(BACKEND_API + "/auth/add", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(acc),
	});

	if (!response.ok) throw new Error(await response.text());
	return response.json();
}

export async function AuthenticateAccount(acc: Account): Promise<Profile> {
	const response = await fetch(BACKEND_API + "/auth/login", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify(acc),
	});

	if (!response.ok) throw new Error(await response.text());
	return response.json();
}
