import { NavLink } from "react-router";
import { useCookies } from "react-cookie";
import Microphone from "../assets/microphone.png";

function Header() {
	const [cookies] = useCookies(["jwt"]);
	const isLoggedIn = !!cookies.jwt; // boolean conversion

	return (
		<header>
			<NavLink to={"/"} end>
				<img src={Microphone} alt="Microphone Logo" width={32} height={32} />
			</NavLink>
			{isLoggedIn ? (
				<NavLink to={"/profile"}>Profile</NavLink>
			) : (
				<>
					<NavLink to={"/login"}>Log In</NavLink>
					<NavLink to={"/signup"}>Sign Up</NavLink>
				</>
			)}
		</header>
	);
}

export default Header;
