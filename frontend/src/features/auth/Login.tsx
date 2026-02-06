import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useState, type SubmitEventHandler } from "react";
import { AuthenticateAccount } from "../../services/AuthService";

export default function Login() {
	const queryClient = useQueryClient();
	const [account, setAccount] = useState<Account>({
		username: "",
		password: "",
	});

	const { data, mutateAsync, status, error } = useMutation<
		Profile,
		Error,
		Account
	>({
		mutationFn: AuthenticateAccount,
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ["users"] });
		},
	});

	const handleSubmit: SubmitEventHandler<HTMLFormElement> = async (
		e: React.ChangeEvent<HTMLFormElement>,
	) => {
		e.preventDefault();
		await mutateAsync(account);
		if (status == "success") setAccount({ username: "", password: "" });
	};

	return (
		<section>
			<form method="POST" onSubmit={handleSubmit}>
				<fieldset>
					<legend>Login To Account</legend>
					<label>
						Username:
						<input
							type="text"
							name="username"
							min={4}
							value={account.username}
							onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
								setAccount((prev: Account) => {
									return {
										...prev,
										username: e.target.value,
									};
								})
							}
						/>
					</label>
					<label>
						Password:
						<input
							type="password"
							name="password"
							min={8}
							value={account.password}
							onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
								setAccount((prev: Account) => {
									return {
										...prev,
										password: e.target.value,
									};
								})
							}
						/>
					</label>
				</fieldset>
				<div>{status == "error" ? error.message : ""}</div>
				<button type="submit">
					{status == "pending" ? "Logging In..." : "Login"}
				</button>
			</form>
			<div>
				<h1>{data?.username}</h1>
				<h2>{data?.userID}</h2>
				<h2>{data?.generateCount}</h2>
			</div>
		</section>
	);
}
