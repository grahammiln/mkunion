package projection

import (
	"github.com/widmogrod/mkunion/x/schema"
	"github.com/widmogrod/mkunion/x/storage/predicate"
	"math"
)

//go:generate go run ../../cmd/mkunion/main.go

//go:tag mkunion:"TriggerDescription"
type (
	//AtPeriod struct {
	//	Duration time.Duration
	//}
	//AtWindowItemSize struct {
	//	Number int
	//}
	AtWatermark struct{}
	//AnyOf struct {
	//	Triggers []TriggerDescription
	//}
	//AllOf struct {
	//	Triggers []TriggerDescription
	//}
)

func TriggerDescriptionToWhere(trigger TriggerDescription) (*predicate.WherePredicates, error) {
	return MatchTriggerDescriptionR2(
		trigger,
		func(x *AtWatermark) (*predicate.WherePredicates, error) {
			return predicate.Where(
				"Data.Window.End <= :watermark",
				predicate.ParamBinds{
					// Placeholder for watermark
					":watermark": schema.MkInt(math.MaxInt64),
				},
			)
		},
	)
}
