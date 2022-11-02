package matching

type Result struct {
	Charset    string
	Language   string
	Confidence int
	BOM        bool
	LTR        bool
}

type ResultConfig func(r *Result)

func CreateResult(charset string, rc ...ResultConfig) Result {
	r := &Result{Charset: charset}

	for _, c := range rc {
		c(r)
	}

	return *r
}

func WithBOM(bv bool) ResultConfig {
	return func(r *Result) {
		r.BOM = bv
	}
}

func WithConfidence(confidence int) ResultConfig {
	return func(r *Result) {
		r.Confidence = confidence
	}
}

func WithLeftToRight(ltr bool) ResultConfig {
	return func(r *Result) {
		r.LTR = ltr
	}
}
