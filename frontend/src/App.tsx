import "./App.css";
import { Route, Routes } from "react-router";
import Layout from "./layouts/Layout";
import HomePage from "./pages/HomePage";
import LoginPage from "./pages/LoginPage";
import SignupPage from "./pages/SignupPage";
import ProfilePage from "./pages/ProfilePage";

function App() {
	return (
		<Routes>
			<Route element={<Layout />}>
				<Route path="/" element={<HomePage />} />
				<Route path="/login" element={<LoginPage />} />
				<Route path="/signup" element={<SignupPage />} />
				<Route path="/profile" element={<ProfilePage />} />
			</Route>
		</Routes>
	);
}

export default App;
