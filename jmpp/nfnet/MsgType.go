package nfnet

type BaseMsg struct {
	MST int //send message ,milliseconds
	Msg string //message
	RB bool  //message back to the executive
	UL int //level 1 ~ 7 strong to weak
	V int //versions
}

//welcome
type  MW01 struct {
	BaseMsg
}

//
type MS01 struct {
	BaseMsg
}

//fetch ID
type MI01 struct {
	Id string
	Imei string
	Phone string
	BaseMsg
}
//Fetch ID request
type MG01 struct {
	BaseMsg
}

//error returned
type ME01 struct {
	CMD string
	BaseMsg
}
//message back to the executive
type MR01 struct {
	RL int  //message length
	BaseMsg
}

//error message
type ME02 struct {
	CMD string
	BaseMsg
}
