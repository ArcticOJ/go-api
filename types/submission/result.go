package submission

type (
	ResponseType   = string
	ResultResponse struct {
		Type string      `json:"type"`
		Data interface{} `json:"data"`
	}
	CaseResult struct {
		ID       uint16  `mapstructure:"id"`
		Message  string  `mapstructure:"message"`
		Verdict  Verdict `mapstructure:"verdict"`
		Memory   uint32  `mapstructure:"memory"`
		Duration float32 `mapstructure:"duration"`
	}
	Verdict     string
	FinalResult struct {
		CompilerOutput string  `mapstructure:"compilerOutput"`
		Verdict        Verdict `mapstructure:"verdict"`
		Points         float32 `mapstructure:"points"`
	}
	Metadata struct {
		SubmissionID uint32  `mapstructure:"submissionID"`
		TestCount    uint16  `mapstructure:"testCount"`
		MaxPoints    float64 `mapstructure:"maxPoints"`
	}
)

const (
	VerdictNone                Verdict = ""
	VerdictAccepted            Verdict = "AC"
	VerdictPartiallyAccepted   Verdict = "PA"
	VerdictWrongAnswer         Verdict = "WA"
	VerdictInternalError       Verdict = "IE"
	VerdictRejected            Verdict = "RJ"
	VerdictCancelled           Verdict = "CL"
	VerdictRuntimeError        Verdict = "RTE"
	VerdictTimeLimitExceeded   Verdict = "TLE"
	VerdictMemoryLimitExceeded Verdict = "MLE"
	VerdictOutputLimitExceeded Verdict = "OLE"
	VerdictStackLimitExceeded  Verdict = "SLE"
	VerdictCompileError        Verdict = "CE"
)

var verdictStrings = map[Verdict]string{
	VerdictAccepted:            "Answer correct",
	VerdictWrongAnswer:         "Wrong answer",
	VerdictPartiallyAccepted:   "Partially accepted",
	VerdictInternalError:       "Internal error",
	VerdictRejected:            "Rejected",
	VerdictCancelled:           "Cancelled",
	VerdictRuntimeError:        "Runtime error",
	VerdictTimeLimitExceeded:   "Time limit exceeded",
	VerdictMemoryLimitExceeded: "Memory limit exceeded",
	VerdictOutputLimitExceeded: "Output limit exceeded",
	VerdictStackLimitExceeded:  "Stack limit exceeded",
	VerdictCompileError:        "Compile error",
}

func (v Verdict) LongString() string {
	if s, ok := verdictStrings[v]; ok {
		return s
	}
	return "<unknown>"
}

const (
	ResponseTypeCase     ResponseType = "case"
	ResponseTypeFinal    ResponseType = "final"
	ResponseTypeAck      ResponseType = "ack"
	ResponseTypeMetadata ResponseType = "metadata"
)
