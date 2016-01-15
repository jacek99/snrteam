namespace go model

// developmental model
//struct DevelopmentalModelArea {
//	VERBAL,
//	FLEXIBILITY,
//	OTHER_1,
//	OTHER_2
//}
//
//struct DevelopmentalModelGoal {
//	goal_id: Long,
//	area: DevelopmentalModelArea,
//	stage: Int32,
//	goal_name: String,
//	goal_description: String
//	goal_order: Int32 // order of goal within an area
//}


// general
enum UserType {
	ADMIN,
	PARENT,
	FACILITATOR
}

enum CellPlatform {
	ANDROID,
	IOS,
	OTHER
}

enum ContactPreference {
	EMAIL,
	PHONE_TEXT,
	PHONE_NOTIFICATION
}

struct Date {
    1: required i16 year,
    2: required i16 month,
    3: required i16 day
}

struct User {
	1: required i64 user_id,
	2: required string user_name,
	3: required string email_address,
	4: required string first_name,
	5: required string last_name,
    6: required Date birth_date,
    7: required Date creation_date,
    8: optional binary photo
//    user_type: UserType,
//    cell_phone: String,
//    cell_phone_type: CellPlatform,
//    home_phone: String,
//    contact_preference: ContactPreference
}

struct Family {
    1: required i64 family_id,
    2: required set<User> administrators,
    3: required string name,
    4: required Date creation_date
}

struct Child {
    1: required i64 child_id,
    2: required string first_name
    3: required Date birth_date,
    4: required Date program_start_date,
    5: optional binary photo
}

//struct ChildProgram {
//	program_id: Long,
//	child: Child,
//	start_date: Date,
//	photo: byte[]
//}
//
//struct ChildProgramDevelopmentalGoal {
//	program_id: Long
//	goal_id: Long,
//	level_achieved: Int32, // 1-5
//	completed: Boolean,
//	date_started: Date,
//	date_accomplished: Date
//}
//
//struct ChildProgramParticipant {
//	program_id: Long,
//	user_id: Long,
//	active: Boolean,
//	start_date: Date,
//	end_date: Date // optional
//}
//
//struct Country {
//	country_id: String,
//	country_name: String
//}
//
//struct StateProvince {
//	state_province_id: String,
//	country_id: String,
//	state_province_name: String,
//	timezone: String
//}
//
//struct Language {
//	language_id: String,
//	name: Strings
//}
//
//struct Family {
//	family_id: Long,
//	name: String,
//	programs: List<ChildProgram>,
//	language: Language,
//	country: Country,
//	state_province: StateProvince,
//	postal_code: String,
//	public_visible: Boolean
//	city: String,
//}
//
//struct PlaySession {
//	session_id: Long,
//	child_program_id: Long,
//	user_id: Long,
//	date: Date,
//	start_time: Time,
//	end_time: Time
//	cancelled: Boolean
//}
//
//struct PlaySessionGame {
//
//}
//
//// communication
//
//struct TeamMeeting {
//	meeting_id: Long,
//	program_ids: List<ChildProgram>,
//	date_time: DateTime,
//	duration_in_minutes: Int32,
//	participants: List<ChildProgramParticipant>, // only those that actually participated
//	meeting_notes: String
//}
//
//struct TeamMessage {
//	message_id: Long,
//	date_time: DateTime,
//	text: String
//	participans: List<ChildProgramParticipant>, // optional, if empty means ALL
//}
//
//
//
//// TODO: play session feedback based on
//enum PlaySessionFeedbackType {
//	// TODO
//	JOURNAL
//}
//
//struct PlaySessionNotes {
//	session_id: Long,
//	notes: String
//}
//
