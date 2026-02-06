type Account = {
	username: string;
	password: string;
};

type Profile = {
	userID: number;
	username: string;
	generateCount: number;
};

type Credentials = Account & {
	userID: number;
}
