export type RegistrationType = {
	UserName: string;
	Email: string;
	Password: string;
};
export type CookieInfo1Type = {
	Email: string;
	Name: string;
	UserID: string;
	UserURL: string;
	AccountType: string;
};
export type UserDataType = {
	UserID: string;
	UserURL: string;
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
export type UserFlairDataType = {
	UserID: string;
	UserURL: string;
	UserName: string;
	UserImage: string;
	UserTitle: string;
	Badges: {
		Reputation: number;
		Gold: number;
		Silver: number;
		Bronze: number;
	};
	Location: string;
	Aboutme: string;
};

export type QuestionDataType = {
	ID: string;
	QuesTitle: string;
	QuesViewed: number;
	QuesUpvote: number;
	QuesDownvote: number;
	QuesVotes: {
		ID : string;
		UserID : string;
		Vote : number;
		VoteTime : string;
	}[]
	QuesBookmark: {
		ID : string;
		UserID : string;
		BookmarkTime : string;
		
	}[]
	QuesTags: string[];
	QuesGroup: string[];
	QuesAnsAccepted: string; // accepted answer ID
	
	QuesAskedBy: string;
	QuesAskedTime: string;

	QuesEditedBy: string;
	QuesEditedTime: string;
	
	QuesDescription: string;
	QuesComment: string[];
	Answers: {
		ID: string;
		UpVote: number;
		DownVote: number;
		AnsweredBy: string 
		Comment: string[];
	}[];
};

export type AnswerDataType= {
    ID: string;
	QuesID : string 
    AnsweredBy: string;
    AnsweredTime: string;
    EditedBy: string;
    EditedTime: string;
    Upvote: number;
    Downvote: number;
	Votes: {
		ID : string;
		UserID : string;
		Vote : number;
		VoteTime : string;
	}[]
    Bookmark: {
		ID : string;
		UserID : string;
		BookmarkTime : string;
	}[]
    Accepted: boolean;
    Description: string;
    Comment: {
		ID: string;
        UserID: string;
        UserName: string;
        Upvote: number;
        Downvote: number;
        CommentTime: string;
        Comment: string;
    }[];
}


export type BlogDataType = {
	ID     : string 
	Title     :  string 
	MetaTitle : string 
	Image :	 string 
	Description : string 
	WrittenBy : 	string[] 
	WrittenTime : string 
	EditedBy : string
	EditedTime : string 
	Tags 	:  string []
	Comments : string []
	Upvote 	 : number
	Downvote  : number
	Votes 	 : {
		ID    : string 
		UserID	 : string 
		Vote : number 
		VoteTime : string 
	
	}[]
	Bookmark : {
		ID 	: string 
		UserID	  : string 
		BookmarkTime : string 
	}[]
	Views 	 : number
	ViewedBy  : string [] 
}


export type QuesArrWithMetadataType = {
	QuesList: QuestionDataType[]
	Metadata: {
		Length:number 
	}
}
export type HashDataType ={
	ID : string
	Name : string
	MetaTitle : string
	Image : string
	BannerImage : string
	NumOfFollower : number 
	NumOfQuestion : number
	NumOfBlog : number
	NumOfAnswer : number
	About : string

}
export type HashArrWithMetadata ={
	HashList : HashDataType[]
	Metadata: {
		Length:number
	}
}
//////////////////////////////////////

// export type BlogListType = {
// 	ID: string;
// 	Title: string;
// };

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





export type NotificationDataType = {
	type: string;
	time: string;
	title: string;
	description?: string;
	point?: number;
};