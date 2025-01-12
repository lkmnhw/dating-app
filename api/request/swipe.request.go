package request

type Swipe struct {
	TargetProfileID string `json:"target_profile_id" validate:"required"`
	Action          string `json:"action" validate:"required,oneof=like pass"`
}
