package recommendor

type IRecommendor interface {
}

type Recommendor struct {
	clubInterestVectors map[string][]float64
}

// func (r *Recommendor) RecommendClubs(interestVector []uint32) {

// 	for clubName, iv := range r.clubInterestVectors {

// 	}
// }
