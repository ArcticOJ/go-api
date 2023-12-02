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

var verdictStrings = map[Verdict]string{
	"AC":  "Answer correct",
	"WA":  "Wrong answer",
	"PA":  "Partially accepted",
	"IE":  "Internal error",
	"RJ":  "Rejected",
	"CL":  "Cancelled",
	"RTE": "Runtime error",
	"TLE": "Time limit exceeded",
	"MLE": "Memory limit exceeded",
	"OLE": "Output limit exceeded",
	"SLE": "Stack limit exceeded",
	"CE":  "Compile error",
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
