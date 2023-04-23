export type UserDataType = {
	ID : string;
	UserID: string;
	UserName: string;
	UserTitle: string;
	Email: string;
	Password: string;
	UserImage: string;
	Badges: {
		Reputation: number;
		Gold: number;
		Silver: number;
		Bronze: number;
	};
	Follower: string[];
	Following: string[];
	Location: string;
	MembershipTime: string;
	LastSeen: string;
	Aboutme: string;
	Mysite: string;
	Github: string;
	Twitter: string;
	Linkedin: string;
	TopTags: {
		Name: string;
		Score: number;
		NumberOfPost: number;
	}[];

	SelectedPanel: string;
	AccountType: string;
};

export type NotificationDataType = {
	type: string;
	time: string;
	title: string;
	description?: string;
	point?: number;
};

export type QuestionDataType = {
	ID: string;
	QuesTitle: string;
	QuesViewed: number;
	QuesUpvote: number;
	QuesDownvote: number;
	QuesBookmark: number;
	QuesTags: string[];
	QuesAnsAccepted: string;

	QuesAskedBy: string;
	QuesAskedTime: string;

	QuesEditedBy: string;
	QuesEditedTime: string;

	QuesDescription: string;
	QuesComment: string[];
	Answers: {
		ID: string;
		Vote: number;
		Comment: string[];
	}[];
};

export type CookieInfo1Type = {
	Email: string;
	Name: string;
	UserID: string;
	UUID: string;
	AccountType: string;
};

export type BlogListType = {
	ID: string;
	Title: string;
};

export type GroupListType = {
	ID: string;
	Name: string;
	Image: string;
	Description: string;
	CoverImage: string;
	NumberOfMember: number;
};
export type FavouriteHashListType = {
	ID: string;
	Name: string;
	Image: string;
	Description: string;
	CoverImage: string;
};

export type HotQuesListType = {
	ID: string;
	QuesTitle: string;
	ImageUrl: string;
};


export type RegistrationType = {
    UserName: string;
    Email: string;
    Password: string;
}