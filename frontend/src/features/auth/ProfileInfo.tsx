import { useQuery } from "@tanstack/react-query";
import { useCookies } from "react-cookie";
import { GetProfile } from "../../services/AuthService";

export default function Profile() {
	const [cookies] = useCookies(["jwt"]);
	const token = cookies.jwt || "";

	const { data, status, error } = useQuery({
		queryKey: ["profile"],
		queryFn: () => GetProfile(token),
		enabled: token !== "",
	});

	if (!token) {
		return (
			<section>
				<h1>Please log in to view your profile</h1>
			</section>
		);
	}

	return (
		<section>
			{status === "pending" && <p>Loading profile...</p>}
			{status === "error" && <p>Error: {error.message}</p>}
			{status === "success" && data && (
				<div>
					<h1>Profile</h1>
					<p><strong>Username:</strong> {data.username}</p>
					<p><strong>User ID:</strong> {data.userID}</p>
					<p><strong>Generate Count:</strong> {data.generateCount}</p>
				</div>
			)}
		</section>
	);
}