type Account = {
	username: string;
	password: string;
};

type Profile = {
	userID: number;
	username: string;
	generateCount: number;
};

type SignupResponse = Account & {
	userID: number;
	token: string;
};

type LoginResponse = {
	userID: number;
	username: string;
	generateCount: number;
	token: string;
};