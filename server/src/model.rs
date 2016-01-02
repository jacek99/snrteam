// developmental model

struct DevelopmentalModelArea {
	VERBAL,
	FLEXIBILITY,
	OTHER_1,
	OTHER_2
}

struct DevelopmentalModelGoal {
	goal_id: Long,
	area: DevelopmentalModelArea,
	stage: Int32,
	goal_name: String,
	goal_description: String
	goal_order: Int32 // order of goal within an area
}


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

struct User {
	user_id: Long,
	email: String,
	first_name: String,
	last_name: String,
    birth_date: Date,
    user_type: UserType,
    cell_phone: String,
    cell_phone_type: CellPlatform,
    home_phone: String,
    contact_preference: ContactPreference
}

struct Child {
	child_id: Long,
	first_name: String,
	birth_date: Date
}

struct ChildProgram {
	program_id: Long,
	child: Child,
	start_date: Date,
	photo: byte[]
}

struct ChildProgramDevelopmentalGoal {
	program_id: Long
	goal_id: Long,
	level_achieved: Int32, // 1-5
	completed: Boolean,
	date_started: Date,
	date_accomplished: Date
}

struct ChildProgramParticipant {
	program_id: Long,
	user_id: Long,
	active: Boolean,
	start_date: Date,
	end_date: Date // optional
}

struct Country {
	country_id: String,
	country_name: String
}

struct StateProvince {
	state_province_id: String,
	country_id: String,
	state_province_name: String,
	timezone: String
}

struct Language {
	language_id: String,
	name: Strings
}

struct Family {
	family_id: Long,
	name: String,
	programs: List<ChildProgram>,
	language: Language,
	country: Country,
	state_province: StateProvince,
	postal_code: String,
	public_visible: Boolean
	city: String,
}

struct PlaySession {
	session_id: Long,
	child_program_id: Long,
	user_id: Long,
	date: Date,
	start_time: Time,
	end_time: Time
	cancelled: Boolean
}

struct PlaySessionGame {

}

// communication

struct TeamMeeting {
	meeting_id: Long,
	program_ids: List<ChildProgram>,
	date_time: DateTime,
	duration_in_minutes: Int32,
	participants: List<ChildProgramParticipant>, // only those that actually participated
	meeting_notes: String
}

struct TeamMessage {
	message_id: Long,
	date_time: DateTime,
	text: String
	participans: List<ChildProgramParticipant>, // optional, if empty means ALL
}



// TODO: play session feedback based on 
enum PlaySessionFeedbackType {
	// TODO
	JOURNAL
}

struct PlaySessionNotes {
	session_id: Long,
	notes: String
}