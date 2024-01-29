package dto



type CreateLockerRequest struct { 
	Reason string `json:"reason"`
	Responsible string `json:"responsible"`
}

type DisableLockerRequest struct { 
	Responsible string `json:"responsible"`
}